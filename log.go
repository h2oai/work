package work

import "fmt"

var (
	logError = LogErrorFunc
)

// Default error logging function. Logs to stdout using fmt.Printf.
func LogErrorFunc(key string, err error) {
	fmt.Printf("ERROR: %s - %s\n", key, err.Error())
}

// Used to set custom error logging function. Make sure to call before anything else.
func SetCustomLogErrorFunc(f func(key string, err error)) {
	logError = f
}
