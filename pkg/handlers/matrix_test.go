package handlers 

import (
	"os"
	"encoding/csv"
	"testing"
)

func readCSV() (records [][]string, err error) {
    var file *os.File
    file, err = os.Open("test.csv")

    if err != nil {
	return
    }

    records, err = csv.NewReader(file).ReadAll()

    if err != nil {
	return nil, err 
    }

    return
}

func TestEcho(t *testing.T) {
    records, err := readCSV()

    if err != nil {
	t.Errorf("error is %s", err.Error())
    }

    if len(records) == 0 {
	t.Errorf("no matrix")
    }

    if len(records) != len(records[0]) {
	t.Errorf("not a square matrix")
    }
	
    echo(records)
}

func TestInvert(t *testing.T) {
    records, err := readCSV()

    if err != nil {
	t.Errorf("error is %s", err.Error())
    }

    if len(records) == 0 {
	t.Errorf("no matrix")
    }

    if len(records) != len(records[0]) {
	t.Errorf("not a square matrix")
    }
	
    invert(records)
}

func TestFlatten(t *testing.T) {
    records, err := readCSV()

    if err != nil {
	t.Errorf("error is %s", err.Error())
    }

    if len(records) == 0 {
	t.Errorf("no matrix")
    }

    if len(records) != len(records[0]) {
	t.Errorf("not a square matrix")
    }
	
    flatten(records)
}

func TestSum(t *testing.T) {
    records, err := readCSV()

    if err != nil {
	t.Errorf("error is %s", err.Error())
    }

    if len(records) == 0 {
	t.Errorf("no matrix")
    }

    if len(records) != len(records[0]) {
	t.Errorf("not a square matrix")
    }
	
    sum(records)
}

func TestProduct(t *testing.T) {
    records, err := readCSV()

    if err != nil {
	t.Errorf("error is %s", err.Error())
    }

    if len(records) == 0 {
	t.Errorf("no matrix")
    }

    if len(records) != len(records[0]) {
	t.Errorf("not a square matrix")
    }
	
    product(records)
}
