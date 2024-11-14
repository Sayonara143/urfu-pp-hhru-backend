package hh

import (
	"context"
	"time"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

// ResumeByID возвращает резюме по его ID
func (s *Service) ResumeByID(ctx context.Context, id uuid.UUID) (*models.Resume, error) {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	resume, err := s.store.ResumeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return resume, s.store.TxCommit(ctx)
}

// Resumes возвращает список резюме с пагинацией
func (s *Service) Resumes(ctx context.Context, limit, offset int) (int, []*models.Resume, error) {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	total, resumes, err := s.store.Resumes(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, resumes, s.store.TxCommit(ctx)
}

// ResumeInsert добавляет новое резюме
func (s *Service) ResumeInsert(ctx context.Context, resume *models.Resume) error {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	now := time.Now()
	resume.ID = uuid.New()
	resume.CreatedAt = &now
	resume.UpdatedAt = &now

	if err = s.store.ResumeInsert(ctx, resume); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// ResumeUpdate обновляет данные резюме
func (s *Service) ResumeUpdate(ctx context.Context, resume *models.Resume) error {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	now := time.Now()
	resume.UpdatedAt = &now

	if err = s.store.ResumeUpdate(ctx, resume); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// ResumeDelete удаляет резюме по его ID
func (s *Service) ResumeDelete(ctx context.Context, id uuid.UUID) error {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err = s.store.ResumeDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
