package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/adyasa2004/Autonomous-Travel-Disruption-Concierge/internal/auth"
	"github.com/adyasa2004/Autonomous-Travel-Disruption-Concierge/internal/db"
	"github.com/adyasa2004/Autonomous-Travel-Disruption-Concierge/internal/models"
	"github.com/adyasa2004/Autonomous-Travel-Disruption-Concierge/internal/validation"
)

// GetInsuranceClaims handles GET /v1/insurance-claims/:user_id.
func GetInsuranceClaims(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	if !auth.CheckOwnership(w, r, userID) {
		return
	}

	ctx := r.Context()

	rows, err := db.Pool.Query(ctx,
		`SELECT id, disruption_event_id, eligible, claim_type, amount, status, created_at
		 FROM insurance_claims
		 WHERE user_id = $1
		 ORDER BY created_at DESC`, userID)
	if err != nil {
		validation.WriteError(w, http.StatusInternalServerError, "internal_error", "Failed to load insurance claims.")
		return
	}
	defer rows.Close()

	var claims []models.InsuranceClaim
	for rows.Next() {
		var c models.InsuranceClaim
		if err := rows.Scan(
			&c.ID, &c.DisruptionEventID, &c.Eligible,
			&c.ClaimType, &c.Amount, &c.Status, &c.CreatedAt,
		); err != nil {
			validation.WriteError(w, http.StatusInternalServerError, "internal_error", "Failed to scan insurance claim.")
			return
		}
		claims = append(claims, c)
	}

	if claims == nil {
		claims = []models.InsuranceClaim{}
	}

	validation.WriteJSON(w, http.StatusOK, claims)
}
