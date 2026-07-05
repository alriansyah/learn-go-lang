package golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New() // pointer logger

	logger.Println("Hello Logger")
	fmt.Println("Hello Logger")
}
