package jobq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/mongodb/mongo-go-driver/mongo/mongoopt"
)

type JobQueueWorker struct {
	*JobQueue
	lastFetch  time.Time
	shouldQuit bool
}

const MIN_WAIT = 10
const FACTOR float64 = 1.75

func minTime(a, b time.Duration) time.Duration {
	if a < b {
		return a
	} else {
		return b
	}
}

func step(input time.Duration) time.Duration {
	if input < 1*time.Second {
		return time.Duration(1.4 * float64(input))
	}
	if input < 5*time.Second && time.Duration(1.7*float64(input)) < 5*time.Second {
		return time.Duration(1.7 * float64(input))
	}
	return 5 * time.Second
}

func (j *JobQueueWorker) backoff() {
	diff := step(time.Now().Sub(j.lastFetch))
	diff = minTime(time.Duration(FACTOR*float64(diff)), 5*time.Second)
	fmt.Println(diff)
	time.Sleep(diff)
}

func (j *JobQueueWorker) getJob() *JobDesc {
	for {
		select {
		case <-j.stop:
			j.wg.Done()
			j.shouldQuit = true
			return nil
		}
		cur := j.c.FindOneAndUpdate(context.Background(),
			bson.NewDocument(bson.EC.Null("startTime")),
			bson.NewDocument(bson.EC.SubDocumentFromElements("$set", bson.EC.Time("startTime", time.Now()))),
			findopt.ReturnDocument(mongoopt.After),
			findopt.Sort(map[string]int{"createdAt": 1, "priority": -1}),
		)
		var ret JobDesc
		err := cur.Decode(&ret)
		if err == mongo.ErrNoDocuments {
			j.backoff()
			continue
		}
		if err != nil {
			log.Println(err)
			continue
		}
		j.lastFetch = time.Now()
		return &ret
	}
}

func (j *JobQueueWorker) finishJob(job *JobDesc) {
	j.c.UpdateOne(context.Background(),
		bson.NewDocument(bson.EC.ObjectID("_id", job.ID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.Time("endTime", time.Now()),
			),
		),
	)
}

func (j *JobQueueWorker) Work() {
	for {
		log.Printf("started worker")
		job := j.getJob()
		if job != nil {
			handler, ok := j.hanlders[job.Operation]
			if ok {
				log.Println("found handler")
				err := handler(job.Payload)
				if err != nil {
					log.Printf("[ERROR] worker %s %s", job.Operation, job.Payload)
				}
			} else {
				log.Printf("can not found handler for %s", job.Operation)
			}
		} else {
			return
		}
	}
}
