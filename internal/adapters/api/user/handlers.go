package user

import (
	"clean-micro/internal/adapters/api"
	"clean-micro/pkg/helpers"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (handler *userHandler) UserAuth(ctx context.Context, in *api.UserAuthRequest) (item *api.UserAuthResponse, err error) {

	clog := handleGrpc("UserAuth", ctx)

	username := in.GetUsername()
	password := in.GetPassword()

	m := make(map[string]string)
	m["username"] = username
	m["password"] = password

	names, err1 := helpers.VerifyMinLen(m)
	if err1 != nil {
		eMsg := fmt.Sprintf("FormValue's <%s> is empty", names)
		clog.Error(eMsg)
		err = status.Error(codes.Code(401), eMsg)
		return
	}

	res, err := handler.apiService.Auth(ctx, username, password)
	if err != nil {
		eMsg := "error in handler.apiService.Auth()"
		clog.WithError(err).Error(eMsg)
		return
	}

	item = &api.UserAuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}

	err = status.Error(codes.OK, "OK")
	return
}

func (handler *userHandler) UserCreate(ctx context.Context, in *api.UserCreateRequest) (item *api.UserCreateResponse, err error) {

	clog := handleGrpc("UserCreate", ctx)

	username := in.GetUsername()
	password := in.GetPassword()

	m := make(map[string]string)
	m["username"] = username
	m["password"] = password

	names, err1 := helpers.VerifyMinLen(m)
	if err1 != nil {
		eMsg := fmt.Sprintf("FormValue's <%s> is empty", names)
		clog.Error(eMsg)
		err = status.Error(codes.Code(400), eMsg)
		return
	}

	err = handler.apiService.Create(ctx, username, password)
	if err != nil {
		eMsg := "error in handler.apiService.Create()"
		clog.WithError(err).Error(eMsg)
		return
	}

	err = status.Error(codes.OK, "OK")
	item = &api.UserCreateResponse{
		Success: true,
	}

	return
}

func (handler *userHandler) UserAccess(ctx context.Context, in *api.UserAccessRequest) (item *api.UserAccessResponse, err error) {

	clog := handleGrpc("UserAccess", ctx)

	accessToken := in.GetAccessToken()
	if len(accessToken) == 0 {
		eMsg := fmt.Sprintf("FormValue's <access_token> is empty")
		clog.Error(eMsg)
		err = status.Error(codes.Code(400), eMsg)
		return
	}

	username, err := handler.apiService.Access(ctx, accessToken)
	if err != nil {
		eMsg := "error in handler.apiService.Access()"
		clog.WithError(err).Error(eMsg)
		return
	}

	err = status.Error(codes.OK, "OK")
	item = &api.UserAccessResponse{
		Username: username,
	}

	return
}

func (handler *userHandler) UserUpdatePassword(ctx context.Context, in *api.UserUpdatePasswordRequest) (item *api.UserUpdatePasswordResponse, err error) {

	clog := handleGrpc("UserUpdatePassword", ctx)

	username := in.GetUsername()
	oldPassword := in.GetOldPassword()
	newPassword := in.GetNewPassword()

	m := make(map[string]string)
	m["username"] = username
	m["old_password"] = oldPassword
	m["new_password"] = newPassword

	names, err1 := helpers.VerifyMinLen(m)
	if err1 != nil {
		eMsg := fmt.Sprintf("FormValue's <%s> is empty", names)
		clog.Error(eMsg)
		err = status.Error(codes.Code(400), eMsg)
		return
	}

	err = handler.apiService.UpdatePassword(ctx, username, oldPassword, newPassword)
	if err != nil {
		eMsg := "error in handler.apiService.UpdatePassword"
		clog.WithError(err).Error(eMsg)
		return
	}

	item = &api.UserUpdatePasswordResponse{
		Success: true,
	}

	return
}

func (handler *userHandler) UserUpdateUsername(ctx context.Context, in *api.UserUpdateUsernameRequest) (item *api.UserUpdateUsernameResponse, err error) {

	clog := handleGrpc("UserUpdateUsername", ctx)

	oldUsername := in.GetOldUsername()
	newUsername := in.GetNewUsername()
	password := in.GetPassword()

	m := make(map[string]string)
	m["oldUsername"] = oldUsername
	m["newUsername"] = newUsername
	m["password"] = password

	names, err1 := helpers.VerifyMinLen(m)
	if err1 != nil {
		eMsg := fmt.Sprintf("FormValue's <%s> is empty", names)
		clog.Error(eMsg)
		err = status.Error(codes.Code(400), eMsg)
		return
	}

	err = handler.apiService.UpdateUsername(ctx, password, oldUsername, newUsername)
	if err != nil {
		eMsg := "error in handler.apiService.UpdateUsername"
		clog.WithError(err).Error(eMsg)
		return
	}

	item = &api.UserUpdateUsernameResponse{
		Success: true,
	}

	return
}

func (handler *userHandler) UserDelete(ctx context.Context, in *api.UserDeleteRequest) (item *api.UserDeleteResponse, err error) {

	clog := handleGrpc("UserDelete", ctx)

	username := in.GetUsername()

	if len(username) == 0 {
		eMsg := fmt.Sprintf("FormValue's <username> is empty")
		clog.Error(eMsg)
		err = status.Error(codes.Code(400), eMsg)
		return
	}

	err = handler.apiService.Delete(ctx, username)
	if err != nil {
		eMsg := "error in handler.apiService.Delete"
		clog.WithError(err).Error(eMsg)
		return
	}

	item = &api.UserDeleteResponse{
		Success: true,
	}

	return
}
