package todoapp_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/obrkn/todoapp/pkg"
)

func TestIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	got := httptest.NewRecorder()
	pkg.Index(got, req)

	if got.Code != http.StatusOK {
		t.Errorf("not status OK")
	}
	fmt.Println(got.Code)
}

// func TestPost(t *testing.T) {
// 	req :=
// }
