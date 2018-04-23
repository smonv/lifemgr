package api

// Nutrient ...
type Nutrient struct {
	ID    string
	Name  string
	Unit  string
	Group string
}

// NutrientStore ...
type NutrientStore interface {
	GetNutrients() (nutrients []*Nutrient, err error)
}
