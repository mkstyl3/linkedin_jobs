# README
This is a personal project developed for learning purposes. Looks for integrate and display job positions offered by other job-seach websites in a more friendly way through graphical charts-

Status: Proof of concept.

Below useful snippets. 

Snipped for create table "users" by migrations and populate with users:

```go
cost := bcrypt.DefaultCost
password := "t43f34ffb78j89"
hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
if err != nil {
    log.Fatal().Err(err).Str("service", "encrypt_service").Msgf("Cannot encrypt pw")
}

var (
    people = []models.User{
        {Name: "marc", Password: hash},
        {Name: "lavanya", Password: hash},
    }
)

// create User table
db.AutoMigrate(&models.User{})
// Create user marc and lavanya
db.Create(&people)
```

Snippet for create multiple Experiences 

```go 
// Connect to DB
	err = models.ConnectToDb()
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to db.")
	} else {
		log.Info().Msg("Successfully connected to db")
	}

	var xps []models.Experience
	for i := uint(4); i < 100; i++ {
		//debo llenar la lista
		xps = append(xps, models.Experience{Years: i})
	}
	if err := models.Db.Create(xps).Error; err != nil {
		fmt.Printf("tmp")
	}
```
