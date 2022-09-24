package cli_test

import (
	"fmt"
	"testing"

	"github.com/ffelipelimao/ports-adapters-architecture/adapter/cli"
	mock_application "github.com/ffelipelimao/ports-adapters-architecture/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Run(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	name := "test"
	price := 36.99
	status := "enable"
	ID := "abc"

	productMock := mock_application.NewMockIProduct(controller)
	productMock.EXPECT().GetID().Return(ID).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()

	service := mock_application.NewMockIProductService(controller)
	service.EXPECT().Create(name, price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(ID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("product ID %s with the name %s has been created with price %f", ID, name, price)

	result, err := cli.Run(service, "create", "", name, price)
	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("product %s has been enable", name)
	result, err = cli.Run(service, "enable", ID, "", 0)
	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("product %s has been disable", name)
	result, err = cli.Run(service, "disable", ID, "", 0)
	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("product ID %s, the name %s and the price %f", ID, name, price)
	result, err = cli.Run(service, "get", ID, "", 0)
	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)

}
