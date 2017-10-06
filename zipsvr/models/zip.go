package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Zip struct {
	Code  string // fields need to be exported because in model folder, so needs to be uppercase
	City  string
	State string
}

type ZipSlice []*Zip //* pointer to Zip struct

type ZipIndex map[string]ZipSlice

func LoadZips(filename string) (ZipSlice, error) { //return multiple functions
	f, err := os.Open(filename) //opens a file
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err) //%v appends the error message
	}
	reader := csv.NewReader(f)
	_, err = reader.Read() // err has already been declared
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	zips := make(ZipSlice, 0, 43000)
	for { // no conditions until we decide to exit
		fields, err := reader.Read()
		if err == io.EOF { //done with the input stream
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z)
	}
}
