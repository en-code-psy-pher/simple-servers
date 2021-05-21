package handlers 

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"errors"
)

// Note:
// functions only handle square matrices

// first class function, pass around as type
// all matrix actions have the same function signature
type action func(records [][]string) (response string)

// reads the matrix from a .csv file
func readMatrix(r *http.Request) ([][]string, error) {
    file, _, err := r.FormFile("file")

    if err != nil {
	return nil, err 
    }

    defer file.Close()

    var records [][]string
    records, err = csv.NewReader(file).ReadAll()

    if err != nil {
	return nil, err 
    }

    if len(records) == 0 {
	return nil, errors.New("no matrix")
    }

    if len(records) != len(records[0]) {
	return nil, errors.New("not a square matrix")
    }

    return records, nil
}

// handler function for all matrix actions
func matrixHandler(w http.ResponseWriter, r *http.Request, matrixAction action) {
    records, err := readMatrix(r)

    if err != nil {
	    w.Write([]byte(fmt.Sprintf("error %s", err.Error())))

	    return
    }

    response := matrixAction(records)

    fmt.Fprint(w, response)
}

// prints the matrix in string format
// output:
// 1,2,3
// 4,5,6
// 7,8,9
func echo(records [][]string) (response string) {
    for _, row := range records {
	    response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
    }

    return
}

func EchoMatrix(w http.ResponseWriter, r *http.Request) {
    matrixHandler(w, r, echo)
}

// prints the inverted matrix in string format
// output:
// 1,4,7
// 2,5,8
// 3,6,9

// transposes or inverts a string matrix
// rows become columns, and columns become rows
func transpose(records [][]string) [][]string {
    newRecords := make([][]string, len(records))

    for i := 0; i < len(records); i++ {
        for j := 0; j < len(records[0]); j++ {
            newRecords[j] = append(newRecords[j], records[i][j])
        }
    }
    return newRecords
}

func invert(records [][]string) (response string) {
    newRecords := transpose(records)

    for _, row := range newRecords {
	    response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
    }

    return
}

func InvertMatrix(w http.ResponseWriter, r *http.Request) {
    matrixHandler(w, r, invert)
}

// prints the matrix inline in string format
// output:
// 1,2,3,4,5,6,7,8,9
func flatten(records [][]string) (response string) {
    for _, row := range records {
	    for _, val := range row {
		    if response == "" {
			response = fmt.Sprintf("%s", val)
		    } else {
			response = fmt.Sprintf("%s,%s", response, val)
		    }
	    }
    }

    return
}

func FlattenMatrix(w http.ResponseWriter, r *http.Request) {
    matrixHandler(w, r, flatten)
}

// prints the sum of all integers in the matrix
func sum(records [][]string) (response string) {
    result := 0

    for _, row := range records {
	    for _, val := range row {
		    i, _ := strconv.Atoi(val)
		    result += i 
	    }
    }

    response = strconv.Itoa(result)

    return
}

func SumMatrix(w http.ResponseWriter, r *http.Request) {
    matrixHandler(w, r, sum)
}

// prints the product of all integers in the matrix
// moving from left to right off each row
func product(records [][]string) (response string) {
    result := 1

    for _, row := range records {
	    for _, val := range row {
		    i, _ := strconv.Atoi(val)
		    result *= i 
	    }
    }

    response = strconv.Itoa(result)

    return
}

func ProductMatrix(w http.ResponseWriter, r *http.Request) {
    matrixHandler(w, r, product)
}
