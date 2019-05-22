package main

import (
	"encoding/binary"
	"errors"
	"github.com/boltdb/bolt"
	"time"
)

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func Create(todo Todo)error{
	db, err := bolt.Open("todos.db", 0600, &bolt.Options{Timeout: 1 * time.Second })
	if err != nil { return err }
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, fail := tx.CreateBucketIfNotExists([]byte("TodoBucket"))
		if fail != nil { return fail }
		id, fail := bucket.NextSequence()
		todo.Id = int(id)
		json, fail := json.Marshal(todo)
		if fail != nil { return fail }
		return bucket.Put(itob(int(id)), json)
	})

	return err
}

func Read()([]Todo, error){
	db, err := bolt.Open("todos.db", 0600, &bolt.Options{Timeout: 1 * time.Second })
	if err != nil { return nil, err }
	defer db.Close()
	var resultList []Todo
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TodoBucket"))
		if b == nil { return errors.New("Bucket didn't exist! Create a new todo")}
		err = b.ForEach(func(k, v []byte) error {
			var todo Todo
			jsonError := json.Unmarshal(v, &todo)
			if jsonError != nil {return jsonError}
			resultList = append(resultList, todo)
			return nil
		})
		return err
	})
	return resultList, err
}

func Update(todo Todo)error{
	db, err := bolt.Open("todos.db", 0600, &bolt.Options{Timeout: 1 * time.Second })
	if err != nil { return err }
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, fail := tx.CreateBucketIfNotExists([]byte("TodoBucket"))
		if fail != nil { return fail }
		json, fail := json.Marshal(todo)
		if fail != nil { return fail }
		return bucket.Put(itob(todo.Id), json)
	})
	return err
}

func Delete(id int)error{
	db, err := bolt.Open("todos.db", 0600, &bolt.Options{Timeout: 1 * time.Second })
	if err != nil { return err }
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, fail := tx.CreateBucketIfNotExists([]byte("TodoBucket"))
		if fail != nil { return fail }
		return bucket.Delete(itob(id))
	})
	return err
}