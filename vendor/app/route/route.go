package route

import (
	"net/http"

	"app/controller"
	hr "app/route/httprouterwrapper"
	"app/route/logrequest"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func LoadHTTP() http.Handler {
	return middleware(routes())
}

func routes() *httprouter.Router {
	r := httprouter.New()

	r.NotFound = alice.
	New().ThenFunc(controller.Error404)

	r.POST("/api/users", hr.Handler(alice.New().ThenFunc(controller.Signup)))

	r.POST("/api/users/stats", hr.Handler(alice.New().ThenFunc(controller.NewStat)))

	r.GET("/api/users/stats/top", hr.Handler(alice.New().ThenFunc(controller.GetTop)))

	return r
}

func middleware(h http.Handler) http.Handler {

	h = logrequest.Handler(h)
	h = context.ClearHandler(h)

	return h
}
