package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
)

type CliArgs struct {
	wordCount int
	lang      string
}

func init() {
	assertAvailablePRNG()
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

func main() {
	args := readCliArgs()
	bip39 := readBip39Words(args)
	words := generateRandomWords(args, bip39)
	for _, word := range words {
		fmt.Println(word)
	}
}

func generateRandomWords(args CliArgs, bip39 []string) []string {
	var words []string
	for i := 0; i < args.wordCount; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(bip39))))
		if err != nil {
			panic(err)
		}
		words = append(words, bip39[num.Int64()])
	}
	return words
}

func readCliArgs() CliArgs {
	args := CliArgs{}
	flag.IntVar(&args.wordCount, "n", 12, "Specify number of BIP39 words to generate. Default is 12.")
	flag.StringVar(&args.lang, "l", "english", "Specify language from which BIP39 list to select from. See list of available languages at https://github.com/bitcoin/bips/tree/master/bip-0039. Default is english.")
	flag.Parse()
	return args
}

func readBip39Words(args CliArgs) []string {
	readFile, err := os.Open(fmt.Sprintf("bips/bip-0039/%v.txt", args.lang))
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var bip39 []string
	for fileScanner.Scan() {
		bip39 = append(bip39, fileScanner.Text())
	}
	readFile.Close()
	return bip39
}
