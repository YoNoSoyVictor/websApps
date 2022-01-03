package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {

	//allow only "/" as home
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	//get 10 posts by id
	posts, err := a.posts.Latest(0)
	if err != nil {
		a.notFound(w)
		return
	}

	a.render(w, r, "home.page.tmpl", &TemplateData{
		Posts: posts,
	})
}

func (a *app) post(w http.ResponseWriter, r *http.Request) {

	//sanitize url
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		a.notFound(w)
		return
	}

	//get post by id
	post, err := a.posts.Get(id)
	if err != nil {
		a.notFound(w)
		return
	}

	a.render(w, r, "post.page.tmpl", &TemplateData{
		Post: post,
	})
}

func (a *app) create(w http.ResponseWriter, r *http.Request) {

	//checking method
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			a.clientError(w, http.StatusBadRequest)
			return
		}

		title := r.PostForm.Get("title")
		if len(strings.TrimSpace(title)) < 5 || len(strings.TrimSpace(title)) > 100 {
			http.Redirect(w, r, "/post/create", http.StatusSeeOther)
			return
		}

		message := r.PostForm.Get("message")
		if len(strings.TrimSpace(message)) < 5 || len(strings.TrimSpace(message)) > 1000 {
			http.Redirect(w, r, "/post/create", http.StatusSeeOther)
			return
		}

		id, err := a.posts.Insert(title, message)
		if err != nil {
			a.serverError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", id), http.StatusSeeOther)
	}

	a.render(w, r, "create.page.tmpl", nil)

}
