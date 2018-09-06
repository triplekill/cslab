package jobq

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type JobQueue struct {
	c        *mongo.Collection
	hanlders map[string]Handler
	wg       sync.WaitGroup
	stop     chan int
}

type JobDesc struct {
	ID        objectid.ObjectID `bson:"_id,omitempty"`
	StartTime time.Time         `bson:"startTime,omitempty"`
	EndTime   time.Time         `bson:"endTime,omitempty"`
	CreatedAt time.Time         `bson:"createdAt"`
	Priority  int               `bson:"priority"`
	Operation string            `bson:"operation"`
	Payload   []interface{}     `bson:"payload,omitempty"`
}

type Handler func(interface{}) error

func NewJobQueue(collection *mongo.Collection) *JobQueue {
	jobq := JobQueue{
		c:        collection,
		hanlders: map[string]Handler{},
	}
	jobq.ensureIndex()
	return &jobq
}

func (j *JobQueue) ensureIndex() {
	iview := j.c.Indexes()
	iview.CreateMany(context.Background(), []mongo.IndexModel{
		mongo.IndexModel{
			Keys: bson.NewDocument(
				bson.EC.Int32("startTime", 1),
				bson.EC.Int32("priority", -1),
				bson.EC.Int32("createdAt", 1),
			),
		},
		mongo.IndexModel{
			Keys: bson.NewDocument(
				bson.EC.Int32("endTime", 1),
			),
			Options: bson.NewDocument(
				bson.EC.Int32("expireAfterSeconds", 3600*24),
			),
		},
	})
}

func (j *JobQueue) CreateWorkerAndRun(num int) {
	j.stop = make(chan int, num)
	for i := 0; i < num; i++ {
		q := JobQueueWorker{j, time.Now(), false}
		j.wg.Add(1)
		go q.Work()
	}
}

func (j *JobQueue) StopAllWorker() {
	for i := 0; i < cap(j.stop); i++ {
		j.stop <- 1
	}
	close(j.stop)
}

func (j *JobQueue) WaitAll() {
	j.wg.Wait()
}

func (j *JobQueue) EnqueueJob(name string, priority int, payload ...interface{}) error {
	job := JobDesc{
		ID:        objectid.New(),
		CreatedAt: time.Now(),
		Priority:  priority,
		Payload:   payload,
		Operation: name,
	}
	_, err := j.c.InsertOne(context.Background(), job)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (j *JobQueue) Register(name string, handler Handler) error {
	if _, ok := j.hanlders[name]; ok {
		log.Println("[WARN] duplicate handler found!")
	}
	j.hanlders[name] = handler
	return nil
}
