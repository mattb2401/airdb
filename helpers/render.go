package helpers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func toJson(p interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func toXML(p interface{}) ([]byte, error) {
	b, err := xml.Marshal(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func RenderJSON(p interface{}, w http.ResponseWriter) {
	js, _ := toJson(p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func RenderXML(p interface{}, w http.ResponseWriter) {
	xml, _ := toXML(p)
	w.Header().Set("Content-Type", "application/xml")
	w.Write(xml)
}
