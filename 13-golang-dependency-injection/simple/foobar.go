package simple

type FooBarService struct {
	*FooService
	*BarService
}

// Function Provider / Constructor
func NewFooBarService(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{FooService: fooService, BarService: barService}
}
