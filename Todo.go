package main

type Entry struct {
	Id   int    `json:"id"`
	Data interface{} `json:"data"`
}
