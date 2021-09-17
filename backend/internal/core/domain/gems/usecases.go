package gems

type UseCases struct {
	repository Repository
}

func NewUseCases(repository Repository) *UseCases {
	return &UseCases{
		repository: repository,
	}
}

func (u *UseCases) Drop(*Gem gem) (*Gem, error) {
	return u.repository.Add(gem)
}

func (u *UseCases) Get(id int64) (*Gem, error) {
	return u.repository.Get(id)
}

func (u *UseCases) GetPendingCollectionByUser(userId int64) ([]Gem, error) {
	return u.repository.GetPendingCollectionByUser(userId)
}
