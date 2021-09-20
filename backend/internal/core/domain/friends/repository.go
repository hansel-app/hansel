package friends

type Repository interface {
	Get(id int64) ([]int64, error)
	Add(friend1_id int64, friend2_id int64) error
	Remove(id int64) error
}
