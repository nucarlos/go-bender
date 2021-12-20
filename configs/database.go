package configs

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB Configuration
func Database() {
	// Define connection string MongoDB
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	// Define database por API
	database := os.Getenv("MONGODB_DATABASE")
	if len(database) == 0 {
		database = "bender"
	}

	// Define username to MongoDB
	username := os.Getenv("MONGODB_USERNAME")
	if len(username) == 0 {
		username = "bender"
	}

	// Define password to MongoDB
	password := os.Getenv("MONGODB_PASSWORD")
	if len(password) == 0 {
		password = "123456"
	}

	// Setting mechanism authentication
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      username,
		Password:      password,
	}

	// Setting config
	err := mgm.SetDefaultConfig(nil, database, options.Client().ApplyURI(connectionString).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}
}
