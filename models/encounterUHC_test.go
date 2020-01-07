package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	fhir "github.com/Universal-Health-Chain/fhir-r4-go-models/models"
	"github.com/stretchr/testify/assert"
)

func TestEncounterUHC(t *testing.T) {
	path := "../examplesForTesting/encounter-example-f203.json"
	encounterFhire, err := readEncounterJson(path)
	assert.Equal(t, err, nil, "OK encounter loaded")
	assert.NotEmpty(t, encounterFhire.Identifier, "OK encounter loaded")
}

func readEncounterJson(path string) (encounter fhir.Encounter, err error) {
	encounter = fhir.Encounter{}
	file, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		fmt.Println(err1)
		return encounter, err1
	}
	err2 := json.Unmarshal([]byte(file), &encounter)
	if err2 != nil {
		fmt.Println(err2)
		return encounter, err2
	}
	return encounter, nil
}
