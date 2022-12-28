// Maria plays college basketball and wants to go pro. Each season she maintains a record of her play. She tabulates the number of times she breaks her season record for most points and least points in a game. Points scored in the first game establish her record for the season, and she begins counting from there.

// Example

// Scores are in the same order as the games played. She tabulates her results as follows:

//                                      Count
//     Game  Score  Minimum  Maximum   Min Max
//      0      12     12       12       0   0
//      1      24     12       24       0   1
//      2      10     10       24       1   1
//      3      24     10       24       1   1

// Given the scores for a season, determine the number of times Maria breaks her records for most and least points scored during the season.

// Returns

//     int[2]: An array with the numbers of times she broke her records. Index

// is for breaking most points records, and index

//     is for breaking least points records.

// Input Format

// The first line contains an integer N, the number of games.
// The second line contains n 	space-separated integers describing the respective values of 0, 1, ... ,n-1 scores.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	var seasonRecord,
		currentMinRecord,
		currentMaxRecord,
		minRecordBreaks,
		maxRecordBreaks int32

	seasonRecord = scores[0]
	currentMinRecord = seasonRecord
	currentMaxRecord = seasonRecord
	minRecordBreaks = 0
	maxRecordBreaks = 0

	for i := 1; i < len(scores); i++ {
		if scores[i] > currentMaxRecord {
			currentMaxRecord = scores[i]
			maxRecordBreaks++
		} else if scores[i] < currentMinRecord {
			currentMinRecord = scores[i]
			minRecordBreaks++
		}
	}

	return []int32{maxRecordBreaks, minRecordBreaks}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	scoresTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var scores []int32

	for i := 0; i < int(n); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	result := breakingRecords(scores)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
