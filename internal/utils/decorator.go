package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func MeasureExecutionTime(f func()) func() {
	return func() {
		now := time.Now()

		// If something goes wrong, a SIGINT could be raised
		// and it will be intercepted here. Thus, this will
		// end the program, and print the execution time.
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT)
		// The magic happens here, in a separated goroutine
		go func() {
			<-sigs

			d := time.Since(now)
			fmt.Printf("Executed in %s\n", d)
			os.Exit(0)
		}()

		f()

		d := time.Since(now)
		fmt.Printf("Executed in %s\n", d)
	}

}
