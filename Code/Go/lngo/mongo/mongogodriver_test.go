package mongo

import (
	"context"
	"log"
	"testing"

	"github.com/exfly/lngo/model"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func Test_MongoGoDriver(t *testing.T) {
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("lngo")
	coll := db.Collection("user")
	db.RunCommand(
		context.Background(),
		bson.NewDocument(bson.EC.Int32("dropDatabase", 1)),
	)
	coll.InsertOne(
		nil,
		model.User{
			ID:   "asdfasdf",
			Name: "newname",
			Pet:  []string{"p2", "p3"},
			Toy:  []model.Toy{model.Toy{ID: "toyid"}},
		},
	)
	cur := coll.FindOne(context.Background(), map[string]interface{}{"_id": "asdfasdf"})
	if err != nil {
		t.Error(err)
	}
	var user model.User
	cur.Decode(&user)
	t.Error(user)
	coll.DeleteOne(context.Background(), nil)
}

func Test_MongoGoDriver_plus(t *testing.T) {
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("lngo")
	coll := db.Collection("user")
	db.RunCommand(
		context.Background(),
		bson.NewDocument(bson.EC.Int32("dropDatabase", 1)),
	)
	coll.InsertOne(
		nil,
		model.User{
			ID:   "asdfasdf",
			Name: "newname",
			Pet:  []string{"p2", "p3"},
			Toy:  []model.Toy{model.Toy{ID: "toyid"}},
		},
	)
	cur := coll.FindOne(context.Background(), M{"_id": "asdfasdf"})
	if err != nil {
		t.Error(err)
	}
	var user model.User
	cur.Decode(&user)
	t.Error(user)
	coll.DeleteOne(context.Background(), nil)
}
