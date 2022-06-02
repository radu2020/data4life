package token

import (
	"../utils"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const (
	spaceLength = 30
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// generateToken takes length of string n and returns a random generated token
func generateToken(tl int) string {
	b := make([]rune, tl)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}



// space just prints a spacing element to std out
func space() {
	fmt.Printf("%s\n", strings.Repeat("-", spaceLength))
}

// GenerateTokens creates a file which contains randomly generated tokens.
// fn - file name tl - token length ta - token amount bs - batch size
// dbs - db batch size sd - sleep duration
func GenerateTokens(fn string, tl, ta, bs, dbs, sd int) {
	space()
	fmt.Printf("" +
		"FileName:\t%s\n" +
		"TokenLength:\t%d\n" +
		"TokenAmount:\t%d\n" +
		"BatchSize:\t%d\n" +
		"DBBatchSize:\t%d\n" +
		"sleepDuration:\t%d\n",
		fn, tl, ta, bs, dbs, sd)
	space()

	// create file
	f, err := os.Create(fn)
	utils.Must(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	// Generate tokens and write them to file
	// Print out progress
	fmt.Println("Writing to file progress:")
	space()
	for i:=1; i<=ta; i++ {
		tokenStr := generateToken(tl) + "\n"
		_, err := w.WriteString(tokenStr)
		utils.Must(err)

		// batch process tokens
		if i % bs == 0 {
			fmt.Printf("%d/%d\n", i, ta)
		}

		// process last batch if batch number not equal to total amount
		if i % bs != 0 && i == ta {
			fmt.Printf("%d/%d\n", i, ta)
		}

		if i == ta {
			err = w.Flush()
			utils.Must(err)
			space()
			fmt.Println("Finished writing to file.")
			space()
		}
	}
}