package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "golang.org/x/oauth2"
	_ "golang.org/x/oauth2/github"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/wiki/home", http.StatusMovedPermanently)
}

var (
	host = flag.String("b", "0.0.0.0", "bind address")
	port = flag.Int("p", 8080, "port to listen to")
)

func main() {
	flag.Parse()

	cfg := &Config{RepoRoot: "./testrepo"}
	SetConfig(cfg)

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/pages/new", NewPageHandler).Methods("GET")
	r.HandleFunc("/pages/new/{page}", NewNamedPageHandler).Methods("GET")
	r.HandleFunc("/pages/{page}/edit", EditPageHandler).Methods("GET")
	r.HandleFunc("/pages", SavePageHandler).Methods("POST")
	r.HandleFunc("/pages/{page}", UpdatePageHandler).Methods("POST")
	r.HandleFunc("/pages/{page}", DeletePageHandler).Methods("DELETE")
	r.HandleFunc("/pages/{page}/revert/{version}", RevertPageHandler).Methods("GET")

	r.HandleFunc("/wiki", ListPagesHandler).Methods("GET")
	r.HandleFunc("/wiki/{page}", ShowWikiHandler).Methods("GET")
	r.HandleFunc("/wiki/{page}/history", WikiHistoryHandler).Methods("GET")
	r.HandleFunc("/wiki/{page}/{version}", WikiVersionHandler).Methods("GET")
	r.HandleFunc("/wiki/{page}/compare/{revisions}", WikiCompareHandler).Methods("GET")

	r1 := r.PathPrefix("/auth").Subrouter()
	r1.HandleFunc("/login", LoginHandler)
	r1.HandleFunc("/callback", CallbackHandler)
	r1.HandleFunc("/logout", LogoutHandler)

	mux1 := http.NewServeMux()
	mux1.Handle("/", negroni.New(
		negroni.HandlerFunc(AuthMiddleware),
		negroni.Wrap(r),
	))

	mux1.Handle("/auth/", negroni.New(
		negroni.HandlerFunc(LoggingMiddleware),
		negroni.Wrap(r),
	))

	n := negroni.Classic()
	n.UseHandler(mux1)

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("starting tethys @", addr)
	err := http.ListenAndServe(addr, n)

	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
