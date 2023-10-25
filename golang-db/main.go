package main

import (
	"context"
	"golang-db/handlers"
	"golang-db/repositories"
	"golang-db/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://bankthanakit:1234@cluster0.gwdj0mb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return client.Database("BookStore")

}

func main() {

	db := initMongo()
	bookRepo := repositories.NewBookDBRepository(db, "book_stock")

	bookSrv := services.NewBookService(bookRepo)

	bookHand := handlers.NewBookHandler(bookSrv)

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/books", bookHand.GetAllBook)
	app.Post("/book", bookHand.CreateBook)
	app.Put("/update/:id", bookHand.UpdateBook)
	app.Delete("/delete/:id", bookHand.DeleteBook)

	app.Listen(":3000")
}
