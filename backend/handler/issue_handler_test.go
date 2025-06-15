package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateIssue_WithValidUser(t *testing.T) {
	t.Run("정상 생성 - userId가 존재할 경우", func(t *testing.T) {
		payload := map[string]interface{}{
			"title":       "테스트 이슈",
			"description": "설명입니다.",
			"userId":      1,
		}
		jsonData, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/issue", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		CreateIssueHandler(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status 201, got %d", resp.StatusCode)
		}
	})
}

func TestCreateIssue_WithInvalidUser(t *testing.T) {
	t.Run("에러 발생 - 존재하지 않는 userId", func(t *testing.T) {
		payload := map[string]interface{}{
			"title":       "에러 테스트 이슈",
			"description": "설명입니다.",
			"userId":      999, // 존재하지 않는 사용자
		}
		jsonData, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/issue", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		CreateIssueHandler(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", resp.StatusCode)
		}
	})
}

func TestGetIssueList_All(t *testing.T) {
	t.Run("이슈 전체 목록 조회", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/issues", nil)
		w := httptest.NewRecorder()

		ListIssuesHandler(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetIssueDetail_NotFound(t *testing.T) {
	t.Run("존재하지 않는 ID 조회 시 404 반환", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/issue/9999", nil)
		w := httptest.NewRecorder()

		IssueDetailHandler(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", resp.StatusCode)
		}
	})
}
