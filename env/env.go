package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValue(key string) string {
  err := godotenv.Load()
  if err != nil {
    log.Fatal()
  }
  
  return os.Getenv(key)
}
