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

func NewUser(roleDef *RoleDefinition) (*User, error) {

	user := new(User)
	user.userID = uuid.New()
	user.SetRoleDefinition(roleDef)
	

	return user, nil
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