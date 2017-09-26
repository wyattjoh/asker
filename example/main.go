package main

import (
	"fmt"
	"os"

	"github.com/wyattjoh/asker"
)

func main() {
	questions := []asker.Asker{
		asker.TextAsker{
			Question: "What is your favorite colour",
			Required: true,
		},
		asker.ChoiceAsker{
			Question: "Pick the best fruit",
			Choices:  []string{"apple", "orange"},
		},
		asker.ConfirmAsker{
			Question: "Did you like this example",
		},
	}

	answers, err := asker.Ask(os.Stdin, questions...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't help you if you won't asker the question: %s", err.Error())
		os.Exit(1)
	}

	did := "did"
	if !answers[2].BoolResponse {
		did = "did not"
	}

	fmt.Printf("Your favorite colour is %s, you picked %s as the best fruit, and you %s like this example.\n", answers[0].StringResponse, answers[1].StringResponse, did)
}
