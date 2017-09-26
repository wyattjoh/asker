package asker

import (
	"bufio"
	"io"
	"strings"
)

// Prompt will read from the input until a newline is reached.
func Prompt(reader io.Reader) (string, error) {
	text, err := bufio.NewReader(reader).ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}

// Ask will walk over each question and prompt for it, returning all the
// answers in an array.
func Ask(reader io.Reader, questions ...Asker) ([]Answer, error) {

	answers := make([]Answer, len(questions))
	for i, question := range questions {
		answer, err := question.Ask(reader)
		if err != nil {
			return nil, err
		}

		answers[i] = *answer
	}

	return answers, nil
}
