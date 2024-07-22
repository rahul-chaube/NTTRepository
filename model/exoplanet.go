package model

import (
	"fmt"
)

type Exoplanet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    int     `json:"distance"`       // in light years
	Radius      float64 `json:"radius"`         // in Earth-radius units
	Mass        float64 `json:"mass,omitempty"` // in Earth-mass units, only for Terrestrial
	Type        string  `json:"type"`           // GasGiant or Terrestrial
}

func (ex Exoplanet) IsValid() (bool, error) {

	if ex.Distance <= 10 || ex.Distance >= 1000 {
		return false, fmt.Errorf("distance must be between 10 and 1000 light years")

	}
	if ex.Radius <= 0.1 || ex.Radius >= 10 {
		return false, fmt.Errorf("radius must be between 0.1 and 10 Earth-radius units")
	}
	if ex.Type == "Terrestrial" {
		if ex.Mass <= 0.1 || ex.Mass >= 10 {
			return false, fmt.Errorf("mass must be between 0.1 and 10 Earth-mass units")
		}
	}

	return true, nil
}
