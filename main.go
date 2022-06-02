package main

import (
	"./token"
	"os"
)

const (
	tokenAmount = 1000 // tokens
	tokenLength = 7 // characters
	fileName = "tokens.txt"
    fileWriteBatchSize = 100 // tokens
    dbWriteBatchSize = 100 // records
	sleepDuration = 2 // seconds
)

func main () {
	os.Remove("./data4life.db") // Remove file on startup

	token.GenerateTokens(
		fileName,
		tokenLength,
		tokenAmount,
		fileWriteBatchSize,
		dbWriteBatchSize,
		sleepDuration)

	token.WriteTokensToDB(fileName, dbWriteBatchSize, sleepDuration)
}
