package mongo

import (
	"fmt"
	"log"
	"time"

	"github.com/exfly/cslab/Code/Go/lngo/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
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

// TestTxn tansaction implement
// core code at
// https://blog.labix.org/2012/08/22/multi-doc-transactions-for-mongodb
// 经典两阶段提交 https://acupple.github.io/2016/08/09/MongoDB%E4%B8%A4%E9%98%B6%E6%AE%B5%E6%8F%90%E4%BA%A4%E5%AE%9E%E7%8E%B0%E4%BA%8B%E5%8A%A1/
// 分布式事务同时使用这种方式就可以
func TestTxn() {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Println(err)
	}
	colle := sess.DB("").C("account")

	p1 := model.People{ID: "A", Account: 100}
	p2 := model.People{ID: "B", Account: 100}
	colle.Insert(p1)
	colle.Insert(p2)
	runner := txn.NewRunner(sess.DB("").C("account_txn"))
	ops := []txn.Op{{
		C:      "account",
		Id:     p1.ID,
		Assert: bson.M{"account": bson.M{"$gte": 10}},
		Update: bson.M{"$inc": bson.M{"account": -10}},
	}, {
		C:      "account",
		Id:     p2.ID,
		Update: bson.M{"$inc": bson.M{"account": 10}},
	}}
	id := bson.NewObjectId() // Optional
	err = runner.Run(ops, id, nil)
	if err != nil {
		log.Println(err)
	}
}
