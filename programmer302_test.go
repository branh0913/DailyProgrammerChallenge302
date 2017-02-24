package main_test



import (
	"github.com/branh0913/DailyProgrammerChallenge302"
	"testing"
)

func TestSymLenth( t *testing.T)  {
	fileString := "element.csv"

	lst := main.GetAtomicSymLen(fileString, 1)

	if lst[0]["b"] == "baron"{
		t.Log("Test passed")
	}

}

func TestMatchSingle(t *testing.T) {

	newMap := map[string]string{"b": "baron"}

	testArrayMap := make([]map[string]string, 0)

	testArrayMap = append(testArrayMap, newMap)

	testsingleCheck := main.MatchSingle("b", testArrayMap)

	if testsingleCheck == "baron"{
		t.Log("Pass")
	}
}

