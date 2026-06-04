package model

import (
	"fmt"
	"time"

	"teuponto.com.br/backend/identity_service/internal/domain/exceptions"
)

//Um objeto auditavel presenta as informações de criação e manipulação das entidades
//
//Estrutura basilar do sistema: serve de base para quase todas as entidades do sistema
type Auditable struct {
	createdAt *time.Time
	updatedAt *time.Time
	deletedAt *time.Time //soft delete
}

func (a *Auditable) GetCreatedAt() *time.Time{
	return a.createdAt
}

func (a *Auditable) GetUpdatedAt() *time.Time{
	return a.updatedAt
}

func (a *Auditable) GetDeletedAt() *time.Time{
	return a.deletedAt
}

func (a *Auditable) SetUpdatedAt(updatedAt *time.Time) error {
	location, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		return fmt.Errorf("%w: %w", exceptions.ErrLoadLocation, err)
	}

	updatedAtWithLocation := updatedAt.In(location)
	
	a.updatedAt = &updatedAtWithLocation
	return nil
}

func (a *Auditable) SetDeletedAt(deletedAt *time.Time) error {
	location, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		return fmt.Errorf("%w: %w", exceptions.ErrLoadLocation, err)
	}

	deletedAtWithLocation := deletedAt.In(location)
	a.deletedAt = &deletedAtWithLocation
	return nil
}

func NewAuditable() (*Auditable, error) {
	//timezone america/saoPaulo
	location, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		return nil, fmt.Errorf("%w: %w", exceptions.ErrLoadLocation, err)
	}

	now := time.Now().In(location)
	
	return &Auditable{
		createdAt: &now,
		updatedAt: nil,
		deletedAt: nil,
	}, nil
}