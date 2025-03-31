package healthcheck

import "net/http"

// HealthcheckHandler is a simple HTTP handler that responds with a 200 OK status and a JSON body indicating the service is healthy.
// It is typically used for health checks in a microservices architecture.
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
