package gems

type Repository interface {
	Get(id int64) (*Gem, error)
	GetAllForUser(userId int64) ([]Gem, error)
	GetPendingCollectionByUser(userId int64) ([]Gem, error)
	Add(gem *Gem) (int64, error)
	Update(gem *Gem) (*Gem, error)
	Remove(id int64) error
	GetGemLogs(id int64) (map[int64]([]Gem), error)
}
