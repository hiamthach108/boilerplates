package dtos

import "boilerplates/shared/enums"

type RoleDto struct {
	Id          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Status      enums.GeneralStatus `json:"status,omitempty"`
	Permissions []PermissionDto     `json:"permissions,omitempty"`
}

type RolePermissionDto struct {
	RoleId       string `json:"roleId"`
	PermissionId string `json:"permissionId"`
}
