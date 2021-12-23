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
	verbose   bool
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
	numbers := generateRandomNumbers(args.wordCount, len(bip39))
	for _, num := range numbers {
		if args.verbose {
			fmt.Printf("%4v:", num)
		}
		fmt.Println(bip39[num])
	}
}

func generateRandomNumbers(count int, max int) []int {
	var numbers []int
	for i := 0; i < count; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, int(num.Int64()))
	}
	return numbers
}

func readCliArgs() CliArgs {
	args := CliArgs{}
	flag.IntVar(&args.wordCount, "n", 12, "Specify number of BIP39 words to generate. Default is 12.")
	flag.StringVar(&args.lang, "l", "english", "Specify language from which BIP39 list to select from. See list of available languages at https://github.com/bitcoin/bips/tree/master/bip-0039. Default is english.")
	flag.BoolVar(&args.verbose, "v", false, "Print indexes of each word in the BIP39 alongside the word. Default is false.")
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
