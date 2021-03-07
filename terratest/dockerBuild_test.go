package test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/docker"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
)

func TestDockerBuild(t *testing.T) {
	t.Parallel()
	// Configure the tag to use on the Docker image.
	tag := "versus/go-upload:test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	// Build the Docker image.
	docker.Build(t, "../", buildOptions)
}
func TestDockerComposeBuild(t *testing.T) {
	t.Parallel()
	dockerComposeFile := "../docker-compose.yml"

	docker.RunDockerCompose(t, &docker.Options{}, "-f", dockerComposeFile, "build")

	dockerOptions := &docker.Options{}

	docker.RunDockerCompose(t, dockerOptions, "up", "-d")
	defer docker.RunDockerCompose(t, dockerOptions, "down")

	serverPort := 8080
	maxRetries := 5
	timeBetweenRetries := 2 * time.Second
	url := fmt.Sprintf("http://localhost:%d/health", serverPort)
	tlsConfig := tls.Config{}

	http_helper.HttpGetWithRetry(t, url, &tlsConfig, 200, "* Connect checker: OK", maxRetries, timeBetweenRetries)
}
