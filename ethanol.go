package ethanol

import "net/http"

type EthanolRouter struct {
}

func (e *EthanolRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
