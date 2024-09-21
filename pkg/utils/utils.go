package utils

import (
    "crypto/rand"
    "encoding/hex"
    "time"
)

// GenerateRandomID creates a random string ID with a given length.
func GenerateRandomID(length int) (string, error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// FormatTime returns the time formatted in a specific layout.
func FormatTime(t time.Time, layout string) string {
    return t.Format(layout)
}

// StringInSlice checks if a string is in a slice.
func StringInSlice(str string, list []string) bool {
    for _, v := range list {
        if v == str {
            return true
        }
    }
    return false
}
