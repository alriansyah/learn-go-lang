package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

// Function Provider
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{
		Error: true,
	}
}

type SimpleService struct {
	*SimpleRepository
}

// Function Provider
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
