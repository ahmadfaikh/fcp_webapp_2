package repository

import (
	"a21hc3NpZ25tZW50/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	result := u.db.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) DeleteSession(token string) error {
	result := u.db.Where("token=?", token).Delete(&model.Session{})
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	result := u.db.Model(&model.Session{}).Select("token", "expiry").Where("email=?", session.Email).UpdateColumns(map[string]interface{}{
		"token":  session.Token,
		"expiry": session.Expiry,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	usr := model.Session{}
	result := u.db.First(&usr).Where("email=?", email)
	if result.Error != nil {
		return usr, result.Error
	}

	return usr, nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	usr := model.Session{}
	result := u.db.First(&usr).Where("token=?", token)
	if result.Error != nil {
		return usr, result.Error
	}
	return usr, nil // TODO: replace this
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
