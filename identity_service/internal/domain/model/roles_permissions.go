package model

import (
	"fmt"

	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type Role string
type Permission string

// sistem roles
const (
	MANAGER_ROLE    Role = "MANAGER"
	SUPERVISOR_ROLE Role = "SUPERVISOR"
	WORKER_ROLE     Role = "WORKER"
	OUTSOURCED_ROLE Role = "OUTSOURCED"
)

//general permissions
const (
	// admin level
	ADJUST_REQUESTED_TIME_CLOCK         Permission = "ADJUST_REQUESTED_TIME_CLOCK"
	ADJUST_MANUAL_TIME_CLOCK            Permission = "ADJUST_MANUAL_TIME_CLOCK"
	CREATE_EMPLOYEE_ENTRIES             Permission = "CREATE_EMPLOYEE_ENTRIES"
	DELETE_EMPLOYEE_ENTRIES             Permission = "DELETE_EMPLOYEE_ENTRIES"
	UPDATE_EMPLOYEE_ENTRIES             Permission = "UPDATE_EMPLOYEE_ENTRIES"
	VIEW_EMPLOYEE_ENTRIES               Permission = "VIEW_EMPLOYEE_ENTRIES"
	VIEW_GENERAL_DASHBOARD              Permission = "VIEW_GENERAL_DASHBOARD"
	VIEW_GENERAL_REPORT                 Permission = "VIEW_GENERAL_REPORT"
	GENERATE_GENERAL_PDF_REPORT         Permission = "GENERATE_GENERAL_PDF_REPORT"
	CREATE_EMPLOYEE                     Permission = "CREATE_EMPLOYEE"
	DELETE_EMPLOYEE                     Permission = "DELETE_EMPLOYEE"
	UPDATE_EMPLOYEE                     Permission = "UPDATE_EMPLOYEE"
	VIEW_EMPLOYEE_REGISTRATION          Permission = "VIEW_EMPLOYEE_REGISTRATION"
	CREATE_WORK_JOURNEY                 Permission = "CREATE_WORK_JOURNEY"
	DELETE_WORK_JOURNEY                 Permission = "DELETE_WORK_JOURNEY"
	UPDATE_WORK_JOURNEY                 Permission = "UPDATE_WORK_JOURNEY"
	VIEW_WORK_JOURNEY_REGISTRATION      Permission = "VIEW_WORK_JOURNEY_REGISTRATION"
	CREATE_SCHEDULE                     Permission = "CREATE_SCHEDULE"
	DELETE_SCHEDULE                     Permission = "DELETE_SCHEDULE"
	UPDATE_SCHEDULE                     Permission = "UPDATE_SCHEDULE"
	VIEW_SCHEDULE_REGISTRATION          Permission = "VIEW_SCHEDULE_REGISTRATION"
	CREATE_HOLIDAY                      Permission = "CREATE_HOLIDAY"
	DELETE_HOLIDAY                      Permission = "DELETE_HOLIDAY"
	UPDATE_HOLIDAY                      Permission = "UPDATE_HOLIDAY"
	VIEW_HOLIDAY_REGISTRATION           Permission = "VIEW_HOLIDAY_REGISTRATION"

	// supervisor level
	ADJUST_SUBORDINATE_REQUESTED_TIME_CLOCK Permission = "ADJUST_SUBORDINATE_REQUESTED_TIME_CLOCK"
	ADJUST_SUBORDINATE_MANUAL_TIME_CLOCK    Permission = "ADJUST_SUBORDINATE_MANUAL_TIME_CLOCK"
	CREATE_SUBORDINATE_ENTRIES              Permission = "CREATE_SUBORDINATE_ENTRIES"
	DELETE_SUBORDINATE_ENTRIES              Permission = "DELETE_SUBORDINATE_ENTRIES"
	UPDATE_SUBORDINATE_ENTRIES              Permission = "UPDATE_SUBORDINATE_ENTRIES"
	VIEW_SUBORDINATE_ENTRIES                Permission = "VIEW_SUBORDINATE_ENTRIES"
	VIEW_SUBORDINATE_DASHBOARD              Permission = "VIEW_SUBORDINATE_DASHBOARD"
	VIEW_SUBORDINATE_REPORT                 Permission = "VIEW_SUBORDINATE_REPORT"
	GENERATE_SUBORDINATE_PDF_REPORT         Permission = "GENERATE_SUBORDINATE_PDF_REPORT"

	// operator level
	CLOCK_IN_OUT                Permission = "CLOCK_IN_OUT"
	VIEW_DASHBOARD              Permission = "VIEW_DASHBOARD"
	VIEW_REPORT                 Permission = "VIEW_REPORT"
	VIEW_HOLIDAYS               Permission = "VIEW_HOLIDAYS"
	GENERATE_PDF_REPORT         Permission = "GENERATE_PDF_REPORT"
	REQUEST_TIME_CLOCK_ADJUSTMENT Permission = "REQUEST_TIME_CLOCK_ADJUSTMENT"
	REQUEST_ENTRY               Permission = "REQUEST_ENTRY"
)

// permissions descriptions
var PermissionDescriptions = map[Permission]string{
	ADJUST_REQUESTED_TIME_CLOCK:         "Permissão para ajustar ponto solicitado",
	ADJUST_MANUAL_TIME_CLOCK:            "Permissão para ajuste manual de ponto",
	CREATE_EMPLOYEE_ENTRIES:             "Permissão para criar lançamentos dos empregados",
	DELETE_EMPLOYEE_ENTRIES:             "Permissão para deletar lançamentos dos empregados",
	UPDATE_EMPLOYEE_ENTRIES:             "Permissão para atualizar lançamentos dos empregados",
	VIEW_EMPLOYEE_ENTRIES:               "Permissão para ver lançamentos dos empregados",
	VIEW_GENERAL_DASHBOARD:              "Permissão para ver dashboard geral da empresa",
	VIEW_GENERAL_REPORT:                 "Permissão para ver relatório geral da empresa",
	GENERATE_GENERAL_PDF_REPORT:         "Permissão para gerar relatório geral da empresa em PDF",
	CREATE_EMPLOYEE:                     "Permissão para cadastrar de funcionários",
	DELETE_EMPLOYEE:                     "Permissão para deletar cadastros de funcionários",
	UPDATE_EMPLOYEE:                     "Permissão para atualizar cadastros de funcionários",
	VIEW_EMPLOYEE_REGISTRATION:          "Permissão para ver cadastros de funcionários",
	CREATE_WORK_JOURNEY:                 "Permissão para cadastrar jornadas",
	DELETE_WORK_JOURNEY:                 "Permissão para deletar cadastros de jornadas",
	UPDATE_WORK_JOURNEY:                 "Permissão para atualizar cadastros de jornadas",
	VIEW_WORK_JOURNEY_REGISTRATION:      "Permissão para ver cadastros de jornadas",
	CREATE_SCHEDULE:                     "Permissão para cadastrar horários",
	DELETE_SCHEDULE:                     "Permissão para deletar cadastros de horários",
	UPDATE_SCHEDULE:                     "Permissão para atualizar cadastros de horários",
	VIEW_SCHEDULE_REGISTRATION:          "Permissão para ver cadastros de horários",
	CREATE_HOLIDAY:                      "Permissão para cadastrar feriados",
	DELETE_HOLIDAY:                      "Permissão para deletar cadastros de feriados",
	UPDATE_HOLIDAY:                      "Permissão para atualizar cadastros de feriados",
	VIEW_HOLIDAY_REGISTRATION:           "Permissão para ver cadastros de feriados",

	ADJUST_SUBORDINATE_REQUESTED_TIME_CLOCK: "Permissão para ajustar ponto solicitado por subordinados",
	ADJUST_SUBORDINATE_MANUAL_TIME_CLOCK:    "Permissão para ajustar manualmente ponto de subordinados",
	CREATE_SUBORDINATE_ENTRIES:              "Permissão para cadastrar lançamentos de funcionários subordinados",
	DELETE_SUBORDINATE_ENTRIES:              "Permissão para deletar lançamentos de funcionários subordinados",
	UPDATE_SUBORDINATE_ENTRIES:              "Permissão para atualizar lançamentos de funcionários subordinados",
	VIEW_SUBORDINATE_ENTRIES:                "Permissão para ver lançamentos de funcionários subordinados",
	VIEW_SUBORDINATE_DASHBOARD:              "Permissão para ver dashboard dos subordinados",
	VIEW_SUBORDINATE_REPORT:                 "Permissão para ver relatório dos subordinados",
	GENERATE_SUBORDINATE_PDF_REPORT:         "Permissão para gerar relatório dos subordinados em PDF",

	CLOCK_IN_OUT:                "Permissão para bater ponto",
	VIEW_DASHBOARD:              "Permissão para ver próprio dashboard",
	VIEW_REPORT:                 "Permissão para ver próprio relatório",
	VIEW_HOLIDAYS:               "Permissão para ver feriados",
	GENERATE_PDF_REPORT:         "Permissão para gerar próprio relatório em PDF",
	REQUEST_TIME_CLOCK_ADJUSTMENT: "Permissão para solicitar ajuste de ponto",
	REQUEST_ENTRY:               "Permissão para solicitar lançamentos",
}

// permissions set
type PermissionSet map[Permission]struct{}

func (ps PermissionSet) Has(permission Permission) bool {
	_, exists := ps[permission]
	return exists
}

func (ps PermissionSet) Add(permission Permission) {
	ps[permission] = struct{}{}
}

func (ps PermissionSet) Remove(permission Permission) {
	delete(ps, permission)
}

func (ps PermissionSet) Clone() PermissionSet {
	clone := make(PermissionSet, len(ps))

	for permission := range ps {
		clone[permission] = struct{}{}
	}

	return clone
}

func (ps PermissionSet) GetPermissions() []Permission {
	permissions := make([]Permission, 0)

	for permission := range ps {
		permissions = append(permissions, permission)
	}

	return permissions
}

// factory
func NewPermissionSet(permissions ...Permission) PermissionSet {
	set := make(PermissionSet)

	for _, permission := range permissions {
		set[permission] = struct{}{}
	}

	return set
}

// relationship between role and permission.
var allowedPermissionsByRole = map[Role]PermissionSet{
	MANAGER_ROLE: NewPermissionSet(
		ADJUST_REQUESTED_TIME_CLOCK,
		ADJUST_MANUAL_TIME_CLOCK,
		CREATE_EMPLOYEE_ENTRIES,
		DELETE_EMPLOYEE_ENTRIES,
		UPDATE_EMPLOYEE_ENTRIES,
		VIEW_EMPLOYEE_ENTRIES,
		VIEW_GENERAL_DASHBOARD,
		VIEW_GENERAL_REPORT,
		GENERATE_GENERAL_PDF_REPORT,
		CREATE_EMPLOYEE,
		DELETE_EMPLOYEE,
		UPDATE_EMPLOYEE,
		VIEW_EMPLOYEE_REGISTRATION,
		CREATE_WORK_JOURNEY,
		DELETE_WORK_JOURNEY,
		UPDATE_WORK_JOURNEY,
		VIEW_WORK_JOURNEY_REGISTRATION,
		CREATE_SCHEDULE,
		DELETE_SCHEDULE,
		UPDATE_SCHEDULE,
		VIEW_SCHEDULE_REGISTRATION,
		CREATE_HOLIDAY,
		DELETE_HOLIDAY,
		UPDATE_HOLIDAY,
		VIEW_HOLIDAY_REGISTRATION,
	),

	SUPERVISOR_ROLE: NewPermissionSet(
		ADJUST_SUBORDINATE_REQUESTED_TIME_CLOCK,
		ADJUST_SUBORDINATE_MANUAL_TIME_CLOCK,
		CREATE_SUBORDINATE_ENTRIES,
		DELETE_SUBORDINATE_ENTRIES,
		UPDATE_SUBORDINATE_ENTRIES,
		VIEW_SUBORDINATE_ENTRIES,
		VIEW_SUBORDINATE_DASHBOARD,
		VIEW_SUBORDINATE_REPORT,
		GENERATE_SUBORDINATE_PDF_REPORT,

		CLOCK_IN_OUT,
		VIEW_DASHBOARD,
		VIEW_REPORT,
		VIEW_HOLIDAYS,
		GENERATE_PDF_REPORT,
		REQUEST_TIME_CLOCK_ADJUSTMENT,
		REQUEST_ENTRY,
	),

	WORKER_ROLE: NewPermissionSet(
		CLOCK_IN_OUT,
		VIEW_DASHBOARD,
		VIEW_REPORT,
		VIEW_HOLIDAYS,
		GENERATE_PDF_REPORT,
		REQUEST_TIME_CLOCK_ADJUSTMENT,
		REQUEST_ENTRY,
	),

	OUTSOURCED_ROLE: NewPermissionSet(
		CREATE_EMPLOYEE_ENTRIES,
		UPDATE_EMPLOYEE_ENTRIES,
		VIEW_EMPLOYEE_ENTRIES,
		VIEW_GENERAL_DASHBOARD,
		VIEW_GENERAL_REPORT,
		GENERATE_GENERAL_PDF_REPORT,
		CREATE_EMPLOYEE,
		UPDATE_EMPLOYEE,
		VIEW_EMPLOYEE_REGISTRATION,
		CREATE_WORK_JOURNEY,
		UPDATE_WORK_JOURNEY,
		VIEW_WORK_JOURNEY_REGISTRATION,
		CREATE_SCHEDULE,
		UPDATE_SCHEDULE,
		VIEW_SCHEDULE_REGISTRATION,
		CREATE_HOLIDAY,
		UPDATE_HOLIDAY,
		VIEW_HOLIDAY_REGISTRATION,
		REQUEST_TIME_CLOCK_ADJUSTMENT,
		REQUEST_ENTRY,
	),
}

// final object that represents all the logic of role and permission
type RoleDefinition struct {
	role        Role
	permissions PermissionSet
}

func (rd *RoleDefinition) GetRole() Role {
	return rd.role
}

func (rd *RoleDefinition) GetPermissionSet() PermissionSet {
	return rd.permissions
}

func (rd *RoleDefinition) SetPermissionSet(permissionSet PermissionSet) error {
	allowedPermissions, exists := allowedPermissionsByRole[rd.role]
	if !exists {
		return fmt.Errorf("%w: %s", exceptions.ErrInvalidRole, rd.role)
	}
	
	for permission, _ := range permissionSet {
		if !allowedPermissions.Has(permission) {
			return fmt.Errorf(
				"%w: %s",
				exceptions.ErrPermissionNotAllowed,
				rd.role,
			)
		}
	}

	rd.permissions = permissionSet

	return nil
}

// factories
// how to alter role? here, create a brand new role and delegate it's permissions
func NewRoleDefinition(role Role, selectedPermissions ...Permission) (*RoleDefinition, error) {
	rd := new(RoleDefinition)

	//create a aux var to store permissionSet
	permissionSet := NewPermissionSet()

	//for each permission from the func parameter, adds to aux var
	for _, permission := range selectedPermissions {
		permissionSet.Add(permission)
	}

	if err := rd.SetPermissionSet(permissionSet); err != nil {
		return nil, err
	}
	
	return &RoleDefinition{
		role:        role,
		permissions: permissionSet,
	}, nil
}

func NewCompanyDefaultRoleDefinition() *RoleDefinition {
	return &RoleDefinition{
		role:        MANAGER_ROLE,
		permissions: allowedPermissionsByRole[MANAGER_ROLE],
	}
}

func NewSupervisorDefaultRoleDefinition() *RoleDefinition {
	return &RoleDefinition{
		role:        SUPERVISOR_ROLE,
		permissions: allowedPermissionsByRole[SUPERVISOR_ROLE],
	}
}

func NewWorkerDefaultRoleDefinition() *RoleDefinition {
	return &RoleDefinition{
		role:        WORKER_ROLE,
		permissions: allowedPermissionsByRole[WORKER_ROLE],
	}
}

func NewOutsourcedDefaultRoleDefinition() *RoleDefinition {
	return &RoleDefinition{
		role: OUTSOURCED_ROLE,
		permissions: allowedPermissionsByRole[OUTSOURCED_ROLE],
	}
}
