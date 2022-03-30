package routes

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"time"

	"github.com/mkstyl3/linkedin_jobs/helpers"
	"github.com/mkstyl3/linkedin_jobs/models"
	"github.com/mkstyl3/linkedin_jobs/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// If we put ID=0 on the entity's FK gorm creates them if not found.
var (
	CREATE_ENTITY_ID uint = 0
)

// Handlers alphabetically sorted
func addJobsGetHandler(w http.ResponseWriter, r *http.Request) {
	type Parcel struct {
		Chart         template.HTML
		CompanySizes  interface{}
		Schedules     interface{}
		EnglishLevels interface{}
	}
	chartTemplate := createBarChart(w)
	// Preload company sizes
	sizes := []models.CompanySize{}
	preloadFormData(&sizes)
	// Preload job schedules
	schedules := []models.Schedules{}
	preloadFormData(&schedules)
	// Preload english_levels
	englishLevels := []models.EnglishLevel{}
	preloadFormData(&englishLevels)

	parcel := Parcel{Chart: chartTemplate, CompanySizes: sizes, Schedules: schedules, EnglishLevels: englishLevels}
	utils.ExecuteTemplate(w, "add-jobs.html", parcel)
}

type Parcel struct {
	Msg    string
	Status int
	Err    error
}

func sendJsonResponse(w http.ResponseWriter, parcel Parcel) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(parcel.Status)
	parcel_json, err := json.Marshal(parcel)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(parcel_json))
}

func addJobPostHandler(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()

	err := helpers.DecodeJSONBody(w, r, &job)
	if err != nil {
		var mr *helpers.MalformedRequest
		if errors.As(err, &mr) {
			log.Error().Msg(mr.Error())
			sendJsonResponse(w, Parcel{Msg: mr.Error(), Status: mr.Status})
			// http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Error().Err(err).Msg("")
			// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			sendJsonResponse(w, Parcel{Msg: "Internal Server Error", Status: http.StatusInternalServerError, Err: err})
		}
		return
	}

	log.Info().Msgf("job: %+v", job)

	// Check that company id is correct in the case that user select one from the input dropdown,
	// and later type another name without selecting it.

	// In release of Go 1.19 probably we will use generics with struct's field accesses and pass multiple structs as arguments
	// https://github.com/golang/go/issues/48522

	var company models.Company

	if err := models.GetByAttr(&company, "name", job.Company.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			job.Company.ID = CREATE_ENTITY_ID
		}
	} else {
		job.Company.ID = company.ID
	}

	var publisher models.Publisher
	if err := models.GetByAttr(&publisher, "name", job.Publisher.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			job.Publisher.ID = CREATE_ENTITY_ID
		}
	} else {
		job.Publisher.ID = publisher.ID
	}

	for i := 0; i < len(job.ProgrammingSkills); i++ {
		var progSkillFromDb models.ProgrammingSkill
		if err := models.GetByAttr(&progSkillFromDb, "name", job.ProgrammingSkills[i].Name); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				job.ProgrammingSkills[i].ID = CREATE_ENTITY_ID
			}
		} else {
			job.ProgrammingSkills[i].ID = progSkillFromDb.ID
		}
	}

	for i := 0; i < len(job.PersonalSkills); i++ {
		var persSkillFromDb models.PersonalSkill
		if err := models.GetByAttr(&persSkillFromDb, "name", job.PersonalSkills[i].Name); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				job.PersonalSkills[i].ID = CREATE_ENTITY_ID
			}
		} else {
			job.PersonalSkills[i].ID = persSkillFromDb.ID
		}
	}

	if err := models.Insert(&job); err != nil {
		log.Fatal().Msg("Job was not inserted")
		sendJsonResponse(w, Parcel{Status: http.StatusInternalServerError, Err: err})
	} else {
		successful_msg := "Job was successfully inserted"
		log.Info().Msg(successful_msg)
		sendJsonResponse(w, Parcel{Msg: successful_msg, Status: http.StatusOK, Err: err})
	}
}
