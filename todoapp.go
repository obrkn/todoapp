package todoapp

import (
	"net/http"

	"github.com/obrkn/todoapp/pkg"

	_ "github.com/go-sql-driver/mysql"
)

func Handler() {
	http.HandleFunc("/", pkg.Index)
	http.ListenAndServe(":8080", nil)
}
