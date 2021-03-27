package models

import (
	"portal/database"
	"time"
)

type Page struct {
	ID          uint      `json:id`
	Href        string    `json:"href"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	Created     time.Time `json:"created"`
}

func (p *Page) SaveNewPage() error {
	db := database.Connect(&database.SQLConfig{
		User:     "root",
		Password: "password",
		Database: "isak_tech_portal",
	})
	defer db.Close()

	_, err := db.Exec("INSERT INTO pages (href, description, thumbnail) VALUES (?, ?, ?)",
		p.Href, p.Description, p.Thumbnail)

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
			&page.Href,
			&page.Description,
			&page.Thumbnail,
			&page.Created); err != nil {
			return nil, err
		}

		pages = append(pages, page)
	}

	return pages, nil
}
