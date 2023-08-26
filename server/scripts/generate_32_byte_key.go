package main

import (
	"crypto/rand"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	logrus.WithField("key", fmt.Sprintf("%x", key)).Info("generated 32 byte key")
}
