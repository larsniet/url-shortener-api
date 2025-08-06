package health

import (
	"fmt"
	"net/http"
)

// HealthCheckHandler godoc
// @Summary Check if the server is running
// @Tags Health Check
// @Description Returns a simple message indicating the server is up and running
// @Accept json
// @Produce text/plain
// @Success 200 {string} string "Server is up and running!"
// @Router /health-check [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running!")
}
