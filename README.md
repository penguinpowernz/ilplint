# ilplint - InfluxDB Line Protocol Validator

A simple command-line tool to validate InfluxDB line protocol from stdin.

## Features
- Real-time validation of InfluxDB line protocol
- Color-coded output (green for valid, red for errors)
- Uses [Telegraf's Influx parser](https://github.com/influxdata/telegraf) for accurate validation and sanity checking of ILP going to telegraf

## Usage

Build it:

    go mod tidy
    go build ./cmd/ilplint

Use it:

```bash
# Pipe input into the tool
echo "measurement,tag1=value1 field1=123 1625145600000000000" | ./ilplint

# Validate from file
cat measurements.lp | ./ilplint
```

## Why

I was ripping my hair out with errors from telegraf about ILP produced by my exec inputs.  It turned out to be that I didn't explicitly
specify `data_format = "influx"` but still, this can be useful for spotting errors in ILP output.
