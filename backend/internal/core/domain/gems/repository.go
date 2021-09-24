package gems

type Repository interface {
	GetPendingCollectionByUser(userId int64) ([]Gem, error)
	Add(gem *Gem) (int64, error)
	PickUpGem(id int64) error
	GetGemLogs(userId int64) (map[int64]([]Gem), error)
}
