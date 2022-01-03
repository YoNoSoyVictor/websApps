package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *app) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *app) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *app) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *app) render(w http.ResponseWriter, r *http.Request, name string, td *TemplateData) {

	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, td)
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}
