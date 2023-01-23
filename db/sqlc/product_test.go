package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// função para a criação de um produto aleatório
// retorna o Product criado no models.go
// CreateProductParams: tipagem definina no product.sql.go
func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Name:  "any_name",
		Price: 10,
	}

	// passagem de parâmetros do CreateProduct definido no product.sql.go
	product, err := testQueries.CreateProduct(context.Background(), arg)

	// utilização dos métodos do testify
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.NotEmpty(t, product.ID)
	require.NotEmpty(t, product.CreatedAt)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Price, product.Price)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	productRandomCreated := createRandomProduct(t)
	productFinded, err := testQueries.GetProduct(context.Background(), productRandomCreated.ID)

	require.NoError(t, err)
	require.NotEmpty(t, productRandomCreated)
	require.NotEmpty(t, productFinded)

	require.Equal(t, productRandomCreated.ID, productFinded.ID)
	require.Equal(t, productRandomCreated.Name, productFinded.Name)
	require.Equal(t, productRandomCreated.Price, productFinded.Price)
	require.Equal(t, productRandomCreated.CreatedAt, productFinded.CreatedAt)
}

func TestDeleteProduct(t *testing.T) {
	productRandomCreated := createRandomProduct(t)
	err := testQueries.DeleteProduct(context.Background(), productRandomCreated.ID)

	require.NoError(t, err)
}

func TestUpdateProduct(t *testing.T) {
	productRandomCreated := createRandomProduct(t)

	arg := UpdateProductParams{
		ID:    productRandomCreated.ID,
		Name:  "updated_name",
		Price: 20,
	}

	productUpdated, err := testQueries.UpdateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, productUpdated)

	require.Equal(t, productRandomCreated.ID, productUpdated.ID)
	require.Equal(t, arg.Name, productUpdated.Name)
	require.Equal(t, arg.Price, productUpdated.Price)
	require.Equal(t, productRandomCreated.CreatedAt, productUpdated.CreatedAt)
}

func TestGetProducts(t *testing.T) {
	productRandomCreated := createRandomProduct(t)
	productsFinded, err := testQueries.GetProducts(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, productRandomCreated)
	require.NotEmpty(t, productsFinded)

	for _, product := range productsFinded {
		require.NotEmpty(t, product.ID)
		require.NotEmpty(t, product.Name)
		require.NotEmpty(t, product.Price)
		require.NotEmpty(t, product.CreatedAt)
	}
}
