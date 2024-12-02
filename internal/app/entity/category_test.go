package entity

import (
	"testing"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

func TestCategoryValidateNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		category    Category
		expectedErr error
	}{
		{
			name: "ShouldNotReturnError",
			category: Category{
				Name:   "category",
				UserID: uint32(1),
			},
		},
		{
			name: "ShouldReturnErrorIfNameIsEmpty",
			category: Category{
				Name:   "",
				UserID: uint32(1),
			},
			expectedErr: errors.AggregatedError{errors.NewEmptyParameterError("name")},
		},
		{
			name: "ShouldReturnErrorIfUserIDIsEmpty",
			category: Category{
				Name:   "task",
				UserID: 0,
			},
			expectedErr: errors.AggregatedError{errors.NewEmptyParameterError("user_id")},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.category.ValidateNew()

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
