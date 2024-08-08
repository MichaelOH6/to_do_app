package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Task struct {
	Name string `json:"Name"`
}

type TaskList struct {
	sync.Mutex
	Tasks []Task
}

var (
	tasks = TaskList{
		Tasks: make([]Task, 0, 10),
	}
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/update", updateHandler) // Add this line
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/load", readHandler)
	http.HandleFunc("/deleteAll", deleteAllHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
	fmt.Println("server started on port http://localhost:8080/")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tasks.Lock()
	defer tasks.Unlock()
	tmpl.Execute(w, tasks.Tasks)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		taskName := r.FormValue("task")
		if taskName != "" {
			tasks.Lock()
			if len(tasks.Tasks) < 10 {
				tasks.Tasks = append(tasks.Tasks, Task{Name: taskName})
			}
			tasks.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		indexStr := r.FormValue("index")
		index, err := strconv.Atoi(indexStr)
		if err == nil {
			tasks.Lock()
			if index >= 0 && index < len(tasks.Tasks) {
				tasks.Tasks = append(tasks.Tasks[:index], tasks.Tasks[index+1:]...)
			}
			tasks.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		indexStr := r.FormValue("index")
		newName := r.FormValue("newName")
		index, err := strconv.Atoi(indexStr)
		if err == nil && newName != "" {
			tasks.Lock()
			if index >= 0 && index < len(tasks.Tasks) {
				tasks.Tasks[index].Name = newName
			}
			tasks.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tasks.Lock()
		tasks.Tasks = nil
		tasks.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tasks.Lock()
		item, _ := json.MarshalIndent(tasks, "", "  ")
		tasks.Unlock()

		file, err2 := os.Create("savedTasks.json")
		check(err2)
		file.Close()

		err := os.WriteFile("savedTasks.json", item, 0644)
		check(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tasks.Lock()
		defer tasks.Unlock()

		if len(tasks.Tasks) == 0 {
			dat, err := os.ReadFile("savedTasks.json")
			check(err)

			var taskList struct {
				Tasks []Task `json:"Tasks"`
			}

			err = json.Unmarshal(dat, &taskList)
			check(err)

			for _, task := range taskList.Tasks {
				if len(tasks.Tasks) < 10 {
					tasks.Tasks = append(tasks.Tasks, task)
				} else {
					break
				}
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
