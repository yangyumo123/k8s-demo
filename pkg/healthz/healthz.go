package healthz

import (
	"net/http"
)

func init() {
	http.HandleFunc("/healthz", handleHealthz)
}
func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
