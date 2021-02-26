package models

import (
	"portal/database"
	"time"
)

type Page struct {
	ID          uint      `json:id`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imgUrl"`
	Created     time.Time `json:"created"`
}

func (p *Page) SaveNewPage() error {
	db := database.Connect(&database.SQLConfig{
		User:     "root",
		Password: "password",
		Database: "isak_tech_portal",
	})
	defer db.Close()

	_, err := db.Exec("INSERT INTO pages (url, description, imageurl) VALUES (?, ?, ?)",
		p.URL, p.Description, p.ImageURL)

	return err
}

func RemovePage(ID int64) error {
	db := database.Connect(&database.SQLConfig{
		User:     "root",
		Password: "password",
		Database: "isak_tech_portal",
	})
	defer db.Close()

	_, err := db.Exec("DELETE FROM pages WHERE id = ?", ID)

	return err
}

func GetAllPages() ([]Page, error) {
	db := database.Connect(&database.SQLConfig{
		User:     "root",
		Password: "password",
		Database: "isak_tech_portal",
	})
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pages")

	if err != nil {
		return nil, err
	}

	var pages []Page

	for rows.Next() {
		var page Page

		if err := rows.Scan(
			&page.ID,
			&page.URL,
			&page.Description,
			&page.ImageURL,
			&page.Created); err != nil {
			return nil, err
		}

		pages = append(pages, page)
	}

	return pages, nil
}
