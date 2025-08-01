package converter

import (
	"database/sql"

	"github.com/UraharaKiska/go-auth/internal/model"
	modelRepo "github.com/UraharaKiska/go-auth/internal/repository/user/model"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserSecureFromService(user *model.User) *desc.UserInfoSecure {
	var updateAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updateAt = timestamppb.New(user.UpdatedAt.Time)
	}
	roleValue := desc.EnumRole_value[user.Info.Role]
	return &desc.UserInfoSecure{
		Name:   user.Info.Name,
		Email:   user.Info.Email,
		Role:    desc.EnumRole(roleValue),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updateAt,
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:            info.Name,
		Email:           info.Email,
		Password:        info.Password,
		PasswordConfirm: info.PasswordConfirm,
		Role:            info.GetRole().String(),
	}
}

func ToUserInfoFromDescLoginRequest(info *desc.LoginRequest) *model.UserInfo {
	return &model.UserInfo{
		Email:           info.Email,
		Password:        info.Password,
	}
}


func ToUserUpdateInfoFromService(info *model.UserUpdateInfo) *modelRepo.UpdateUserInfo {
	updateUserInfo := &modelRepo.UpdateUserInfo{}
	if info.Name.Valid {
		updateUserInfo.Name = sql.NullString{String: info.Name.Value, Valid: true}
	} else {
		updateUserInfo.Name = sql.NullString{Valid: false}
	}
	if info.Email.Valid {
		updateUserInfo.Email = sql.NullString{String: info.Email.Value, Valid: true}
	} else {
		updateUserInfo.Email = sql.NullString{Valid: false}
	}
	return updateUserInfo
}

func ToUserUpdateInfoFromDesc(info *desc.UpdateUserInfo) *model.UserUpdateInfo {
	userUpdateInfo := &model.UserUpdateInfo{}
	if info.GetName() != nil {
		userUpdateInfo.Name = model.OptionString{Value: info.GetName().GetValue(), Valid: true}
	} else {
		userUpdateInfo.Name = model.OptionString{Valid: false}
	}
	if info.GetEmail() != nil {
		userUpdateInfo.Email = model.OptionString{Value: info.GetEmail().GetValue(), Valid: true}
	} else {
		userUpdateInfo.Email = model.OptionString{Valid: false}
	}
	return userUpdateInfo
}


func ToUserAuthFromDesc(user *desc.LoginRequest) *model.UserAuth {
	return &model.UserAuth{
		Email: user.GetEmail(),
		Password: user.GetPassword(),
	}
}