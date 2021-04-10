package db

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1* time.Second})
	if err != nil {
		fmt.Println("error here", err)
		return err
	}

	fn := func(tx *bolt.Tx) error {
		_, err:= tx.CreateBucketIfNotExists(taskBucket)

		return err
	}

	return db.Update(fn)
}


func CreateTask(task string) (int, error)  {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err!= nil {
		return -1, err
	}
	return id, nil
}
func AllTasks() ([]Task, error)  {
	var tasks []Task
	err:= db.View(func(tx *bolt.Tx) error {
		b :=tx.Bucket(taskBucket)
		c :=b.Cursor()
		for k, v :=c.First(); k != nil; k, v=c.Next() {
			tasks = append(tasks, Task{
				Key: btoi(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTasks(key int)  error {
	 err := db.Update(func(tx *bolt.Tx) error {
		 b := tx.Bucket(taskBucket)
		 return b.Delete(itob(key))
	 })
	 if err != nil {
	 	return err
	 }
	 return nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}