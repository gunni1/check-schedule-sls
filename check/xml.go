package main

import "encoding/xml"

//Schedule is the top level of the response xml
type Schedule struct {
	XMLName xml.Name `xml:"vp"`
	Head    Head     `xml:"kopf"`
}

//Head is some meta information about schedule changes
type Head struct {
	Titel      string `xml:"titel"`
	UploadDate string `xml:"datum"`
	Info       Info   `xml:"kopfinfo"`
}

//Info is called kopfinfo in xml. It holds overview information about current schedule changes
type Info struct {
	ChangesTeacher string `xml:"aenderungl"`
}
