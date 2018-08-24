package mongo

import (
	"log"
	"testing"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func TestBSONOmitempty(t *testing.T) {
	sess, err := mgo.Dial("127.0.0.1")
	c := sess.DB("lngo").C("home")
	if err != nil {
		log.Panic(err)
	}
	id := bson.NewObjectId().Hex()

	obj1 := &Home{
		ID:         id,
		Name:       "name",
		InsertTime: time.Now(),
	}
	err = c.Insert(obj1)
	if err != nil {
		log.Panic(err)
	}
	log.Println(obj1)

	obj2 := &Home{
		ID:   id,
		Name: "newName",
	}
	// c.UpdateId(id, obj)

	c.UpdateId(id, bson.M{
		"$set": obj2,
	})

	res := Home{}
	c.FindId(id).Limit(1).One(&res)
	if !TimeIsEqual(res.InsertTime, obj1.InsertTime) {
		log.Println(res)
		t.Errorf("set failed! id=%v", id)
	}
}

func TimeIsEqual(t1, t2 time.Time) bool {
	return t1.Unix() == t2.Unix()
}
