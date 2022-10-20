package products

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto/usuario/transacción específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
// Crear un mock de Storage, dicho mock debe contener en su data un producto/usuario/transacción específico cuyo nombre puede ser “Before Update”.
// El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase.
// Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto/usuario/transacción mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test.
// TODO implements mockedData to the MockStore struct
// TODO implements directly casting instead Marshal/Unmarshal
type MockStore struct {
	ReadWasCalled bool
}

func (m *MockStore) Read(data interface{}) error {
	m.ReadWasCalled = true
	product1 := Product{
		ID:        1,
		Name:      "car",
		Color:     "red",
		Price:     0,
		Stock:     0,
		Code:      "",
		Published: false,
		CreatedAt: time.Time{},
	}
	product2 := Product{
		ID:        2,
		Name:      "bycicle",
		Color:     "yellow",
		Price:     0,
		Stock:     0,
		Code:      "",
		Published: false,
		CreatedAt: time.Time{},
	}
	products := []Product{product1, product2}
	MarshalData, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return json.Unmarshal(MarshalData, &data)
}

func (m *MockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateName(t *testing.T) {
	// Arrange
	myMockStore := MockStore{}
	repository := NewRepository(&myMockStore)
	expectedResult := Product{
		ID:        1,
		Name:      "updatedName",
		Color:     "red",
		Price:     0,
		Stock:     0,
		Code:      "",
		Published: false,
		CreatedAt: time.Time{},
	}

	// Execute
	result, err := repository.UpdateName(1, "updatedName")

	assert.Equal(t, expectedResult, result)
	assert.True(t, myMockStore.ReadWasCalled)
	assert.Nil(t, err)
}
