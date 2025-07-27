package converter

import (
	"database/sql"

	"github.com/UraharaKiska/go-auth/internal/model"
	modelRepo "github.com/UraharaKiska/go-auth/internal/repository/user/model"
	desc "github.com/UraharaKiska/go-auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updateAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updateAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(&user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updateAt,
	}
}

func ToUserInfoFromService(info *model.UserInfo) *desc.UserInfo {
	roleValue := desc.EnumRole_value[info.Role]
	return &desc.UserInfo{
		Name:            info.Name,
		Email:           info.Email,
		Password:        info.Password,
		PasswordConfirm: info.PasswordConfirm,
		Role:            desc.EnumRole(roleValue),
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
