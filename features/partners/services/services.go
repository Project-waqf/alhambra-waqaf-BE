package services

import (
	"wakaf/features/partners/domain"

	"go.uber.org/zap"
)

type PartnersServices struct {
	PartnerRepo domain.RepoInterface
	logger      *zap.Logger
}

func New(data domain.RepoInterface, logger *zap.Logger) domain.UseCaseInterface {
	return &PartnersServices{
		PartnerRepo: data,
		logger:      logger,
	}
}
