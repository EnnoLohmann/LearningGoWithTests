package main

import (
	"learning_go_with_tests/mocking"
	"os"
	"time"
)

func main() {
	mocking.Countdown(os.Stdout, &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep})
}
