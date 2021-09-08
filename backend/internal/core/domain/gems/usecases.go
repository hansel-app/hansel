package gems

type UseCases struct {
	repository Repository
}

func NewUseCases(repository Repository) *UseCases {
	return &UseCases{
		repository: repository,
	}
}

func (u *UseCases) Drop(message string) (*Gem, error) {
	gem := &Gem{
		Message: message,
	}
	return u.repository.Add(gem)
}

func (u *UseCases) PickUp(id int64) (*Gem, error) {
	return u.repository.Get(id)
}
