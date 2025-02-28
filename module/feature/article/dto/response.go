package dto

import (
	"github.com/capstone-kelompok-7/backend-disappear/module/entities"
	"time"
)

type ArticleFormatter struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Photo     string    `json:"photo"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatArticle(article entities.Articles) ArticleFormatter {
	articleFormatter := ArticleFormatter{}
	articleFormatter.ID = article.ID
	articleFormatter.Title = article.Title
	articleFormatter.Content = article.Content
	articleFormatter.Photo = article.Photo
	articleFormatter.CreatedAt = article.CreatedAt

	return articleFormatter
}

func FormatterArticle(articles []entities.Articles) []ArticleFormatter {
	var articleFormatter []ArticleFormatter

	for _, article := range articles {
		formatArticle := FormatArticle(article)
		articleFormatter = append(articleFormatter, formatArticle)
	}

	return articleFormatter
}
