package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
)

func (s *Storage) SearchLogInsert(ctx context.Context, log *models.SearchLog) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(log).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) SearchLogs(ctx context.Context, limit, offset int) (int, []*models.SearchLog, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var logs []*models.SearchLog
	total, err := tx.NewSelect().
		Model(&logs).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, logs, nil
}
