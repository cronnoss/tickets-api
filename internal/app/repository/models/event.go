package models

import "github.com/uptrace/bun"

type Event struct {
	bun.BaseModel `bun:"table:events"`
	ID            int    `bun:",pk,autoincrement"`
	ShowID        int    `bun:",notnull"`
	Date          string `bun:",notnull"`
	CreatedAt     string `bun:",nullzero"`
	UpdatedAt     string `bun:",nullzero"`
}
