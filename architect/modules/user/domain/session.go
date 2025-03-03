package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	id           uuid.UUID
	userId       uuid.UUID
	refreshToken string
	accessExpAt  time.Time
	refreshExpAt time.Time
}

func NewSession(id uuid.UUID, userId uuid.UUID, refreshToken string, accessExpAt time.Time, refreshExpAt time.Time) *Session {
	return &Session{
		id,
		userId,
		refreshToken,
		accessExpAt,
		refreshExpAt,
	}
}

func (s Session) Id() uuid.UUID {
	return s.id
}

func (s Session) UserId() uuid.UUID {
	return s.userId
}

func (s Session) RefreshToken() string {
	return s.refreshToken
}

func (s Session) RefreshExpAt() time.Time {
	return s.refreshExpAt
}

func (s Session) AccessExpAt() time.Time {
	return s.accessExpAt
}
