package app

import (
	"net/http"
	"strings"
)

type route struct {
	path       string
	method     string
	authorized bool
	handler    func(http.ResponseWriter, *http.Request)
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	OPTION = "OPTION"
	HEADER = "HEADER"
)

var (
	routes map[string]route
)

func Router() {
	//http.HandleFunc("/", indexHandler)
	routes = map[string]route{}
	routes["/"] = route{path: "/", method: GET, authorized: false, handler: indexHandler}
	routes["/view/res/*"] = route{path: "/", method: GET, authorized: false, handler: resourcesHandler}
	routes["/user"] = route{path: "/user", method: GET, authorized: true, handler: indexHandler}
	routes["/add"] = route{path: "/add", method: GET, authorized: true, handler: addHandler}
	routes["/save"] = route{path: "/edit", method: POST, authorized: true, handler: addSaveHandler}
	routes["/view"] = route{path: "/view", method: GET, authorized: false, handler: viewHandler}
	routes["/sign/in"] = route{path: "/sign/up", method: GET, authorized: false, handler: signInHandler}
	routes["/sign/up"] = route{path: "/sign/up", method: GET, authorized: false, handler: signUpHandler}
	routes["/doSignIn"] = route{path: "/doSignIn", method: POST, authorized: false, handler: signInSaveHandler}
	routes["/doSignUp"] = route{path: "/doSignUp", method: POST, authorized: false, handler: signUpSaveHandler}
}

func NewRouter(key string) (r route, ok bool) {
	if strings.Contains(key, "/view/res/") {
		key = "/view/res/*"
	}
	r, err := routes[key]
	return r, err
}
