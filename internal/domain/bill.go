package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type Bill struct {
	id      uuid.UUID
	name    string
	balance int
	isOpen  bool
	UserID  uuid.UUID
	Cards   []uuid.UUID
}

func NewBill(name string, userID uuid.UUID) *Bill {
	return &Bill{
		id:      uuid.Must(uuid.NewV7()),
		name:    name,
		UserID:  userID,
		balance: 0,
		isOpen:  false,
	}
}

func (b *Bill) ID() uuid.UUID { return b.id }
func (b *Bill) Name() string  { return b.name }
func (b *Bill) Balance() int  { return b.balance }
func (b *Bill) IsOpen() bool  { return b.isOpen }

func (b *Bill) Close() { b.isOpen = false }

type BillRepository interface {
	Save(ctx context.Context, bill *Bill) error
	GetAll(ctx context.Context) ([]Bill, error)
	GetByName(ctx context.Context, name string) (Bill, error)
}

func (b *Bill) Validate() error {
	if b.name == "" {
		return &EmptyFieldError{field: "name"}
	}

	return nil
}
