package test

import (
	"bytes"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/cmd/server/handler"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/domain"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/products"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/test/mocks"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockStore mocks.MockStorage) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	repo := products.NewRepository(&mockStore)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Save())
	pr.GET("/", p.GetAll())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllProducts(t *testing.T) {
	// arrange
	mockStorage := mocks.MockStorage{
		DataMock: []domain.Product{{
			ID:    1,
			Name:  "Product1",
			Type:  "Test",
			Count: 10,
			Price: 10.5,
		}, {
			ID:    2,
			Name:  "Product2",
			Type:  "Test",
			Count: 5,
			Price: 1.5,
		},
		},
	}
	var resp []domain.Product
	r := createServer(mockStorage)
	req, rr := createRequestTest(http.MethodGet, "/products/", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, len(mockStorage.DataMock), len(resp))
}

func TestSaveProduct(t *testing.T) {
	// arrrange
	mockStorage := mocks.MockStorage{
		DataMock: []domain.Product{{
			ID:    1,
			Name:  "Product1",
			Type:  "Test",
			Count: 10,
			Price: 10.5,
		}, {
			ID:    2,
			Name:  "Product2",
			Type:  "Test",
			Count: 5,
			Price: 1.5,
		},
		},
	}
	var resp domain.Product
	r := createServer(mockStorage)
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
        "nombre": "Tester","tipo": "Funcional","cantidad": 10,"precio": 9
    }`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)
	// assert.Equal(t, resp, rr.Code)
}
