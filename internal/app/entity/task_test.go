package entity

import (
	"testing"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

func TestTaskValidateNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		task        Task
		expectedErr error
	}{
		{
			name: "ShouldNotReturnError",
			task: Task{
				Name:        "task",
				Description: "new task",
				UserID:      uint32(1),
			},
		},
		{
			name: "ShouldReturnErrorIfNameIsEmpty",
			task: Task{
				Name:        "",
				Description: "new task",
				UserID:      uint32(1),
			},
			expectedErr: errors.AggregatedError{errors.NewEmptyParameterError("name")},
		},
		{
			name: "ShouldReturnErrorIfDescriptionIsEmpty",
			task: Task{
				Name:        "task",
				Description: "",
				UserID:      uint32(1),
			},
			expectedErr: errors.AggregatedError{errors.NewEmptyParameterError("description")},
		},
		{
			name: "ShouldReturnErrorIfUserIDIsEmpty",
			task: Task{
				Name:        "task",
				Description: "new task",
				UserID:      0,
			},
			expectedErr: errors.AggregatedError{errors.NewEmptyParameterError("user_id")},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.task.ValidateNew()

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
