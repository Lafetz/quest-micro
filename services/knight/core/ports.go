package knight

import (
	"context"
)

type KnightServiceApi interface {
	GetKnight(context.Context, string) (*Knight, error)
	AddKnight(context.Context, Knight) (*Knight, error)
}
type KnightRepository interface {
	GetKnight(context.Context, string) (*Knight, error)
	AddKnight(context.Context, Knight) (*Knight, error)
}
