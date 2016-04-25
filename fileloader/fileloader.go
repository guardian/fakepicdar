package fileloader

import (
	"encoding/xml"
	"fmt"
	"github.com/guardian/fakepicdar/schema"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func filterXMLFiles(inFiles []os.FileInfo) (outFiles []os.FileInfo) {
	for _, v := range inFiles {
		if strings.HasSuffix(v.Name(), ".xml") {
			outFiles = append(outFiles, v)
		}
	}

	return outFiles
}

func GetRecords(baseDir string, startRange int, endRange int) (record []schema.Record) {
	len := endRange - startRange
	records := make([]schema.Record, len)
	xmlDir := baseDir + "/xmls"

	files, err := ioutil.ReadDir(xmlDir)
	if err != nil {
		log.Fatal(err)
	}

	xmlFiles := filterXMLFiles(files)
	currentSlice := xmlFiles[startRange:endRange]

	for _, file := range currentSlice {
		filename := xmlDir + "/" + file.Name()

		fmt.Printf("\nLoading %s\n", filename)

		response := ReadResponse(filename)

		records = append(records, response)
	}

	return records
}
