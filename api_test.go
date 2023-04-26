package api

import (
	"testing"
)

func TestAPI(t *testing.T) {
	api := New(Config{BaseURL: "http://192.168.222.135:7860/"})
	memoryResult, err := api.Memory()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ram: %+v\ncuda: %+v", memoryResult.RAM, memoryResult.Cuda)
}
