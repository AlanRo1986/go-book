package app

import (
	"book/tmpl"
	"book/util"
	"net/http"
)

func init() {

}

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("init")
//
//	_, err := fmt.Fprint(w, "Hello, World!")
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	util.Output(w, tmpl.Index(r), util.PUT_HTML)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	util.Output(w, tmpl.SignIn(r), util.PUT_HTML)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	util.Output(w, tmpl.SignUp(r), util.PUT_HTML)
}

func signInSaveHandler(w http.ResponseWriter, r *http.Request) {
	util.Output(w, tmpl.SignInSave(r), util.PUT_JSON)
}

func signUpSaveHandler(w http.ResponseWriter, r *http.Request) {
	util.Output(w, tmpl.SignUpSave(r), util.PUT_JSON)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	util.Output(w, tmpl.AddArticel(r), util.PUT_HTML)
}
func addSaveHandler(w http.ResponseWriter, r *http.Request) {
	util.Output(w, tmpl.AddSaveArticel(r), util.PUT_JSON)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	util.Output(w, tmpl.GetArticle(r), util.PUT_HTML)
}
