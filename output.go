package main

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func output(ec2s []Ec2Info) {
	buffer := &bytes.Buffer{}
	w := new(tabwriter.Writer)
	w.Init(buffer, 1, 4, 8, '\t', 0)
	for _, ec2info := range ec2s {
		fmt.Fprintf(w, "%s\t%s\t%s\n", ec2info.PrivateIP, ec2info.ID, ec2info.Name)
	}
	w.Flush()
	fmt.Println(buffer)
}
