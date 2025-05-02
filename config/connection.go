package config

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func ConnectionDB() {
	// 🛠 Read env vars inside the function
	cassandraHost := os.Getenv("CASSANDRA_HOST")
	cassandraUser := os.Getenv("CASSANDRA_USER")
	cassandraPassword := os.Getenv("CASSANDRA_PASSWORD")

	if cassandraHost == "" {
		log.Fatal("CASSANDRA_HOST not set")
	}

	cluster := gocql.NewCluster(cassandraHost)
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{
		NumRetries: 3,
	}
	cluster.Keyspace = "library2"
	cluster.Consistency = gocql.One
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cassandraUser,
		Password: cassandraPassword,
	}

	var err error
	// Now assign the session to the global Session variable
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Error connecting to Cassandra: %v", err)
	}

	// Successfully connected
	log.Println("Connected OK")
}
