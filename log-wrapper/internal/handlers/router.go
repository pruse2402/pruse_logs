package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	_ "pruse_logs/log-wrapper/docs"
	"pruse_logs/log-wrapper/internal/config"
	"pruse_logs/log-wrapper/internal/config/globals"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

// GetRouter creates a router and registers all the routes for the
// service and returns it.
func GetRouter() http.Handler {
	router := httprouter.New()
	router.PanicHandler = PanicHandler
	setPingRoutes(router)
	setLogRoutes(router)
	setPrometheusRoutes(router)

	router.Handler("GET", "/swagger", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one/:two", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one/:two/:three", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one/:two/:three/:four", httpSwagger.WrapHandler)

	return router
}

func tokenAuthentication(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userName, pass, ok := r.BasicAuth()
		// Check if contains ewaybill service token in Basic Auth or as parameters
		if (ok && pass == config.SERVICE_TOKEN && userName == globals.COMPANY_NAME) || r.FormValue("token") == config.SERVICE_TOKEN {
			h(w, r, ps)
			return
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	log.Printf("Recovering from panic, Reason: %+v", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
