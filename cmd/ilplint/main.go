package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	// "github.com/influxdata/influxdb1-client/models"
	"github.com/influxdata/telegraf/plugins/parsers/influx"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	p := influx.Parser{}
	p.Init()

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			color.Red("LINE ERROR: %v\n", err)
			return
		}

		// _, err = models.ParsePointsString(line)
		_, err = p.ParseLine(line)
		if err != nil {
			color.Red("ERROR: %v\n", err)
			continue
		}

		color.Green("%s", strings.TrimSpace(line))
	}
}
