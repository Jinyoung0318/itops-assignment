package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"itops-assignment/model"
	"itops-assignment/repository"
)

func errorResponse(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": message,
		"code":  code,
	})
}

func CreateIssueHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		UserID      *uint  `json:"userId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Title == "" || req.Description == "" {
		errorResponse(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	var user *model.User
	status := "PENDING"
	if req.UserID != nil {
		u := repository.GetUserByID(*req.UserID)
		if u == nil {
			errorResponse(w, "Invalid userId", http.StatusBadRequest)
			return
		}
		user = u
		status = "IN_PROGRESS"
	}

	now := time.Now()
	issue := &model.Issue{
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
		User:        user,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	repository.CreateIssue(issue)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(issue)
}

func ListIssuesHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status != "" && !model.ValidStatuses[status] {
		errorResponse(w, "Invalid status", http.StatusBadRequest)
		return
	}
	issues := repository.ListIssues(status)
	json.NewEncoder(w).Encode(map[string]interface{}{"issues": issues})
}

func IssueDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/issue/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}
	issue := repository.GetIssueByID(id)
	if issue == nil {
		errorResponse(w, "Issue not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(issue)
}

func UpdateIssueHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/issue/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}
	existing := repository.GetIssueByID(id)
	if existing == nil {
		errorResponse(w, "Issue not found", http.StatusNotFound)
		return
	}
	if existing.Status == "COMPLETED" || existing.Status == "CANCELLED" {
		errorResponse(w, "Cannot update completed or cancelled issue", http.StatusBadRequest)
		return
	}

	var req struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
		UserID      *uint   `json:"userId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Title != nil {
		existing.Title = *req.Title
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}

	if req.UserID != nil {
		if *req.UserID == 0 {
			existing.User = nil
			existing.Status = "PENDING"
		} else {
			user := repository.GetUserByID(*req.UserID)
			if user == nil {
				errorResponse(w, "Invalid userId", http.StatusBadRequest)
				return
			}
			existing.User = user
			if existing.Status == "PENDING" && (req.Status == nil || *req.Status == "") {
				existing.Status = "IN_PROGRESS"
			}
		}
	}

	if req.Status != nil {
		if !model.ValidStatuses[*req.Status] {
			errorResponse(w, "Invalid status", http.StatusBadRequest)
			return
		}
		if existing.User == nil && (*req.Status == "IN_PROGRESS" || *req.Status == "COMPLETED") {
			errorResponse(w, "Cannot set status without assignee", http.StatusBadRequest)
			return
		}
		existing.Status = *req.Status
	}

	existing.UpdatedAt = time.Now()
	json.NewEncoder(w).Encode(existing)
}
