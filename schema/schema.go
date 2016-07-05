package schema

import (
	"encoding/xml"
)

type MogulResponse struct {
	XMLName      xml.Name `xml.MogulResponse`
	SourceAction string   `xml:"sourceAction,attr"`
	ResponseData ResponseData
	Response     Response
}

type Field struct {
	XMLName xml.Name `xml.Field`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type Record struct {
	XMLName xml.Name `xml.Record`
	MMRef   string
	Field   []Field
	VURL    []string
}

type ResponseData struct {
	XMLName    xml.Name `xml.ResponseData`
	MatchCount string
	Match      []Record
	Record     Record
	SearchID   string
}

type Response struct {
	XMLName xml.Name `xml.Response`
	Result  string   `xml:"result,attr"`
	MAK     string
}

type ActionType string

type ActionData struct {
	UserName, Password string
}

type MogulAction struct {
	XMLName    xml.Name `xml.MogulAction`
	ActionType ActionType
	ActionData ActionData
}
