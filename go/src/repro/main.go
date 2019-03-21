package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type targetStruct struct {
	SomeList []string `yaml:"some_list"`
}

func main() {
	someYaml := `
some_list:
  - item1
  - item2
`

	target := targetStruct{}
	_ = yaml.Unmarshal([]byte(someYaml), &target)
	fmt.Printf("%+v", target)
}

