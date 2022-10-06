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

func TestCountCharacters(t *testing.T) {
	res, _ := countCharacters("AABBabc")
	// TODO: better way to test slices?
	if res[0] != 3 || res[1] != 3 || res[2] != 1 {
		t.Errorf("Wrong count! got %v wanted {3,3,1}", res)
	}
}

func TestStringToSet(t *testing.T) {
	res := stringToSet("Aa")
	if len(res) != 1 || res[0] != "a" {
		t.Errorf("Wrong set! got %v wanted {a}", res)
	}
}

func TestStringOccurencesToString(t *testing.T) {
	res, _ := stringOccurencesToString("AaBbac")
	if res != "321" {
		t.Errorf("Wrong occurences counted! Got %s wanted '321'", res)
	}
}
