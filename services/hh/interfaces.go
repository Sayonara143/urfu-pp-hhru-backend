package hh

import (
	"context"

	"github.com/google/uuid"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
)

type (
	Storage interface {
		// Методы для работы с транзакциями
		CtxWithTx(ctx context.Context) (context.Context, error)
		TxCommit(ctx context.Context) error
		TxRollback(ctx context.Context) error

		// Методы для работы с пользователями
		UserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
		Users(ctx context.Context, limit, offset int) (int, []*models.User, error)
		UserInsert(ctx context.Context, user *models.User) error
		UserUpdate(ctx context.Context, user *models.User) error
		UserDelete(ctx context.Context, id uuid.UUID) error

		// Методы для работы с профилями
		StudentProfileByID(ctx context.Context, id uuid.UUID) (*models.StudentProfile, error)
		StudentProfiles(ctx context.Context, limit, offset int) (int, []*models.StudentProfile, error)
		StudentProfileInsert(ctx context.Context, profile *models.StudentProfile) error
		StudentProfileUpdate(ctx context.Context, profile *models.StudentProfile) error
		StudentProfileDelete(ctx context.Context, id uuid.UUID) error

		EmployerProfileByID(ctx context.Context, id uuid.UUID) (*models.EmployerProfile, error)
		EmployerProfiles(ctx context.Context, limit, offset int) (int, []*models.EmployerProfile, error)
		EmployerProfileInsert(ctx context.Context, profile *models.EmployerProfile) error
		EmployerProfileUpdate(ctx context.Context, profile *models.EmployerProfile) error
		EmployerProfileDelete(ctx context.Context, id uuid.UUID) error

		// Методы для работы с резюме
		ResumeByID(ctx context.Context, id uuid.UUID) (*models.Resume, error)
		Resumes(ctx context.Context, limit, offset int) (int, []*models.Resume, error)
		ResumeInsert(ctx context.Context, resume *models.Resume) error
		ResumeUpdate(ctx context.Context, resume *models.Resume) error
		ResumeDelete(ctx context.Context, id uuid.UUID) error

		// Методы для работы с вакансиями
		JobVacancyByID(ctx context.Context, id uuid.UUID) (*models.JobVacancy, error)
		JobVacancies(ctx context.Context, limit, offset int) (int, []*models.JobVacancy, error)
		JobVacancyInsert(ctx context.Context, vacancy *models.JobVacancy) error
		JobVacancyUpdate(ctx context.Context, vacancy *models.JobVacancy) error
		JobVacancyDelete(ctx context.Context, id uuid.UUID) error

		// Методы для работы с заявками
		JobApplicationByID(ctx context.Context, id uuid.UUID) (*models.JobApplication, error)
		JobApplications(ctx context.Context, limit, offset int) (int, []*models.JobApplication, error)
		JobApplicationInsert(ctx context.Context, application *models.JobApplication) error
		JobApplicationUpdate(ctx context.Context, application *models.JobApplication) error
		JobApplicationDelete(ctx context.Context, id uuid.UUID) error

		// Методы для работы с отзывами
		ReviewByID(ctx context.Context, id uuid.UUID) (*models.Review, error)
		Reviews(ctx context.Context, limit, offset int) (int, []*models.Review, error)
		ReviewInsert(ctx context.Context, review *models.Review) error
		ReviewUpdate(ctx context.Context, review *models.Review) error
		ReviewDelete(ctx context.Context, id uuid.UUID) error

		// Методы для работы с собеседованиями
		InterviewByID(ctx context.Context, id uuid.UUID) (*models.Interview, error)
		Interviews(ctx context.Context, limit, offset int) (int, []*models.Interview, error)
		InterviewInsert(ctx context.Context, interview *models.Interview) error
		InterviewUpdate(ctx context.Context, interview *models.Interview) error
		InterviewDelete(ctx context.Context, id uuid.UUID) error

		// Методы для работы с чёрным списком
		BlacklistEntryByID(ctx context.Context, id uuid.UUID) (*models.BlacklistEntry, error)
		BlacklistEntries(ctx context.Context, limit, offset int) (int, []*models.BlacklistEntry, error)
		BlacklistEntryInsert(ctx context.Context, entry *models.BlacklistEntry) error
		BlacklistEntryUpdate(ctx context.Context, entry *models.BlacklistEntry) error
		BlacklistEntryDelete(ctx context.Context, id uuid.UUID) error

		// Методы для логов поиска
		SearchLogInsert(ctx context.Context, log *models.SearchLog) error
		SearchLogs(ctx context.Context, limit, offset int) (int, []*models.SearchLog, error)

		// Методы для сохранённых поисков
		SavedSearchByID(ctx context.Context, id uuid.UUID) (*models.SavedSearch, error)
		SavedSearches(ctx context.Context, limit, offset int) (int, []*models.SavedSearch, error)
		SavedSearchInsert(ctx context.Context, search *models.SavedSearch) error
		SavedSearchDelete(ctx context.Context, id uuid.UUID) error
	}

	Paginator interface {
		Limit() int
		Offset() int
		SetItemsTotal(itemsTotal int)
	}
	Pagination interface {
		Get(ctx context.Context) Paginator
	}
)
