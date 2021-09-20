package gems

type Repository interface {
	Get(id int64) (*Gem, error)
	GetAllForUser(userId int64) ([]Gem, error)
	GetPendingCollectionByUser(userId int64) ([]Gem, error)
	Add(gem *Gem) (string, error)
	Update(gem *Gem) (*Gem, error)
	Remove(id int64) error
}
