package app

import (
	"book/conf"
	"log"
	"net/http"
	"sync"
)

const (
	ERROR_NOT_FOUND    = "ERROR_NOT_FOUND"
	ERROR_NOT_METHOD   = "ERROR_NOT_METHOD"
	ERROR_AUTH_INVALID = "ERROR_AUTH_INVALID"
)

var (
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func init() {
	wg.Add(100)
}

func HttpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header.Get("Accept"))
		log.Println(r.Header.Get("User-Agent"))
		log.Println(r.Header)
		log.Println(r.Proto, r.Host, r.Method, r.URL.Path)

		token, _ := r.Cookie("token")
		log.Println("token", token)

		//if strings.Index(r.URL.Path,"/view/res/") == 0 {
		//	resourcesHandler(w,r)
		//	return
		//}
		route, ok := NewRouter(r.URL.Path)
		if !ok {
			errorHandler(w, r, ERROR_NOT_FOUND)
			return
		}
		if r.Method != route.method {
			errorHandler(w, r, ERROR_NOT_METHOD)
			return
		}
		if route.authorized && token != nil && len(token.Value) < 32 {
			errorHandler(w, r, ERROR_AUTH_INVALID)
			return
		}

		route.handler(w, r)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, s string) {
	if s == ERROR_NOT_FOUND {
		http.NotFound(w, r)
		return
	}
	if s == ERROR_NOT_METHOD {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	http.Error(w, http.StatusText(http.StatusNonAuthoritativeInfo), http.StatusNonAuthoritativeInfo)
	return
}

func resourcesHandler(w http.ResponseWriter, r *http.Request) {
	filePath := conf.ROOT + string([]byte(r.URL.Path)[1:])

	//file, err := os.OpenFile(filePath, os.O_RDONLY, 066)
	//defer file.Close()
	//if err != nil {
	//	fmt.Println("not found", filePath)
	//	return
	//}
	http.ServeFile(w, r, filePath)
}
