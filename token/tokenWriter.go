package token

import (
	"../db"
	"../utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// fileToTokens opens a file and returns tokens from the file reader
// fp - filePath
func fileToTokens(fp string) ([]string, error) {
	f, err := os.Open(fp)
	utils.Must(err)
	defer f.Close()
	return tokensFromReader(f)
}

// tokensFromReader takes an io reader and returns a list of in-memory read tokens
// r - file read stream
func tokensFromReader(r io.Reader) ([]string, error) {
	var tokens []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tokens, nil
}

// histogram takes an array of tokens and returns a map of each token to its frequency
// tokens - tokens array
func histogram(tokens []string) map[string]int{
	tf := map[string]int{}
	for _, token := range tokens {
		if v, ok := tf[token]; ok {
			tf[token] = v+1
		} else {
			tf[token] = 1
		}
	}
	return tf
}

// dbWrite takes a token histogram, opens a DB conn and writes tokens to DB
// tf - token frequency histogram bs - transaction batch size s - sleep duration
func dbWrite(tf map[string]int, bs int, s int) {
	// Open DB Connection
	conn, err := db.OpenDbConn()
	utils.Must(err)
	defer conn.Close()

	// Create insert statement
	statement, err := conn.Prepare(
		"INSERT INTO tokens (" +
			"token, repeated, frequency" +
			") VALUES (?,?,?)")
	utils.Must(err)

	// Write tokens to DB and track progress
	fmt.Println("Writing to DB progress:")
	space()
	counter := 0
	for k, v := range tf {
		repeated := false
		frequency := 1
		if v != 1 {
			repeated = true
			frequency = v
		}

		// Insert record
		_, err = statement.Exec(k, repeated, frequency)
		utils.Must(err)

		// Sleep at interval to reduce load on the DB
		counter++
		if counter % bs == 0 {
			fmt.Printf("%d/%d\n", counter, len(tf))
			time.Sleep(time.Duration(s)*time.Second)
		}

		// Last batch
		if counter % bs != 0 && counter == len(tf) {
			fmt.Printf("%d/%d\n", counter, len(tf))
		}

		if counter == len(tf) {
			space()
			fmt.Println("Finished writing to database.")
			space()
		}
	}
}

// WriteTokensToDB reads tokens from file and batch writes data to DB
// path - file path bs - batch size
func WriteTokensToDB(path string, bs int, s int) {
	tokens, err := fileToTokens(path)
	utils.Must(err)
	dbWrite(histogram(tokens), bs, s)
}