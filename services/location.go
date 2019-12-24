package services

import (
	"github.com/moeen/atlas-dns/config"
)

// FindLocation calculates the location of given coordinates and velocity
func FindLocation(x, y, z, vel float32) float32 {
	// Since the return type of the function is float32, we need to convert the sector ID to float32 as well
	sectorId := float32(config.GetSectorId())

	return (x * sectorId) + (y * sectorId) + (z * sectorId) + vel
}
