package go_koans

import (
	"bytes"
	"strings"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		out.WriteString(in.String())
		/*
		   Your code goes here.
		   Hint, use these resources:
		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		str := strings.Split(in.String(), " ")
		out.WriteString(str[0])

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
