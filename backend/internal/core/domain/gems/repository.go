package gems

type Repository interface {
	Get(id int64) (*Gem, error)
	GetAllForUser(userId int64) ([]Gem, error)
	Add(message string) (*Gem, error)
	Update(gem *Gem) (*Gem, error)
	Remove(id int64) error
}
