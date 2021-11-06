package todoapp_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/obrkn/todoapp/pkg"
)

func TestGetIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	got := httptest.NewRecorder()
	pkg.Index(got, req)

	if got.Code != http.StatusOK {
		t.Errorf("not status OK")
	}
	fmt.Println(got.Code)
}

func TestPostIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)

	rand.Seed(time.Now().UnixNano())
	testNum := fmt.Sprint("テスト番号: ", rand.Intn(100000))
	q := req.URL.Query()
	q.Add("content", testNum)
	req.URL.RawQuery = q.Encode()

	got := httptest.NewRecorder()
	pkg.Index(got, req)

	if got.Code != http.StatusOK {
		t.Errorf("not status OK")
	}
	fmt.Println(got.Code)

	if got := got.Body.String(); !strings.Contains(got, testNum) {
		t.Errorf("not match")
	}
	fmt.Println(testNum)
}
