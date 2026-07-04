package simple

type FooRepository struct {
}

// Function Provider / Constructor
func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

// Function Provider / Constructor
func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{FooRepository: fooRepository}
}
