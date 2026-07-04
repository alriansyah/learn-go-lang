package simple

type BarRepository struct {
}

// Function Provider / Constructor
func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

// Function Provider / Constructor
func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{BarRepository: barRepository}
}
