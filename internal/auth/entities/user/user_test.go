package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneNumber(t *testing.T) {
	testCases := []struct {
		name          string
		phone         int64
		expectedValid bool
	}{
		{
			name:          "valid viet nam phone number",
			phone:         84913174196,
			expectedValid: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			got := isValidVNPhone(tc.phone)
			assert.Equal(tt, tc.expectedValid, got)
		})
	}
}
