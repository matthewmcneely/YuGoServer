package template_data

/*
	Loads the static template data file (json-based), used in serving templates.
	You can find an example in the test_fixtures folder.
*/

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var TemplateData interface{}

func Load(dir string) {
	buf, err := ioutil.ReadFile(dir + "/template_data.json")
	if err != nil {
		TemplateData = nil
	} else {
		err = json.Unmarshal(buf, &TemplateData)
		if err != nil {
			log.Printf("Error unmarshaling template_data file\n-----\n%s\n", err)
			TemplateData = nil
		}
	}
}
