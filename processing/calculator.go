package processing

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"home/edsjr/work/testes/mathematics"
)

var (
	highPrecedenceOperationTokens = "*/"
	operationTokens               = "+-"
	tokens                        = highPrecedenceOperationTokens + operationTokens
	sentenceParts                 []string
)

type Calculator struct {
	mathematics mathematics.IMathematics
}

func (c *Calculator) Calculate(sentence string) (result float64, err error) {
	if !isValid(sentence) {
		return result, errors.New("invalid sentence")
	}

	for {
		notProcessedParts := make([]string, 0)
		total, processedParts, highPrecedenceErr := processHighPrecedence(c.mathematics)
		if highPrecedenceErr != nil {
			err = highPrecedenceErr
			return
		}

		if len(processedParts) != 0 {

			notProcessedParts = append(notProcessedParts, sentenceParts[:processedParts[0]]...)
			notProcessedParts = append(notProcessedParts, strconv.FormatFloat(total, 'f', 15, 64))
			notProcessedParts = append(notProcessedParts, sentenceParts[processedParts[2]+1:]...)

			sentenceParts = notProcessedParts
		} else {
			break
		}
	}

	for {
		notProcessedParts := make([]string, 0)
		total, processedParts, operationsErr := processOtherOperations(c.mathematics)
		if operationsErr != nil {
			err = operationsErr
			return
		}

		if len(processedParts) != 0 {

			notProcessedParts = append(notProcessedParts, sentenceParts[:processedParts[0]]...)
			notProcessedParts = append(notProcessedParts, strconv.FormatFloat(total, 'f', 15, 64))
			notProcessedParts = append(notProcessedParts, sentenceParts[processedParts[2]+1:]...)

			sentenceParts = notProcessedParts
		} else {
			break
		}
	}

	result, err = strconv.ParseFloat(sentenceParts[0], 64)

	return
}

func isValid(sentence string) bool {
	pattern := `[0-9*/+-.]`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(sentence) {
		return false
	} else {
		if strings.Contains(tokens, string(sentence[0])) || strings.Contains(tokens, string(sentence[len(sentence)-1])) {
			return false
		}
	}

	getSentenceParts(sentence)

	for idx, part := range sentenceParts {
		if strings.Contains(tokens, part) {
			if idx > 0 && strings.Contains(tokens, sentenceParts[idx-1]) {
				return false
			}
			continue
		}

		if len(strings.Replace(part, ".", "", -1)) < len(part)-1 {
			return false
		}
	}

	return true
}

func getSentenceParts(sentence string) {
	sentenceParts = make([]string, 0)
	part := ""
	for _, c := range sentence {
		if strings.Contains(tokens, string(c)) {
			if part != "" {
				sentenceParts = append(sentenceParts, part)
			}
			sentenceParts = append(sentenceParts, string(c))
			part = ""
			continue
		}

		part += string(c)
	}

	sentenceParts = append(sentenceParts, part)
}

func processHighPrecedence(m mathematics.IMathematics) (total float64, processedParts []int, err error) {
	processedParts = make([]int, 0)
	for i := 0; i < len(sentenceParts); i++ {
		if strings.Contains(highPrecedenceOperationTokens, sentenceParts[i]) {
			valueA, _ := strconv.ParseFloat(sentenceParts[i-1], 64)
			valueB, _ := strconv.ParseFloat(sentenceParts[i+1], 64)

			switch sentenceParts[i] {
			case "*":
				total, err = m.Multiplication(valueA, valueB)
			case "/":
				total, err = m.Division(valueA, valueB)
			}

			processedParts = append(processedParts, []int{i - 1, i, i + 1}...)

			if err != nil {
				return
			}

			break
		}
	}

	return
}

func processOtherOperations(m mathematics.IMathematics) (total float64, processedParts []int, err error) {
	processedParts = make([]int, 0)
	for i := 0; i < len(sentenceParts); i++ {
		if strings.Contains(operationTokens, sentenceParts[i]) {
			valueA, _ := strconv.ParseFloat(sentenceParts[i-1], 64)
			valueB, _ := strconv.ParseFloat(sentenceParts[i+1], 64)

			switch sentenceParts[i] {
			case "+":
				total, err = m.Addition(valueA, valueB)
			case "-":
				total, err = m.Subtraction(valueA, valueB)
			}

			processedParts = append(processedParts, []int{i - 1, i, i + 1}...)

			if err != nil {
				return
			}

			break
		}
	}

	return
}
