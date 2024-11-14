package api

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/Sayonara143/urfu-pp-hhru-backend/services/auth"
	"github.com/google/uuid"
)

type (
	// AuthService описывает методы аутентификации
	AuthService interface {
		RegisterUser(ctx context.Context, user *models.User) error
		LoginUser(ctx context.Context, email, password string) (string, error)
		VerifyToken(ctx context.Context, tokenString string) (*auth.Claims, error)
	}
	HHService interface {
		UserService
		StudentProfileService
		EmployerProfileService
		ResumeService
		JobVacancyService
		JobApplicationService
		ReviewService
		InterviewService
		SearchLogService
		BlacklistService
		SavedSearchService
	}

	// UserService описывает методы для управления пользователями
	UserService interface {
		UserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
		Users(ctx context.Context, limit, offset int) (int, []*models.User, error)
		UserInsert(ctx context.Context, user *models.User) error
		UserUpdate(ctx context.Context, user *models.User) error
		UserDelete(ctx context.Context, id uuid.UUID) error
	}

	// StudentProfileService описывает методы для профилей студентов
	StudentProfileService interface {
		StudentProfileByID(ctx context.Context, id uuid.UUID) (*models.StudentProfile, error)
		StudentProfiles(ctx context.Context, limit, offset int) (int, []*models.StudentProfile, error)
		StudentProfileInsert(ctx context.Context, profile *models.StudentProfile) error
		StudentProfileUpdate(ctx context.Context, profile *models.StudentProfile) error
		StudentProfileDelete(ctx context.Context, id uuid.UUID) error
	}

	// EmployerProfileService описывает методы для профилей работодателей
	EmployerProfileService interface {
		EmployerProfileByID(ctx context.Context, id uuid.UUID) (*models.EmployerProfile, error)
		EmployerProfiles(ctx context.Context, limit, offset int) (int, []*models.EmployerProfile, error)
		EmployerProfileInsert(ctx context.Context, profile *models.EmployerProfile) error
		EmployerProfileUpdate(ctx context.Context, profile *models.EmployerProfile) error
		EmployerProfileDelete(ctx context.Context, id uuid.UUID) error
	}

	// ResumeService описывает методы для работы с резюме
	ResumeService interface {
		ResumeByID(ctx context.Context, id uuid.UUID) (*models.Resume, error)
		Resumes(ctx context.Context, limit, offset int) (int, []*models.Resume, error)
		ResumeInsert(ctx context.Context, resume *models.Resume) error
		ResumeUpdate(ctx context.Context, resume *models.Resume) error
		ResumeDelete(ctx context.Context, id uuid.UUID) error
	}

	// JobVacancyService описывает методы для работы с вакансиями
	JobVacancyService interface {
		JobVacancyByID(ctx context.Context, id uuid.UUID) (*models.JobVacancy, error)
		JobVacancies(ctx context.Context, limit, offset int) (int, []*models.JobVacancy, error)
		JobVacancyInsert(ctx context.Context, vacancy *models.JobVacancy) error
		JobVacancyUpdate(ctx context.Context, vacancy *models.JobVacancy) error
		JobVacancyDelete(ctx context.Context, id uuid.UUID) error
	}

	// JobApplicationService описывает методы для работы с заявками
	JobApplicationService interface {
		JobApplicationByID(ctx context.Context, id uuid.UUID) (*models.JobApplication, error)
		JobApplications(ctx context.Context, limit, offset int) (int, []*models.JobApplication, error)
		JobApplicationInsert(ctx context.Context, application *models.JobApplication) error
		JobApplicationUpdate(ctx context.Context, application *models.JobApplication) error
		JobApplicationDelete(ctx context.Context, id uuid.UUID) error
	}

	// ReviewService описывает методы для работы с отзывами
	ReviewService interface {
		ReviewByID(ctx context.Context, id uuid.UUID) (*models.Review, error)
		Reviews(ctx context.Context, limit, offset int) (int, []*models.Review, error)
		ReviewInsert(ctx context.Context, review *models.Review) error
		ReviewUpdate(ctx context.Context, review *models.Review) error
		ReviewDelete(ctx context.Context, id uuid.UUID) error
	}

	// InterviewService описывает методы для работы с собеседованиями
	InterviewService interface {
		InterviewByID(ctx context.Context, id uuid.UUID) (*models.Interview, error)
		Interviews(ctx context.Context, limit, offset int) (int, []*models.Interview, error)
		InterviewInsert(ctx context.Context, interview *models.Interview) error
		InterviewUpdate(ctx context.Context, interview *models.Interview) error
		InterviewDelete(ctx context.Context, id uuid.UUID) error
	}

	// BlacklistService описывает методы для работы с чёрным списком
	BlacklistService interface {
		BlacklistEntryByID(ctx context.Context, id uuid.UUID) (*models.BlacklistEntry, error)
		BlacklistEntries(ctx context.Context, limit, offset int) (int, []*models.BlacklistEntry, error)
		BlacklistEntryInsert(ctx context.Context, entry *models.BlacklistEntry) error
		BlacklistEntryUpdate(ctx context.Context, entry *models.BlacklistEntry) error
		BlacklistEntryDelete(ctx context.Context, id uuid.UUID) error
	}

	// SearchLogService описывает методы для работы с логами поиска
	SearchLogService interface {
		SearchLogInsert(ctx context.Context, log *models.SearchLog) error
		SearchLogs(ctx context.Context, limit, offset int) (int, []*models.SearchLog, error)
	}

	// SavedSearchService описывает методы для работы с сохранёнными поисками
	SavedSearchService interface {
		SavedSearchByID(ctx context.Context, id uuid.UUID) (*models.SavedSearch, error)
		SavedSearches(ctx context.Context, limit, offset int) (int, []*models.SavedSearch, error)
		SavedSearchInsert(ctx context.Context, search *models.SavedSearch) error
		SavedSearchDelete(ctx context.Context, id uuid.UUID) error
	}

	Logger interface {
		SendWithFields(fields map[string]interface{})
		Error(err error, msg string)
	}
)
