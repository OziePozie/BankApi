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

type Model struct {
	id       uuid.UUID
	name     string
	email    string
	password string
}

func NewUserRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) Save(ctx context.Context, user *domain.User) error {

	log.Print(user.Email())

	_, err := r.pool.Exec(ctx, "INSERT INTO accounts (acc_uuid,first_name, email, password) values ($1,$2,$3, $4);",
		user.ID(), user.Name(), user.Email(), user.PasswordHash())
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil

}

func (r *Repository) FindByName(ctx context.Context, email string) (*domain.User, error) {

	//query := `SELECT account_id, first_name, password FROM accounts WHERE accounts.email=$1::TEXT;`

	row := r.pool.QueryRow(ctx, "SELECT acc_uuid, first_name, email, password FROM accounts WHERE email=$1", email)

	var m Model

	err := row.Scan(
		&m.id,
		&m.name,
		&m.email,
		&m.password)
	if err != nil {
		log.Print(err)
	}

	log.Print("postgres pass = ", m.password)

	user := m.modelToDomain()

	return user, nil
}

func (m *Model) modelToDomain() *domain.User {
	var model domain.User

	model.SetId(m.id)
	model.SetEmail(m.email)
	model.SetName(m.name)
	model.SetPasswordHash([]byte((m.password)))

	return &model
}
