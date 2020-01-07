package models

import (
	"fmt"

	fhir "github.com/Universal-Health-Chain/fhir-r4-go-models/models"
)

type EncounterUHC struct {
	DomainResource string `bson:",inline"`
}

func (encounter *EncounterUHC) CreateFromFhireR4(encounterFhire *fhir.Encounter) {
	fmt.Println(encounterFhire)
	encounter = &EncounterUHC{}
}
