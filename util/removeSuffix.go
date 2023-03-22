package util

import (
	"fmt"
	"strings"
)

func removeSuffix(s string) (string, error) {
	//cerca l'ultima occorrenza di .
	index := strings.LastIndex(s, ".")
	if index == -1 {
		return "", fmt.Errorf("no suffix found")
	}
	return s[:index], nil
}
