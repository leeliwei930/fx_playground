package services

import (
	"fmt"

	"github.com/leeliwei930/fx_playground/repositories"
)

type TaskReportService struct {
	repo repositories.TodoRepository
}

func NewTaskReportService(repo repositories.TodoRepository) *TaskReportService {
	return &TaskReportService{
		repo: repo,
	}
}

func (s *TaskReportService) GetTaskReport() {
	fmt.Printf("Completed Tasks %d, Uncompleted Tasks %d\n", len(s.repo.GetCompletedTasks()), len(s.repo.GetUncompletedTasks()))
}
