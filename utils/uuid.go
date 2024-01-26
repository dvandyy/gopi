package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Generate unique ID based on timestamp and UUID
func GenerateUniqueID(prefix string) string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond) // current timestamp in milliseconds
	randomPart := uuid.New().String()[:6]                        // 6 characters of a random UUID
	return fmt.Sprintf("%s%d%s", prefix, timestamp, randomPart)
}
