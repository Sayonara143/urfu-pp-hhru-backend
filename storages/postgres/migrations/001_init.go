package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	MigrationSet.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.Exec(`

CREATE TABLE users (
	id           UUID PRIMARY KEY,
	email        TEXT UNIQUE NOT NULL,
	full_name    TEXT NOT NULL,
	role         TEXT NOT NULL,
	created_at   TIMESTAMPTZ NOT NULL,
	updated_at   TIMESTAMPTZ NOT NULL
);

CREATE TABLE student_profiles (
	id           UUID PRIMARY KEY,
	user_id      UUID REFERENCES users(id) ON DELETE CASCADE,
	phone        TEXT,
	university   TEXT,
	faculty      TEXT,
	course       INT,
	skills       TEXT,
	languages    TEXT,
	about        TEXT,
	created_at   TIMESTAMPTZ NOT NULL,
	updated_at   TIMESTAMPTZ NOT NULL
);

CREATE TABLE employer_profiles (
	id                  UUID PRIMARY KEY,
	user_id             UUID REFERENCES users(id) ON DELETE CASCADE,
	company_name        TEXT NOT NULL,
	company_description TEXT,
	phone               TEXT,
	website             TEXT,
	created_at          TIMESTAMPTZ NOT NULL,
	updated_at          TIMESTAMPTZ NOT NULL
);

CREATE TABLE resumes (
	id                 UUID PRIMARY KEY,
	student_profile_id UUID REFERENCES student_profiles(id) ON DELETE CASCADE,
	title              TEXT NOT NULL,
	experience         TEXT,
	education          TEXT,
	skills             TEXT,
	languages          TEXT,
	created_at         TIMESTAMPTZ NOT NULL,
	updated_at         TIMESTAMPTZ NOT NULL
);

CREATE TABLE job_vacancies (
	id               UUID PRIMARY KEY,
	employer_id      UUID REFERENCES employer_profiles(id) ON DELETE CASCADE,
	title            TEXT NOT NULL,
	description      TEXT,
	requirements     TEXT,
	employment_type  TEXT NOT NULL,
	salary_range     TEXT,
	created_at       TIMESTAMPTZ NOT NULL,
	updated_at       TIMESTAMPTZ NOT NULL
);

CREATE TABLE job_applications (
	id                 UUID PRIMARY KEY,
	job_vacancy_id     UUID REFERENCES job_vacancies(id) ON DELETE CASCADE,
	student_profile_id UUID REFERENCES student_profiles(id) ON DELETE CASCADE,
	resume_id          UUID REFERENCES resumes(id) ON DELETE CASCADE,
	cover_letter       TEXT,
	status             TEXT NOT NULL,
	created_at         TIMESTAMPTZ NOT NULL,
	updated_at         TIMESTAMPTZ NOT NULL
);

CREATE TABLE reviews (
	id                 UUID PRIMARY KEY,
	student_profile_id UUID REFERENCES student_profiles(id) ON DELETE CASCADE,
	employer_id        UUID REFERENCES employer_profiles(id) ON DELETE CASCADE,
	review_type        TEXT NOT NULL,
	rating             INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
	text               TEXT,
	created_at         TIMESTAMPTZ NOT NULL
);

CREATE TABLE interviews (
	id                 UUID PRIMARY KEY,
	job_application_id UUID REFERENCES job_applications(id) ON DELETE CASCADE,
	start_time         TIMESTAMPTZ NOT NULL,
	end_time           TIMESTAMPTZ NOT NULL,
	status             TEXT NOT NULL,
	location           TEXT,
	notes              TEXT,
	created_at         TIMESTAMPTZ NOT NULL,
	updated_at         TIMESTAMPTZ NOT NULL
);

CREATE TABLE blacklist_entries (
	id           UUID PRIMARY KEY,
	type         TEXT NOT NULL CHECK (type IN ('employer_to_student', 'student_to_employer')),
	initiator_id UUID REFERENCES users(id) ON DELETE CASCADE,
	target_id    UUID REFERENCES users(id) ON DELETE CASCADE,
	reason       TEXT,
	created_at   TIMESTAMPTZ NOT NULL,
	updated_at   TIMESTAMPTZ NOT NULL
);

CREATE TABLE blacklist (
	id           UUID PRIMARY KEY,
	user_id      UUID REFERENCES users(id) ON DELETE CASCADE,
	reason       TEXT NOT NULL,
	banned_by    UUID REFERENCES users(id) ON DELETE SET NULL,
	permanent    BOOLEAN DEFAULT FALSE,
	created_at   TIMESTAMPTZ NOT NULL,
	updated_at   TIMESTAMPTZ NOT NULL
);

CREATE TABLE search_logs (
	id        UUID PRIMARY KEY,
	user_id   UUID REFERENCES users(id) ON DELETE CASCADE,
	query     TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE saved_searches (
	id        UUID PRIMARY KEY,
	user_id   UUID REFERENCES users(id) ON DELETE CASCADE,
	query     TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);

`)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.Exec(`
DROP TABLE saved_searches CASCADE;
DROP TABLE search_logs CASCADE;
DROP TABLE global_blacklist CASCADE;
DROP TABLE blacklist_entries CASCADE;
DROP TABLE interviews CASCADE;
DROP TABLE reviews CASCADE;
DROP TABLE job_applications CASCADE;
DROP TABLE job_vacancies CASCADE;
DROP TABLE resumes CASCADE;
DROP TABLE employer_profiles CASCADE;
DROP TABLE student_profiles CASCADE;
DROP TABLE users CASCADE;
`)
		return err
	})
}
