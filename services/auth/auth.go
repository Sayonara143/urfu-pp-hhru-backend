package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	SecretKey             = []byte("your-secret-key")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidToken       = errors.New("invalid token")
)

type Service struct {
	store Storage
}

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

type Storage interface {
	UserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	UserByEmail(ctx context.Context, email string) (*models.User, error)
	UserInsert(ctx context.Context, user *models.User) error
	UserUpdate(ctx context.Context, user *models.User) error

	CtxWithTx(ctx context.Context) (context.Context, error)
	TxCommit(ctx context.Context) error
	TxRollback(ctx context.Context) error
}

// New создает новый экземпляр сервиса аутентификации
func New(store Storage) *Service {
	return &Service{store: store}
}

// RegisterUser регистрирует нового пользователя с использованием соли и хэширования
func (s *Service) RegisterUser(ctx context.Context, user *models.User) error {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	// Генерация соли
	salt, err := generateSalt(16)
	if err != nil {
		return err
	}

	// Хэширование пароля
	hashedPassword, err := hashPassword(user.Password, salt)
	if err != nil {
		return err
	}

	user.PasswordHash = hashedPassword
	user.PasswordSalt = salt

	if err = s.store.UserInsert(ctx, user); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// LoginUser выполняет аутентификацию и возвращает JWT-токен
func (s *Service) LoginUser(ctx context.Context, email, password string) (string, error) {
	ctx, err := s.store.CtxWithTx(ctx)
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	user, err := s.store.UserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrUserNotFound
	}

	// Проверка пароля
	if !comparePassword(password, user.PasswordSalt, user.PasswordHash) {
		return "", ErrInvalidCredentials
	}

	token, err := s.generateToken(user)
	if err != nil {
		return "", err
	}

	return token, s.store.TxCommit(ctx)
}

// hashPassword хэширует пароль с солью
func hashPassword(password, salt string) (string, error) {
	saltedPassword := password + salt
	hash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// comparePassword проверяет введённый пароль с хранимым хэшем
func comparePassword(password, salt, hash string) bool {
	saltedPassword := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedPassword))
	return err == nil
}

// generateSalt генерирует случайную соль
func generateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// generateToken создает JWT-токен с использованием jwt/v5
func (s *Service) generateToken(user *models.User) (string, error) {
	claims := &Claims{
		UserID: *user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

// VerifyToken проверяет JWT-токен и возвращает данные пользователя
func (s *Service) VerifyToken(ctx context.Context, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
