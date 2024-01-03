package lib

import (
	"bytes"
	"os"
	"text/tabwriter"
)

func PrintMatrix(data [][]byte) {
	tw := tabwriter.NewWriter(os.Stdout, 1, 2, 1, ' ', 0)

	for _, row := range data {
		buf := bytes.Buffer{}

		for _, value := range row {
			buf.WriteByte(value)
			buf.WriteByte('\t')
		}

		buf.WriteByte('\n')
		tw.Write(buf.Bytes())
	}

	_ = tw.Flush()
	os.Stdout.WriteString("\n")
}
