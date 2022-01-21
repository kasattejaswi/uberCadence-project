package localworker

import "fmt"

func StartWorker(path string) {
	fmt.Println("Reading configuration at:", path)
	fmt.Println("Starting worker")
}
