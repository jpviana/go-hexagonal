package application_test

import (
	"github.com/jpviana/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T){
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T){
	product := application.Product{}
	product.id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.IsValid()
	require.Nil(t,err)

	product.Status = "INVALID"
	err = product.IsValid()
	require.Equal(t,"the status must be enabled ou disabled", err.Error())

	product.Status = application.ENABLED
	err = product.IsValid()
	require.Equal(t,"the status must be enabled ou disabled", err.Error())
	
	product.Price = -10
	err = product.IsValid()
	require.Equal(t,"the price must be greater than zero to enable the product", err.Error())
}
