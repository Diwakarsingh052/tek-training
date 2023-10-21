package core

// Storer interface declares the behavior this package needs to persists and
// retrieve data.
//
//go:generate mockgen -source core.go -destination mockcore/core_mock.go -package mockcore
type Storer interface {
	Create(usr User) error
	Update(usr User) error
	Delete(usr User) error
}

var StorerInterface Storer

// Store manages the set of APIs for data access.
type Store struct {
	Storer
}

// NewStore constructs a core for data api access.
func NewStore(storer Storer) *Store {
	//initialize the store struct
	return &Store{Storer: storer}

}
