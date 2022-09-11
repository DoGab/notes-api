package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dogab/notes-api/pkg/utils"
	"github.com/joho/godotenv"
)

var DbPort uint64
var DbHost, DbName, DbUser, DbPassword string
var SecretKey []byte

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	// Return the value of the variable
	return os.Getenv(key)
}

func InitEnvs() {
	DbHost = utils.GetEnv("DB_HOST", "localhost")
	DbUser = utils.GetEnv("DB_USER", "postgres")
	DbName = utils.GetEnv("DB_NAME", "notes-api")
	DbPassword = utils.GetEnv("DB_PASSWORD", "postgres")
	port := utils.GetEnv("DB_PORT", "5432")
	var err error
	DbPort, err = strconv.ParseUint(port, 10, 32)
	if err != nil {
		log.Println("Could not parse DB_PORT to int")
	}
	SecretKey = []byte(utils.GetEnv("SECRET_KEY", "secretsigningkey"))
}
