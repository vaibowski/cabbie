package customer

import (
	"cabbie/models"
	"cabbie/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewCustomer(t *testing.T) {
	tests := []struct {
		name             string
		customer         models.Customer
		expectedResponse string
		expectedError    error
	}{
		{
			name: "customer is successfully created",
			customer: models.Customer{
				Name:     "Naruto",
				Email:    "lordseventh@gmail.com",
				Password: "dattebayo",
				Phone:    "1234",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := repository.CustomerRepository{MapDatastore: map[string]models.Customer{}}
			svc := Service{CustomerRepository: &repo}
			resp, err := svc.CreateNewCustomer(test.customer)
			assert.NotEmpty(t, resp)
			assert.NoError(t, err)
		})
	}
}
