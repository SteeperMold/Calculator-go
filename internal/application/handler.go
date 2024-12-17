package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SteeperMold/Calculator-go/pkg/calculation"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	result, err := calculation.Calculate(request.Expression)
	if err != nil {
		if errors.Is(err, calculation.ErrInvalidExpression) {
			http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"result": %f}`, result)
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Address, nil)
}
