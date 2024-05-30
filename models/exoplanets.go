package models

type Exoplanet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    int     `json:"distance"`       // in light years
	Radius      float64 `json:"radius"`         // in Earth-radius units
	Mass        float64 `json:"mass,omitempty"` // in Earth-mass units, only for Terrestrial
	Type        string  `json:"type"`
}
