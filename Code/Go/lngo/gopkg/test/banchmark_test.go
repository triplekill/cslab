package test

import (
	"log"
	"os"
	"runtime/pprof"
	"testing"
)

func Test_unmarshalt(t *testing.T) {
	f, err := os.Create("unmarshalt.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 100000; i++ {
		unmarshalt()
	}
}
