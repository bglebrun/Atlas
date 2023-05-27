package badwords

import (
	"reflect"
	"strings"
	"testing"
)

func TestBadWordsFilter(t *testing.T) {
	filterList := []string{"Apple", "banana", "carrot", "pineapple"}
	file := strings.Join(filterList, "\n")
	reader := strings.NewReader(file)

	testingList, err := ReadList(reader)
	if err != nil {
		t.Errorf("Test produced error %s", err.Error())
	}

	if !reflect.DeepEqual(filterList, testingList.words) {
		t.Errorf("Testing did not initialize word list")
	}

	testString := "Eating very delicous apples today"
	want := "Eating very delicous *****s today"
	got := testingList.Filter(testString)
	if want != got {
		t.Errorf("Words not filtered, got: %s", got)
	}
}
