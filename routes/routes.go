package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mkstyl3/linkedin_jobs/middleware"
	"github.com/mkstyl3/linkedin_jobs/models"
	"github.com/mkstyl3/linkedin_jobs/sessions"
	"github.com/mkstyl3/linkedin_jobs/utils"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	r.HandleFunc("/", middleware.AuthRequired(indexPostHandler)).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return r
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := models.GetCommentsDb()

	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte("Internal Server Error"))
		return
	}
	utils.ExecuteTemplate(w, "index.html", comments)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input_str := r.PostForm.Get("comment")
	_, err := models.PostCommentsDb(models.Comment{Text: input_str})

	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte("Internal Server Error"))
		return
	}
	http.Redirect(w, r, "/", 302)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.AuthenticateUser(username, password)
	if err != nil {
		switch err {
		case models.ErrUserNotFound:
			utils.ExecuteTemplate(w, "login.html", "unknown user")
		case models.ErrInvalidLogin:
			utils.ExecuteTemplate(w, "login.html", "invalid login")
		default:
			w.WriteHeader((http.StatusInternalServerError))
			w.Write([]byte("Internal Server Error"))
		}
		return
	}
	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
