package entity

import (
	"testing"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

func TestUserValidateNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		user        User
		expectedErr error
	}{
		{
			name: "ShouldNotReturnError",
			user: User{
				Name:  "User",
				Email: "user@email.com",
			},
		},
		{
			name: "ShouldReturnErrorIfNameIsEmpty",
			user: User{
				Name:  "",
				Email: "user@email.com",
			},
			expectedErr: errors.AggregatedError{errors.NewEmptyParameterError("name")},
		},
		{
			name: "ShouldReturnErrorIfEmailIsEmpty",
			user: User{
				Name:  "User",
				Email: "",
			},
			expectedErr: errors.AggregatedError{errors.NewEmptyParameterError("email")},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.user.ValidateNew()

			if tt.expectedErr == nil && err != nil {
				t.Errorf("should return no err but return %v", err)
			} else if tt.expectedErr != nil && err != nil && tt.expectedErr.Error() != errors.Cause(err).Error() {
				t.Errorf("should return %v but return %v", tt.expectedErr, err)
			} else if tt.expectedErr != nil && err == nil {
				t.Errorf("should return %v but non return", tt.expectedErr)
			}
		})
	}
}
