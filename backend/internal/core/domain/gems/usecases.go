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
	return u.repository.Add(message)
}

func (u *UseCases) Get(id int64) (*Gem, error) {
	return u.repository.Get(id)
}
