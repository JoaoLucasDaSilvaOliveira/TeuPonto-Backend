package model

import "github.com/google/uuid"

type User struct {
	Auditable //embedded
	userID uuid.UUID
	roleDefinition RoleDefinition 
}

func (u *User) getUserID() uuid.UUID{
	return u.userID	
}

func (u *User) getRole() Role{
	return u.roleDefinition.role
}

func (u *User) getPermissions() []Permission {
	permissionSlice := make([]Permission, 0)

	for key, _ := range u.roleDefinition.permissions {
		permissionSlice = append(permissionSlice, key)
	}

	return permissionSlice
}

