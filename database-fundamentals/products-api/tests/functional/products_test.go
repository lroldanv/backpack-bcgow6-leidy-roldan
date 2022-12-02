package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/cmd/server/handler"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/internal/domain"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/internal/product"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/pkg/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		panic("can't connect to database")
	}
}
func TestStoreProductOk(t *testing.T) {
	new := domain.Product{
		Name:  "producto nuevo",
		Type:  "producto tipo",
		Count: 3,
		Price: 84.4,
	}

	product, err := json.Marshal(new)
	require.Nil(t, err) // require provides same global fucntion but terminate the current test

	req, rr := createRequest(http.MethodPost, "/api/v1/products/", product)
	s := createServer()
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// struct for assertion
	p := struct{ Data domain.Product }{}
	err = json.Unmarshal(rr.Body.Bytes(), &p)
	require.Nil(t, err)

	new.ID = p.Data.ID
	assert.Equal(t, new, p.Data)
}

func TestGetByNameOk(t *testing.T) {
	req, rr := createRequest(http.MethodGet, "/api/v1/products/", []byte(`{"name":"new product}`))
	s := createServer()
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Create server
func createServer() *gin.Engine {

	engine, db := db.ConnectDatabase()
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	gin.SetMode(gin.ReleaseMode)

	pr := engine.Group("api/v1/products")
	pr.GET("/", handler.GetByName())

	return engine
}

// Create test request
func createRequest(method string, url string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()

}
