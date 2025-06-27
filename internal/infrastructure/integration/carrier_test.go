package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCarrier(t *testing.T) {
	t.Run("should create carrier successfully", func(t *testing.T) {
		regions := []CarrierRegion{
			{
				Region:        "sudeste",
				EstimatedDays: 5,
				PricePerKg:    10.0,
			},
		}

		carrier := NewCarrier("test-carrier", "Test Carrier", regions)

		assert.Equal(t, "test-carrier", carrier.ID)
		assert.Equal(t, "Test Carrier", carrier.Name)
		assert.Len(t, carrier.Regions, 1)
		assert.Equal(t, "sudeste", carrier.Regions[0].Region)
		assert.Equal(t, 5, carrier.Regions[0].EstimatedDays)
		assert.Equal(t, 10.0, carrier.Regions[0].PricePerKg)
	})
}

func TestCarrier_GetRegionInfo(t *testing.T) {
	carrier := &Carrier{
		ID:   "test-carrier",
		Name: "Test Carrier",
		Regions: []CarrierRegion{
			{
				Region:        "sudeste",
				EstimatedDays: 5,
				PricePerKg:    10.0,
			},
			{
				Region:        "sul",
				EstimatedDays: 7,
				PricePerKg:    8.0,
			},
		},
	}

	t.Run("should return region info when region exists", func(t *testing.T) {
		region, exists := carrier.GetRegionInfo("sudeste")

		assert.True(t, exists)
		assert.NotNil(t, region)
		assert.Equal(t, "sudeste", region.Region)
		assert.Equal(t, 5, region.EstimatedDays)
		assert.Equal(t, 10.0, region.PricePerKg)
	})

	t.Run("should return nil when region does not exist", func(t *testing.T) {
		region, exists := carrier.GetRegionInfo("norte")

		assert.False(t, exists)
		assert.Nil(t, region)
	})
}

func TestCarrier_CalculateShipping(t *testing.T) {
	carrier := &Carrier{
		ID:   "test-carrier",
		Name: "Test Carrier",
		Regions: []CarrierRegion{
			{
				Region:        "sudeste",
				EstimatedDays: 5,
				PricePerKg:    10.0,
			},
		},
	}

	t.Run("should calculate shipping for valid region", func(t *testing.T) {
		price, days, ok := carrier.CalculateShipping("sudeste", 2.0)

		assert.True(t, ok)
		assert.Equal(t, 20.0, price) // 2.0 * 10.0
		assert.Equal(t, 5, days)
	})

	t.Run("should return minimum price for light packages", func(t *testing.T) {
		price, days, ok := carrier.CalculateShipping("sudeste", 0.5)

		assert.True(t, ok)
		assert.Equal(t, 10.0, price) // Minimum price (price per kg)
		assert.Equal(t, 5, days)
	})

	t.Run("should return false for invalid region", func(t *testing.T) {
		price, days, ok := carrier.CalculateShipping("norte", 2.0)

		assert.False(t, ok)
		assert.Equal(t, 0.0, price)
		assert.Equal(t, 0, days)
	})
}

func TestCarrier_IsAvailableForRegion(t *testing.T) {
	carrier := &Carrier{
		ID:   "test-carrier",
		Name: "Test Carrier",
		Regions: []CarrierRegion{
			{
				Region:        "sudeste",
				EstimatedDays: 5,
				PricePerKg:    10.0,
			},
		},
	}

	t.Run("should return true for available region", func(t *testing.T) {
		available := carrier.IsAvailableForRegion("sudeste")
		assert.True(t, available)
	})

	t.Run("should return false for unavailable region", func(t *testing.T) {
		available := carrier.IsAvailableForRegion("norte")
		assert.False(t, available)
	})
}

func TestCarrier_GetName(t *testing.T) {
	carrier := &Carrier{
		ID:   "test-carrier",
		Name: "Test Carrier",
	}

	t.Run("should return carrier name", func(t *testing.T) {
		name := carrier.GetName()
		assert.Equal(t, "Test Carrier", name)
	})
}

func TestCarrier_GetID(t *testing.T) {
	carrier := &Carrier{
		ID:   "test-carrier",
		Name: "Test Carrier",
	}

	t.Run("should return carrier ID", func(t *testing.T) {
		id := carrier.GetID()
		assert.Equal(t, "test-carrier", id)
	})
}

func TestCarrierRepositoryImpl(t *testing.T) {
	repo := NewCarrierRepository()

	t.Run("should return all carriers", func(t *testing.T) {
		carriers := repo.GetAll()

		assert.NotEmpty(t, carriers)
		assert.Len(t, carriers, 3) // We have 3 carriers in the mock data

		// Verify we have the expected carriers
		carrierNames := make(map[string]bool)
		for _, carrier := range carriers {
			carrierNames[carrier.Name] = true
		}

		assert.True(t, carrierNames["Nebulix Logística"])
		assert.True(t, carrierNames["RotaFácil Transportes"])
		assert.True(t, carrierNames["Moventra Express"])
	})

	t.Run("should return carrier by ID", func(t *testing.T) {
		carrier, err := repo.GetByID("nebulix")

		assert.NoError(t, err)
		assert.NotNil(t, carrier)
		assert.Equal(t, "nebulix", carrier.ID)
		assert.Equal(t, "Nebulix Logística", carrier.Name)
	})

	t.Run("should return error for non-existent carrier", func(t *testing.T) {
		carrier, err := repo.GetByID("nonexistent")

		assert.Error(t, err)
		assert.Nil(t, carrier)
		assert.Contains(t, err.Error(), "Carrier not found")
	})
}

func TestGetAvailableCarriers(t *testing.T) {
	carriers := getAvailableCarriers()

	t.Run("should return correct number of carriers", func(t *testing.T) {
		assert.Len(t, carriers, 3)
	})

	t.Run("should have Nebulix Logística with correct data", func(t *testing.T) {
		nebulix := carriers[0]
		assert.Equal(t, "nebulix", nebulix.ID)
		assert.Equal(t, "Nebulix Logística", nebulix.Name)
		assert.Len(t, nebulix.Regions, 2)

		// Check regions
		southRegion, exists := nebulix.GetRegionInfo("sul")
		assert.True(t, exists)
		assert.Equal(t, 4, southRegion.EstimatedDays)
		assert.Equal(t, 5.90, southRegion.PricePerKg)

		southeastRegion, exists := nebulix.GetRegionInfo("sudeste")
		assert.True(t, exists)
		assert.Equal(t, 4, southeastRegion.EstimatedDays)
		assert.Equal(t, 5.90, southeastRegion.PricePerKg)
	})

	t.Run("should have RotaFácil Transportes with correct data", func(t *testing.T) {
		rotafacil := carriers[1]
		assert.Equal(t, "rotafacil", rotafacil.ID)
		assert.Equal(t, "RotaFácil Transportes", rotafacil.Name)
		assert.Len(t, rotafacil.Regions, 4)
	})

	t.Run("should have Moventra Express with correct data", func(t *testing.T) {
		moventra := carriers[2]
		assert.Equal(t, "moventra", moventra.ID)
		assert.Equal(t, "Moventra Express", moventra.Name)
		assert.Len(t, moventra.Regions, 2)
	})
}
