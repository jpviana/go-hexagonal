package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (s *ProductService) Save(p ProductInterface) (ProductInterface, error) {
	product, err := s.Persistence.Save(p)
	if err != nil {
		return nil, err
	}
	return product, err
}
