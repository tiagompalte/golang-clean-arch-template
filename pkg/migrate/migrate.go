package migrate

import "context"

type Migrate interface {
	Create(ctx context.Context, name string) error
	Up(ctx context.Context) error
	Down(ctx context.Context) error
	Fix(ctx context.Context, version int) error
}
