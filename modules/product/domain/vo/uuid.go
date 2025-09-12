package vo

import (
	"github.com/google/uuid"
	"github.com/valdinei-santos/product-details/infra/logger"
)

type ID uuid.UUID

func NewUUID(log logger.Logger) (ID, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Error("Falha ao gerar UUID", err)
	}
	return ID(uuid), nil
}
