package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

// TestLoadConfig tests setting and reading environment variables
func TestLoadConfig(t *testing.T) {
	// LoadConfig should return an error when we didn't set any sector ID
	t.Run("no-sector_id", func(t *testing.T) {
		err := LoadConfig()
		assert.NotNil(t, err, "DNS_SECTOR_ID is required")
	})

	sectorIds := []int{10, 99, 100, 15000}

	// Trying to set sector IDs above to environment, load the config and after that, trying to call GetSectorId and it
	// should return the same value that we set
	for _, sectorId := range sectorIds {
		t.Run(fmt.Sprintf("sector-id-%d", sectorId), func(t *testing.T) {
			err := os.Setenv("DNS_SECTOR_ID", strconv.Itoa(sectorId))
			assert.Nil(t, err, "failed to set the env variable")

			err = LoadConfig()
			assert.Nil(t, err, "failed to read from env")

			got := GetSectorId()
			assert.Equal(t, sectorId, got, "key values should be the same")
		})

	}

	// GetPort should return 8080 when we didn't set any port
	t.Run("no-port", func(t *testing.T) {
		want := 8080
		got := GetPort()
		assert.Equal(t, want, got, "port should be 8080 when it's not specified")
	})

	ports := []int{8000, 8080, 9000}

	// Trying to set ports above to environment, load the config and after that, trying to call GetPort and it should
	// return the same value that we set
	for _, port := range ports {
		t.Run(fmt.Sprintf("port-%d", port), func(t *testing.T) {
			err := os.Setenv("DNS_PORT", strconv.Itoa(port))
			assert.Nil(t, err, "failed to set the env variable")

			err = LoadConfig()
			assert.Nil(t, err, "failed to read from env")

			got := GetPort()
			assert.Equal(t, port, got, "key values should be the same")
		})
	}
}
