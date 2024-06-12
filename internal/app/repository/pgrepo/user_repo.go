package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maximus969/users-app/internal/app/domain"
	"github.com/maximus969/users-app/internal/app/repository/models"
	"github.com/maximus969/users-app/internal/pkg/pg"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r UserRepo) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	dbUser := domainToUser(user)

	newId := uuid.New()
	created := time.Now()

	fmt.Println("new id: ", newId)
	fmt.Println("created: ", created)

	dbUser.Id = newId
	dbUser.Created = created

	fmt.Printf("user: %+v\n", user)
	fmt.Printf("dbUser: %+v\n", dbUser)

	var insertedUser models.User
	err := r.db.NewInsert().Model(&dbUser).Returning("*").Scan(ctx, &insertedUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to insert a user: %w", err)
	}

	domainUser, err := userToDomain(insertedUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create domain user: %w", err)
	}

	return domainUser, nil
}

func (r UserRepo) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var dbUser models.User
	err := r.db.NewSelect().Model(&dbUser).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("failed to get a user: %w", err)
	}

	user, err := userToDomain(dbUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create domain user: %w", err)
	}

	return user, nil
}

func (r UserRepo) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	dbUser := domainToUser(user)

	var updatedUser models.User
	err := r.db.NewUpdate().
		Model(&dbUser).
		Where("id = ?", dbUser.Id).
		ExcludeColumn("created", "id").
		Returning("*").
		Scan(ctx, &updatedUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to update a user: %w", err)
	}

	domainUser, err := userToDomain(updatedUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create domain user: %w", err)
	}

	return domainUser, nil
}

func (r UserRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.NewDelete().Model((*models.User)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete a user: %w", err)
	}

	return nil
}

func (r UserRepo) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []models.User
	query := r.db.NewSelect().Model(&users)

	err := query.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	domainUsers := make([]domain.User, len(users))
	for i, user := range users {
		domainUser, err := userToDomain(user)
		if err != nil {
			return nil, fmt.Errorf("failed to create domain user: %w", err)
		}

		domainUsers[i] = domainUser
	}

	return domainUsers, nil
}
