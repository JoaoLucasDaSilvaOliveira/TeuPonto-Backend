package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	*Auditable     //embedded
	userID         uuid.UUID
	roleDefinition *RoleDefinition
}

func (u *User) GetUserID() uuid.UUID {
	return u.userID
}

func (u *User) GetRoleDefinition() *RoleDefinition {
	return u.roleDefinition
}

func (u *User) SetRoleDefinition(roleDef *RoleDefinition, setUpdatedAt bool) {
	if setUpdatedAt {
		now := time.Now()
		u.SetUpdatedAt(&now)
	}
	
	u.roleDefinition = roleDef
}

func NewUser(roleDef *RoleDefinition) (*User, error) {
	auditable, err := NewAuditable()

	if err != nil {
		return nil, err
	}

	user := new(User)
	user.Auditable = auditable
	user.userID = uuid.New()
	user.SetRoleDefinition(roleDef, false)
	

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