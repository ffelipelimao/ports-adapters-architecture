package application

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProduct_Enable(t *testing.T) {
	product := Product{
		Name:   "Hello",
		Status: DISABLE,
		Price:  10,
	}

	err := product.Enable()
	assert.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	assert.Equal(t, "the price should be greater than zero", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := Product{
		Name:   "Hello",
		Status: ENABLE,
		Price:  0,
	}

	err := product.Disable()
	assert.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	assert.Equal(t, "the price should be zero to disable product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := Product{
		ID:     uuid.NewString(),
		Name:   "Hello",
		Status: DISABLE,
		Price:  float64(10),
	}

	_, err := product.IsValid()
	assert.Nil(t, err)

	product.Status = "fail"
	_, err = product.IsValid()
	assert.Equal(t, "must be enable or disable", err.Error())

	product.Status = ENABLE
	product.Price = 0
	_, err = product.IsValid()
	assert.Equal(t, "Price: must be greater than zero.", err.Error())

}
