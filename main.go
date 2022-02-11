// main.go
// A basic server

// main.go file belongs to the main package
package main

// the net/http and os packages are imported into the file
import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var client *redis.Client
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
var ctx = context.Background()
var templates *template.Template

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", AuthRequired(indexGetHandler)).Methods("GET")
	r.HandleFunc("/", AuthRequired(indexPostHandler)).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)

}

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		_, ok := session.Values["username"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	
	comments, err := client.LRange(ctx, "comments", 0, 10).Result()

	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte("Internal Server Error"))
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	err := client.LPush(ctx, "comments", comment).Err()
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte("Internal Server Error"))
		return
	}
	http.Redirect(w, r, "/", 302)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	hash, err := client.Get(ctx, "user:"+username).Bytes()
	if err == redis.Nil {
		templates.ExecuteTemplate(w, "login.html", "unknown user")
		return
	} else if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte("Internal Server Error"))
		return
	}
	
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		templates.ExecuteTemplate(w, "login.html", "invalid password")
	}
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte("Internal Server Error"))
		return
	}
	err = client.Set(ctx, "user:"+username, hash, 0).Err()
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte("Internal Server Error"))
		return
	}
	http.Redirect(w, r, "/login", 302)
}
