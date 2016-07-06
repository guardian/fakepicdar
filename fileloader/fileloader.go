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

func xmlDir(baseDir string) (xmlDirectory string) {
	return baseDir + "/xmls"
}

func getXMLFiles(baseDir string) (xmlFiles []os.FileInfo) {
	xmlDir := baseDir + "/xmls"

	files, err := ioutil.ReadDir(xmlDir)
	if err != nil {
		log.Fatal(err)
	}

	return filterXMLFiles(files)
}

func getRecordsFromFiles(baseDir string, inFiles []os.FileInfo) (records []schema.Record) {
	for _, file := range inFiles {
		records = append(records, getRecordFromFile(baseDir, file.Name()))
	}

	return records
}

func getRecordFromFile(baseDir string, filename string) (record schema.Record) {
	filename = xmlDir(baseDir) + "/" + filename
	fmt.Printf("\nLoading %s\n", filename)

	return ReadResponse(filename)
}

func GetAllRecords(baseDir string) (records []schema.Record) {
	xmlFiles := getXMLFiles(baseDir)

	return getRecordsFromFiles(baseDir, xmlFiles)
}

func GetRecord(baseDir string, key string) (record schema.Record) {
	filename := key + ".xml"

	return getRecordFromFile(baseDir, filename)
}

func GetRecords(baseDir string, startRange int, endRange int) (records []schema.Record) {
	xmlFiles := getXMLFiles(baseDir)
	currentSlice := xmlFiles[startRange:endRange]

	return getRecordsFromFiles(baseDir, currentSlice)
}
