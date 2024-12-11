package main

import (
	"github.com/leeliwei930/fx_playground/models"
	"github.com/leeliwei930/fx_playground/repositories"
	"github.com/leeliwei930/fx_playground/services"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(func() models.DefaultTasks {
			return models.DefaultTasks{
				{ID: 1, Title: "Task 1", Description: "Description 1", Completed: true},
				{ID: 2, Title: "Task 2", Description: "Description 2", Completed: false},
				{ID: 3, Title: "Task 3", Description: "Description 3", Completed: true},
			}
		}),
		fx.Provide(services.NewTaskReportService),
		fx.Provide(repositories.NewTodoRepository),
		fx.Invoke(func(s *services.TaskReportService, shutdowner fx.Shutdowner) {
			s.GetTaskReport()
			shutdowner.Shutdown()
		},
		),
	).Run()
}
