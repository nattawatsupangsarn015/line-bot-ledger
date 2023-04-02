package main

import (
	"log"
	"os"
	"testing"

	"example/line-bot-ledger/config"
	"example/line-bot-ledger/test"

	"github.com/joho/godotenv"
)

func Test(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	NODE_ENV := os.Getenv("NODE_ENV")

	log.Print("connecting db ..")
	config.ConnectMongoDB()
	log.Print("connecting db success !")

	test.TestRouterWithSuccess(t, NODE_ENV)

	// test.TestRouterAdminsWithSuccess(t, NODE_ENV)
	// test.TestRouterChabotWithSuccess(t, NODE_ENV)
	// test.TestRouterDishesWithSuccess(t, NODE_ENV)
	// test.TestRouterIngredientsWithSuccess(t, NODE_ENV)
	// test.TestRouterLineWithSuccess(t, NODE_ENV)
}
