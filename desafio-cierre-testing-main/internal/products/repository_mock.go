package products

type MockRepository struct {
	DataMock []Product
	Error    error
}

func (m *MockRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.DataMock, nil
}
