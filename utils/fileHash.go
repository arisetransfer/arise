package utils

import (
  "crypto/sha256"
  "io"
  "os"
  "log"
  "encoding/hex"
)

func FileHash(name string) (string) {
  f, err := os.Open(name)
	if err != nil {
		log.Println("File Not Found!")
		return ""
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
    return ""
	}
	hash := hex.EncodeToString(h.Sum(nil))
  return hash
}
