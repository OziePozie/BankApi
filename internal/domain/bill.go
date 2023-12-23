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

func (b *Bill) SetId(id uuid.UUID) {
	b.id = id
}

func (b *Bill) SetName(name string) {
	b.name = name
}

func (b *Bill) SetBalance(balance int) {
	b.balance = balance
}

func (b *Bill) SetIsOpen(isOpen bool) {
	b.isOpen = isOpen
}

func (b *Bill) SetUserID(UserID uuid.UUID) {
	b.UserID = UserID
}

func (b *Bill) SetCards(Cards []uuid.UUID) {
	b.Cards = Cards
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
	FindAllByUserWithILIKE(ctx context.Context, billName string, offset int, limit int, userId uuid.UUID) (*[]Bill, error)
	DepositAmount(ctx context.Context, billID uuid.UUID, amount int) (newBalance int, err error)
	GetBillByBillIDAndUserIDEquals(ctx context.Context, userID uuid.UUID, billID uuid.UUID) (*Bill, error)
}

func (b *Bill) Validate() error {
	if b.name == "" {
		return &EmptyFieldError{field: "name"}
	}

	return nil
}
