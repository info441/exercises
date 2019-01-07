package models

import (
	"encoding/csv"
	"fmt"
	"io"
)

//Zip represents a zip code record.
//The `json:"..."` field tags allow us to change
//the name of the field when it is encoded into JSON
//see https://drstearns.github.io/tutorials/gojson/
type Zip struct {
	Code  string `json:"code,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

//ZipSlice is a slice of *Zip,
//that is, pointers to Zip struct instances
type ZipSlice []*Zip

//ZipIndex is a map from strings to ZipSlices
type ZipIndex map[string]ZipSlice

//LoadZips loads the zip code records from CSV stream
//returning a ZipSlice or an error. The expectedNumber
//should be set to the expected number of records, or
//zero if you don't know how many records there will be.
func LoadZips(r io.Reader, expectedNumber int) (ZipSlice, error) {
	//create a new CSV reader over the input stream
	reader := csv.NewReader(r)

	//read the header row and ensure that there are at least 7 fields
	fields, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}
	if len(fields) < 7 {
		return nil, fmt.Errorf("CSV has %d fields but we require %d", len(fields), 7)
	}

	//make a slice of *Zip with enough capacity
	//to hold expectedNumber. This is just an
	//optimization so that we can avoid reallocations
	//as we append to the slice
	zips := make(ZipSlice, 0, expectedNumber)

	//loop until we return...
	for {
		//read a row from the CSV file
		fields, err := reader.Read()
		//if we got io.EOF as an error, we are
		//done reading the input stream (End Of File)
		if err == io.EOF {
			return zips, nil
		}
		//if we got some other error, return it
		if err != nil {
			return nil, fmt.Errorf("Error parsing CSV: %v", err)
		}
		//create a *Zip and initialize the fields
		// These numbers represent the respective fields that we are interested in
		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		//append the *Zip to the slice
		zips = append(zips, z)
	}
}
