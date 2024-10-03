package tasktracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type todo struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt time.Time
}

type List []todo

func (l *List) Add(task string) {
	ls := *l
	count := len(ls) + 1
	status := "todo"
	t := todo{
		Id:          count,
		Description: task,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) InProgress(id int) error {
	ls := *l
	if id <= 0 || id > len(ls) {
		return fmt.Errorf("the error is ID: %d, doesn't exist", id)
	}

	status := "inprogress"

	ls[id-1].Status = status
	ls[id-1].UpdatedAt = time.Now()
	return nil
}

func (l *List) Complete(id int) error {

	ls := *l
	if id <= 0 || id > len(ls) {
		return fmt.Errorf("the error is ID: %d, doesn't exist", id)
	}
	ls[id-1].Status = "done"
	ls[id-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(id int) error {
	ls := *l
	if id <= 0 || id > len(ls) {
		return fmt.Errorf("the error is ID: %d, doesn't exist", id)
	}

	*l = append(ls[:id-1], ls[id:]...)
	return nil
}

func (l *List) Save(filepath string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formated := ""
	for k, v := range *l {
		prefix := " "
		if v.Status == "done" {
			prefix = "X"
		} else if v.Status == "inprogress" {
			prefix = "%"
		}
		formated += fmt.Sprintf("%s%d: %s\n", prefix, k+1, v.Description)
	}

	return formated
}

func (l *List) DoneOnly() string {
	formatted := ""
	for k, v := range *l {
		prefix := "X"
		if v.Status != "done" {
			continue
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, v.Description)
	}
	return formatted
}

func (l *List) NotDone() string {
	formatted := ""
	for k, v := range *l {
		if v.Status != "todo" {
			continue
		}
		formatted = fmt.Sprintf("%d: %s\n", k+1, v.Description)
	}

	return formatted
}

func (l *List) IProgress() string {
	formatted := ""
	for k, v := range *l {
		if v.Status != "inprogress" {
			continue
		}
		formatted = fmt.Sprintf("%d: %s\n", k+1, v.Description)
	}

	return formatted
}
