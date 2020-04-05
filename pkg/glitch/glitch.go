package glitch

import "fmt"

// LogError -
func LogError(err error) {
	fmt.Errorf("%v", err)
}
