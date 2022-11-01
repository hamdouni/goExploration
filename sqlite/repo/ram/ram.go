package ram

import (
	"fmt"
	"goExploration/sqlite/task"
)

// List is a simple in memory slice for storing tasks.
// It implements task.repository interface.
type List struct {
	repo []task.Item
}

func (l *List) Save(t task.Item) (ID int) {
	t.ID = len(l.repo)
	l.repo = append(l.repo, t)
	return t.ID
}
func (l List) Get(ID int) (t task.Item, err error) {
	for _, it := range l.repo {
		if it.ID == ID {
			return it, nil
		}
	}
	return t, fmt.Errorf("Could not found ID %d", ID)
}
func (l List) GetAll() []task.Item {
	return l.repo
}
func (l List) GetOpened() []task.Item {
	var items []task.Item
	for _, it := range l.repo {
		if it.State == task.Opened {
			items = append(items, it)
		}
	}
	return items
}
func (l List) GetClosed() []task.Item {
	var items []task.Item
	for _, it := range l.repo {
		if it.State == task.Closed {
			items = append(items, it)
		}
	}
	return items
}
func (l *List) Update(item task.Item) {
	l.repo[item.ID] = item
}
