package httpHandlers

import (
	"net/http"

	"github.com/cr1ng3T34m1337/arch-lab3-api/api/dataProviders"
)

func HandlersProvider(store *dataProviders.Store) map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/balancers": handleBalancersFunc(store),
		"/machines":  handleMachinesFunc(store),
	}
}
