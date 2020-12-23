package main

import (
	"os"
	"testing"
)

func Test_configure(t *testing.T) {
	const (
		Endpoint = "localhost:9092"
		Topic    = "test"
		GroupId  = "test-consumer"
	)
	os.Setenv("ENDPOINT", Endpoint)
	os.Setenv("TOPIC", Topic)
	os.Setenv("GROUP_ID", GroupId)

	type args struct {
		c *Configuration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"read environment variable configuration",
			args {
				&Configuration{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
