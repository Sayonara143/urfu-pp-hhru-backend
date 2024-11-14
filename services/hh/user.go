package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

// UserByID получает пользователя по его ID.
func (s *Service) UserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	user, err := s.store.UserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, s.store.TxCommit(ctx)
}

// Users возвращает список пользователей с пагинацией.
func (s *Service) Users(ctx context.Context, limit, offset int) (int, []*models.User, error) {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	total, users, err := s.store.Users(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, users, s.store.TxCommit(ctx)
}

// UserInsert вставляет нового пользователя в базу данных.
func (s *Service) UserInsert(ctx context.Context, user *models.User) error {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err := s.store.UserInsert(ctx, user); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// UserUpdate обновляет данные пользователя.
func (s *Service) UserUpdate(ctx context.Context, user *models.User) error {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err := s.store.UserUpdate(ctx, user); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// UserDelete удаляет пользователя из базы данных по его ID.
func (s *Service) UserDelete(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()
	if err := s.store.UserDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
