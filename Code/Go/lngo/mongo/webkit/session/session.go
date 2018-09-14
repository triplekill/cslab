package session

import (
	"context"
	"crypto/hmac"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/exfly/cslab/Code/Go/lngo/mongo/webkit/mongoplus"
	"github.com/exfly/cslab/Code/Go/lngo/mongo/webkit/utils"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Session struct {
	ID        string      `bson:"_id"`
	UserID    *string     `bson:"userId,omitempty"`
	Expire    time.Time   `bson:"expire"`
	Created   time.Time   `bson:"created"`
	SessionID string      `bson:"sessionId"`
	Data      mongoplus.M `bson:"data,omitempty"`
	dirty     bool        `bson:"-"`
}

const hmac_length = 32

type User interface {
	GetUserID() string
}

type SessionConfig struct {
	Collection     *mongo.Collection
	Key            string
	Secret         []byte
	LoginedTimeout time.Duration
	GuestTimeout   time.Duration
	UserLoader     func(userId string) (User, error)
}

var cfg *SessionConfig

func SetConfig(c *SessionConfig) {
	cfg = c
}

func Set(ctx context.Context, key string, value interface{}) {
	session := ctxSession(ctx)
	session.Data[key] = value
	session.dirty = true
}

func Get(ctx context.Context, key string) (value interface{}) {
	session := ctxSession(ctx)
	return session.Data[key]
}

func ctxSession(ctx context.Context) *Session {
	session, _ := ctx.Value("session").(*Session)
	return session
}

func SetSessionUser(w http.ResponseWriter, ctx context.Context, user User) {
	if user == nil {
		log.Println("warn: try set session with nil user")
		return
	}
	session := ctxSession(ctx)
	if session.UserID == nil {
		setSecureCookie(cfg.Key, session.SessionID, 30*24*3600, w)
	}
	userID := user.GetUserID()
	session.UserID = &userID
	session.dirty = true
}

func setSecureCookie(name, payload string, maxage int, w http.ResponseWriter) {
	toHash := []byte(fmt.Sprintf("%s|%s", name, payload))
	hash := utils.Sha256HMAC(toHash, cfg.Secret)
	final := append([]byte(payload), hash...)
	cookie := http.Cookie{
		Name:   name,
		Value:  base64.StdEncoding.EncodeToString(final),
		MaxAge: maxage,
		Path:   "/",
		// Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func getSecureCookie(name string, r *http.Request) (payload string) {
	payload = ""
	cookie, err := r.Cookie(name)
	if err != nil {
		return
	}
	// check if expired
	h, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return
	}
	payloadByte := h[:len(h)-hmac_length]
	toHash := []byte(fmt.Sprintf("%s|%s", name, payloadByte))
	if hmac.Equal(utils.Sha256HMAC(toHash, cfg.Secret), h[len(h)-hmac_length:len(h)]) {
		payload = string(payloadByte)
	}
	return
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session := ctxSession(r.Context())
	session.Expire = time.Now().Add(10 * time.Minute)
	session.UserID = nil
	cfg.Collection.ReplaceOne(
		r.Context(),
		mongoplus.M{"_id": session.ID},
		session,
	)
	setSecureCookie(cfg.Key, session.SessionID, int(cfg.GuestTimeout.Seconds()), w)
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			session Session
			err     error
		)
		ctx := r.Context()
		defer func() {
			if err != nil {
				// raven.CaptureError(err, nil)
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				// FIXME: redirect to login page ?
			}
		}()
		sessionId := getSecureCookie(cfg.Key, r)

		if sessionId != "" {
			cur := cfg.Collection.FindOne(ctx, mongoplus.M{"sessionId": sessionId})
			cur.Decode(&session)
		}
		// if session is not in database, create one
		if session.ID == "" {
			now := time.Now()
			session = Session{
				ID:        objectid.New().Hex(),
				Created:   now,
				SessionID: utils.RandomString(24),
				Expire:    now.Add(time.Minute * 10),
				dirty:     true,
				Data:      make(map[string]interface{}),
			}
			_, err := cfg.Collection.InsertOne(ctx, session)
			if err != nil {
				return
			}
			setSecureCookie(cfg.Key, session.SessionID, int(cfg.GuestTimeout.Seconds()), w)
		} else {
			if session.Data == nil {
				session.Data = make(mongoplus.M)
			}
		}

		ctx = context.WithValue(r.Context(), "session", &session)

		if session.UserID != nil {
			user, err := cfg.UserLoader(*session.UserID)
			if err == nil {
				ctx = context.WithValue(ctx, "userId", *session.UserID)
				ctx = context.WithValue(ctx, "user", user)
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))

		if session.dirty {
			_, err := cfg.Collection.ReplaceOne(ctx, mongoplus.M{"_id": session.ID}, session)
			if err != nil {
				log.Println(err)
				// raven.CaptureError(err, nil)
			}
		}
	})
}
