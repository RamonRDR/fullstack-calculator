package httpapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/RamonRDR/fullstack-calculator/backend/internal/calculator"
)

type calculationRequest struct {
	Operation string   `json:"operation"`
	A         *float64 `json:"a"`
	B         *float64 `json:"b"`
}

type resultResponse struct {
	Result float64 `json:"result"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// NewHandler returns the HTTP routes exposed by the calculator service.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/calculate", calculateHandler)
	return mux
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var request calculationRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if request.Operation == "" {
		writeError(w, http.StatusBadRequest, "operation is required")
		return
	}
	if request.A == nil || request.B == nil {
		writeError(w, http.StatusBadRequest, "both operands are required")
		return
	}

	result, err := calculator.Calculate(request.Operation, *request.A, *request.B)
	if err != nil {
		switch {
		case errors.Is(err, calculator.ErrDivisionByZero),
			errors.Is(err, calculator.ErrUnsupportedOperation),
			errors.Is(err, calculator.ErrInvalidOperand),
			errors.Is(err, calculator.ErrInvalidResult):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	writeJSON(w, http.StatusOK, resultResponse{Result: result})
}

func decodeJSON(r *http.Request, destination any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(destination); err != nil {
		return err
	}

	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("request body must contain a single JSON object")
	}

	return nil
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
