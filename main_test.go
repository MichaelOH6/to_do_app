package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddHandler(t *testing.T) {
	// Initialize the task list
	tasks = TaskList{
		Tasks: make([]Task, 0, 10),
	}

	// Create a POST request with a task name
	form := url.Values{}
	form.Add("task", "Test Task")
	req, err := http.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Check if the task was added
	if len(tasks.Tasks) != 1 {
		t.Errorf("handler did not add task: got %v want %v", len(tasks.Tasks), 1)
	}

	// Check the task name
	if tasks.Tasks[0].Name != "Test Task" {
		t.Errorf("handler added wrong task name: got %v want %v", tasks.Tasks[0].Name, "Test Task")
	}
}

func TestDeleteHandler(t *testing.T) {
	// Initialize the task list with some tasks
	tasks = TaskList{
		Tasks: []Task{
			{Name: "Task 1"},
			{Name: "Task 2"},
			{Name: "Task 3"},
		},
	}

	// Create a POST request to delete the task at index 1
	form := url.Values{}
	form.Add("index", "1")
	req, err := http.NewRequest(http.MethodPost, "/delete", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Check if the task was deleted
	if len(tasks.Tasks) != 2 {
		t.Errorf("handler did not delete task: got %v want %v", len(tasks.Tasks), 2)
	}

	// Check the remaining tasks to ensure the correct one was deleted
	if tasks.Tasks[0].Name != "Task 1" || tasks.Tasks[1].Name != "Task 3" {
		t.Errorf("handler deleted wrong task: remaining tasks %v", tasks.Tasks)
	}
}

func TestDeleteAllHandler(t *testing.T) {
	// Initialize the task list with some tasks
	tasks = TaskList{
		Tasks: []Task{
			{Name: "Task 1"},
			{Name: "Task 2"},
			{Name: "Task 3"},
		},
	}

	// Create a POST request to delete all the tasks
	form := url.Values{}
	form.Add("index", "1")
	req, err := http.NewRequest(http.MethodPost, "/deleteAll", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteAllHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Check if the tasks was deleted
	if len(tasks.Tasks) != 0 {
		t.Errorf("handler did not delete task: got %v want %v", len(tasks.Tasks), 0)
	}
}
