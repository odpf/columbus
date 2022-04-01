package user

//go:generate mockery --name Repository --outpkg mocks --output ../lib/mocks/ --with-expecter --structname UserRepository --filename user_repository.go
import (
	"context"
	"time"

	compassv1beta1 "github.com/odpf/columbus/api/proto/odpf/compass/v1beta1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// User is a basic entity of a user
type User struct {
	ID        string    `json:"id,omitempty" diff:"-" db:"id"`
	Email     string    `json:"email" diff:"email" db:"email"`
	Provider  string    `json:"provider" diff:"-" db:"provider"`
	CreatedAt time.Time `json:"-" diff:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" diff:"-" db:"updated_at"`
}

// ToProto transforms struct with some fields only to proto
func (d User) ToProto() *compassv1beta1.User {
	if d.ID == "" {
		return nil
	}

	return &compassv1beta1.User{
		Id:    d.ID,
		Email: d.Email,
	}
}

// ToFullProto transforms struct with all fields to proto
func (d User) ToFullProto() *compassv1beta1.User {
	if d.ID == "" {
		return nil
	}

	var createdAtPB *timestamppb.Timestamp
	if !d.CreatedAt.IsZero() {
		createdAtPB = timestamppb.New(d.CreatedAt)
	}

	var updatedAtPB *timestamppb.Timestamp
	if !d.UpdatedAt.IsZero() {
		updatedAtPB = timestamppb.New(d.UpdatedAt)
	}

	return &compassv1beta1.User{
		Id:        d.ID,
		Email:     d.Email,
		Provider:  d.Provider,
		CreatedAt: createdAtPB,
		UpdatedAt: updatedAtPB,
	}
}

// NewFromProto transforms proto to struct
func NewFromProto(proto *compassv1beta1.User) User {
	var createdAt time.Time
	if proto.GetCreatedAt() != nil {
		createdAt = proto.GetCreatedAt().AsTime()
	}

	var updatedAt time.Time
	if proto.GetUpdatedAt() != nil {
		updatedAt = proto.GetUpdatedAt().AsTime()
	}

	return User{
		ID:        proto.GetId(),
		Email:     proto.GetEmail(),
		Provider:  proto.GetProvider(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// Validate validates a user is valid or not
func (u *User) Validate() error {
	if u == nil {
		return ErrNoUserInformation
	}

	if u.Email == "" || u.Provider == "" {
		return InvalidError{Email: u.Email, Provider: u.Provider}
	}

	return nil
}

// Repository contains interface of supported methods
type Repository interface {
	Create(ctx context.Context, u *User) (string, error)
	GetID(ctx context.Context, email string) (string, error)
}
