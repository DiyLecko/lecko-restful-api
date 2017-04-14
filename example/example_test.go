package example

import (
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestRestAPI(t *testing.T) {
	StartExampleRest()
}
