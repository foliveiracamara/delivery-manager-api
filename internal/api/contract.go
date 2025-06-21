package api

import "context"

type Api interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
