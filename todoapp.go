package todoapp

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Handler() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root@/todoapp?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if r.Method == "POST" {
		result, err := db.Exec(`INSERT INTO tasks(content) VALUES(?)`, "Bomb")
		if err != nil {
			log.Fatal(err)
		}

		affected, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(affected)

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(lastInsertID)
	}

	rows, err := db.Query(`SELECT id, content, created_at, updated_at FROM tasks`)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var content string
		var createdAt *time.Time
		var updatedAt *time.Time
		err = rows.Scan(&id, &content, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, content, createdAt, updatedAt)
	}
	item := struct {
		data string
	}{
		"test",
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, item)
	if err != nil {
		panic(err)
	}
}
