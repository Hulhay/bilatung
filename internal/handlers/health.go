package handlers

import (
	"bilatung/internal/models"
)

func (h *handler) GetHealtcheck() (*models.Health, error) {
	var health = &models.Health{
		Status: "UP",
	}

	return health, nil
}
