package datastore

import "github.com/smonv/lifemgr/api"

// Nutrient ...
type Nutrient interface {
	GetNutrients() (nutrients []*api.Nutrient, err error)
}
