package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/invopop/jsonschema"
)

func DumpToJson(t any, out io.Writer) {
	schema := jsonschema.Reflect(t)
	jsonBytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling json schema: %s\n", err)
	}
	fmt.Fprint(out, string(jsonBytes))
}
