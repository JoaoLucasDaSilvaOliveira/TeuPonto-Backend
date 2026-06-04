package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
	"teuponto.com.br/backend/identity_service/internal/domain/valueobject"
)

type Company struct {
	*model.Auditable //embedded
	id                      uuid.UUID
	idEndereco              uuid.UUID
	idPagamento             uuid.UUID
	cnpj                    valueobject.CNPJ
	razaoSocial             string
	nomeFantasia            string
	politicaGeofence        valueobject.PoliticaGeofence
	email                   valueobject.Email
	senhaAcesso             string
	cnpjTerceirizado        valueobject.CNPJ
	senhaAcessoTerceirizado string
	rolesTerceirizado       model.RoleDefinition
}

func (c *Company) GetID() uuid.UUID {
	return c.id
}

func (c *Company) GetIDEndereco() uuid.UUID {
	return c.idEndereco
}

func (c *Company) SetIDEndereco(idEndereco uuid.UUID) {
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.idEndereco = idEndereco
}

func (c *Company) GetIDPagamento() uuid.UUID {
	return c.idPagamento
}

func (c *Company) SetIDPagamento(idPagamento uuid.UUID) {
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.idPagamento = idPagamento
}

func (c *Company) GetCNPJ() valueobject.CNPJ {
	return c.cnpj
}

func (c *Company) SetCNPJ(rawCNPJ string) error {
	cnpj, err := valueobject.NewCNPJ(rawCNPJ)
	if err != nil {
		return err
	}

	c.cnpj = cnpj
	
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	return nil
}

func (c *Company) GetRazaoSocial() string {
	return c.razaoSocial
}

func (c *Company) SetRazaoSocial(razaoSocial string) error {
	trimmed := strings.TrimSpace(razaoSocial)
	if trimmed == "" {
		return fmt.Errorf("%w: razao social vazia", exceptions.ErrEmptyString)
	}

	c.razaoSocial = trimmed
	
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	return nil
}

func (c *Company) GetNomeFantasia() string {
	return c.nomeFantasia
}

func (c *Company) SetNomeFantasia(nomeFantasia string) {
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.nomeFantasia = strings.TrimSpace(nomeFantasia)
}

func (c *Company) GetPoliticaGeofence() valueobject.PoliticaGeofence {
	return c.politicaGeofence
}

func (c *Company) SetPoliticaGeofence(politicaGeofence valueobject.PoliticaGeofence) {
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.politicaGeofence = politicaGeofence
}

func (c *Company) GetEmail() valueobject.Email {
	return c.email
}

func (c *Company) SetEmail(rawEmail string) error {
	email, err := valueobject.NewEmail(rawEmail)
	if err != nil {
		return err
	}

	c.email = email

	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	return nil
}

func (c *Company) GetSenhaAcesso() string {
	return c.senhaAcesso
}

func (c *Company) SetSenhaAcesso(senhaAcesso string) {
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.senhaAcesso = senhaAcesso
}

func (c *Company) GetCNPJTerceirizado() valueobject.CNPJ {
	return c.cnpjTerceirizado
}

func (c *Company) SetCNPJTerceirizado(rawCNPJ string) error {
	cnpj, err := valueobject.NewCNPJ(rawCNPJ)
	if err != nil {
		return err
	}

	c.cnpjTerceirizado = cnpj

	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	return nil
}

func (c *Company) GetSenhaAcessoTerceirizado() string {
	return c.senhaAcessoTerceirizado
}

func (c *Company) SetSenhaAcessoTerceirizado(senhaAcessoTerceirizado string) {
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.senhaAcessoTerceirizado = senhaAcessoTerceirizado
}

func (c *Company) GetRolesTerceirizado() model.RoleDefinition {
	return c.rolesTerceirizado
}

func (c *Company) SetRolesTerceirizado(rolesTerceirizado model.RoleDefinition) {
	//altera o updated_at para auditoria
	now := time.Now()
	c.SetUpdatedAt(&now)

	c.rolesTerceirizado = rolesTerceirizado
}

func NewCompany(
	idEndereco uuid.UUID,
	idPagamento uuid.UUID,
	rawCNPJ string,
	razaoSocial string,
	nomeFantasia string,
	politicaGeofence valueobject.PoliticaGeofence,
	rawEmail string,
	senhaAcesso string,
	rawCNPJTerceirizado string,
	senhaAcessoTerceirizado string,
	rolesTerceirizado model.RoleDefinition,
) (*Company, error) {
	auditable, err := model.NewAuditable()
	if err != nil {
		return nil, err
	}

	cnpj, err := valueobject.NewCNPJ(rawCNPJ)
	if err != nil {
		return nil, err
	}

	email, err := valueobject.NewEmail(rawEmail)
	if err != nil {
		return nil, err
	}

	cnpjTerceirizado, err := valueobject.NewCNPJ(rawCNPJTerceirizado)
	if err != nil {
		return nil, err
	}

	trimmedRazaoSocial := strings.TrimSpace(razaoSocial)
	if trimmedRazaoSocial == "" {
		return nil, fmt.Errorf("%w: razao social vazia", exceptions.ErrEmptyString)
	}

	return &Company{
		Auditable:               auditable,
		id:                      uuid.New(),
		idEndereco:              idEndereco,
		idPagamento:             idPagamento,
		cnpj:                    cnpj,
		razaoSocial:             trimmedRazaoSocial,
		nomeFantasia:            strings.TrimSpace(nomeFantasia),
		politicaGeofence:        politicaGeofence,
		email:                   email,
		senhaAcesso:             senhaAcesso,
		cnpjTerceirizado:        cnpjTerceirizado,
		senhaAcessoTerceirizado: senhaAcessoTerceirizado,
		rolesTerceirizado:       rolesTerceirizado,
	}, nil
}
