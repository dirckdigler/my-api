package routers

import (
	"fmt"
	"net/http"
)

// IndexRoute provides welcome message.
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}
