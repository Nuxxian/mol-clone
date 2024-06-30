package models

type Question struct {
    Prompt string `json:"prompt"`
    Options []string `json:"options"`
    Answer []string `json:"answer"`
}

type Quiz struct {
    Name string `json:"name"`
    Questions []Question `json:"questions"`
}

type QuizEntry struct {
    Name string `json:"name"`
    Answers []string `json:"answers"`
    Time int `json:"time"`
    Date string `json:"date"`
}
