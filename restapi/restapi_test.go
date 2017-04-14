package restapi

import (
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestRestAPI(t *testing.T) {
	api := Init()
	api.Start("3000")
}
