package httpHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cr1ng3T34m1337/arch-lab3-api/api/dataProviders"
	"github.com/cr1ng3T34m1337/arch-lab3-api/tools"
)

type SetWorkingStatusRequest struct {
	Id            int    `json:"id"`
	WorkingStatus string `json:"workingStatus"`
}

func handleMachinesFunc(store *dataProviders.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			setMachineWorkingStatus(rw, r, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func setMachineWorkingStatus(rw http.ResponseWriter, r *http.Request, store *dataProviders.Store) {
	var req SetWorkingStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("Error decoding SetWorkingStatusRequest input: %s", err)
		tools.WriteJsonBadRequest(rw, fmt.Errorf("bad JSON payload: %s", err))
		return
	}
	if err := store.SetWorkingStatus(req.Id, req.WorkingStatus); err != nil {
		fmt.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw, err)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
}
