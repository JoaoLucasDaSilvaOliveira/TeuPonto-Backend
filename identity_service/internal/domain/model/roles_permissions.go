package model

import (
	sistemError "teuponto.com.br/backend/identity_service/internal/domain/error"
	"fmt"
)

type Role string
type Permission string

// roles do sistema
const (
	COMPANY_ROLE    Role = "COMPANY"
	SUPERVISOR_ROLE Role = "SUPERVISOR"
	WORKER_ROLE     Role = "WORKER"
)

// permissões gerais
const (
	// nível admin
	AJUSTAR_PONTO_SOLICITADO      Permission = "AJUSTAR_PONTO_SOLICITADO"
	AJUSTAR_PONTO_MANUAL          Permission = "AJUSTAR_PONTO_MANUAL"
	CRUD_LANCAMENTOS_EMPREGADOS   Permission = "CRUD_LANCAMENTOS_EMPREGADOS"
	VER_DASHBOARD_GERAL           Permission = "VER_DASHBOARD_GERAL"
	VER_RELATORIO_GERAL           Permission = "VER_RELATORIO_GERAL"
	GERAR_RELATORIO_PDF_GERAL     Permission = "GERAR_RELATORIO_PDF_GERAL"
	CRUD_FUNCIONARIO              Permission = "CRUD_FUNCIONARIO"
	CRUD_JORNADA                  Permission = "CRUD_JORNADA"
	CRUD_HORARIO                  Permission = "CRUD_HORARIO"
	CRUD_FERIADO                  Permission = "CRUD_FERIADO"

	// nível supervisor
	AJUSTAR_PONTO_SOLICITADO_SUBORDINADO    Permission = "AJUSTAR_PONTO_SOLICITADO_SUBORDINADO"
	AJUSTAR_PONTO_MANUAL_SUBORDINADO        Permission = "AJUSTAR_PONTO_MANUAL_SUBORDINADO"
	CRUD_LANCAMENTOS_SUBORDINADO            Permission = "CRUD_LANCAMENTOS_SUBORDINADO"
	VER_DASHBOARD_SUBORDINADO               Permission = "VER_DASHBOARD_SUBORDINADO"
	VER_RELATORIO_SUBORDINADO               Permission = "VER_RELATORIO_SUBORDINADO"
	GERAR_RELATORIO_PDF_SUBORDINADO         Permission = "GERAR_RELATORIO_PDF_SUBORDINADO"

	// nível worker
	BATER_PONTO              Permission = "BATER_PONTO"
	VER_DASHBOARD            Permission = "VER_DASHBOARD"
	VER_RELATORIO            Permission = "VER_RELATORIO"
	VER_FERIADOS             Permission = "VER_FERIADOS"
	GERAR_RELATORIO_PDF      Permission = "GERAR_RELATORIO_PDF"
	SOLICITAR_AJUSTE_PONTO   Permission = "SOLICITAR_AJUSTE_PONTO"
	SOLICITAR_LANCAMENTO     Permission = "SOLICITAR_LANCAMENTO"
)

//descricao das permissoes
var permissionDescriptions = map[Permission]string{
	AJUSTAR_PONTO_SOLICITADO:   "Permissão para ajustar ponto solicitado",
	AJUSTAR_PONTO_MANUAL:       "Permissão para ajuste manual de ponto",
	CRUD_LANCAMENTOS_EMPREGADOS: "Permissão para criar, atualizar ou remover lançamentos de empregados",
	VER_DASHBOARD_GERAL:        "Permissão para ver dashboard geral da empresa",
	VER_RELATORIO_GERAL:        "Permissão para ver relatório geral da empresa",
	GERAR_RELATORIO_PDF_GERAL:  "Permissão para gerar relatório geral da empresa em PDF",
	CRUD_FUNCIONARIO:           "Permissão para criar, atualizar ou remover funcionários",
	CRUD_JORNADA:               "Permissão para criar, atualizar ou remover jornadas",
	CRUD_HORARIO:               "Permissão para criar, atualizar ou remover horários",
	CRUD_FERIADO:               "Permissão para criar, atualizar ou remover feriados",

	AJUSTAR_PONTO_SOLICITADO_SUBORDINADO: "Permissão para ajustar ponto solicitado por subordinados",
	AJUSTAR_PONTO_MANUAL_SUBORDINADO:     "Permissão para ajustar manualmente ponto de subordinados",
	CRUD_LANCAMENTOS_SUBORDINADO:         "Permissão para gerenciar lançamentos de subordinados",
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

//conjunto de permissoes
type PermissionSet map[Permission]struct{}

//factory
func NewPermissionSet(permissions ...Permission) PermissionSet {
	set := make(PermissionSet)

	for _, permission := range permissions {
		set[permission] = struct{}{}
	}

	return set
}

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

//relacionamento entre role -> permission
var allowedPermissionsByRole = map[Role]PermissionSet{
	COMPANY_ROLE: NewPermissionSet(
		AJUSTAR_PONTO_SOLICITADO,
		AJUSTAR_PONTO_MANUAL,
		CRUD_LANCAMENTOS_EMPREGADOS,
		VER_DASHBOARD_GERAL,
		VER_RELATORIO_GERAL,
		GERAR_RELATORIO_PDF_GERAL,
		CRUD_FUNCIONARIO,
		CRUD_JORNADA,
		CRUD_HORARIO,
		CRUD_FERIADO,
	),

	SUPERVISOR_ROLE: NewPermissionSet(
		AJUSTAR_PONTO_SOLICITADO_SUBORDINADO,
		AJUSTAR_PONTO_MANUAL_SUBORDINADO,
		CRUD_LANCAMENTOS_SUBORDINADO,
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
}

//objeto final que representa toda a logica de role e permissao
type RoleDefinition struct {
	role        Role
	permissions PermissionSet
}

//factory
func NewRoleDefinition(role Role, selectedPermissions ...Permission) (RoleDefinition, error) {
	allowedPermissions, exists := allowedPermissionsByRole[role]
	if !exists {
		return RoleDefinition{}, fmt.Errorf("%w: %s", sistemError.ErrInvalidRole, role)
	}

	permissions := NewPermissionSet()

	for _, permission := range selectedPermissions {
		if !allowedPermissions.Has(permission) {
			return RoleDefinition{}, fmt.Errorf(
				"%w: %s",
				sistemError.ErrPermissionNotAllowed,
				role,
			)
		}

		permissions.Add(permission)
	}

	return RoleDefinition{
		role:        role,
		permissions: permissions,
	}, nil
}