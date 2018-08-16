package main

import (
	"context"
	"log"
	"time"

	"github.com/globalsign/mgo"
)

func main() {
	// mongo.Test1()
	start := time.Now()
	for i := 0; i < 100; i++ {
		mgo.Dial("127.0.0.1")
	}
	log.Println(time.Now().Sub(start))

	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Panic(err)
	}
	start = time.Now()
	for i := 0; i < 100000; i++ {
		s := sess.Copy()
		s.Close()
	}
	log.Println(time.Now().Sub(start))

	ctx := context.Background()
	ctxs := context.WithValue(ctx, "session", sess)

	start = time.Now()
	for i := 0; i < 100000; i++ {
		_ = ctxs.Value("session").(*mgo.Session)
	}
	log.Println(time.Now().Sub(start))
	log.Println()
}
