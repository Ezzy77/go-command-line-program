package todo_test

import (
	"os"
	"testing"

	todo "github.com/ezzy77/command-go"
)

func TestAdd(t *testing.T) {

	// arrange
	l := todo.List{}

	taskName := "New Task"

	//act

	l.Add(taskName)

	//assert
	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	// arrange
	l := todo.List{}

	taskName := "New Task"

	//act
	l.Add(taskName)
	//l.Complete(1)

	//assert
	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("new task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("new task should be completed.")
	}

}

func TestDelete(t *testing.T) {
	//arrange
	l := todo.List{}

	tasks := []string{
		"New Task1",
		"New Task2",
		"New Task3",
	}

	//act

	for _, v := range tasks {
		l.Add(v)

	}

	//assert
	if l[0].Task != tasks[0] {
		t.Errorf("expected %q, got %q instead", tasks[0], l[0].Task)

	}
	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("expected %q, got %q instead.", tasks[2], len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("expected %q, got %q instead.", tasks[2], l[1].Task)
	}

}

func TestSaveGet(t *testing.T) {

	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf(" error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
