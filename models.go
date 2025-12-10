package main

import "time"

// Task represents a simple to-do item.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// In-memory store (simple, not concurrent-safe for heavy load)
var store = NewMemoryStore()

// MemoryStore holds tasks in memory.
type MemoryStore struct {
	tasks []*Task
	next  int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{tasks: make([]*Task, 0), next: 1}
}

func (m *MemoryStore) All() []*Task {
	return m.tasks
}

func (m *MemoryStore) Create(title string) *Task {
	t := &Task{ID: m.next, Title: title, Done: false, CreatedAt: time.Now()}
	m.next++
	m.tasks = append(m.tasks, t)
	return t
}

func (m *MemoryStore) Toggle(id int) *Task {
	for _, t := range m.tasks {
		if t.ID == id {
			t.Done = !t.Done
			return t
		}
	}
	return nil
}

func (m *MemoryStore) Reset() {
	m.tasks = make([]*Task, 0)
	m.next = 1
}
