package knight

import (
	"context"
)

type KnightServiceApi interface {
	KnightStatus(context.Context, string) (bool, error)
	AddKnight(context.Context, *Knight) (*Knight, error)
	UpdateStatus(context.Context, string, bool) error
}
type KnightRepository interface {
	GetKnight(context.Context, string) (*Knight, error)
	AddKnight(context.Context, *Knight) (*Knight, error)
	UpdateStatus(context.Context, string, bool) error
}
