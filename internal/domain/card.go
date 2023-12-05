package domain

import "github.com/gofrs/uuid"

type Card struct {
	id      uuid.UUID
	name    string
	balance int
	isOpen  bool
	BillID  uuid.UUID
}

func (c *Card) ID() uuid.UUID {
	return c.id
}

func (c *Card) Name() string { return c.name }
func (c *Card) Balance() int { return c.balance }
func (c *Card) IsOpen() bool { return c.isOpen }
