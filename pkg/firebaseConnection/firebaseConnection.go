package firebaseConnection

import (
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var (
	ctx = &gin.Context{}
)

func InitializeFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")

	conf := &firebase.Config{
		ProjectID:     "dnd-dicelogger",
	}

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	fmt.Println("conectado a firebase")
	return app
}

func CreateFirebaseClient() *auth.Client {
	app := InitializeFirebaseApp()

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return client
}
