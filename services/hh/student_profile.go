package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

// Получение профиля студента по ID
func (s *Service) StudentProfileByID(ctx context.Context, id uuid.UUID) (*models.StudentProfile, error) {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	profile, err := s.store.StudentProfileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return profile, s.store.TxCommit(ctx)
}

// Получение списка профилей студентов с пагинацией
func (s *Service) StudentProfiles(ctx context.Context, limit, offset int) (int, []*models.StudentProfile, error) {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	total, profiles, err := s.store.StudentProfiles(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, profiles, s.store.TxCommit(ctx)
}

// Вставка нового профиля студента
func (s *Service) StudentProfileInsert(ctx context.Context, profile *models.StudentProfile) error {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err := s.store.StudentProfileInsert(ctx, profile); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// Обновление профиля студента
func (s *Service) StudentProfileUpdate(ctx context.Context, profile *models.StudentProfile) error {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err := s.store.StudentProfileUpdate(ctx, profile); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// Удаление профиля студента
func (s *Service) StudentProfileDelete(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err := s.store.StudentProfileDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
