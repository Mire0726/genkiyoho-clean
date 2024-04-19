package uuid

import (
	"github.com/google/uuid"
)

// UUIDGenerator defines an interface for generating UUIDs
type UUIDGenerator interface {
    Generate() string
}

// UUIDService is a service for generating UUIDs
type UUIDService struct{}

// NewUUIDService creates a new instance of UUIDService
func NewUUIDService() *UUIDService {
    return &UUIDService{}
}

// Generate generates a new UUID string
func (s *UUIDService) Generate() string {
    return uuid.NewString()
}