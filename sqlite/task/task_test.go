package task_test

import (
	"fmt"
	"goExploration/sqlite/task"
	"testing"
)

// Fake repo of 6 tasks : 3 opened and 3 closed
func initFakeRepo() list {
	taskRepo := list{}
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
	// Create a repo for saving tasks
	taskRepo := initFakeRepo()
	id := 2
	item, err := taskRepo.Get(id)
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
	openedItems := taskRepo.GetOpened()
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
	closedItems := taskRepo.GetClosed()
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

// Fake Repo as a simple slice
type list struct {
	repo []task.Item
}

func (l *list) Save(t task.Item) (ID int) {
	t.ID = len(l.repo)
	l.repo = append(l.repo, t)
	return t.ID
}
func (l list) Get(ID int) (t task.Item, err error) {
	for _, it := range l.repo {
		if it.ID == ID {
			return it, nil
		}
	}
	return t, fmt.Errorf("Could not found ID %d", ID)
}
func (l list) GetAll() []task.Item {
	return l.repo
}
func (l list) GetOpened() []task.Item {
	var items []task.Item
	for _, it := range l.repo {
		if it.State == task.Opened {
			items = append(items, it)
		}
	}
	return items
}
func (l list) GetClosed() []task.Item {
	var items []task.Item
	for _, it := range l.repo {
		if it.State == task.Closed {
			items = append(items, it)
		}
	}
	return items
}
func (l *list) Update(item task.Item) {
	l.repo[item.ID] = item
}
