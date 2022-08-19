package handler

import (
	"go-projects/03.choose_your_own_adventure/model"
	"html/template"
	"net/http"
)

func HTMLHandler(pages map[string]model.Page, fallback http.Handler) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[1:]

		if page, ok := pages[path]; ok {
			t := template.Must(template.ParseFiles("page.html"))
			t.Execute(response, page)
		} else {
			fallback.ServeHTTP(response, request)
		}
	}
}
