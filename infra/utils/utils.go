package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// ParseHTML returns a HTML at string form
func ParseHTML(path string) (body string, err error) {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", wd, path))
	if err != nil {
		return body, err
	}
	body = string(dat[:])

	return body, nil
}

// ReadCSFile retorn csv file parsed
func ReadCSFile(path string) (entities []string, err error) {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path = fmt.Sprintf("%s/%s", wd, path)
	csvfile, err := os.Open(path)
	if err != nil {
		return entities, err
	}

	r := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		entities = append(entities, record[0])
	}

	return entities, err
}
