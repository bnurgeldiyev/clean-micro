package user

import (
	"clean-micro/internal/domain/user"
	"clean-micro/pkg/helpers"
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	sqlUserGetByUsername     = `SELECT id, create_ts, update_ts FROM tbl_user WHERE username=$1`
	sqlUserGetByID           = `SELECT username, create_ts, update_ts FROM tbl_user WHERE id=$1`
	sqlPasswordGetByUsername = `SELECT password FROM tbl_user WHERE username=$1`
	sqlPasswordChangeByID    = `UPDATE tbl_user SET password=$1, update_ts=$2 WHERE id=$3`
	sqlUsernameChangeByID    = `UPDATE tbl_user SET username=$1, update_ts=$2 WHERE id=$3`
	sqlUserCreate            = `INSERT INTO tbl_user(username, password) VALUES($1, $2)`
	sqlUserDeleteByID        = `DELETE FROM tbl_user WHERE id=$1`
)

func (us *userStorage) GetByUsername(ctx context.Context, username string) (item *user.User, err error) {

	clog := log.WithFields(log.Fields{
		"method":   "userStorage.GetByUsername",
		"username": username,
	})

	var usr user.User
	err1 := us.db.Get(&usr, sqlUserGetByUsername, username)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		eMsg := "error in us.db.Get"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
	}

	item = &user.User{
		ID:       usr.ID,
		Username: username,
		CreateTS: usr.CreateTS,
		UpdateTS: usr.UpdateTS,
	}

	return
}

func (us *userStorage) GetByID(ctx context.Context, id uuid.UUID) (item *user.User, err error) {

	clog := log.WithFields(log.Fields{
		"method": "userStorage.GetByID",
		"id":     id,
	})

	var usr user.User
	err1 := us.db.Get(&usr, sqlUserGetByID, id)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		eMsg := "error in us.db.Get"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
	}

	item = &user.User{
		ID:       id,
		Username: usr.Username,
		CreateTS: usr.CreateTS,
		UpdateTS: usr.UpdateTS,
	}

	return
}

func (us *userStorage) UpdatePasswordByID(ctx context.Context, pwdHash string, id uuid.UUID) (err error) {

	clog := log.WithFields(log.Fields{
		"method": "userStorage.ChangePasswordByID",
		"id":     id,
	})

	tx := us.db.MustBegin()
	tx.MustExec(sqlPasswordChangeByID, pwdHash, helpers.GetTimeNowString(), id)

	err = tx.Commit()
	if err != nil {
		eMsg := "error in tx.Commit()"
		clog.WithError(err).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	return
}

func (us *userStorage) UpdateUsernameByID(ctx context.Context, username string, id uuid.UUID) (err error) {

	clog := log.WithFields(log.Fields{
		"method": "userStorage.ChangeUsernameByID",
		"id":     id,
	})

	tx := us.db.MustBegin()
	tx.MustExec(sqlUsernameChangeByID, username, helpers.GetTimeNowString(), id)

	err = tx.Commit()
	if err != nil {
		eMsg := "error in tx.Commit()"
		clog.WithError(err).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	return
}

func (us *userStorage) GetPasswordByUsername(ctx context.Context, username string) (pwdHash string, err error) {

	clog := log.WithFields(log.Fields{
		"method":   "userStorage.GetPasswordByUsername",
		"username": username,
	})

	var usr user.User
	err1 := us.db.Get(&usr, sqlPasswordGetByUsername, username)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			pwdHash = ""
			err = nil
			return
		}

		eMsg := "error in us.db.Get"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
	}

	pwdHash = usr.Password

	return
}

func (us *userStorage) CreateUser(ctx context.Context, username, pwdHash string) (err error) {

	clog := log.WithFields(log.Fields{
		"method":   "userStorage.CreateUser",
		"username": username,
	})

	tx := us.db.MustBegin()
	tx.MustExec(sqlUserCreate, username, pwdHash)

	err = tx.Commit()
	if err != nil {
		eMsg := "error in tx.Commit()"
		clog.WithError(err).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	return
}

func (us *userStorage) DeleteByID(ctx context.Context, id uuid.UUID) (err error) {

	clog := log.WithFields(log.Fields{
		"method": "userStorage.DeleteByID",
		"id":     id,
	})

	tx := us.db.MustBegin()
	tx.MustExec(sqlUserDeleteByID, id)

	err = tx.Commit()
	if err != nil {
		eMsg := "error in tx.Commit()"
		clog.WithError(err).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	return
}
