package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShippingQuote(t *testing.T) {
	t.Run("should create shipping quote successfully", func(t *testing.T) {
		shipping := NewShippingQuote("Test Carrier", "test-carrier", 25.50, 5)

		assert.Equal(t, "Test Carrier", shipping.CarrierName)
		assert.Equal(t, "test-carrier", shipping.CarrierID)
		assert.Equal(t, 25.50, shipping.EstimatedPrice)
		assert.Equal(t, 5, shipping.EstimatedDays)
	})
}

func TestNewShippingRequest(t *testing.T) {
	t.Run("should create shipping request successfully", func(t *testing.T) {
		request := NewShippingRequest(2.5, "SP")

		assert.Equal(t, 2.5, request.WeightKg)
		assert.Equal(t, "SP", request.DestinationState)
	})
}

func TestShippingRequest_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		request  ShippingRequest
		expected bool
	}{
		{
			name: "valid request",
			request: ShippingRequest{
				WeightKg:         2.5,
				DestinationState: "SP",
			},
			expected: true,
		},
		{
			name: "invalid - zero weight",
			request: ShippingRequest{
				WeightKg:         0,
				DestinationState: "SP",
			},
			expected: false,
		},
		{
			name: "invalid - negative weight",
			request: ShippingRequest{
				WeightKg:         -1.0,
				DestinationState: "SP",
			},
			expected: false,
		},
		{
			name: "invalid - empty state",
			request: ShippingRequest{
				WeightKg:         2.5,
				DestinationState: "",
			},
			expected: false,
		},
		{
			name: "invalid - both invalid",
			request: ShippingRequest{
				WeightKg:         0,
				DestinationState: "",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.request.IsValid()
			assert.Equal(t, tt.expected, result)
		})
	}
}
