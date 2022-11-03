package task_test

import (
	"fmt"
	"goExploration/sqlite/repo/ram"
	"goExploration/sqlite/task"
	"testing"
)

// Fake repo of 6 tasks : 3 opened and 3 closed
func initFakeRepo() ram.List {
	taskRepo := ram.List{}
	task.Init(&taskRepo)
	for i := 0; i < 3; i++ {
		task.Create(fmt.Sprintf("write a book %v", i))
	}
	for i := 3; i < 6; i++ {
		task.Create(fmt.Sprintf("finish a book %v", i))
		task.Close(i)
	}
	return taskRepo
}

func TestSave(t *testing.T) {
	taskRepo := initFakeRepo()
	id := 2
	item, err := taskRepo.GetByID(id)
	if err != nil {
		t.Fatalf("could not get task id %v: %v", id, err)
	}
	want := "write a book 2"
	got := item.Description
	if got != want {
		t.Fatalf("expected %v got %v", want, got)
	}
}

func TestGetAll(t *testing.T) {
	taskRepo := initFakeRepo()
	allItems := taskRepo.GetAll()
	size := len(allItems)
	if size != 6 {
		t.Fatalf("expected size 6 got %v", size)
	}
	for i := 0; i < 3; i++ {
		want := fmt.Sprintf("write a book %v", i)
		got := allItems[i].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i, want, got)
		}
	}
	for i := 3; i < 6; i++ {
		want := fmt.Sprintf("finish a book %v", i)
		got := allItems[i].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i, want, got)
		}
	}
}

func TestGetOpened(t *testing.T) {
	taskRepo := initFakeRepo()
	openedItems := taskRepo.GetByState(task.Opened)
	size := len(openedItems)
	if size != 3 {
		t.Fatalf("expected size 3 got %v", size)
	}
	for i := 0; i < 3; i++ {
		want := fmt.Sprintf("write a book %v", i)
		got := openedItems[i].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i, want, got)
		}
	}
}

func TestGetClosed(t *testing.T) {
	taskRepo := initFakeRepo()
	closedItems := taskRepo.GetByState(task.Closed)
	size := len(closedItems)
	if size != 3 {
		t.Fatalf("expected size 3 got %v", size)
	}
	for i := 0; i < 3; i++ {
		want := fmt.Sprintf("finish a book %v", i+3)
		got := closedItems[i].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i+3, want, got)
		}
	}
}

func TestCreateTask(t *testing.T) {
	repo := ram.List{}
	task.Init(&repo)
	want := "Only test the parts of the application that you want to work"
	id := task.Create(want)
	got, err := repo.GetByID(id)
	if err != nil {
		t.Fatalf("expected item id %d exists got %s", id, err)
	}
	if got.Description != want {
		t.Fatalf("expected %s got %s", want, got)
	}
}

func TestCloseTask(t *testing.T) {
	repo := ram.List{}
	task.Init(&repo)
	empty := len(repo.GetAll())
	if 0 != empty {
		t.Fatalf("expected empty repo but got %d", empty)
	}
	want := "The only way to get more done is to have less to do"
	id := task.Create(want)
	closed := len(repo.GetByState(task.Closed))
	if 0 != closed {
		t.Fatalf("expected no closed item in repo got %d", closed)
	}
	task.Close(id)
	closed = len(repo.GetByState(task.Closed))
	if 1 != closed {
		t.Fatalf("expected 1 closed item in repo got %d", closed)
	}
	got := repo.GetByState(task.Closed)[0].Description
	if got != want {
		t.Fatalf("expected closed item to be '%s' got '%s'", want, got)
	}
}
