package model

import (
	"github.com/google/uuid"
)

type User struct {
	userID         uuid.UUID
	roleDefinition *RoleDefinition
}

func (u *User) GetUserID() uuid.UUID {
	return u.userID
}

func (u *User) GetRoleDefinition() *RoleDefinition {
	return u.roleDefinition
}

func (u *User) SetRoleDefinition(roleDef *RoleDefinition) {
	u.roleDefinition = roleDef
}

func NewUser(roleDef *RoleDefinition) *User {
	user := new(User)
	user.userID = uuid.New()
	user.SetRoleDefinition(roleDef)
	
	return user
}
	
func NewDefaultCompanyUser() *User {
	newUser := NewUser(NewCompanyDefaultRoleDefinition())
	return newUser
}

func NewDefaultSupervisorUser() *User {
	newUser := NewUser(NewSupervisorDefaultRoleDefinition())
	return newUser
}

func NewDefaultWorkerUser() *User {
	newUser := NewUser(NewWorkerDefaultRoleDefinition())
	return newUser
}