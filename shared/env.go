package shared

import (
	"fmt"
	"syscall"
)

func EnvString(key, fallback string) string {
	fmt.Println("------------------")
	if val, ok := syscall.Getenv(key); ok {
		fmt.Println(val)
		return ":" + val
	}
	return fallback
}
