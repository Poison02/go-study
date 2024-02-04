package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("Hello Go Module")
	logrus.Println(uuid.NewString())
}
