package discussion

//go:generate mockery --name Repository --outpkg mocks --output ../lib/mocks/ --with-expecter --structname DiscussionRepository --filename discussion_repository.go

import (
	"context"
	"fmt"
	"strings"
	"time"

	compassv1beta1 "github.com/odpf/columbus/api/proto/odpf/compass/v1beta1"
	"github.com/odpf/columbus/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const MAX_ARRAY_FIELD_NUM = 10

type Repository interface {
	GetAll(ctx context.Context, filter Filter) ([]Discussion, error)
	Create(ctx context.Context, discussion *Discussion) (string, error)
	Get(ctx context.Context, did string) (Discussion, error)
	Patch(ctx context.Context, discussion *Discussion) error
	GetAllComments(ctx context.Context, discussionID string, filter Filter) ([]Comment, error)
	CreateComment(ctx context.Context, cmt *Comment) (string, error)
	GetComment(ctx context.Context, commentID string, discussionID string) (Comment, error)
	UpdateComment(ctx context.Context, cmt *Comment) error
	DeleteComment(ctx context.Context, commentID string, discussionID string) error
}

type Discussion struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Type      Type      `json:"type"`
	State     State     `json:"state"`
	Labels    []string  `json:"labels"`
	Assets    []string  `json:"assets"`
	Assignees []string  `json:"assignees"`
	Owner     user.User `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToProto transforms struct to proto
func (d Discussion) ToProto() *compassv1beta1.Discussion {
	return &compassv1beta1.Discussion{
		Id:        d.ID,
		Title:     d.Title,
		Body:      d.Body,
		Type:      d.Type.String(),
		State:     d.State.String(),
		Labels:    d.Labels,
		Assets:    d.Assets,
		Assignees: d.Assignees,
		Owner:     d.Owner.ToProto(),
		CreatedAt: timestamppb.New(d.CreatedAt),
		UpdatedAt: timestamppb.New(d.UpdatedAt),
	}
}

// NewFromProto transforms proto to struct
func NewFromProto(proto *compassv1beta1.Discussion) Discussion {
	return Discussion{
		ID:        proto.Id,
		Title:     proto.Title,
		Body:      proto.Body,
		Type:      GetTypeEnum(proto.Type),
		State:     GetStateEnum(proto.State),
		Labels:    proto.Labels,
		Assets:    proto.Assets,
		Assignees: proto.Assignees,
		Owner:     user.NewFromProto(proto.Owner),
		CreatedAt: proto.CreatedAt.AsTime(),
		UpdatedAt: proto.UpdatedAt.AsTime(),
	}
}

// IsEmpty returns true if all fields inside discussion are considered empty
func (d Discussion) IsEmpty() bool {
	if len(strings.TrimSpace(d.Title)) > 0 {
		return false
	}

	if len(strings.TrimSpace(d.Body)) > 0 {
		return false
	}

	if len(strings.TrimSpace(d.Type.String())) > 0 {
		return false
	}

	if len(strings.TrimSpace(d.State.String())) > 0 {
		return false
	}

	if d.Labels != nil {
		return false
	}

	if d.Assets != nil {
		return false
	}

	if d.Assignees != nil {
		return false
	}

	return true
}

// Validate checks emptyness required fields and constraint in discussion and return error if the required is empty
func (d Discussion) Validate() error {
	if len(strings.TrimSpace(d.Title)) == 0 {
		return fmt.Errorf("title cannot be empty")
	}

	if len(strings.TrimSpace(d.Body)) == 0 {
		return fmt.Errorf("body cannot be empty")
	}

	if len(strings.TrimSpace(d.Type.String())) == 0 {
		return fmt.Errorf("type must be specified")
	}

	return d.ValidateConstraint()
}

// ValidateConstraint checks whether non empty/nil fields fulfill the contract
func (d Discussion) ValidateConstraint() error {
	if len(strings.TrimSpace(d.Type.String())) > 0 && !IsTypeStringValid(d.Type.String()) {
		return ErrInvalidType
	}

	if len(strings.TrimSpace(d.State.String())) > 0 && !IsStateStringValid(d.State.String()) {
		return ErrInvalidState
	}

	if d.Assignees != nil && len(d.Assignees) > MAX_ARRAY_FIELD_NUM {
		return fmt.Errorf("assignees cannot be more than %d", MAX_ARRAY_FIELD_NUM)
	}

	if d.Assets != nil && len(d.Assets) > MAX_ARRAY_FIELD_NUM {
		return fmt.Errorf("assets cannot be more than %d", MAX_ARRAY_FIELD_NUM)
	}

	if d.Labels != nil && len(d.Labels) > MAX_ARRAY_FIELD_NUM {
		return fmt.Errorf("labels cannot be more than %d", MAX_ARRAY_FIELD_NUM)
	}
	return nil
}
