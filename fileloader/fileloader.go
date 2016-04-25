package fileloader

import (
	"encoding/xml"
	"fmt"
	"github.com/guardian/fakepicdar/schema"
	"io/ioutil"
	"log"
)

func ReadResponse(filename string) (record schema.Record) {
	v := schema.MogulResponse{}

	buf, readErr := ioutil.ReadFile(filename)

	if readErr != nil {
		fmt.Printf("error: %v", readErr)
	}

	parseErr := xml.Unmarshal(buf, &v)

	if parseErr != nil {
		fmt.Printf("error: %v", parseErr)
		return
	}

	return v.ResponseData.Record
}

func GetRecords(baseDir string, startRange int, endRange int) (record []schema.Record) {
	len := endRange - startRange
	records := make([]schema.Record, len)
	xmlDir := baseDir + "/xmls"

	files, err := ioutil.ReadDir(xmlDir)
	if err != nil {
		log.Fatal(err)
	}

	currentSlice := files[startRange:endRange]

	for _, file := range currentSlice {
		filename := xmlDir + "/" + file.Name()
		response := ReadResponse(filename)

		records = append(records, response)
	}

	return records
}
