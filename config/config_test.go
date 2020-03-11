package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	var temp *int

	defer func(t *int) {
		fmt.Println(*t)
	}(temp)

	var n = 12
	temp = &n
}
