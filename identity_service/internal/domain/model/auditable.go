package model

import (
	"fmt"
	"time"

	sistem_errors "teuponto.com.br/backend/identity_service/internal/domain/error"
)

//Um objeto auditavel presenta as informações de criação e manipulação das entidades
//
//Estrutura basilar do sistema: serve de base para quase todas as entidades do sistema
type Auditable struct {
	createdAt *time.Time
	updatedAt *time.Time
	deletedAt *time.Time //soft delete
}

func (a *Auditable) getCreatedAt() *time.Time{
	return a.createdAt
}

func (a *Auditable) getUpdatedAt() *time.Time{
	return a.updatedAt
}

func (a *Auditable) getDeletedAt() *time.Time{
	return a.deletedAt
}

func (a *Auditable) setUpdatedAt(updatedAt *time.Time){
	a.updatedAt = updatedAt
}

func (a *Auditable) setDeletedAt(deletedAt *time.Time){
	a.deletedAt = deletedAt
}

func NewAuditable() (*Auditable, error) {
	//timezone america/saoPaulo
	location, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		return nil, fmt.Errorf("%w: %w", sistem_errors.ErrLoadLocation, err)
	}

	now := time.Now().In(location)
	
	return &Auditable{
		createdAt: &now,
		updatedAt: nil,
		deletedAt: nil,
	}, nil
}