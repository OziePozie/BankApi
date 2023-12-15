package user

import (
	"BankApi/internal/domain"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) Save(ctx context.Context, user *domain.User) error {

	log.Print(user.Email())

	_, err := r.pool.Exec(ctx, "INSERT INTO accounts (first_name, email, password) values ($1,$2,$3);",
		user.Name(), user.Email(), user.PasswordHash())
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil

}

func (r *Repository) FindByName(ctx context.Context, email string) (*domain.User, error) {

	//query := `SELECT account_id, first_name, password FROM accounts WHERE accounts.email=$1::TEXT;`

	row := r.pool.QueryRow(ctx, "SELECT account_id, first_name, password FROM accounts WHERE accounts.email='timak';", email)

	var model domain.User

	log.Println(row)

	var m struct {
		id       uuid.UUID
		name     string
		password string
	}

	row.Scan(
		&m.id,
		&m.name,
		&m.password)

	log.Print(m.password)

	model.SetId(m.id)
	model.SetName(m.name)
	model.SetPasswordHash([]byte((m.password)))

	return &model, nil
}

func modelToDomain() {

}
