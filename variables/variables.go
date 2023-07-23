package variables

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func getEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

var RABBITMQ_DEFAULT_USER = getEnv("RABBITMQ_DEFAULT_USER")
var RABBITMQ_DEFAULT_PASS = getEnv("RABBITMQ_DEFAULT_PASS")
var RABBITMQ_DEFAULT_HOST = getEnv("RABBITMQ_DEFAULT_HOST")
var RABBITMQ_DEFAULT_PORT = getEnv("RABBITMQ_DEFAULT_PORT")
var RABBITMQ_MAIN_QUEUE = getEnv("RABBITMQ_MAIN_QUEUE")
var RABBITMQ_URL = fmt.Sprintf(
															 "amqp://%s:%s@%s:%s", 
															 RABBITMQ_DEFAULT_USER,
															 RABBITMQ_DEFAULT_PASS,
															 RABBITMQ_DEFAULT_HOST,
															 RABBITMQ_DEFAULT_PORT,
															)