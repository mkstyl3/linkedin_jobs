package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/gorilla/mux"
	"github.com/mkstyl3/linkedin_jobs/middleware"
	"github.com/mkstyl3/linkedin_jobs/models"
	"github.com/mkstyl3/linkedin_jobs/sessions"
	"github.com/mkstyl3/linkedin_jobs/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func preloadFormData(data interface{}) {
	err := models.GetAll(data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info().Err(err).Msg(fmt.Sprint("No models found"))
		} else {
			log.Fatal().Err(err).Msg(fmt.Sprint("Error retrieving all models"))
		}
	}
	log.Info().Msg(fmt.Sprintf("Retrieved models: %+v", data))
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	r.HandleFunc("/", middleware.AuthRequired(indexPostHandler)).Methods("POST")
	r.HandleFunc("/add-job", addJobsGetHandler).Methods("GET")
	r.HandleFunc("/addjob", addJobPostHandler).Methods("POST")
	r.HandleFunc("/bar", barGetHandler).Methods("GET")
	r.HandleFunc("/companies", companyNamesGetHandler).Methods("GET")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/personal-skill-names", personalSkillNamesGetHandler).Methods("GET")
	r.HandleFunc("/programming-skill-names", programmingSkillNamesGetHandler).Methods("GET")
	r.HandleFunc("/publishers", publishersGetHandler).Methods("GET")
	fs := http.FileServer(http.Dir("./dist/"))
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))
	return r
}

func barGetHandler(w http.ResponseWriter, r *http.Request) {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first bar chart generated by go-echarts",
		Subtitle: "It's extremely easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	// Where the magic happens
	//bar.Render(w)
	bar.Renderer = utils.NewSnippetRenderer(bar, bar.Validate)
	// bar.Render(w)
	chart := utils.RenderToHtml(bar)

	utils.ExecuteTemplate(w, "index.html", chart)

}

func companyNamesGetHandler(w http.ResponseWriter, r *http.Request) {
	companies := []models.Company{}
	models.GetAll(&companies)
	// var company_names []string
	// for _, n := range companies {
	// 	company_names = append(company_names, n.Name)
	// }
	bytes, err := json.Marshal(companies)

	if err != nil {
		log.Error().Msg("Error Marshaling list")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func createBarChart(w http.ResponseWriter) template.HTML {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first bar chart generated by go-echarts",
		Subtitle: "It's extremely easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	// Where the magic happens
	//bar.Renderer = utils.NewSnippetRenderer(bar, bar.Validate)
	return utils.RenderToHtml(bar)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	// comments, err := models.GetCommentsDb()

	// if err != nil {
	// 	w.WriteHeader((http.StatusInternalServerError))
	// 	w.Write([]byte("Internal Server Error"))
	// 	return
	// }

	// create User table

	// t, err := template.ParseFiles("dist/index.html")
	// if err != nil {
	// 	return
	// }

	// err = t.Execute(f, comments)
	// if err != nil {
	// 	return
	// }
	// f.Close()

	// Output HTML
	// t, err := template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	return
	// }
	// f, err := os.Create("index.html")
	// err = t.Execute(f, nil)

	utils.ExecuteTemplate(w, "dashboard.html", nil)

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

func personalSkillNamesGetHandler(w http.ResponseWriter, r *http.Request) {
	all_models := []models.PersonalSkill{}
	models.GetAll(&all_models)
	bytes, err := json.Marshal(all_models)

	if err != nil {
		log.Error().Msg("Error Marshaling list")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func programmingSkillNamesGetHandler(w http.ResponseWriter, r *http.Request) {
	all_models := []models.ProgrammingSkill{}
	models.GetAll(&all_models)
	bytes, err := json.Marshal(all_models)

	if err != nil {
		log.Error().Msg("Error Marshaling list")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func publishersGetHandler(w http.ResponseWriter, r *http.Request) {
	all_models := []models.Publisher{}
	models.GetAll(&all_models)
	bytes, err := json.Marshal(all_models)

	if err != nil {
		log.Error().Msg("Error Marshaling list")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
