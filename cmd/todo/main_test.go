package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(t *testing.T) {
	task := "test task number 1"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)
	t.Run("AddNewTask", func(t *testing.T) {

		cmd := exec.Command(cmdPath, "-t", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-l")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := task + "\n"
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

}
