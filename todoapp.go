package todoapp

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Handler() {
	db, err := sql.Open("mysql", "root@/todoapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec(`INSERT INTO tasks(content) VALUES('oppai')`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// 	http.HandleFunc("/", index)
	// 	http.ListenAndServe(":8080", nil)
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	item := struct {
// 		data string
// 	}{
// 		"test",
// 	}
// 	tmpl, err := template.ParseFiles("templates/index.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = tmpl.Execute(w, item)
// 	if err != nil {
// 		panic(err)
// 	}
// }
