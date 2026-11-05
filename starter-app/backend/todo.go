package main

import "sync"

// Todo is the shape of one item. The `json:"..."` tags control the
// field names in the JSON that the frontend receives.
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Store keeps todos in memory. Data is lost when the server restarts;
// that's fine for now (Step 6 moves it into PostgreSQL).
type Store struct {
	mu     sync.Mutex // protects the fields below: many requests run at the same time
	todos  []Todo
	nextID int
}

func NewStore() *Store {
	return &Store{
		todos: []Todo{
			{ID: 1, Title: "Read docs/GUIDE.md", Done: true},
			{ID: 2, Title: "Implement POST /api/todos", Done: false},
		},
		nextID: 3,
	}
}
