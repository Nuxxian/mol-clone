package database

import (
	"context"
	"fmt"
	"log"
	"nuxxian/mol-clone/server/models"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
	// models "nuxxian/mol-clone/server/models"
)

var (
	database   *db.Client
	ctx        context.Context
)


func Database() {
	ctx = context.Background()
	opt := option.WithCredentialsFile("demol-clone-firebase-adminsdk-1bqp9-0cc265b15f.json")
	config := &firebase.Config{ProjectID: "demol-clone", DatabaseURL: "https://demol-clone-default-rtdb.europe-west1.firebasedatabase.app/"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	database, err = app.Database(ctx)
	if err != nil {
		fmt.Println("error creating database", err.Error())
		return
	}


    // AddEntry(models.QuizEntry{Name:"Jorne"}, "-O0QH6QPabiKx1p6BMI8")
    fmt.Println(GetEntries("-O0QH6QPabiKx1p6BMI8"))
}

func AddQuiz(name string){
    newQuiz := models.Quiz{}
    newQuiz.Name = name
    newQuiz.Questions = []models.Question{}
    ref := database.NewRef("/")
    ref.Push(ctx, newQuiz)
}

func AddEntry(entry models.QuizEntry, ID string) error {
    ref := database.NewRef(ID + "/entries")
    _, err := ref.Push(ctx, entry)
    return err
}

func GetEntries(ID string) ([]models.QuizEntry, error) {
    ref := database.NewRef(ID + "/entries")
    results, err := ref.OrderByValue().GetOrdered(ctx)
    if err != nil {
        return []models.QuizEntry{}, err
    }
    var res []models.QuizEntry
    for _, r := range results {
        var d models.QuizEntry
        if err := r.Unmarshal(&d); err != nil {
            return []models.QuizEntry{}, err
        }
        res = append(res, d)
    }
    return res, nil
}



