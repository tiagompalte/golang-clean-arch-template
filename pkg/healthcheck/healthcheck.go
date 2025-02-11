package healthcheck

import "context"

type HealthCheck interface {
	IsHealthy(ctx context.Context) (bool, error)
}
