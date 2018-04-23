package main

import (
	"context"
	"log"
	"os"

	"github.com/smonv/lifemgr/api"
	"github.com/smonv/lifemgr/api/firestore"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	ctx = context.Background()
)

func main() {
	firebaseSecret := os.Getenv("FIREBASE_SECRET")
	options := option.WithCredentialsFile(firebaseSecret)
	conf := &firebase.Config{
		ProjectID: "lifemgr-dev",
	}

	app, err := firebase.NewApp(ctx, conf, options)
	if err != nil {
		log.Fatalln(err)
	}

	fs, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	ns := firestore.NewNutrientStore(fs)
	server := api.NewServer(ns)

	server.GetNutrients()
}
