package proto

import (
	"testing"

	"github.com/dmartzol/goapi/goapi"
	"github.com/stretchr/testify/assert"
)

func TestCoreAccount(t *testing.T) {
	testCases := []struct {
		description string
		pbAccount   Account
		should      goapi.Account
	}{
		{
			description: "simple test",
			pbAccount: Account{
				Id:        "635f86cf-3f9e-40c7-8d17-1015f497621b",
				FirstName: "dani",
				LastName:  "martinez",
				Email:     "dani@example.com",
			},
			should: goapi.Account{
				FirstName: "dani",
				LastName:  "martinez",
				Email:     "dani@example.com",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			got, err := CoreAccount(&tt.pbAccount)
			assert.Nil(t, err)
			assert.Equal(t, got.FirstName, tt.should.FirstName)
			assert.Equal(t, got.LastName, tt.should.LastName)
			assert.Equal(t, got.Email, tt.should.Email)
		})
	}
}
