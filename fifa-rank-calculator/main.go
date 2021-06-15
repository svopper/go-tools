package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getPbefore(reader *bufio.Reader) float32 {
	fmt.Printf("Enter the team's number of points before the game\n")
	val, _ := reader.ReadString('\n')
	parsedVal, _ := strconv.ParseFloat(strings.TrimSpace(val), 32)
	return float32(parsedVal)
}

func getImportanceCoeff(reader *bufio.Reader) float32 {
	coeffString := `Enter importance coefficient:
(1) friendlies played outside the International Match Calendar windows
(2) friendlies played within the International Match Calendar windows
(3) Nations League matches (group stage)
(4) Nations League matches (play-offs and finals)
(5) Confederations' final competitions qualifiers, FIFA World Cup qualifiers
(6) Confederations' final competitions matches (before quarter-finals)
(7) Confederations' final competitions matches (quarter-finals and later)
(8) FIFA World Cup matches (before quarter-finals)
(9) FIFA World Cup matches (quarter-finals and later)
	`
	fmt.Println(coeffString)
	fmt.Println()
	val, _ := reader.ReadString('\n')
	return parseCoeffChoice(val)
}

func getGameResult(reader *bufio.Reader) float32 {
	resultString := `Enter result type:
(1) loss after regular or extra time
(2) draw or loss in a penalty shootout
(3) win in a penalty shootout
(4) win after regular or extra time
	`
	fmt.Println(resultString)
	fmt.Println()
	val, _ := reader.ReadString('\n')
	return parseResultChoice(val)
}

func parseCoeffChoice(choice string) float32 {
	trimmed := strings.TrimSpace(choice)
	fmt.Println(trimmed)
	switch trimmed {
	case "1":
		return 5
	case "2":
		return 10
	case "3":
		return 15
	case "4":
		fallthrough
	case "5":
		return 25
	case "6":
		return 35
	case "7":
		return 40
	case "8":
		return 50
	case "9":
		return 60
	default:
		panic("PANICC!!")
	}

}

func parseResultChoice(choice string) float32 {
	trimmed := strings.TrimSpace(choice)
	switch trimmed {
	case "1":
		return 0.
	case "2":
		return 0.5
	case "3":
		return 0.75
	case "4":
		return 1.
	default:
		panic("PANICC")
	}
}

func calculateRank(reader *bufio.Reader) float32 {
	pBeforeTeam := getPbefore(reader)
	pBeforeOpponent := getPbefore(reader)
	importanceCoefficient := getImportanceCoeff(reader)
	gameResult := getGameResult(reader)

	dr := math.Abs(float64(pBeforeTeam - pBeforeOpponent))
	exp := (float64)(math.Abs(dr) / 600.)
	we := 1 / (math.Pow(10, exp) + 1)

	return pBeforeTeam + importanceCoefficient*(gameResult-float32(we))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result := calculateRank(reader)
	fmt.Println(result)
}
