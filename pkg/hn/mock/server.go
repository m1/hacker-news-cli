package mock

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)

func NewTestServer() *httptest.Server {
	os.Setenv("HN_MAX_POSTS", strconv.Itoa(100))

	testServer := http.NewServeMux()
	testServer.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		p := r.URL.Query().Get("p")
		if p != "" && p == "2" {
			fmt.Fprint(w, string(HtmlPage2))
			return
		}

		fmt.Fprint(w, string(HtmlPage1))
	})

	testServer.HandleFunc("/news?p=2", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, string(HtmlPage2))
	})

	return httptest.NewServer(testServer)
}
