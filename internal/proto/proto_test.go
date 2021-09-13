package proto

import (
	"testing"

	"github.com/dmartzol/goapi/goapi"
	"github.com/stretchr/testify/assert"
)

func TestGoapiAccount(t *testing.T) {
	testCases := []struct {
		description    string
		pbAccount      Account
		expectedResult goapi.Account
	}{
		{
			description: "simple test",
			pbAccount: Account{
				Id:        "635f86cf-3f9e-40c7-8d17-1015f497621b",
				FirstName: "dani",
				LastName:  "martinez",
				Email:     "dani@example.com",
			},
			expectedResult: goapi.Account{
				FirstName: "dani",
				LastName:  "martinez",
				Email:     "dani@example.com",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			expected := tt.expectedResult
			actual, err := GoapiAccount(&tt.pbAccount)
			assert.Nil(t, err)
			assert.Equal(t, actual.FirstName, expected.FirstName)
			assert.Equal(t, actual.LastName, expected.LastName)
			assert.Equal(t, actual.Email, expected.Email)
		})
	}
}
