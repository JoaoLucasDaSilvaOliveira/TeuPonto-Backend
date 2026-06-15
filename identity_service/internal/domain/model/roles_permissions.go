package model

import (
	"fmt"

	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

type Role string
type Permission string

// roles do sistema
const (
	MANAGER_ROLE    Role = "MANAGER"
	SUPERVISOR_ROLE Role = "SUPERVISOR"
	WORKER_ROLE     Role = "WORKER"
	OUTSOURCED_ROLE Role = "OUTSOURCED"
)

// permissões gerais
const (
	// nível admin
	AJUSTAR_PONTO_SOLICITADO         Permission = "AJUSTAR_PONTO_SOLICITADO"
	AJUSTAR_PONTO_MANUAL             Permission = "AJUSTAR_PONTO_MANUAL"
	CRIAR_LANCAMENTOS_EMPREGADOS     Permission = "CRIAR_LANCAMENTOS_EMPREGADOS"
	DELETAR_LANCAMENTOS_EMPREGADOS   Permission = "DELETAR_LANCAMENTOS_EMPREGADOS"
	ATUALIZAR_LANCAMENTOS_EMPREGADOS Permission = "ATUALIZAR_LANCAMENTOS_EMPREGADOS"
	VER_LANCAMENTOS_EMPREGADOS       Permission = "VER_LANCAMENTOS_EMPREGADOS"
	VER_DASHBOARD_GERAL              Permission = "VER_DASHBOARD_GERAL"
	VER_RELATORIO_GERAL              Permission = "VER_RELATORIO_GERAL"
	GERAR_RELATORIO_PDF_GERAL        Permission = "GERAR_RELATORIO_PDF_GERAL"
	CRIAR_FUNCIONARIO                Permission = "CRIAR_FUNCIONARIO"
	DELETAR_FUNCIONARIO              Permission = "DELETAR_FUNCIONARIO"
	ATUALIZAR_FUNCIONARIO            Permission = "ATUALIZAR_FUNCIONARIO"
	VER_CADASTRO_FUNCIONARIO         Permission = "VER_CADASTRO_FUNCIONARIO"
	CRIAR_JORNADA                    Permission = "CRIAR_JORNADA"
	DELETAR_JORNADA                  Permission = "DELETAR_JORNADA"
	ATUALIZAR_JORNADA                Permission = "ATUALIZAR_JORNADA"
	VER_CADASTRO_JORNADA             Permission = "VER_CADASTRO_JORNADA"
	CRIAR_HORARIO                    Permission = "CRIAR_HORARIO"
	DELETAR_HORARIO                  Permission = "DELETAR_HORARIO"
	ATUALIZAR_HORARIO                Permission = "ATUALIZAR_HORARIO"
	VER_CADASTRO_HORARIO             Permission = "VER_CADASTRO_HORARIO"
	CRIAR_FERIADO                    Permission = "CRIAR_FERIADO"
	DELETAR_FERIADO                  Permission = "DELETAR_FERIADO"
	ATUALIZAR_FERIADO                Permission = "ATUALIZAR_FERIADO"
	VER_CADASTRO_FERIADO             Permission = "VER_CADASTRO_FERIADO"

	// nível supervisor
	AJUSTAR_PONTO_SOLICITADO_SUBORDINADO Permission = "AJUSTAR_PONTO_SOLICITADO_SUBORDINADO"
	AJUSTAR_PONTO_MANUAL_SUBORDINADO     Permission = "AJUSTAR_PONTO_MANUAL_SUBORDINADO"
	CRIAR_LANCAMENTOS_SUBORDINADO        Permission = "CRIAR_LANCAMENTOS_SUBORDINADO"
	DELETAR_LANCAMENTOS_SUBORDINADO      Permission = "DELETAR_LANCAMENTOS_SUBORDINADO"
	ATUALIZAR_LANCAMENTOS_SUBORDINADO    Permission = "ATUALIZAR_LANCAMENTOS_SUBORDINADO"
	VER_LANCAMENTOS_SUBORDINADO          Permission = "VER_LANCAMENTOS_SUBORDINADO"
	VER_DASHBOARD_SUBORDINADO            Permission = "VER_DASHBOARD_SUBORDINADO"
	VER_RELATORIO_SUBORDINADO            Permission = "VER_RELATORIO_SUBORDINADO"
	GERAR_RELATORIO_PDF_SUBORDINADO      Permission = "GERAR_RELATORIO_PDF_SUBORDINADO"

	// nível operador
	BATER_PONTO            Permission = "BATER_PONTO"
	VER_DASHBOARD          Permission = "VER_DASHBOARD"
	VER_RELATORIO          Permission = "VER_RELATORIO"
	VER_FERIADOS           Permission = "VER_FERIADOS"
	GERAR_RELATORIO_PDF    Permission = "GERAR_RELATORIO_PDF"
	SOLICITAR_AJUSTE_PONTO Permission = "SOLICITAR_AJUSTE_PONTO"
	SOLICITAR_LANCAMENTO   Permission = "SOLICITAR_LANCAMENTO"
)

// descricao das permissoes
var PermissionDescriptions = map[Permission]string{
	AJUSTAR_PONTO_SOLICITADO:         "Permissão para ajustar ponto solicitado",
	AJUSTAR_PONTO_MANUAL:             "Permissão para ajuste manual de ponto",
	CRIAR_LANCAMENTOS_EMPREGADOS:     "Permissão para criar lançamentos dos empregados",
	DELETAR_LANCAMENTOS_EMPREGADOS:   "Permissão para deletar lançamentos dos empregados",
	ATUALIZAR_LANCAMENTOS_EMPREGADOS: "Permissão para atualizar lançamentos dos empregados",
	VER_LANCAMENTOS_EMPREGADOS:       "Permissão para ver lançamentos dos empregados",
	VER_DASHBOARD_GERAL:              "Permissão para ver dashboard geral da empresa",
	VER_RELATORIO_GERAL:              "Permissão para ver relatório geral da empresa",
	GERAR_RELATORIO_PDF_GERAL:        "Permissão para gerar relatório geral da empresa em PDF",
	CRIAR_FUNCIONARIO:                "Permissão para cadastrar de funcionários",
	DELETAR_FUNCIONARIO:              "Permissão para deletar cadastros de funcionários",
	ATUALIZAR_FUNCIONARIO:            "Permissão para atualizar cadastros de funcionários",
	VER_CADASTRO_FUNCIONARIO:         "Permissão para ver cadastros de funcionários",
	CRIAR_JORNADA:                    "Permissão para cadastrar jornadas",
	DELETAR_JORNADA:                  "Permissão para deletar cadastros de jornadas",
	ATUALIZAR_JORNADA:                "Permissão para atualizar cadastros de jornadas",
	VER_CADASTRO_JORNADA:             "Permissão para ver cadastros de jornadas",
	CRIAR_HORARIO:                    "Permissão para cadastrar horários",
	DELETAR_HORARIO:                  "Permissão para deletar cadastros de horários",
	ATUALIZAR_HORARIO:                "Permissão para atualizar cadastros de horários",
	VER_CADASTRO_HORARIO:             "Permissão para ver cadastros de horários",
	CRIAR_FERIADO:                    "Permissão para cadastrar feriados",
	DELETAR_FERIADO:                  "Permissão para deletar cadastros de feriados",
	ATUALIZAR_FERIADO:                "Permissão para atualizar cadastros de feriados",
	VER_CADASTRO_FERIADO:             "Permissão para ver cadastros de feriados",

	AJUSTAR_PONTO_SOLICITADO_SUBORDINADO: "Permissão para ajustar ponto solicitado por subordinados",
	AJUSTAR_PONTO_MANUAL_SUBORDINADO:     "Permissão para ajustar manualmente ponto de subordinados",
	CRIAR_LANCAMENTOS_SUBORDINADO:        "Permissão para cadastrar lançamentos de funcionários subordinados",
	DELETAR_LANCAMENTOS_SUBORDINADO:      "Permissão para deletar lançamentos de funcionários subordinados",
	ATUALIZAR_LANCAMENTOS_SUBORDINADO:    "Permissão para atualizar lançamentos de funcionários subordinados",
	VER_LANCAMENTOS_SUBORDINADO:          "Permissão para ver lançamentos de funcionários subordinados",
	VER_DASHBOARD_SUBORDINADO:            "Permissão para ver dashboard dos subordinados",
	VER_RELATORIO_SUBORDINADO:            "Permissão para ver relatório dos subordinados",
	GERAR_RELATORIO_PDF_SUBORDINADO:      "Permissão para gerar relatório dos subordinados em PDF",

	BATER_PONTO:            "Permissão para bater ponto",
	VER_DASHBOARD:          "Permissão para ver próprio dashboard",
	VER_RELATORIO:          "Permissão para ver próprio relatório",
	VER_FERIADOS:           "Permissão para ver feriados",
	GERAR_RELATORIO_PDF:    "Permissão para gerar próprio relatório em PDF",
	SOLICITAR_AJUSTE_PONTO: "Permissão para solicitar ajuste de ponto",
	SOLICITAR_LANCAMENTO:   "Permissão para solicitar lançamentos",
}

// conjunto de permissoes
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

// relacionamento entre role -> permission
var allowedPermissionsByRole = map[Role]PermissionSet{
	MANAGER_ROLE: NewPermissionSet(
		AJUSTAR_PONTO_SOLICITADO,
		AJUSTAR_PONTO_MANUAL,
		CRIAR_LANCAMENTOS_EMPREGADOS,
		DELETAR_LANCAMENTOS_EMPREGADOS,
		ATUALIZAR_LANCAMENTOS_EMPREGADOS,
		VER_LANCAMENTOS_EMPREGADOS,
		VER_DASHBOARD_GERAL,
		VER_RELATORIO_GERAL,
		GERAR_RELATORIO_PDF_GERAL,
		CRIAR_FUNCIONARIO,
		DELETAR_FUNCIONARIO,
		ATUALIZAR_FUNCIONARIO,
		VER_CADASTRO_FUNCIONARIO,
		CRIAR_JORNADA,
		DELETAR_JORNADA,
		ATUALIZAR_JORNADA,
		VER_CADASTRO_JORNADA,
		CRIAR_HORARIO,
		DELETAR_HORARIO,
		ATUALIZAR_HORARIO,
		VER_CADASTRO_HORARIO,
		CRIAR_FERIADO,
		DELETAR_FERIADO,
		ATUALIZAR_FERIADO,
		VER_CADASTRO_FERIADO,
	),

	SUPERVISOR_ROLE: NewPermissionSet(
		AJUSTAR_PONTO_SOLICITADO_SUBORDINADO,
		AJUSTAR_PONTO_MANUAL_SUBORDINADO,
		CRIAR_LANCAMENTOS_SUBORDINADO,
		DELETAR_LANCAMENTOS_SUBORDINADO,
		ATUALIZAR_LANCAMENTOS_SUBORDINADO,
		VER_LANCAMENTOS_SUBORDINADO,
		VER_DASHBOARD_SUBORDINADO,
		VER_RELATORIO_SUBORDINADO,
		GERAR_RELATORIO_PDF_SUBORDINADO,

		BATER_PONTO,
		VER_DASHBOARD,
		VER_RELATORIO,
		VER_FERIADOS,
		GERAR_RELATORIO_PDF,
		SOLICITAR_AJUSTE_PONTO,
		SOLICITAR_LANCAMENTO,
	),

	WORKER_ROLE: NewPermissionSet(
		BATER_PONTO,
		VER_DASHBOARD,
		VER_RELATORIO,
		VER_FERIADOS,
		GERAR_RELATORIO_PDF,
		SOLICITAR_AJUSTE_PONTO,
		SOLICITAR_LANCAMENTO,
	),

	OUTSOURCED_ROLE: NewPermissionSet(
		CRIAR_LANCAMENTOS_EMPREGADOS,
		ATUALIZAR_LANCAMENTOS_EMPREGADOS,
		VER_LANCAMENTOS_EMPREGADOS,
		VER_DASHBOARD_GERAL,
		VER_RELATORIO_GERAL,
		GERAR_RELATORIO_PDF_GERAL,
		CRIAR_FUNCIONARIO,
		ATUALIZAR_FUNCIONARIO,
		VER_CADASTRO_FUNCIONARIO,
		CRIAR_JORNADA,
		ATUALIZAR_JORNADA,
		VER_CADASTRO_JORNADA,
		CRIAR_HORARIO,
		ATUALIZAR_HORARIO,
		VER_CADASTRO_HORARIO,
		CRIAR_FERIADO,
		ATUALIZAR_FERIADO,
		VER_CADASTRO_FERIADO,
		SOLICITAR_AJUSTE_PONTO,
		SOLICITAR_LANCAMENTO,
	),
}

// objeto final que representa toda a logica de role e permissao
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
