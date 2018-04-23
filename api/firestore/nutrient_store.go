package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/smonv/lifemgr/api"
	"google.golang.org/api/iterator"
)

var (
	ctx = context.Background()
)

// NutrientStore ...
type NutrientStore struct {
	client *firestore.Client
}

// NewNutrientStore ...
func NewNutrientStore(c *firestore.Client) NutrientStore {
	return NutrientStore{
		client: c,
	}
}

// GetNutrients ...
func (ns NutrientStore) GetNutrients() (nutrients []*api.Nutrient, err error) {
	iter := ns.client.Collection("nutrients").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		nutrient := &api.Nutrient{
			Id: doc.Ref.ID,
		}
		doc.DataTo(&nutrient)
		nutrients = append(nutrients, nutrient)
	}

	return nutrients, nil
}
