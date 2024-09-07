package models

import "github.com/uptrace/bun"

type Show struct {
	bun.BaseModel `bun:"table:shows"`
	ID            int `bun:",pk,autoincrement"`
	Name          string
	CreatedAt     string `bun:",nullzero"`
	UpdatedAt     string `bun:",nullzero"`
}
