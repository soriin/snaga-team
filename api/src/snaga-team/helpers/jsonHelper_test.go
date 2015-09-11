package helpers

import (
	"testing"
	"strings"
	"bytes"
)

type testJson struct {
	Name, OtherName string
}

func TestReadJson(t *testing.T) {
	testJsonString := "{\"Name\": \"test\", \"OtherName\": \"otherTest\"}"
	var testJsonOutput testJson

	err := ReadJson(strings.NewReader(testJsonString), &testJsonOutput)

	if err != nil {
		t.Errorf("Unexpected error thrown, %v", err.Error())
	}

	if testJsonOutput.Name != "test" {
		t.Errorf("Expected Name to be 'test' but got %v", testJsonOutput.Name)
	}

	if testJsonOutput.OtherName != "otherTest" {
		t.Errorf("Expected OtherName to be 'otherTest' but got %v", testJsonOutput.OtherName)
	}
}

func TestSendJson(t *testing.T) {
	testJsonObj := testJson{"test", "otherTest"}
	var buffer bytes.Buffer
	want := "{\"Name\":\"test\",\"OtherName\":\"otherTest\"}"

	err := SendJson(&buffer, testJsonObj)
	if err != nil {
		t.Errorf("Unexpected error thrown, %v", err.Error())
	}

	bufferString := buffer.String()

	if bufferString != want {
		t.Errorf("Expected bufferString to be %v but got %v", want, bufferString)
	}
}
