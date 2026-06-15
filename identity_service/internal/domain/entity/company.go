package entity

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
	"teuponto.com.br/backend/identity_service/internal/domain/model"
	"teuponto.com.br/backend/identity_service/internal/domain/valueobject"
)

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
	companyOutsourceds       []Outsourced
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

func (c *Company) GetCompanyOutsourceds() []Outsourced {
	return c.companyOutsourceds
}

func (c *Company) SetCompanyOutsourceds(companyOutsourceds []Outsourced) {
	c.companyOutsourceds = companyOutsourceds
}

func (c *Company) AddCompanyOutsourced(companyOutsourced Outsourced) error {
	//check for duplicated associations
	for _, outsourced := range c.companyOutsourceds {
		if outsourced.GetID() == companyOutsourced.GetID() {
			return fmt.Errorf("%w: terceirizado já cadastrado para a empresa", exceptions.ErrDuplicatedElement)
		}
	}
	
	c.companyOutsourceds = append(c.companyOutsourceds, companyOutsourced)

	return nil
}

func (c *Company) RemoveCompanyOutsourced(id uuid.UUID) error {
	/* biblicaly accurated remove element from the slice method
		for i, outsourced := range c.companyOutsourceds {
			if outsourced.GetID() == id {
				// get the slice's position of the object and makes a copy of the slice but without the refered object/position
				// move the elements to the back, filling the slice's gap 
				copy(c.companyOutsourceds[i:], c.companyOutsourceds[i+1:])
		
				// replace the last position by struct's "zero value" 
				// this will make the pointer *model.User be considere as nil, releasing the ref to the GC
				c.companyOutsourceds[len(c.companyOutsourceds)-1] = Outsourced{} 
		
				// reduces the slice's size
				c.companyOutsourceds = c.companyOutsourceds[:len(c.companyOutsourceds)-1]
				break
		
				//why all that?
				// a lil explanation is that go is so fkng simple that the casualy created a native method which makes possible a memory leak
				// when u gonna delete something from the slice which contains pointers (interfaces, structs, strings, maps, slices, channels,...), using the append() method, u may think that it just returns the slice changed, copying the elements to a brand new slice, BUT NO, why? idk. So basicaly go just moves the bytes of the array to the left, causing that the last position of the slice is still there ponting to some space in the ram memory :-) and even if u dare to think that ur safe by just reduce the slice size, ur wrong -> go still grant that space in memory to the pointer, even if it's no in the range of the array.
			}
		}
	*/
	/* or just create an simple new slice and add everything u want and use/return the new box ;-) */
	companyOutsourcedsSlice := make([]Outsourced, 0, len(c.companyOutsourceds)-1)
	
	for _, outsourced := range c.companyOutsourceds {
		if outsourced.GetID() != id {
			companyOutsourcedsSlice = append(companyOutsourcedsSlice, outsourced)
		}
	}

	//check the slices's size
	if len(companyOutsourcedsSlice) == len(c.companyOutsourceds) {
		return fmt.Errorf("%w: terceirizado não cadastrado para a empresa", exceptions.ErrNoSuchElement)
	}

	c.companyOutsourceds = companyOutsourcedsSlice

	return nil
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
	companyOutsourceds []Outsourced,
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
