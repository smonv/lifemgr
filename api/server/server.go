package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/smonv/lifemgr/api"
	"github.com/smonv/lifemgr/api/datastore"
)

// Server ...
type Server struct {
	nutrientStore datastore.Nutrient
}

// New ...
func New(
	ns datastore.Nutrient,
) Server {
	return Server{
		nutrientStore: ns,
	}
}

// GetNutrients ...
func (s Server) GetNutrients(_ context.Context, _ *empty.Empty) (*api.GetNutrientsResponse, error) {
	nutrients, err := s.nutrientStore.GetNutrients()
	if err != nil {
		return nil, err
	}

	return &api.GetNutrientsResponse{
		Nutrients: nutrients,
	}, nil
}
