package dto

import "samsamoohooh-mini-api/internal/core/domain"

type CreateUserResponse struct {
	ID         uint              `json:"id"`
	Name       string            `json:"name"`
	Resolution string            `json:"resolution"`
	Role       domain.RoleType   `json:"role"`
	Sub        string            `json:"sub"`
	Social     domain.SocialType `json:"social"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	DeletedAt  string            `jsno:"deletedAt"`
}

func NewCreateUserResponse(user domain.User) (res *CreateUserResponse) {
	res = &CreateUserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       user.Role,
		Sub:        user.Sub,
		Social:     user.Social,
		CreatedAt:  user.CreatedAt.Format(CreatedAtFormat),
		UpdatedAt:  user.UpdatedAt.Format(UpdatedAtFormat),
	}

	// if deletedAt is not null
	if user.DeletedAt.Valid {
		res.DeletedAt = user.DeletedAt.Time.Format(DeletedAtFormat)
	}

	return res
}

type GetUserByIDResponse struct {
	ID         uint              `json:"id"`
	Name       string            `json:"name"`
	Resolution string            `json:"resolution"`
	Role       domain.RoleType   `json:"role"`
	Sub        string            `json:"sub"`
	Social     domain.SocialType `json:"social"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	DeletedAt  string            `json:"deletedAt"`
}

func NewGetUserByIDResponse(user domain.User) (res *GetUserByIDResponse) {
	res = &GetUserByIDResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       user.Role,
		Sub:        user.Sub,
		Social:     user.Social,
		CreatedAt:  user.CreatedAt.Format(CreatedAtFormat),
		UpdatedAt:  user.UpdatedAt.Format(UpdatedAtFormat),
	}

	if user.DeletedAt.Valid {
		res.DeletedAt = user.DeletedAt.Time.Format(DeletedAtFormat)
	}

	return res
}
