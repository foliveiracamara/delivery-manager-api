package persistence

import (
	"testing"

	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryPackageRepository(t *testing.T) {
	repo := NewInMemoryPackageRepository()

	t.Run("should save and retrieve package", func(t *testing.T) {
		pkg, err := domain.NewPackage("Test Product", "SP", 2.5, domain.DestinationRegionSoutheast)
		require.NoError(t, err)

		// Save package
		err = repo.Save(pkg)
		assert.NoError(t, err)

		// Retrieve package
		retrieved, err := repo.GetByID(pkg.ID)
		assert.NoError(t, err)
		assert.Equal(t, pkg.ID, retrieved.ID)
		assert.Equal(t, pkg.Product, retrieved.Product)
		assert.Equal(t, pkg.DestinationState, retrieved.DestinationState)
		assert.Equal(t, pkg.WeightKg, retrieved.WeightKg)
	})

	t.Run("should return error when package not found", func(t *testing.T) {
		_, err := repo.GetByID("nonexistent-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Package not found")
	})

	t.Run("should update existing package", func(t *testing.T) {
		pkg, err := domain.NewPackage("Original Product", "SP", 2.5, domain.DestinationRegionSoutheast)
		require.NoError(t, err)

		// Save original package
		err = repo.Save(pkg)
		require.NoError(t, err)

		// Update package
		pkg.Product = "Updated Product"
		err = repo.Save(pkg)
		assert.NoError(t, err)

		// Retrieve and verify update
		retrieved, err := repo.GetByID(pkg.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Product", retrieved.Product)
	})
}
