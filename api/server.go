package api

import (
	"fmt"
	"log"
)

// Server ...
type Server struct {
	nutrientStore NutrientStore
}

// NewServer ...
func NewServer(
	ns NutrientStore,
) Server {
	return Server{
		nutrientStore: ns,
	}
}

// GetNutrients ...
func (s Server) GetNutrients() {
	nutrients, err := s.nutrientStore.GetNutrients()
	if err != nil {
		log.Fatalln(err)
	}
	for _, n := range nutrients {
		fmt.Printf("%#v\n", n)
	}
}
