package repositories

import (
	"iter"
	"slices"

	"github.com/leeliwei930/fx_playground/models"
)

type TodoRepository interface {
	GetAllTasks() []*models.Task
	GetCompletedTasks() []*models.Task
	GetUncompletedTasks() []*models.Task
}

func NewTodoRepository(tasks models.DefaultTasks) TodoRepository {
	return &todoRepository{
		tasks: tasks,
	}
}

type todoRepository struct {
	tasks []*models.Task
}

func (r *todoRepository) GetAllTasks() []*models.Task {
	return r.tasks
}

func completedTasks(tasks []*models.Task) iter.Seq[*models.Task] {
	return func(yield func(*models.Task) bool) {
		for _, task := range tasks {
			if task.Completed {
				yield(task)
				if !yield(task) {
					break
				}
			}
		}
	}
}

func uncompletedTasks(tasks []*models.Task) iter.Seq[*models.Task] {
	return func(yield func(*models.Task) bool) {
		for _, task := range tasks {
			if !task.Completed {
				yield(task)
				if !yield(task) {
					break
				}
			}
		}
	}
}
func (r *todoRepository) GetCompletedTasks() []*models.Task {
	return slices.Collect(completedTasks(r.tasks))
}

func (r *todoRepository) GetUncompletedTasks() []*models.Task {
	return slices.Collect(uncompletedTasks(r.tasks))
}
