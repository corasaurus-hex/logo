package main

import "testing"

func TestProjectBuilding(t *testing.T) {
	project, error := newProject()

	if error != nil {
		t.Fatal("LOL", error, project)
	}
}

func TestYourMomBuilding(t *testing.T) {
	project, error := newProject()

	if error != nil {
		t.Log("LOL", error, project)
	}
}
