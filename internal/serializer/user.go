package serializer

import (
	"crm/gen/auth"
	"crm/gen/user"
	"crm/internal/model"
)

func ModelUser2User(modelUser *model.User) *user.User {
	superior := user.Superior{
		ID:   modelUser.Superior.ID,
		Name: modelUser.Superior.Name,
	}
	group := user.Group{
		ID:   modelUser.Group.ID,
		Name: modelUser.Group.Name,
	}
	res := user.User{
		ID:       modelUser.ID,
		Username: modelUser.Username,
		Name:     modelUser.Name,
		Mobile:   modelUser.Mobile,
		Email:    modelUser.Email,
		Jobs:     modelUser.Jobs,
		IsAdmin:  modelUser.IsAdmin,
		Superior: &superior,
		Group:    &group,
	}

	return &res
}

func ModelUser2AuthUser(modelUser *model.User) *auth.User {
	superior := auth.Superior{
		ID:   modelUser.Superior.ID,
		Name: modelUser.Superior.Name,
	}
	group := auth.Group{
		ID:   modelUser.Group.ID,
		Name: modelUser.Group.Name,
	}
	res := auth.User{
		ID:       modelUser.ID,
		Username: modelUser.Username,
		Name:     modelUser.Name,
		Mobile:   modelUser.Mobile,
		Email:    modelUser.Email,
		Jobs:     modelUser.Jobs,
		IsAdmin:  modelUser.IsAdmin,
		Superior: &superior,
		Group:    &group,
	}
	return &res
}

func ModelUsers2Users(modelUsers []*model.User) []*user.User {
	res := make([]*user.User, 0, len(modelUsers))
	for i := 0; i < len(modelUsers); i++ {
		res = append(res, ModelUser2User(modelUsers[i]))
	}

	return res
}
