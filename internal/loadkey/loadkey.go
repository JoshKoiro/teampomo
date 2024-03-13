package loadkey

import (
	"fmt"
	"os"
)

func LoadKey() (string, error) {
	key, err := os.ReadFile("key.txt")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(key), nil
}
