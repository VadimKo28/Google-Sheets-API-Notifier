package domain

type GoogleSheetElement struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Complete bool `json:"complete"`
}