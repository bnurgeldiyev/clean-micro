package user

import (
	"clean-micro/internal/adapters/api/user"
	"clean-micro/pkg/helpers"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) user.Service {
	return &service{storage: storage}
}

func (s *service) Auth(ctx context.Context, username, password string) (item *user.Auth, err error) {

	clog := log.WithContext(ctx).WithFields(log.Fields{
		"method":   "service.Auth",
		"username": username,
	})

	usr, err1 := s.storage.GetByUsername(ctx, username)
	if err1 != nil {
		eMsg := "error in s.GetByUsername()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	if usr == nil {
		eMsg := fmt.Sprintf("User with username=%v not found", username)
		clog.Error(eMsg)
		err = status.Error(codes.Code(401), eMsg)
		return
	}

	pwdHash, err1 := s.storage.GetPasswordByUsername(ctx, username)
	if err1 != nil {
		eMsg := "error in s.storage.GetPasswordByUsername()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	err1 = bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(password))
	if err1 != nil {
		eMsg := "error in bcrypt.CompareHashAndPassword()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(401), eMsg)
		return
	}

	accessToken, err1 := helpers.GenerateAccessToken(usr.ID)
	if err1 != nil {
		eMsg := "error in bcrypt.GenerateAccessToken()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	refreshToken, err1 := helpers.GenerateRefreshToken()
	if err1 != nil {
		eMsg := "error in bcrypt.GenerateRefreshToken()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	item = &user.Auth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.String(),
	}

	return
}

func (s *service) Access(ctx context.Context, accessToken string) (username string, err error) {

	clog := log.WithContext(ctx).WithFields(log.Fields{
		"method":   "service.Access",
		"username": username,
	})

	id, err1 := helpers.VerifyAccessToken(accessToken)
	if err1 != nil {
		eMsg := "error in helpers.VerifyAccessToken"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(401), eMsg)
		return
	}

	usr, err1 := s.storage.GetByID(ctx, id)
	if err1 != nil {
		eMsg := "error in api.access.UserGetByUsername"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	if usr == nil {
		eMsg := fmt.Sprintf("User with id=<%s> not found", id)
		clog.Error(eMsg)
		err = status.Error(codes.Code(401), eMsg)
		return
	}

	username = usr.Username

	return
}

func (s *service) Create(ctx context.Context, username, password string) (err error) {

	clog := log.WithContext(ctx).WithFields(log.Fields{
		"method":   "service.Create",
		"username": username,
	})

	usr, err1 := s.storage.GetByUsername(ctx, username)
	if err1 != nil {
		eMsg := "error in s.storage.GetByUsername"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	if usr != nil {
		eMsg := fmt.Sprintf("User with username=%v already exists", username)
		clog.Error(eMsg)
		err = status.Error(codes.Code(409), eMsg)
		return
	}

	pwdHashBytes, err1 := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err1 != nil {
		eMsg := "error in bcrypt.GenerateFromPassword"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	pwdHash := string(pwdHashBytes)

	err = s.storage.CreateUser(ctx, username, pwdHash)

	return
}

func (s *service) Delete(ctx context.Context, username string) (err error) {

	clog := log.WithContext(ctx).WithFields(log.Fields{
		"method":   "service.Delete",
		"username": username,
	})

	usr, err1 := s.storage.GetByUsername(ctx, username)
	if err1 != nil {
		eMsg := "error in s.storage.GetByUsername"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	if usr != nil {
		eMsg := fmt.Sprintf("User with username=%v already exists", username)
		clog.Error(eMsg)
		err = status.Error(codes.Code(409), eMsg)
		return
	}

	err = s.storage.DeleteByID(ctx, usr.ID)

	return
}

func (s *service) UpdateUsername(ctx context.Context, password, oldUsername, newUsername string) (err error) {

	clog := log.WithContext(ctx).WithFields(log.Fields{
		"method":       "service.UpdateUsername",
		"old_username": oldUsername,
	})

	usr, err1 := s.storage.GetByUsername(ctx, oldUsername)
	if err1 != nil {
		eMsg := "error in s.storage.GetByUsername"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	if usr == nil {
		eMsg := fmt.Sprintf("User with username=%v not found", oldUsername)
		clog.Error(eMsg)
		err = status.Error(codes.Code(404), eMsg)
		return
	}

	pwdHash, err1 := s.storage.GetPasswordByUsername(ctx, oldUsername)
	if err1 != nil {
		eMsg := "error in s.storage.GetPasswordByUsername()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	err1 = bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(password))
	if err1 != nil {
		eMsg := "error in bcrypt.CompareHashAndPassword()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(403), eMsg)
		return
	}

	err = s.storage.UpdateUsernameByID(ctx, newUsername, usr.ID)

	return
}

func (s *service) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) (err error) {

	clog := log.WithContext(ctx).WithFields(log.Fields{
		"method":   "service.UpdatePassword",
		"username": username,
	})

	usr, err1 := s.storage.GetByUsername(ctx, username)
	if err1 != nil {
		eMsg := "error in s.storage.GetByUsername"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	if usr == nil {
		eMsg := fmt.Sprintf("User with username=%v not found", username)
		clog.Error(eMsg)
		err = status.Error(codes.Code(404), eMsg)
		return
	}

	oldPwdHash, err1 := s.storage.GetPasswordByUsername(ctx, username)
	if err1 != nil {
		eMsg := "error in s.storage.GetPasswordByUsername()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	err1 = bcrypt.CompareHashAndPassword([]byte(oldPwdHash), []byte(oldPassword))
	if err1 != nil {
		eMsg := "error in bcrypt.CompareHashAndPassword()"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(401), eMsg)
		return
	}

	pwdHashBytes, err1 := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err1 != nil {
		eMsg := "error in bcrypt.GenerateFromPassword"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	pwdHash := string(pwdHashBytes)

	err = s.storage.UpdatePasswordByID(ctx, pwdHash, usr.ID)

	return
}
