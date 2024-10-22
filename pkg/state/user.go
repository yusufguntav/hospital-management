package state

import (
	"context"

	"github.com/google/uuid"
	"github.com/yusufguntav/hospital-management/pkg/entities"
)

const (
	CurrentUserId         = "CurrentUserId"
	CurrentUserIP         = "CurrentIP"
	CurrentUserROLE       = "CurrentUserRole"
	CurrentUserHospitalID = "CurrentUserHospitalId"
)

func CurrentUserRole(ctx context.Context) entities.AuthRole {
	value := ctx.Value(CurrentUserROLE)
	if value == nil {
		return entities.Staff
	}
	return value.(entities.AuthRole)
}

func CurrentUserHospitalId(ctx context.Context) string {
	value := ctx.Value(CurrentUserHospitalID)
	if value == nil {
		return ""
	}
	return value.(string)
}

func CurrentUser(ctx context.Context) uuid.UUID {
	value := ctx.Value(CurrentUserId)
	if value == nil {
		return uuid.Nil
	}
	return uuid.MustParse(value.(string))
}

func SetCurrentUser(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, CurrentUserId, userID.String())
}
