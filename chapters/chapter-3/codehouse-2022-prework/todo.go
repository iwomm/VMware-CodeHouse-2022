package main

type Todo struct {
	Id      int    `json:id`
	Value   string `json:"value"`
	DueDate string `json:"due_date"`
}
