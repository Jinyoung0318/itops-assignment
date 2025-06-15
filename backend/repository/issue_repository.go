package repository

import (
	_ "errors"
	"itops-assignment/model"
	"sync"
	"time"
)

var (
	issues      = make(map[uint]*model.Issue)
	issuesMutex sync.Mutex
	idCounter   uint = 1
)

var predefinedUsers = []*model.User{
	{ID: 1, Name: "김개발"},
	{ID: 2, Name: "이디자인"},
	{ID: 3, Name: "박기획"},
}

func GetUserByID(id uint) *model.User {
	for _, u := range predefinedUsers {
		if u.ID == id {
			return u
		}
	}
	return nil
}

func CreateIssue(issue *model.Issue) *model.Issue {
	issuesMutex.Lock()
	defer issuesMutex.Unlock()

	issue.ID = idCounter
	idCounter++
	issue.CreatedAt = time.Now()
	issue.UpdatedAt = issue.CreatedAt
	issues[issue.ID] = issue
	return issue
}

func GetIssueByID(id int) *model.Issue {
	issuesMutex.Lock()
	defer issuesMutex.Unlock()

	issue, exists := issues[uint(id)]
	if !exists {
		return nil
	}
	return issue
}

func ListIssues(status string) []*model.Issue {
	issuesMutex.Lock()
	defer issuesMutex.Unlock()

	var list []*model.Issue
	for _, issue := range issues {
		if status == "" || issue.Status == status {
			list = append(list, issue)
		}
	}
	return list
}
