package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/YoNoSoyVictor/ThoughtBox/pkg/models"
	validations "github.com/YoNoSoyVictor/ThoughtBox/pkg/validations/entities"
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

	//render webpage
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

	//render webpage
	a.render(w, r, "post.page.tmpl", &TemplateData{
		Post: post,
	})
}

func (a *app) create(w http.ResponseWriter, r *http.Request) {

	//checking method

	//post
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			a.clientError(w, http.StatusBadRequest)
			return
		}

		title := r.PostForm.Get("title")
		if len(strings.TrimSpace(title)) < 5 || len(strings.TrimSpace(title)) > 100 {
			a.session.Put(r, "flash", "Title must be between 5 and 100 characters")
			http.Redirect(w, r, "/post/create", http.StatusSeeOther)
			return
		}

		message := r.PostForm.Get("message")
		if len(strings.TrimSpace(message)) < 5 || len(strings.TrimSpace(message)) > 1000 {
			a.session.Put(r, "flash", "Message must be between 5 and 1000 characters")
			http.Redirect(w, r, "/post/create", http.StatusSeeOther)
			return
		}

		id, err := a.posts.Insert(title, message)
		if err != nil {
			a.serverError(w, err)
			return
		}

		a.session.Put(r, "flash", "Post successfully created!")
		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", id), http.StatusSeeOther)
		return

	}

	//get
	a.render(w, r, "create.page.tmpl", &TemplateData{})
}

func (a *app) signup(w http.ResponseWriter, r *http.Request) {

	//post
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			a.clientError(w, http.StatusBadRequest)
			return
		}

		name := r.PostForm.Get("name")
		if validations.ValidateName(name) != nil {
			a.session.Put(r, "flash", "Error on field: name")
			http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
			return
		}

		email := r.PostForm.Get("email")
		if validations.ValidateMail(email) != nil {
			a.session.Put(r, "flash", "Error on field: email")
			http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
			return
		}

		password := r.PostForm.Get("password")
		if validations.ValidatePassword(password) != nil {
			a.session.Put(r, "flash", "Error on field: password")
			http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
			return
		}

		err = a.users.Signup(name, email, password)
		if err != nil {

			if err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
				a.session.Put(r, "flash", "Email already in use")
				http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
				return
			}

			a.serverError(w, err)
		}

		a.session.Put(r, "flash", "User created succesfully!")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}

	//get
	a.render(w, r, "signup.page.tmpl", &TemplateData{})
}

func (a *app) login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			a.clientError(w, http.StatusBadRequest)
			return
		}

		email := r.PostForm.Get("email")
		if validations.ValidateMail(email) != nil {
			a.session.Put(r, "flash", "Error on field: email")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}

		password := r.PostForm.Get("password")

		id, err := a.users.Login(email, password)
		if err != nil {
			if errors.Is(err, models.ErrInvalidCredentials) {
				a.session.Put(r, "flash", "Invalid credentials")
				http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			} else {
				a.serverError(w, err)
			}
			return
		}

		a.session.Put(r, "authenticatedUserID", id)
		a.session.Put(r, "flash", "login succesfull")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}

	a.render(w, r, "login.page.tmpl", &TemplateData{})
}

func (a *app) logout(w http.ResponseWriter, r *http.Request) {
	a.session.Remove(r, "authenticatedUserID")
	a.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
