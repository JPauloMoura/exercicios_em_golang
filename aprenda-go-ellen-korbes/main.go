package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type O struct {
	Nome string
}

func main() {
	o := O{"joao"}
	j, err := json.Marshal(o)

	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(j)
}
