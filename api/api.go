package api

import (
	"Artify/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var (
	articles = map[string]database.Article{}
)

func createArticle(c *fiber.Ctx) error {
	article := new(database.Article)

	err := c.BodyParser(&article)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	article.ID = uuid.New().String()

	articles[article.ID] = *article

	c.Status(200).JSON(&fiber.Map{
		"article": article,
	})

	return nil
}

func readArticle(c *fiber.Ctx) error {
	id := c.Params("ID")

	if article, ok := articles[id]; ok {
		c.Status(200).JSON(&fiber.Map{
			"article": article,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
	}
	return nil
}

func readArticles(c *fiber.Ctx) error {
	c.Status(200).JSON(&fiber.Map{
		"article": articles,
	})
	return nil
}

func updateArticles(c *fiber.Ctx) error {
	updatedArticle := new(database.Article)

	err := c.BodyParser(updatedArticle)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	id := c.Params("ID")
	if article, ok := articles[id]; ok {
		article.Title = updatedArticle.Title
		article.Description = updatedArticle.Description
		article.Ratings = updatedArticle.Ratings
		articles[id] = article
		c.Status(200).JSON(&fiber.Map{
			"article": article,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
	}

	return nil
}

func deleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, ok := articles[id]; ok {
		delete(articles, id)
		c.Status(200).JSON(&fiber.Map{
			"message": "articles deleted successfully",
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
	}

	return nil
}

func SetUpRoute() *fiber.App {
	app := *fiber.New()
	app.Post("/api/v1/articles", createArticle)
	app.Get("/api/v1/articles/:id", readArticle)
	app.Get("/api/v1/articles/", readArticles)
	app.Put("/api/v1/articles/:id", updateArticles)
	app.Delete("/api/v1/articles/:id", deleteArticle)
	return &app
}
