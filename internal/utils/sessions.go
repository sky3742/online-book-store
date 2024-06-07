package utils

import (
	"online-book-store/internal/constant"
	"online-book-store/internal/model"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserSession struct {
	ID string
}

var sessions = map[string]UserSession{}

func CreateSession(userID string) string {
	token := uuid.NewV4().String()

	sessions[token] = UserSession{
		ID: userID,
	}

	time.AfterFunc(constant.SESSION_EXPIRY_DURATION*time.Second, func() {
		delete(sessions, token)
	})

	return token
}

func GetSession(token string) (UserSession, *model.AppError) {
	session, ok := sessions[token]
	if !ok {
		return UserSession{}, constant.ErrInvalidSession
	}
	return session, nil
}

func RemoveSession(token string) {
	delete(sessions, token)
}
