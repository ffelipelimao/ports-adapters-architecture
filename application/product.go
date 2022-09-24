package application

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

//go:generate mockgen -destination=./mocks/application.go -source=./product.go

const (
	DISABLE = "disabled"
	ENABLE  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

type IProduct interface {
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	IsValid() (bool, error)
	Enable() error
	Disable() error
}

type IProductService interface {
	Get(ID string) (IProduct, error)
	Create(name string, price float64) (IProduct, error)
	Enable(product IProduct) (IProduct, error)
	Disable(product IProduct) (IProduct, error)
}

type ProductReader interface {
	Get(ID string) (IProduct, error)
}

type ProductWriter interface {
	Save(product IProduct) (IProduct, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

func NewProduct() *Product {
	return &Product{
		ID:     uuid.NewString(),
		Status: DISABLE,
	}
}

func (p Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLE
	}

	if p.Status != ENABLE && p.Status != DISABLE {
		return false, errors.New("must be enable or disable")
	}

	err := validation.ValidateStruct(&p,
		validation.Field(&p.Price, validation.Min(float64(0)), validation.Required.Error("must be greater than zero")),
		validation.Field(&p.ID, is.UUIDv4),
		validation.Field(&p.Name, validation.Required),
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLE
		return nil
	}
	return errors.New("the price should be greater than zero")

}
func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLE
		return nil
	}
	return errors.New("the price should be zero to disable product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
