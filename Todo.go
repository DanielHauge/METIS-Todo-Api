package main

type Todo struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
	Date string `json:"date"`
	Time string `json:"time"`
}
