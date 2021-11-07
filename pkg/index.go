package pkg

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root@tcp(db)/todoapp?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if r.Method == "POST" {
		content := r.FormValue("content")
		_, err := db.Exec(`INSERT INTO tasks(content) VALUES(?)`, content)
		if err != nil {
			log.Fatal(err)
		}
	}

	rows, err := db.Query(`SELECT id, content, created_at, updated_at FROM tasks ORDER BY id DESC`)
	if err != nil {
		log.Fatal(err)
	}
	type Task struct {
		Id        int
		Content   string
		CreatedAt *time.Time
		UpdatedAt *time.Time
	}
	task := Task{}
	var tasks []Task
	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Content, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, tasks)
	if err != nil {
		panic(err)
	}
}