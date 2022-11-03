package task

type State int

const (
	Opened State = iota
	Closed
)

type Item struct {
	ID          int
	Description string
	State       State
}

type repository interface {
	Save(Item) (ID int)
	Update(Item)
	GetAll() []Item
	GetByID(ID int) (Item, error)
	GetByState(st State) []Item
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
	it, err := config.repo.GetByID(ID)
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
