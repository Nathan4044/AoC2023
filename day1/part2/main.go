package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("../input.txt")
    // file, err := os.Open("test.txt")
    defer file.Close()

    if err != nil {
        log.Fatal(err)
    }

    contents := bufio.NewScanner(file)
    var result int64

    for contents.Scan() {
        number, err := extractNumber(contents.Text())

        if err != nil {
            log.Fatal(err)
        }

        result += number
    }

    fmt.Printf("%d\n", result)
}

func extractNumber(line string) (int64, error) {
    line = replaceWordsWithDigits(line)

    digits := []rune{}

    for _, c := range line {
        if isDigit(c) {
            digits = append(digits, c)
        }
    }

    if len(digits) < 1 {
        return 0, fmt.Errorf("not enough digits found (%s) from line: %s", string(digits), line)
    }

    textResult := string(digits[0]) + string(digits[len(digits)-1])
    result, err := strconv.ParseInt(textResult, 10, 64)

    if err != nil {
        return 0, fmt.Errorf("invalid number: got '%s' from '%s'", textResult, line)
    }

    return result, nil
}

func isDigit(c rune) bool {
    if '0' <= c && c <= '9' {
        return true
    }

    return false
}

func replaceWordsWithDigits(line string) string {

    var result string
    i := 0

    for i < len(line) {
        match, num := matchesNumber(line[i:])

        if match {
            result += num
        }

        i++
    }

    return result
}

func matchesNumber(line string) (bool, string) {
    numbers := map[string]string{
        "zero":  "0",
        "one":   "1",
        "two":   "2",
        "three": "3",
        "four":  "4",
        "five":  "5",
        "six":   "6",
        "seven": "7",
        "eight": "8",
        "nine":  "9",
        "0":     "0",
        "1":     "1",
        "2":     "2",
        "3":     "3",
        "4":     "4",
        "5":     "5",
        "6":     "6",
        "7":     "7",
        "8":     "8",
        "9":     "9",
    }

    for word, digit := range numbers {
        if len(line) < len(word) {
            continue
        }

        if word == line[:len(word)] {
            return true, digit
        }
    }

    return false, ""
}
