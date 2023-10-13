// Package stores manages all operations related to store entities.
package stores

// Storer interface defines common methods for managing user entities.
// This allows us to have different types of storers (say, MySQLStorer, PostgresStorer),
// and as long as they implement these methods, they can be used interchangeably.
type Storer interface {
	Create(usr User) error
	Update(usr User) error
	Delete(usr User) error
}

// StorerInterface is a global variable which will be used by services to interact with store.
// This helps in Dependency Injection and to swap storers without changing code in services.
// this is not recommended because anyone can change the value of the interface
var StorerInterface Storer

// Service struct is our main service that will use a Storer to handle operations.
// We embed the Storer interface inside the Service struct. This means that we can call
// the methods defined in the Storer interface directly on a Service instance.
type Service struct {
	Storer
}

// NewService function is a constructor for Service. It takes an implementer of Storer interface,
// and returns a new Service with the provided Storer. This design allows creating services with
// different storers at runtime.
func NewService(storer Storer) Service {
	s := Service{Storer: storer}
	return s
}
