package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/proyecto-dnd/backend/cmd/server/router"
	"github.com/proyecto-dnd/backend/pkg/firebaseConnection"
	"github.com/proyecto-dnd/backend/pkg/s3"

	_ "github.com/proyecto-dnd/backend/docs"
)

// @title           Dicelogger API
// @version         1.0
// @description     API endpoints documentation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := ConnectDB()
	firebaseApp := firebaseConnection.InitializeFirebaseApp()
	go s3.InitializeS3()

	
	
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())
	
	router := router.NewRouter(engine, db, firebaseApp)
	router.MapRoutes()
	
	//PARA DOCKERIZAR CAMBIAR localhost por 0.0.0.0
	if err := engine.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
	defer db.Close()
}

// func initializeFirebaseApp() *firebase.App {
// 	ctx := &gin.Context{}
// 	opt := option.WithCredentialsFile("../serviceAccountKey.json")

// 	conf := &firebase.Config{
// 		ProjectID:     "logoflegends-19383",
// 		StorageBucket: "logoflegends-19383.appspot.com",
// 	}

// 	app, err := firebase.NewApp(ctx, conf, opt)
// 	if err != nil {
// 		log.Fatalf("error initializing app: %v\n", err)
// 	}

// 	// firestoreClient, err := firebaseApp.Firestore(ctx)
// 	// if err != nil {
// 	// 	log.Fatalf("error creating client: %v\n", err)
// 	// }

// 	// ref := firestoreClient.Collection("testDB").NewDoc()
// 	// result, err := ref.Set(ctx, map[string]interface{}{
// 	// 	"title":       "la primera...",
// 	// 	"description": "Y la segunda...",
// 	// })
// 	// if err != nil {
// 	// 	log.Fatalf("error when creating testDB: %v\n", err)
// 	// }

// 	// log.Printf("Result is: [%v]", result)
// 	fmt.Println("conectado a firebase")
// 	return app
// }

func ConnectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	// Create the data source.
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open the connection.
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
	// Check the connection.
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Se conectó a la base de datos bien")
	return db
}
