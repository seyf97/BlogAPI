package models

import (
	"time"

	"github.com/seyf97/BlogAPI/db"
)

type Article struct {
	ID         int64
	Title      string `binding:"required"`
	Content    string `binding:"required"`
	CreatedAt  time.Time
	User_ID    int64
	Category   string
	LastEdited *time.Time //nil time.Time is not supported so we pass a ptr
}

func GetAllArticlesDB() ([]Article, error) {
	query := "SELECT * FROM articles"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID,
			&article.Title,
			&article.Content,
			&article.CreatedAt,
			&article.User_ID,
			&article.Category,
			&article.LastEdited)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}
	return articles, nil
}

func GetArticleByIdDB(article_id int64) (*Article, error) {
	query := "SELECT * FROM articles WHERE id = ?"

	row := db.DB.QueryRow(query, article_id)

	var article Article

	err := row.Scan(&article.ID,
		&article.Title,
		&article.Content,
		&article.CreatedAt,
		&article.User_ID,
		&article.Category,
		&article.LastEdited)
	if err != nil {
		return nil, err
	}

	return &article, nil

}

func (a Article) DeleteDB() error {
	// Only the ID from Article will be used.
	// Remaining fields are nill
	query := "DELETE FROM articles WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.ID)

	return err
}

func (a *Article) SaveDB() error {
	query := "INSERT INTO articles(title, content, created_at, user_id, category) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.Title, a.Content, a.CreatedAt, a.User_ID, a.Category)
	return err
}

func (a *Article) UpdateDB() error {
	query := `
	UPDATE articles 
	SET title = ?, content = ?, category = ?, last_edited = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.Title, a.Content, a.Category, a.LastEdited, a.ID)
	return err
}
