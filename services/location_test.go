package services

import (
	"fmt"
	"github.com/moeen/atlas-dns/config"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

// TestFindLocation tests calculating location based on given coordinates and velocity
func TestFindLocation(t *testing.T) {
	// Data represents the data that our function needs
	type Data struct {
		X   float32
		Y   float32
		Z   float32
		Vel float32
	}

	cases := []struct {
		In       Data
		SectorId int
		Out      float32
	}{
		{
			In:       Data{X: 0, Y: 0, Z: 0, Vel: 0},
			SectorId: 1,
			Out:      0,
		},
		{
			In:       Data{X: 0, Y: 0, Z: 0, Vel: 20.1},
			SectorId: 460,
			Out:      20.1,
		},
		{
			In:       Data{X: 100, Y: 11.2, Z: 34, Vel: 21},
			SectorId: 1250,
			Out:      181521,
		},
		{
			In:       Data{X: -176.9, Y: 130, Z: -111, Vel: 5.6},
			SectorId: 120,
			Out:      -18942.4,
		},
		{
			In:       Data{X: 0, Y: 124.25, Z: -9.5, Vel: 5},
			SectorId: 1,
			Out:      119.75,
		},
	}

	// In each case we pass the `In` data and set the `SectorId` to environment and we expect the function to return
	// `Out`
	for i, c := range cases {
		t.Run(fmt.Sprintf("location-%d", i), func(t *testing.T) {
			_ = os.Setenv("DNS_SECTOR_ID", strconv.Itoa(c.SectorId))
			_ = config.LoadConfig()

			got := FindLocation(c.In.X, c.In.Y, c.In.Z, c.In.Vel)

			assert.Equalf(t, c.Out, got, "data: %+v sector id: %d", c.In, c.SectorId)
		})
	}

}
