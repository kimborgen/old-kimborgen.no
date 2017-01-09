package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Clearence   int16
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
		0,
	},
	Route{
		"TodoIndex",
		"GET",
		"/api/todos",
		TodoIndex,
		1,
	},
	Route{
		"TodoShow",
		"GET",
		"/api/todos/{todoId}",
		TodoShow,
		1,
	},
	Route{
		"Login",
		"POST",
		"/api/login",
		Login,
		0,
	},
	Route{
		"Kollektiv",
		"GET",
		"/api/kollektiv/{name}",
		Kollektivet,
		0,
	},
	Route{
		"Kollektiv",
		"POST",
		"/api/kollektiv/update",
		KollektivetUpdate,
		0,
	},
	Route{
		"Kollektiv",
		"POST",
		"/api/kollektiv/new",
		KollektivetNew,
		0,
	},
}
