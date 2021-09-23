package gems

type UseCases struct {
	repository Repository
}

func NewUseCases(repository Repository) *UseCases {
	return &UseCases{
		repository: repository,
	}
}

func (u *UseCases) Drop(gem *Gem) (int64, error) {
	return u.repository.Add(gem)
}

func (u *UseCases) Get(id int64) (*Gem, error) {
	return u.repository.Get(id)
}

func (u *UseCases) GetPendingCollectionByUser(userId int64) ([]Gem, error) {
	return u.repository.GetPendingCollectionByUser(userId)
}

func (u *UseCases) PickUpGem(id int64) error {
	return u.repository.PickUpGem(id)

// Returns a map of friend_id to list of gem logs.
// Not returning User object here directly to keep Gem domain independent from User domain.
// UserInfo of friend will be fetched higher up in the service instead.
func (u *UseCases) GetGemLogs(userID int64) (map[int64]([]Gem), error) {
	return u.repository.GetGemLogs(userID)
}
