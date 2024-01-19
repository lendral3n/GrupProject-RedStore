package handler

import (
	"MyEcommerce/features/user"
	"time"
)

type UserResponse struct {
	Name         string `json:"name" form:"name"`
	UserName     string `json:"user_name" form:"user_name"`
	Email        string `json:"email" form:"email"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

type AdminUserResponse struct {
	ID           uint      `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	UserName     string    `json:"user_name" form:"user_name"`
	Email        string    `json:"email" form:"email"`
	Role         string    `json:"role" form:"role"`
	PhotoProfile string    `json:"photo_profile" form:"photo_profile"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
}

func CoreToResponse(data *user.Core) UserResponse {
	var result = UserResponse{
		Name:         data.Name,
		UserName:     data.UserName,
		Email:        data.Email,
		PhotoProfile: data.PhotoProfile,
	}
	return result
}

func CoreToResponseList(data []user.Core) []AdminUserResponse {
	var results []AdminUserResponse
	for _, v := range data {
		var result = AdminUserResponse{
			ID:           v.ID,
			Name:         v.Name,
			UserName:     v.UserName,
			Email:        v.Email,
			Role:         v.Role,
			PhotoProfile: v.PhotoProfile,
			CreatedAt:    v.CreatedAt,
		}
		results = append(results, result)
	}
	return results
}
