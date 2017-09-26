package asker

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/fatih/color"
)

// Asker is a struct which can Ask the user for input.
type Asker interface {
	Ask(reader io.Reader) (*Answer, error)
}

// TextAsker will Ask the user a text based inquiry.
type TextAsker struct {
	Question  string
	Prefix    string
	Required  bool
	Validator func(response string) error
}

// Ask implements the Asker interface.
func (tp TextAsker) Ask(reader io.Reader) (*Answer, error) {
	for {
		fmt.Printf("[%s] %s: %s", color.CyanString("?"), tp.Question, tp.Prefix)
		response, err := Prompt(reader)
		if err != nil {
			return nil, err
		}

		if tp.Required && response == "" {
			color.Red(">> input is required")
			continue
		}

		if tp.Validator != nil {
			err := tp.Validator(response)
			if err != nil {
				color.Red(">> %v", err)
				continue
			}
		}

		if response == "" {
			return &Answer{}, nil
		}

		return &Answer{
			StringResponse: response,
			Provided:       true,
		}, nil
	}
}

// ChoiceAsker will allow the user to select from a list of options.
type ChoiceAsker struct {
	Choices  []string
	Question string
}

// Ask implements the Asker interface.
func (cp ChoiceAsker) Ask(reader io.Reader) (*Answer, error) {
	tp := TextAsker{
		Question: cp.Question,
		Required: true,
		Validator: func(response string) error {
			choice, err := strconv.Atoi(response)
			if err != nil || choice > len(cp.Choices) || choice <= 0 {
				return errors.New("please select one of the provided options")
			}

			return nil
		},
	}

	for i, choice := range cp.Choices {
		fmt.Printf("  %d) %s\n", i+1, choice)
	}

	answer, err := tp.Ask(reader)
	if err != nil {
		return nil, err
	}

	choice, err := strconv.Atoi(answer.StringResponse)
	if err != nil {
		return nil, err
	}

	return &Answer{
		StringResponse: cp.Choices[choice-1],
		Provided:       true,
	}, nil
}

// ConfirmAsker will allow the user to select from y/N.
type ConfirmAsker struct {
	Question string
	Default  bool
}

// Ask implements the Asker interface.
func (cp ConfirmAsker) Ask(reader io.Reader) (*Answer, error) {
	prefix := "(y/N) "
	if cp.Default {
		prefix = "(Y/n) "
	}

	tp := TextAsker{
		Question: cp.Question,
		Prefix:   prefix,
		Validator: func(response string) error {
			if response == "" {
				return nil
			}

			if _, err := strconv.ParseBool(response); err != nil {
				return err
			}

			return nil
		},
	}

	response, err := tp.Ask(reader)
	if err != nil {
		return nil, err
	}

	if !response.Provided {
		return &Answer{
			BoolResponse: cp.Default,
			Provided:     false,
		}, nil
	}

	answer, err := strconv.ParseBool(response.StringResponse)
	if err != nil {
		return nil, err
	}

	return &Answer{
		BoolResponse: answer,
		Provided:     true,
	}, nil
}

// func main() {
// 	questions := []Asker{
// 		TextAsker{
// 			Question: "What is bird",
// 		},
// 		TextAsker{
// 			Question: "No, what is bird",
// 			Required: true,
// 		},
// 		TextAsker{
// 			Question: "Ok, but, bird",
// 			Validator: func(response string) error {
// 				if response != "bird" {
// 					return errors.New("not bird")
// 				}

// 				return nil
// 			},
// 		},
// 		ChoiceAsker{
// 			Choices:  []string{"red", "blue", "orange"},
// 			Question: "Favorite Colour",
// 		},
// 	}

// 	answers, err := Ask(io.Stdin, questions...)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Can't ask the questions: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("%#v\n", answers)
// }
