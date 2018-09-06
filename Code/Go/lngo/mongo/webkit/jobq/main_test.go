package jobq_test

import (
	"context"
	"os"
	"testing"
	"time"

	jobq "github.com/exfly/lngo/mongo/webkit/jobq"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client

func setupDatabase() {
	client, _ = mongo.NewClient("mongodb://localhost:27017/test")
	client.Connect(context.Background())
	client.Database("test").Collection("job").Drop(nil)
}

func TestMain(m *testing.M) {
	setupDatabase()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestSetupWorker(t *testing.T) {
	jq := jobq.NewJobQueue(client.Database("test").Collection("job"))
	jq.Register("test", func(payload interface{}) error {
		t.Log(payload)
		return nil
	})
	jq.CreateWorkerAndRun(2)
	t.Log("created")
	jq.EnqueueJob("test", 10, "test print")
	time.Sleep(1 * time.Second)
	jq.StopAllWorker()
	jq.WaitAll()
}
