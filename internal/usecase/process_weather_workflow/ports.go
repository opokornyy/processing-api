package processweatherworkflow

import (
	"processing-api/internal/domain/recommendation"
)

// NotificationService defines the interface for sending recommendations
type NotificationService interface {
	SendRecommendation(rec *recommendation.Recommendation) error
}
