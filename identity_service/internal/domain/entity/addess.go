package entity

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/paemuri/brdoc"
	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
	"teuponto.com.br/backend/identity_service/internal/domain/valueobject"
)

type Address struct {
	id               uuid.UUID
	cep              valueobject.CEP
	street           string
	number           int
	neighborhood     string
	longitude        float64
	latitude         float64
	codIbge          int
}

//TODO:
// pesquisar e validar o range do cod ibge
// cod ibge e number não podem ser negativos

func (a *Address) GetID() uuid.UUID {
	return a.id
}

func (a *Address) GetCEP() valueobject.CEP {
	return a.cep
}

func (a *Address) SetCEP(rawCEP string, federativeUnit brdoc.FederativeUnit) error {
	cep, err := valueobject.NewCEP(rawCEP, federativeUnit)
	if err != nil {
		return err
	}

	a.cep = cep
	return nil
}

func (a *Address) GetStreet() string {
	return a.street
}

func (a *Address) SetStreet(street string) error {
	trimmed := strings.TrimSpace(street)
	if trimmed == "" {
		return fmt.Errorf("%w: rua vazia", exceptions.ErrEmptyString)
	}

	a.street = trimmed
	return nil
}

func (a *Address) GetNumber() int {
	return a.number
}

func (a *Address) SetNumber(number int) {
	a.number = number
}

func (a *Address) GetNeighborhood() string {
	return a.neighborhood
}

func (a *Address) SetNeighborhood(neighborhood string) error {
	trimmed := strings.TrimSpace(neighborhood)
	if trimmed == "" {
		return fmt.Errorf("%w: bairro vazio", exceptions.ErrEmptyString)
	}

	a.neighborhood = trimmed
	return nil
}

func (a *Address) GetLongitude() float64 {
	return a.longitude
}

func (a *Address) SetLongitude(longitude float64) {
	a.longitude = longitude
}

func (a *Address) GetLatitude() float64 {
	return a.latitude
}

func (a *Address) SetLatitude(latitude float64) {
	a.latitude = latitude
}

func (a *Address) GetCodIbge() int {
	return a.codIbge
}

func (a *Address) SetCodIbge(codIbge int) {
	a.codIbge = codIbge
}

func NewAddress(
	rawCEP string,
	federativeUnit brdoc.FederativeUnit,
	street string,
	number int,
	neighborhood string,
	longitude float64,
	latitude float64,
	codIbge int,
) (*Address, error) {
	address := new(Address)

	address.id = uuid.New()

	address.SetCodIbge(codIbge)
	address.SetLatitude(latitude)
	address.SetLongitude(longitude)
	address.SetNumber(number)
	
	if err := address.SetCEP(rawCEP, federativeUnit); err != nil {
		return nil, err
	}
	
	if err := address.SetNeighborhood(neighborhood); err != nil {
		return nil, err
	}
	
	if err := address.SetStreet(street); err != nil {
		return nil, err
	}
	
	return address, nil
}
