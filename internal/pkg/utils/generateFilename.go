package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

func GenerateFilename(originalName string) string {
	prefix := "ktp_"

	// generate 6 random bytes -> 12 hex chars
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		// fallback to timestamp + sanitized name
		sanitized := strings.ReplaceAll(originalName, " ", "_")
		return prefix + strconv.FormatInt(time.Now().Unix(), 10) + "_" + sanitized
	}
	randStr := hex.EncodeToString(b)

	// include a short timestamp to reduce collision probability
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	sanitized := strings.ReplaceAll(originalName, " ", "_")

	return prefix + randStr + "_" + ts + "_" + sanitized
}
