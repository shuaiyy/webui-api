package api

import (
	"github.com/goccy/go-json"
)

type memory struct {
	RAM struct {
		Free  float64 `json:"free"`
		Used  int64   `json:"used"`
		Total float64 `json:"total"`
	} `json:"ram"`
	Cuda struct {
		System struct {
			Free  int64 `json:"free"`
			Used  int64 `json:"used"`
			Total int64 `json:"total"`
		} `json:"system"`
		Active    MemoryStats `json:"active"`
		Allocated MemoryStats `json:"allocated"`
		Reserved  MemoryStats `json:"reserved"`
		Inactive  MemoryStats `json:"inactive"`
		Events    struct {
			Retries int `json:"retries"`
			Oom     int `json:"oom"`
		} `json:"events"`
	} `json:"cuda"`
}

type MemoryStats struct {
	Current int64 `json:"current"`
	Peak    int64 `json:"peak"`
}

// Memory Get memory info
func (a *api) Memory() (result *memory, err error) {
	resp, erro := a.get(a.Config.Path.Memory)
	if erro != nil {
		err = erro
		return
	}

	err = json.Unmarshal(resp, &result)
	return
}
