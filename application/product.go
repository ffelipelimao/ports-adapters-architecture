package application

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	DISABLE = "disabled"
	ENABLE  = "enabled"
)

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
}

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
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
