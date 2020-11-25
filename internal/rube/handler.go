package rube

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/pprof"
	"time"
)


func Handler() http.Handler {
	routes := mux.NewRouter()
	routes.StrictSlash(true)

	// routes go here

	return http.TimeoutHandler(Service(routes), 25*time.Second, "Timed out performing request")
}

func Service(h http.Handler) http.Handler {
	router := mux.NewRouter()
	router.PathPrefix("/debug/pprof").Handler(pprofHandler())
	// todo: add health, metrics, and swagger
	router.PathPrefix("/").Handler(h)
	return router
}

func pprofHandler() http.Handler {
	routes := mux.NewRouter()
	routes.HandleFunc("/debug/pprof/", pprof.Index)
	routes.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	routes.HandleFunc("/debug/pprof/profile", pprof.Profile)
	routes.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	routes.HandleFunc("/debug/pprof/trace", pprof.Trace)
	routes.HandleFunc("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	routes.HandleFunc("/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
	routes.HandleFunc("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	return routes
}