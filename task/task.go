package task

import (
	"gopkg.in/mgo.v2/bson"
)

// Task It defines structure for task
type Task struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Todo string        `bson:"todo"`
}
