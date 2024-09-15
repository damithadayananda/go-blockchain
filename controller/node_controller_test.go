package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getNodesToBeInformed(t *testing.T) {
	receivedNodes := []string{
		"http://localhost:8080",
		"http://localhost:8081",
	}
	knownNodes := []string{
		"http://localhost:8082",
		"http://localhost:8081",
	}
	node := getNodesToBeInformed(receivedNodes, knownNodes)
	assert.Equal(t, []string{"http://localhost:8082"}, node)
}
