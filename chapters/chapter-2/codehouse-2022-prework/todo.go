package main

type Todo struct {
	ID      int    `json:"id"`
	Value   string `json:"value"`
	DueDate string `json:"due_date"`
}
