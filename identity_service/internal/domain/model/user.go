package model

import "github.com/google/uuid"

type User struct {
	*Auditable //embedded
	userID uuid.UUID
	roleDefinition RoleDefinition 
}

func (u *User) GetUserID() uuid.UUID{
	return u.userID	
}

func (u *User) GetRole() Role{
	return u.roleDefinition.role
}

func (u *User) GetPermissions() []Permission {
	permissionSlice := make([]Permission, 0)

	for key, _ := range u.roleDefinition.permissions {
		permissionSlice = append(permissionSlice, key)
	}

	return permissionSlice
}

func NewUser(roleDef RoleDefinition) (*User, error) {
	auditable, err := NewAuditable()

	if err != nil {
		return nil, err
	}
	
	return &User{
		userID: uuid.New(),
		roleDefinition: roleDef,
		Auditable: auditable,
	}, nil
}

func NewDefaultCompanyUser() (*User, error) {
	newUser, err := NewUser(NewCompanyDefaultRoleDefinition())

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func NewDefaultSupervisorUser() (*User, error) {
	newUser, err := NewUser(NewSupervisorDefaultRoleDefinition())

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func NewDefaultWorkerUser() (*User, error) {
	newUser, err := NewUser(NewWorkerDefaultRoleDefinition())

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

