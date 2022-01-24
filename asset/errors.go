package asset

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyID     = errors.New("asset does not have ID")
	ErrUnknownType = errors.New("unknown type")
)

type NotFoundError struct {
	AssetID   string
	AssetType string
	AssetURN  string
}

func (e NotFoundError) Error() string {
	fields := []string{"no such record"}
	if e.AssetID != "" {
		fields = append(fields, fmt.Sprintf("with asset id \"%s\"", e.AssetID))
	}
	if e.AssetURN != "" {
		fields = append(fields, fmt.Sprintf("with asset urn \"%s\"", e.AssetURN))
	}
	if e.AssetType != "" {
		fields = append(fields, fmt.Sprintf("with asset type \"%s\"", e.AssetType))
	}
	return fmt.Sprintf("{%s}", strings.Join(fields, " "))
}

type InvalidError struct {
	AssetID   string
	AssetType string
	AssetURN  string
}

func (e InvalidError) Error() string {
	fields := []string{"empty asset field"}
	if e.AssetID != "" {
		fields = append(fields, fmt.Sprintf("with asset id \"%s\"", e.AssetID))
	}
	if e.AssetURN != "" {
		fields = append(fields, fmt.Sprintf("with asset urn \"%s\"", e.AssetURN))
	}
	if e.AssetType != "" {
		fields = append(fields, fmt.Sprintf("with asset type \"%s\"", e.AssetType))
	}
	return fmt.Sprintf("{%s}", strings.Join(fields, " "))
}
