package handlers

import (
	"net/http"

	"github.com/adyasa2004/Autonomous-Travel-Disruption-Concierge/internal/db"
	"github.com/adyasa2004/Autonomous-Travel-Disruption-Concierge/internal/validation"
)

// Health handles GET /health.
// Reports 200 only when Postgres is reachable (§12.5).
func Health(w http.ResponseWriter, r *http.Request) {
	if err := db.Ping(r.Context()); err != nil {
		validation.WriteError(w, http.StatusServiceUnavailable, "service_unavailable",
			"Database is not reachable.")
		return
	}

	validation.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
