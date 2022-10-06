package util

import "testing"

func TestCalculateLove(t *testing.T) {
	got, err := CalculateLove("Teppo", "Anna-Reeta")
	if err != nil {
		t.Errorf("Teppo and Anna-Reeta got error %s", err)
	} else if got.Percent != 100 {
		t.Errorf("Teppo and Anna-Reeta score is %d, want 100", got.Percent)
	} else if got.Description != "You are perfect couple!" {
		t.Errorf("Teppo and Anna-Reeta description is\n%s, want \nYou are perfect couple!", got.Description)
	}
}

func TestCalculateLoveEmpty(t *testing.T) {
	_, err := CalculateLove("", "")
	if err == nil {
		t.Error("Empty names did not got error")
	}
}

func TestCalculateLoveIllegal(t *testing.T) {
	_, err := CalculateLove("Teppo", "123")
	if err == nil {
		t.Error("Illegal characters did not got error")
	}
}

func TestGetDescription(t *testing.T) {
	desc := getDescription(70)
	if desc != "There might be some problems but give it a try!" {
		t.Errorf("Got wrong description! Got \n%s\nwanted\n'There might be some problems but give it a try!'", desc)
	}
}
