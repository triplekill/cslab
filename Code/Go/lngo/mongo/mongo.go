package mongo

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Home struct {
	ID         string    `bson:"_id"`
	Name       string    `bson:"name,omitempty"`
	InsertTime time.Time `bson:"insert_time,omitempty"`
}

func Test1() {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Panic(err)
	}
	c := sess.DB("lngo").C("home")
	h := Home{Name: "123", InsertTime: time.Now()}
	c.Upsert(bson.M{"name": "123"}, h)
	c.Find(bson.M{"name": "123"}).One(&h)
	fmt.Println(h.InsertTime.Format("2006-01-02 15:04:05"))
	tz, _ := time.LoadLocation("America/New_York")
	fmt.Println(h.InsertTime.In(tz).Format("2006-01-02 15:04:05"))
}

func BSONOmitempty() {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Panic(err)
	}
	id := bson.NewObjectId().Hex()
	c := sess.DB("lngo").C("home")
	obj := &Home{
		ID:         id,
		Name:       "name",
		InsertTime: time.Now(),
	}
	err = c.Insert(obj)
	if err != nil {
		log.Panic(err)
	}
	log.Println(obj)

	obj = &Home{
		ID:   id,
		Name: "newName",
	}
	// c.UpdateId(id, obj)

	c.UpdateId(id, bson.M{
		"$set": obj,
	})

}
