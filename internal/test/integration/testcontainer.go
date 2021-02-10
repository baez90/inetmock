package integration

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupINetMockContainer(ctx context.Context, tb testing.TB, exposedPorts ...string) (imContainer testcontainers.Container, err error) {
	_, fileName, _, _ := runtime.Caller(0)

	var repoRoot string
	if repoRoot, err = filepath.Abs(filepath.Join(filepath.Dir(fileName), "..", "..", "..")); err != nil {
		return
	}

	var plainHttpPresent = false
	for _, port := range exposedPorts {
		plainHttpPresent = plainHttpPresent || port == "80/tcp"
	}

	if !plainHttpPresent {
		exposedPorts = append(exposedPorts, "80/tcp")
	}

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    repoRoot,
			Dockerfile: filepath.Join("./", "testdata", "integration.dockerfile"),
		},
		ExposedPorts: exposedPorts,
		WaitingFor:   wait.ForListeningPort("80/tcp"),
	}

	imContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return
	}

	tb.Cleanup(func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = imContainer.Terminate(shutdownCtx)
	})

	return
}
