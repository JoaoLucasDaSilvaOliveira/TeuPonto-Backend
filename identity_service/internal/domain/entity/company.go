package entity

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
	"teuponto.com.br/backend/identity_service/internal/domain/valueobject"
)

type CompanyOutsourceds map[valueobject.CNPJ]model.RoleDefinition

func (co CompanyOutsourceds) Has(cnpj valueobject.CNPJ) bool {
	_, exists := co[cnpj]

	return exists
}

func (co CompanyOutsourceds) Add(cnpj valueobject.CNPJ, roleDef model.RoleDefinition) {
	co[cnpj] = roleDef
	//TODO: como atualizar/alterar as permissões???
}

func (co CompanyOutsourceds) Remove(cnpj valueobject.CNPJ) {
	if co.Has(cnpj) {
		delete(co, cnpj)
	}
}

type Company struct {
	id                       uuid.UUID
	idAddress                uuid.UUID
	idPayment                uuid.UUID
	cnpj                     valueobject.CNPJ
	corporateName            string
	businessName             string
	geofencePolicy           valueobject.PoliticaGeofence
	email                    valueobject.Email
	defaultCalculationPolicy model.CalculationPolicy
	companyOutsourceds       CompanyOutsourceds
}

func (c *Company) GetID() uuid.UUID {
	return c.id
}

func (c *Company) GetIDAddress() uuid.UUID {
	return c.idAddress
}

func (c *Company) SetIDAddress(idAddress uuid.UUID) {
	c.idAddress = idAddress
}

func (c *Company) GetIDPayment() uuid.UUID {
	return c.idPayment
}

func (c *Company) SetIDPayment(idPayment uuid.UUID) {
	c.idPayment = idPayment
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

	return nil
}

func (c *Company) GetCorporateName() string {
	return c.corporateName
}

func (c *Company) SetCorporateName(corporateName string) error {
	trimmed := strings.TrimSpace(corporateName)
	if trimmed == "" {
		return fmt.Errorf("%w: razao social vazia", exceptions.ErrEmptyString)
	}

	c.corporateName = trimmed

	return nil
}

func (c *Company) GetBusinessName() string {
	return c.businessName
}

func (c *Company) SetBusinessName(businessName string) error {
	trimmed := strings.TrimSpace(businessName)
	if trimmed == "" {
		return fmt.Errorf("%w: nome fantasia vazio", exceptions.ErrEmptyString)
	}

	c.businessName = trimmed

	return nil
}

func (c *Company) GetGeofecePolicy() valueobject.PoliticaGeofence {
	return c.geofencePolicy
}

func (c *Company) SetGeofecePolicy(geofecePolicy valueobject.PoliticaGeofence) {
	c.geofencePolicy = geofecePolicy
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

	return nil
}

func (c *Company) GetDefaultCalculationPolicy() model.CalculationPolicy {
	return c.defaultCalculationPolicy
}

func (c *Company) SetDefaultCalculationPolicy(calculationPolicy model.CalculationPolicy) error {
	if calculationPolicy == nil {
		return exceptions.ErrEmptyCalculationPolicy
	}

	c.defaultCalculationPolicy = calculationPolicy
	return nil
}

func (c *Company) GetCompanyOutsourceds() CompanyOutsourceds {
	return c.companyOutsourceds
}

func (c *Company) SetCompanyOutsourceds(companyOutsourceds CompanyOutsourceds) {
	c.companyOutsourceds = companyOutsourceds
}

func NewCompany(
	idAddress uuid.UUID,
	idPaymet uuid.UUID,
	rawCNPJ string,
	corporateName string,
	businessName string,
	geofencePolicy valueobject.PoliticaGeofence,
	rawEmail string,
	defaultCalculationPolicy model.CalculationPolicy,
	companyOutsourceds CompanyOutsourceds,
) (*Company, error) {
	company := new(Company)

	company.id = uuid.New()
	
	company.SetIDAddress(idAddress)
	company.SetIDPayment(idPaymet)
	company.SetCompanyOutsourceds(companyOutsourceds)
	company.SetGeofecePolicy(geofencePolicy)

	if err := company.SetCNPJ(rawCNPJ); err != nil {
		return nil, err
	}
	
	if err := company.SetEmail(rawEmail); err != nil {
		return nil, err
	}
	
	if err := company.SetBusinessName(businessName); err != nil {
		return nil, err
	}
	
	if err := company.SetCorporateName(corporateName); err != nil {
		return nil, err
	}
	
	if err := company.SetDefaultCalculationPolicy(defaultCalculationPolicy); err != nil {
		return nil, err
	}	

	return company, nil
}
