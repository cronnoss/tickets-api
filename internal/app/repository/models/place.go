package models

import "github.com/uptrace/bun"

type Place struct {
	bun.BaseModel `bun:"table:places"`
	ID            int     `bun:",pk,autoincrement"`
	X             float64 `bun:",notnull"`
	Y             float64 `bun:",notnull"`
	Width         float64 `bun:",notnull"`
	Height        float64 `bun:",notnull"`
	IsAvailable   bool    `bun:",notnull"`
	CreatedAt     string  `bun:",nullzero"`
	UpdatedAt     string  `bun:",nullzero"`
}
