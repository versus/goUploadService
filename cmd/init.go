package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can't loading .env file")
	}

	secretKey, ok := os.LookupEnv("JWT_SECRET_KEY")
	if !ok || secretKey == "" {
		log.Fatalln("error: JWT_SECRET_KEY variable not found! ")
	}

	tokenValidateDaysEnv, ok := os.LookupEnv("JWT_VALID_DAYS")
	if !ok || tokenValidateDaysEnv == "" {
		tokenValidateDays = defaultTokenValidateDays
	} else {
		tokenValidateDays, err = strconv.Atoi(tokenValidateDaysEnv)
		if err != nil {
			log.Println("variable JWT_VALID_DAYS can't convert to integer")
			tokenValidateDays = defaultTokenValidateDays
		}
	}

	portEnv, ok := os.LookupEnv("PORT")
	if !ok || portEnv == "" {
		port = defaultPort
	} else {
		port, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Println("variable PORT can't convert to integer")
			port = defaultPort
		}
	}

}
