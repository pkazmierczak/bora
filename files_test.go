package main

import "testing"

var testData = `
stackname: exampleStack
region: us-east-1
foobars:
 - foo: 1
   bar:
    - one
    - two
    - three
 
 - foo: 2
   bar:
    - one1
    - two2
    - three3
`

func testYamlParser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func testFileReader(t *testing.T) {}

func testTemplateParser(t *testing.T) {}
