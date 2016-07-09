package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewPageHandler(w http.ResponseWriter, r *http.Request) {
	getTemplate(NEW_PAGE_TEMPLATE).Execute(w, nil)
}

func NewNamedPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["page"]
	title := titleFromName(name)
	getTemplate(NEW_NAMED_PAGE_TEMPLATE).Execute(w, title)
}

func EditPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["page"]
	page, err := LoadPage(name)

	if err != nil {
		log.Println("ERROR loading page:", err)
		http.Redirect(w, r, "/pages/new", http.StatusMovedPermanently)
	} else {
		getTemplate(EDIT_PAGE_TEMPLATE).Execute(w, page)
	}
}

func SavePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SavePageHandler", r.FormValue("title"), r.FormValue("content"))
	title := r.FormValue("title")
	content := r.FormValue("content")
	name := nameFromTitle(title)
	page := &Page{Name: name, Title: title, Content: []byte(content)}
	err := page.Save()

	if err != nil {
		getTemplate(ERROR_PAGE_TEMPLATE).Execute(w, err)
	} else {
		http.Redirect(w, r, page.Url(), http.StatusMovedPermanently)
	}
}

func UpdatePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdatePageHandler")
	vars := mux.Vars(r)
	name := vars["page"]
	page, err := LoadPage(name)

	if err != nil {
		getTemplate(ERROR_PAGE_TEMPLATE).Execute(w, err)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	err = page.Update(title, content)

	if err != nil {
		getTemplate(ERROR_PAGE_TEMPLATE).Execute(w, err)
		return
	}

	http.Redirect(w, r, page.Url(), http.StatusMovedPermanently)
}

func DeletePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeletePageHandler")
}

func RevertPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RevertPageHandler")
}

func ListPagesHandler(w http.ResponseWriter, r *http.Request) {

}

func ShowWikiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["page"]
	log.Println("looking for page:", name)
	page, err := LoadPage(name)

	if err != nil {
		log.Println("ERROR loading page:", err)
		newPageUrl := fmt.Sprintf("/pages/new/%s", name)
		http.Redirect(w, r, newPageUrl, http.StatusMovedPermanently)
	} else {
		getTemplate(SHOW_WIKI_PAGE_TEMPLATE).Execute(w, page)
	}
}

func WikiHistoryHandler(w http.ResponseWriter, r *http.Request) {

}

func WikiVersionHandler(w http.ResponseWriter, r *http.Request) {

}

func WikiCompareHandler(w http.ResponseWriter, r *http.Request) {

}
