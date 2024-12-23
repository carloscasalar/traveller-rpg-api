package payload

import "fmt"

// Error returns a JSON error message payload
func Error(message string) string {
	return fmt.Sprintf(`{"message": "%s"}`, message)
}
