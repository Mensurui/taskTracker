package main

import (
	"flag"
	"fmt"
	"github.com/Mensurui/taskTracker.git"
	"os"
)

var todoFileName = ".todo.json"

func main() {

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	list := flag.Bool("l", false, "Use to view the list of tasks you have.")
	listDone := flag.Bool("ld", false, "Use to view list of tasks you have done only.")
	listNotDone := flag.Bool("ln", false, "Use to view list of tasks you didn't do only.")
	listInProgress := flag.Bool("lp", false, "Use to view list of tasks you have in progress.")
	task := flag.String("t", "", "Use to add to the list of tasks.")
	del := flag.Int("d", 0, "Use to delete a task from the tasks list.")
	ip := flag.Int("ip", 0, "Use to update the task to a progressing task.")
	done := flag.Int("dn", 0, "Use to update the task to completion.")

	flag.Parse()

	l := &tasktracker.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Println(l)
	case *listDone:
		fmt.Println(l.DoneOnly())
	case *listNotDone:
		fmt.Println(l.NotDone())
	case *listInProgress:
		fmt.Println(l.IProgress())
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		l.Delete(*del)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *ip > 0:
		l.InProgress(*ip)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *done > 0:
		l.Complete(*done)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)

	}
}
