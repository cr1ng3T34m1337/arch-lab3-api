package httpHandlers

import (
	"fmt"
	"net/http"

	"github.com/cr1ng3T34m1337/arch-lab3-api/api/dataProviders"
	"github.com/cr1ng3T34m1337/arch-lab3-api/tools"
)

func handleBalancersFunc(store *dataProviders.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			listBalancers(rw, r, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func listBalancers(rw http.ResponseWriter, r *http.Request, store *dataProviders.Store) {
	res, err := store.ListBalancers()
	if err != nil {
		fmt.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw, err)
		return
	}
	tools.WriteJsonOk(rw, &res)
}
