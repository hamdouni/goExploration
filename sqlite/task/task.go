package task

type state int

const (
	Opened state = iota
	Closed
)

type Item struct {
	ID          int
	Description string
	State       state
}

type repository interface {
	Save(Item) (ID int)
	Update(Item)
	Get(ID int) (Item, error)
	GetAll() []Item
	GetOpened() []Item
	GetClosed() []Item
}

var config = struct {
	repo repository
}{}

func Init(r repository) {
	config.repo = r
}

func Create(desc string) int {
	return config.repo.Save(Item{
		Description: desc,
		State:       Opened,
	})
}

func Close(ID int) error {
	it, err := config.repo.Get(ID)
	if err != nil {
		return err
	}
	it.State = Closed
	config.repo.Update(it)
	return nil
}

func (it Item) String() string {
	return it.Description
}
