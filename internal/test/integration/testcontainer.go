package integration

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupINetMockContainer(ctx context.Context, tb testing.TB, exposedPorts ...string) (imContainer testcontainers.Container, err error) {
	_, fileName, _, _ := runtime.Caller(0)

	var repoRoot string
	if repoRoot, err = filepath.Abs(filepath.Join(filepath.Dir(fileName), "..", "..", "..")); err != nil {
		return
	}

	var waits []wait.Strategy

	for _, port := range exposedPorts {
		waits = append(waits, wait.ForListeningPort(nat.Port(port)))
	}

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    repoRoot,
			Dockerfile: filepath.Join("./", "testdata", "integration.dockerfile"),
		},
		SkipReaper:   true,
		ExposedPorts: exposedPorts,
		WaitingFor:   wait.ForAll(waits...),
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
