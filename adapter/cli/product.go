package cli

import (
	"fmt"

	"github.com/ffelipelimao/ports-adapters-architecture/application"
)

func Run(service application.IProductService, action string, ID string, name string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(name, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product ID %s with the name %s has been created with price %f", product.GetID(), product.GetName(), product.GetPrice())

	case "enable":
		product, err := service.Get(ID)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product %s has been enable", res.GetName())

	case "disable":
		product, err := service.Get(ID)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product %s has been disable", res.GetName())

	default:
		product, err := service.Get(ID)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product ID %s, the name %s and the price %f", product.GetID(), product.GetName(), product.GetPrice())
	}

	return result, nil
}
