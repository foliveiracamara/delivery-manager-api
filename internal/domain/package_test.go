package domain

import (
	"testing"
	"time"

	"github.com/foliveiracamara/delivery-manager-api/internal/domain/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPackage(t *testing.T) {
	tests := []struct {
		name             string
		product          string
		destinationState string
		weightKg         float64
		expectedRegion   DestinationRegion
		shouldError      bool
		expectedErrorMsg string
	}{
		{
			name:             "should create package successfully for SP",
			product:          "Smartphone",
			destinationState: "SP",
			weightKg:         0.5,
			expectedRegion:   DestinationRegionSoutheast,
			shouldError:      false,
		},
		{
			name:             "should create package successfully for PR",
			product:          "Notebook",
			destinationState: "PR",
			weightKg:         2.0,
			expectedRegion:   DestinationRegionSouth,
			shouldError:      false,
		},
		{
			name:             "should create package successfully for GO",
			product:          "Mesa",
			destinationState: "GO",
			weightKg:         15.0,
			expectedRegion:   DestinationRegionMidwest,
			shouldError:      false,
		},
		{
			name:             "should create package successfully for PE",
			product:          "Livros",
			destinationState: "PE",
			weightKg:         3.2,
			expectedRegion:   DestinationRegionNortheast,
			shouldError:      false,
		},
		{
			name:             "should create package successfully for AM",
			product:          "Ferramentas",
			destinationState: "AM",
			weightKg:         8.7,
			expectedRegion:   DestinationRegionNorth,
			shouldError:      false,
		},
		{
			name:             "should fail with invalid region",
			product:          "Smartphone",
			destinationState: "SP",
			weightKg:         0.5,
			expectedRegion:   "invalid_region",
			shouldError:      true,
			expectedErrorMsg: "Invalid destination region",
		},
		{
			name:             "should create package with empty product (no validation)",
			product:          "",
			destinationState: "SP",
			weightKg:         0.5,
			expectedRegion:   DestinationRegionSoutheast,
			shouldError:      false,
		},
		{
			name:             "should create package with zero weight (no validation)",
			product:          "Smartphone",
			destinationState: "SP",
			weightKg:         0,
			expectedRegion:   DestinationRegionSoutheast,
			shouldError:      false,
		},
		{
			name:             "should create package with negative weight (no validation)",
			product:          "Smartphone",
			destinationState: "SP",
			weightKg:         -1.0,
			expectedRegion:   DestinationRegionSoutheast,
			shouldError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg, err := NewPackage(tt.product, tt.destinationState, tt.weightKg, tt.expectedRegion)

			if tt.shouldError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
				assert.Nil(t, pkg)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, pkg)
				assert.NotEmpty(t, pkg.ID)
				assert.Equal(t, tt.product, pkg.Product)
				assert.Equal(t, tt.destinationState, pkg.DestinationState)
				assert.Equal(t, tt.weightKg, pkg.WeightKg)
				assert.Equal(t, tt.expectedRegion, pkg.DestinationRegion)
				assert.Equal(t, StatusCreated, pkg.Status)
				assert.NotZero(t, pkg.CreatedAt)
				assert.NotZero(t, pkg.UpdatedAt)
			}
		})
	}
}

func TestPackage_UpdateStatus(t *testing.T) {
	tests := []struct {
		name         string
		status       PackageStatus
		shouldError  bool
		errorMessage string
		setupPackage func(*Package)
	}{
		{
			name:        "should update to created (no carrier needed)",
			status:      StatusCreated,
			shouldError: false,
		},
		{
			name:         "should fail to update to waiting pickup without carrier",
			status:       StatusWaitingPickup,
			shouldError:  true,
			errorMessage: "Package cannot be marked as esperando_coleta without a carrier assigned",
		},
		{
			name:        "should update to waiting pickup with carrier",
			status:      StatusWaitingPickup,
			shouldError: false,
			setupPackage: func(p *Package) {
				shipping := vo.NewShippingQuote("Test Carrier", "test-carrier", 25.50, 5)
				p.Shipping = &shipping
			},
		},
		{
			name:         "should fail to update to collected without carrier",
			status:       StatusCollected,
			shouldError:  true,
			errorMessage: "Package cannot be marked as coletado without a carrier assigned",
		},
		{
			name:        "should update to collected with carrier",
			status:      StatusCollected,
			shouldError: false,
			setupPackage: func(p *Package) {
				shipping := vo.NewShippingQuote("Test Carrier", "test-carrier", 25.50, 5)
				p.Shipping = &shipping
			},
		},
		{
			name:         "should fail to update to shipped without carrier",
			status:       StatusShipped,
			shouldError:  true,
			errorMessage: "Package cannot be marked as enviado without a carrier assigned",
		},
		{
			name:        "should update to shipped with carrier",
			status:      StatusShipped,
			shouldError: false,
			setupPackage: func(p *Package) {
				shipping := vo.NewShippingQuote("Test Carrier", "test-carrier", 25.50, 5)
				p.Shipping = &shipping
			},
		},
		{
			name:         "should fail to update to delivered without carrier",
			status:       StatusDelivered,
			shouldError:  true,
			errorMessage: "Package cannot be marked as entregue without a carrier assigned",
		},
		{
			name:        "should update to delivered with carrier",
			status:      StatusDelivered,
			shouldError: false,
			setupPackage: func(p *Package) {
				shipping := vo.NewShippingQuote("Test Carrier", "test-carrier", 25.50, 5)
				p.Shipping = &shipping
			},
		},
		{
			name:         "should fail to update to lost without carrier",
			status:       StatusLost,
			shouldError:  true,
			errorMessage: "Package cannot be marked as extraviado without a carrier assigned",
		},
		{
			name:        "should update to lost with carrier",
			status:      StatusLost,
			shouldError: false,
			setupPackage: func(p *Package) {
				shipping := vo.NewShippingQuote("Test Carrier", "test-carrier", 25.50, 5)
				p.Shipping = &shipping
			},
		},
		{
			name:         "should fail with invalid status",
			status:       "invalid_status",
			shouldError:  true,
			errorMessage: "Invalid status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh package for each test
			testPkg, err := NewPackage("Test Product", "SP", 1.0, DestinationRegionSoutheast)
			require.NoError(t, err)

			// Setup package if needed
			if tt.setupPackage != nil {
				tt.setupPackage(testPkg)
			}

			originalUpdatedAt := testPkg.UpdatedAt
			time.Sleep(1 * time.Millisecond) // Ensure time difference

			err = testPkg.UpdateStatus(tt.status)

			if tt.shouldError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMessage)
				assert.NotEqual(t, tt.status, testPkg.Status)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.status, testPkg.Status)
				assert.True(t, testPkg.UpdatedAt.After(originalUpdatedAt))
			}
		})
	}
}

func TestPackage_AssignShipping(t *testing.T) {
	pkg, err := NewPackage("Test Product", "SP", 1.0, DestinationRegionSoutheast)
	require.NoError(t, err)

	shipping := vo.NewShippingQuote("Test Carrier", "test-carrier", 25.50, 5)

	t.Run("should assign shipping successfully", func(t *testing.T) {
		originalUpdatedAt := pkg.UpdatedAt
		time.Sleep(1 * time.Millisecond)

		pkg.AssignShipping(shipping)

		assert.Equal(t, &shipping, pkg.Shipping)
		assert.Equal(t, StatusWaitingPickup, pkg.Status)
		assert.True(t, pkg.UpdatedAt.After(originalUpdatedAt))
	})
}

func TestPackage_SortShippingsByDeliveryTime(t *testing.T) {
	pkg, err := NewPackage("Test Product", "SP", 1.0, DestinationRegionSoutheast)
	require.NoError(t, err)

	shippings := []vo.Shipping{
		vo.NewShippingQuote("Slow Carrier", "slow", 30.0, 10),
		vo.NewShippingQuote("Fast Carrier", "fast", 25.0, 3),
		vo.NewShippingQuote("Medium Carrier", "medium", 27.0, 7),
	}

	t.Run("should sort shippings by delivery time", func(t *testing.T) {
		sorted := pkg.SortShippingsByDeliveryTime(shippings)

		assert.Len(t, sorted, 3)
		assert.Equal(t, 3, sorted[0].EstimatedDays)  // Fastest first
		assert.Equal(t, 7, sorted[1].EstimatedDays)  // Medium
		assert.Equal(t, 10, sorted[2].EstimatedDays) // Slowest last
	})
}

func TestIsValidStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   PackageStatus
		expected bool
	}{
		{"valid created", StatusCreated, true},
		{"valid waiting pickup", StatusWaitingPickup, true},
		{"valid collected", StatusCollected, true},
		{"valid shipped", StatusShipped, true},
		{"valid delivered", StatusDelivered, true},
		{"valid lost", StatusLost, true},
		{"invalid status", "invalid", false},
		{"empty status", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidStatus(tt.status)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStateToRegionMapping(t *testing.T) {
	tests := []struct {
		state    string
		expected DestinationRegion
	}{
		{"SP", DestinationRegionSoutheast},
		{"RJ", DestinationRegionSoutheast},
		{"MG", DestinationRegionSoutheast},
		{"ES", DestinationRegionSoutheast},
		{"PR", DestinationRegionSouth},
		{"SC", DestinationRegionSouth},
		{"RS", DestinationRegionSouth},
		{"GO", DestinationRegionMidwest},
		{"MT", DestinationRegionMidwest},
		{"MS", DestinationRegionMidwest},
		{"DF", DestinationRegionMidwest},
		{"BA", DestinationRegionNortheast},
		{"PE", DestinationRegionNortheast},
		{"CE", DestinationRegionNortheast},
		{"MA", DestinationRegionNortheast},
		{"PB", DestinationRegionNortheast},
		{"RN", DestinationRegionNortheast},
		{"AL", DestinationRegionNortheast},
		{"SE", DestinationRegionNortheast},
		{"PI", DestinationRegionNortheast},
		{"AM", DestinationRegionNorth},
		{"PA", DestinationRegionNorth},
		{"AC", DestinationRegionNorth},
		{"RR", DestinationRegionNorth},
		{"RO", DestinationRegionNorth},
		{"AP", DestinationRegionNorth},
		{"TO", DestinationRegionNorth},
	}

	for _, tt := range tests {
		t.Run(tt.state, func(t *testing.T) {
			region, exists := StateToRegionMapping[tt.state]
			assert.True(t, exists)
			assert.Equal(t, tt.expected, region)
		})
	}
}
