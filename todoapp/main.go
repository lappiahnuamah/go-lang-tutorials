package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

var (
	tasks []Task
	mu    sync.Mutex
	id    int
)

type Task struct {
	ID   int
	Name string
	Done bool
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addTaskHandler)
	http.HandleFunc("/done", markDoneHandler)

	fmt.Println("Server is running on port 8082")
	http.ListenAndServe(":8082", nil)
}

var tmpl = template.Must(template.ParseFiles("templates/add.html"))

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		mu.Lock()
		id++
		task := Task{ID: id, Name: r.FormValue("task"), Done: false}
		tasks = append(tasks, task)
		mu.Unlock()
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		tmpl.Execute(w, nil)
	}
}

var listTmpl = template.Must(template.New("list").Parse(`
<h1>Task List</h1>
<ul>
{{range .}}
	<li>{{.Name}} - <a href="/done?id={{.ID}}">Done</a></li>
{{end}}
</ul>
<a href="/add">Add Task</a>
`))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go To-DO App!")
}

func markDoneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	mu.Lock()
	for i := range tasks {
		if fmt.Sprintf("%d", tasks[i].ID) == id {
			tasks[i].Done = true
		}
	}
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
