package tasktracker_test

import (
	"github.com/Mensurui/taskTracker.git"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := tasktracker.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Description != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Description)
	}
	if l[0].Status != "todo" {
		t.Errorf("New task should not be completed.")
	}
	l.InProgress(1)
	if l[0].Status != "inprogress" {
		t.Errorf("New task should be in progress")
	}
	l.Complete(1)
	if l[0].Status != "done" {
		t.Errorf("New task should be completed.")
	}

}

func TestDelete(t *testing.T) {
	l := tasktracker.List{}
	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}
	for _, v := range tasks {
		l.Add(v)
	}
	if l[0].Description != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Description)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}
	if l[1].Description != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Description)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := tasktracker.List{}
	l2 := tasktracker.List{}
	taskName := "New Task"
	l1.Add(taskName)
	if l1[0].Description != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Description)
	}
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
	if l1[0].Description != l2[0].Description {
		t.Errorf("Task %q should match %q task.", l1[0].Description, l2[0].Description)
	}
}
