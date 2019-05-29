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

func Create(data interface{}, bucket string) (int,error){
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout: 1 * time.Second })
	if err != nil { return 0,err }
	defer db.Close()
	var newId int
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, fail := tx.CreateBucketIfNotExists([]byte(bucket))
		if fail != nil { return fail }
		id, fail := bucket.NextSequence()
		integerId := int(id)
		newId = integerId
		entry := Entry{Data:data, Id:integerId}
		json, fail := json.Marshal(entry)
		if fail != nil { return fail }
		return bucket.Put(itob(int(id)), json)
	})

	return newId, err
}

func Read(bucket string)([]Entry, error){
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout: 1 * time.Second })
	if err != nil { return nil, err }
	defer db.Close()
	var resultList []Entry
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil { return errors.New("Bucket didn't exist! Create a new entry in that bucket")}
		err = b.ForEach(func(k, v []byte) error {
			var todo Entry
			jsonError := json.Unmarshal(v, &todo)
			if jsonError != nil {return jsonError}
			resultList = append(resultList, todo)
			return nil
		})
		return err
	})
	return resultList, err
}

func Update(entry Entry, bucket string)error{
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout: 1 * time.Second })
	if err != nil { return err }
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, fail := tx.CreateBucketIfNotExists([]byte(bucket))
		if fail != nil { return fail }
		json, fail := json.Marshal(entry)
		if fail != nil { return fail }
		return bucket.Put(itob(entry.Id), json)
	})
	return err
}

func Delete(id int, bucket string)error{
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout: 1 * time.Second })
	if err != nil { return err }
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, fail := tx.CreateBucketIfNotExists([]byte(bucket))
		if fail != nil { return fail }
		return bucket.Delete(itob(id))
	})
	return err
}