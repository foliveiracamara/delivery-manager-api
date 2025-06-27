package service

import (
	"testing"

	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	"github.com/foliveiracamara/delivery-manager-api/internal/infrastructure/integration"
	apperr "github.com/foliveiracamara/delivery-manager-api/internal/shared/apperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockCarrierRepository implements CarrierRepository for testing
type MockCarrierRepository struct {
	carriers []*integration.Carrier
}

func (m *MockCarrierRepository) GetAll() []*integration.Carrier {
	return m.carriers
}

func (m *MockCarrierRepository) GetByID(id string) (*integration.Carrier, error) {
	for _, carrier := range m.carriers {
		if carrier.ID == id {
			return carrier, nil
		}
	}
	return nil, apperr.NewNotFoundError("Carrier not found")
}

func TestPackageService_Create(t *testing.T) {
	service := &PackageService{}

	t.Run("should create package successfully", func(t *testing.T) {
		pkg := &domain.Package{
			Product:           "Test Product",
			DestinationState:  "SP",
			WeightKg:          2.5,
			DestinationRegion: domain.DestinationRegionSoutheast,
		}

		result, err := service.Create(pkg)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.ID)
		assert.Equal(t, "Test Product", result.Product)
		assert.Equal(t, "SP", result.DestinationState)
		assert.Equal(t, 2.5, result.WeightKg)
		assert.Equal(t, domain.DestinationRegionSoutheast, result.DestinationRegion)
		assert.Equal(t, domain.StatusCreated, result.Status)
	})
}

func TestPackageService_UpdateStatus(t *testing.T) {
	service := &PackageService{}

	t.Run("should update status successfully", func(t *testing.T) {
		pkg, err := domain.NewPackage("Test Product", "SP", 2.5, domain.DestinationRegionSoutheast)
		require.NoError(t, err)

		err = service.UpdateStatus(pkg, domain.StatusWaitingPickup)

		assert.NoError(t, err)
		assert.Equal(t, domain.StatusWaitingPickup, pkg.Status)
	})
}

func TestPackageService_QuoteAvailableShippings(t *testing.T) {
	// Create mock carriers
	mockCarriers := []*integration.Carrier{
		{
			ID:   "carrier1",
			Name: "Fast Carrier",
			Regions: []integration.CarrierRegion{
				{
					Region:        "sudeste",
					EstimatedDays: 3,
					PricePerKg:    10.0,
				},
			},
		},
		{
			ID:   "carrier2",
			Name: "Slow Carrier",
			Regions: []integration.CarrierRegion{
				{
					Region:        "sudeste",
					EstimatedDays: 7,
					PricePerKg:    8.0,
				},
			},
		},
		{
			ID:   "carrier3",
			Name: "Wrong Region Carrier",
			Regions: []integration.CarrierRegion{
				{
					Region:        "sul",
					EstimatedDays: 5,
					PricePerKg:    9.0,
				},
			},
		},
	}

	mockRepo := &MockCarrierRepository{carriers: mockCarriers}
	service := &PackageService{carrierRepo: mockRepo}

	t.Run("should quote available shippings for southeast region", func(t *testing.T) {
		pkg, err := domain.NewPackage("Test Product", "SP", 2.0, domain.DestinationRegionSoutheast)
		require.NoError(t, err)

		shippings, err := service.QuoteAvailableShippings(pkg)

		assert.NoError(t, err)
		assert.Len(t, shippings, 2) // Only carriers that serve southeast

		// Should be sorted by delivery time (fastest first)
		assert.Equal(t, "Fast Carrier", shippings[0].CarrierName)
		assert.Equal(t, 3, shippings[0].EstimatedDays)
		assert.Equal(t, 20.0, shippings[0].EstimatedPrice) // 2.0 * 10.0

		assert.Equal(t, "Slow Carrier", shippings[1].CarrierName)
		assert.Equal(t, 7, shippings[1].EstimatedDays)
		assert.Equal(t, 16.0, shippings[1].EstimatedPrice) // 2.0 * 8.0
	})

	t.Run("should handle minimum price correctly", func(t *testing.T) {
		pkg, err := domain.NewPackage("Test Product", "SP", 0.5, domain.DestinationRegionSoutheast) // Very light package
		require.NoError(t, err)

		shippings, err := service.QuoteAvailableShippings(pkg)

		assert.NoError(t, err)
		assert.Len(t, shippings, 2)

		// Price should be at least the price per kg (minimum price)
		assert.Equal(t, 10.0, shippings[0].EstimatedPrice) // Minimum price, not 0.5 * 10.0
		assert.Equal(t, 8.0, shippings[1].EstimatedPrice)  // Minimum price, not 0.5 * 8.0
	})
}

func TestPackageService_HireCarrier(t *testing.T) {
	// Create mock carriers
	mockCarriers := []*integration.Carrier{
		{
			ID:   "carrier1",
			Name: "Test Carrier",
			Regions: []integration.CarrierRegion{
				{
					Region:        "sudeste",
					EstimatedDays: 5,
					PricePerKg:    10.0,
				},
			},
		},
	}

	mockRepo := &MockCarrierRepository{carriers: mockCarriers}
	service := &PackageService{carrierRepo: mockRepo}

	t.Run("should hire carrier successfully", func(t *testing.T) {
		pkg, err := domain.NewPackage("Test Product", "SP", 2.0, domain.DestinationRegionSoutheast)
		require.NoError(t, err)

		err = service.HireCarrier(pkg, "carrier1")

		assert.NoError(t, err)
		assert.NotNil(t, pkg.Shipping)
		assert.Equal(t, "Test Carrier", pkg.Shipping.CarrierName)
		assert.Equal(t, "carrier1", pkg.Shipping.CarrierID)
		assert.Equal(t, 20.0, pkg.Shipping.EstimatedPrice) // 2.0 * 10.0
		assert.Equal(t, 5, pkg.Shipping.EstimatedDays)
		assert.Equal(t, domain.StatusWaitingPickup, pkg.Status)
	})

	t.Run("should fail when package already has carrier", func(t *testing.T) {
		pkg, err := domain.NewPackage("Test Product", "SP", 2.0, domain.DestinationRegionSoutheast)
		require.NoError(t, err)

		// Assign first carrier
		err = service.HireCarrier(pkg, "carrier1")
		require.NoError(t, err)

		// Try to assign second carrier
		err = service.HireCarrier(pkg, "carrier1")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Package already has a carrier")
	})

	t.Run("should fail when carrier not found", func(t *testing.T) {
		pkg, err := domain.NewPackage("Test Product", "SP", 2.0, domain.DestinationRegionSoutheast)
		require.NoError(t, err)

		err = service.HireCarrier(pkg, "nonexistent")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Carrier not found")
	})

	t.Run("should fail when carrier does not serve region", func(t *testing.T) {
		// Create carrier that only serves south region
		southCarrier := &integration.Carrier{
			ID:   "south-carrier",
			Name: "South Carrier",
			Regions: []integration.CarrierRegion{
				{
					Region:        "sul",
					EstimatedDays: 5,
					PricePerKg:    10.0,
				},
			},
		}

		mockRepoWithSouth := &MockCarrierRepository{carriers: []*integration.Carrier{southCarrier}}
		serviceWithSouth := &PackageService{carrierRepo: mockRepoWithSouth}

		pkg, err := domain.NewPackage("Test Product", "SP", 2.0, domain.DestinationRegionSoutheast) // Southeast region
		require.NoError(t, err)

		err = serviceWithSouth.HireCarrier(pkg, "south-carrier")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Carrier does not serve the destination region")
	})
}
