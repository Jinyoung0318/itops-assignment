package main

import (
	"fmt"
	"itops-assignment/handler"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		//이슈 생성
		case r.URL.Path == "/issue" && r.Method == http.MethodPost:
			handler.CreateIssueHandler(w, r)

		//이슈 목록 조회
		case r.URL.Path == "/issues" && r.Method == http.MethodGet:
			handler.ListIssuesHandler(w, r)

		//이슈 상세 조회
		case strings.HasPrefix(r.URL.Path, "/issue/") && r.Method == http.MethodGet:
			handler.IssueDetailHandler(w, r)

		//이슈 수정
		case strings.HasPrefix(r.URL.Path, "/issue/") && r.Method == http.MethodPatch:
			handler.UpdateIssueHandler(w, r)

		default:
			http.NotFound(w, r)
		}
	})

	fmt.Print("서버 실행 : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
