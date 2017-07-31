package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

func output(cacheinfo map[string]interface{}) {
	columns := strings.Split(cacheinfo["columns"].(string), ",")
	ec2s := cacheinfo["instances"].([]interface{})

	buffer := &bytes.Buffer{}
	w := new(tabwriter.Writer)
	w.Init(buffer, 1, 4, 8, '\t', 0)
	for _, e := range ec2s {
		ee := e.(map[string]interface{})
		values := make([]string, 0)
		for _, c := range columns {
			values = append(values, ee[c].(string))
		}
		fmt.Fprintln(w, strings.Join(values, "\t"))
	}
	w.Flush()
	fmt.Println(buffer)
}
