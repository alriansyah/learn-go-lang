package simple

type SimpleRepository struct {
}

// Function Provider
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	*SimpleRepository
}

// Function Provider
func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}
