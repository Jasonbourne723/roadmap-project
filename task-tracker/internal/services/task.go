package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"roadmap/task-tracker/internal/models"
	"time"
)

var TaskService = new(taskService)

type taskService struct {
}

func (t *taskService) Add(description string) error {
	if len(description) == 0 {
		return errors.New("description is null")
	}

	list := read()

	var id int
	if len(list) == 0 {
		id = 1
	} else {
		id = list[len(list)-1].Id + 1
	}

	list = append(list, models.Task{
		Id:          id,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      "todo",
	})
	return write(list)
}

func (t *taskService) List(status string) []models.Task {
	list := read()
	if len(status) == 0 {
		return list
	}
	return Filter(list, func(t models.Task) bool {
		return t.Status == status
	})
}

func (t *taskService) Delete(id int) {
	list := read()
	deleted := false
	for i, v := range list {
		if id == v.Id {
			if i == len(list)-1 {
				list = list[:len(list)-1]
			} else {
				list = append(list[:i], list[i+1:]...)
			}
			deleted = true
		}
	}
	if deleted {
		write(list)
	}
}

func (t *taskService) Update(id int, desc string) {
	list := read()
	for i, v := range list {
		if v.Id == id {
			list[i].Description = desc
		}
	}
	write(list)
}

func (t *taskService) Mark(id int, status string) {
	list := read()
	for i, v := range list {
		if v.Id == id {
			list[i].Status = status
			list[i].UpdatedAt = time.Now()
		}
	}
	write(list)
}

func read() []models.Task {
	f, err := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer f.Close()
	var list []models.Task
	json.NewDecoder(f).Decode(&list)
	return list
}

func write(list []models.Task) error {
	f, err := os.OpenFile("task.json", os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(list)
}

func Filter(list []models.Task, f func(models.Task) bool) []models.Task {
	var newList []models.Task
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
