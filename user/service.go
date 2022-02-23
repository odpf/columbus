package user

import (
	"context"
	"fmt"
)

type contextKeyType struct{}

var (
	// userContextKey is the key used for user.FromContext and
	// user.NewContext.
	userContextKey = contextKeyType(struct{}{})
)

// Service is a type of service that manages business process
type Service struct {
	repository Repository
	cfg        Config
}

// ValidateUser checks if user information is already in DB
// if exist in DB, return user ID, if not exist in DB, create a new one
func (s *Service) ValidateUser(ctx context.Context, email string) (string, error) {
	if email == "" {
		return "", ErrNoUserInformation
	}

	userID, err := s.repository.GetID(ctx, email)
	if err == nil {
		if userID != "" {
			return userID, nil
		}
		return "", fmt.Errorf("%w, fetched user id from DB is nil with email: %s", ErrNoUserInformation, email)
	}
	user := &User{
		Email:    email,
		Provider: s.cfg.IdentityProviderDefaultName,
	}
	if userID, err = s.repository.Create(ctx, user); err != nil {
		return "", err
	}
	return userID, nil
}

// NewContext returns a new context.Context that carries the provided
// user ID.
func NewContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userContextKey, userID)
}

// FromContext returns the user ID from the context if present, and empty
// otherwise.
func FromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	h, _ := ctx.Value(userContextKey).(string)
	if h != "" {
		return h
	}
	return h
}

// NewService initializes user service
func NewService(repository Repository, cfg Config) *Service {
	return &Service{
		repository: repository,
		cfg:        cfg,
	}
}
