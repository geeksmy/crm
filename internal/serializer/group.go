package serializer

import (
	"crm/gen/group"
	"crm/internal/model"
)

func ModeGroup2Group(groupModel *model.Group) *group.Group {
	res := group.Group{
		ID:   groupModel.ID,
		Name: groupModel.Name,
	}
	return &res
}

func ModeGroups2Groups(groupsModel []*model.Group) []*group.Group {
	res := make([]*group.Group, 0, len(groupsModel))
	for i := 0; i < len(groupsModel); i++ {
		res = append(res, ModeGroup2Group(groupsModel[i]))
	}
	return res
}
