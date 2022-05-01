package chapter1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/* Count duplicate lines that are input from the standard input (stdin) */
func CountDuplicates1() {
	scanner := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)

	for i := 0; i < 10; i++ {
		scanner.Scan()
		counts[scanner.Text()]++
	}

	printCountsMap(counts)
}

/* Count duplicates from file
This function works in "streaming" mode, i.e. input is read and broken into lines as
needed.

Uses os.Open() to open the file.
*/
func CountDuplicates2() {
	counts := make(map[string]int)
	f, err := os.Open("E:\\Abhishek\\Go\\Misc\\TGPL\\src\\chapter-1\\ch1-3sample.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%s", err)
		return
	}

	countLinesInFile(f, counts)
	f.Close()
}

/* Count duplicates from file by reading the whole of it at once
Read the entire input into memory at once, split it into lines at once,
then process the lines.

Uses ioutil.ReadFile()
*/
func CountDuplicates3() {
	counts := make(map[string]int)
	byteData, err := ioutil.ReadFile("E:\\Abhishek\\Go\\Misc\\TGPL\\src\\chapter-1\\ch1-3sample.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%s", err)
		return
	}

	data := strings.Split(string(byteData), "\n")
	for _, line := range data {
		counts[line]++
	}
	printCountsMap(counts)
}

/***** Utility functions **********
An opened-file and stdin (e.g. from keyboard) both are of the type os.File
Hence, we can create a Scanner with either an opened-file or stdin as the parameter.
*/
func countLinesInFile(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	printCountsMap(counts)
}

func printCountsMap(counts map[string]int) {
	for line, count := range counts {
		fmt.Printf("%s\t%d\n", line, count)
	}
}
