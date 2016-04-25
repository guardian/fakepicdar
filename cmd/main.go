package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/guardian/fakepicdar/fileloader"
	"github.com/guardian/fakepicdar/schema"
	"io"
	"log"
	"net/http"
	"os"
)

var baseDir string

func main() {
	flag.StringVar(&baseDir, "basedir", ".", "base dir of picdar images/responses")
	flag.Parse()

	fs := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.HandleFunc("/", serveQuery)

	log.Println("Listening ...")

	http.ListenAndServe(":3000", nil)
}

func serveQuery(w http.ResponseWriter, r *http.Request) {
	v := schema.MogulAction{
		ActionType: "",
		ActionData: schema.ActionData{
			UserName: "", Password: ""}}

	buf := bytes.NewBuffer(nil)

	if r.Body != nil {
		if _, err := io.Copy(buf, r.Body); err != nil {
			log.Fatal(err)
		}

		err := xml.Unmarshal([]byte(buf.String()), &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		responseXml := &schema.MogulResponse{}

		log.Println("\n-----Received-----")
		fmt.Println(buf)

		fmt.Printf("\n--++Received ActionType: %s\n", v.ActionType)

		switch v.ActionType {
		case "Login":
			fmt.Println("\n--++Faking login.")
			responseXml = &schema.MogulResponse{Response: schema.Response{MAK: "fake"}}
		case "Search":
			fmt.Println("\n--++Sending search response.")
			responseXml = &schema.MogulResponse{
				Response:     schema.Response{MAK: "fake", Result: "OK"},
				ResponseData: schema.ResponseData{MatchCount: "5", SearchID: "0"}}
		case "RetrieveResults":
			fmt.Println("\n--++Faking Results.")
			records := fileloader.GetRecords(baseDir, 0, 10)
			responseXml = &schema.MogulResponse{
				Response:     schema.Response{MAK: "fake", Result: "OK"},
				ResponseData: schema.ResponseData{Match: records}}
		default:
			fmt.Println("\n!!--++Unrecognised ActionType.")
			return
		}

		sendResponse(responseXml, w)
	}
}

func sendResponse(d *schema.MogulResponse, w http.ResponseWriter) {
	output, err := xml.MarshalIndent(d, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	log.Println("\n-----Sent-----")
	os.Stdout.Write(output)

	w.Write(output)
}
