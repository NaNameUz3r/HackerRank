// Given a time in

// -hour AM/PM format, convert it to military (24-hour) time.

// Note:
// - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()

}

func timeConversion(s string) string {
	isAM := s[len(s)-2:] == "AM"
	isPM := s[len(s)-2:] == "PM"

	splitTime := strings.Split(s, ":")
	splitTime[2] = splitTime[2][:2]

	if isAM {
		if splitTime[0] == "12" {
			splitTime[0] = "00"
		}
	} else if isPM {
		if splitTime[0] != "12" {
			swapHourInt, _ := strconv.Atoi(splitTime[0])
			swapHourInt += 12
			swapHourStr := strconv.Itoa(swapHourInt)
			splitTime[0] = swapHourStr
		}
	}

	resultTime := splitTime[0] + ":" + splitTime[1] + ":" + splitTime[2]
	return resultTime
}
	if isAM {
		if splitTime[0] == "12" {
			splitTime[0] = "00"
		}
	} else if isPM {
		if splitTime[0] != "12" {
			swapHourInt, _ := strconv.Atoi(splitTime[0])
			swapHourInt += 12
			swapHourStr := strconv.Itoa(swapHourInt)
			splitTime[0] = swapHourStr
		}
	}

	resultTime := splitTime[0] + ":" + splitTime[1] + ":" + splitTime[2]
	return resultTime
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
