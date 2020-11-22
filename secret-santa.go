package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/smallfish/simpleyaml"
)

// Read the YAML into a nice in-memory representation of a map of strings to array
func readYaml(y *simpleyaml.Yaml) (map[string][]string, error) {
	// Get the array size so that we can iterate over each element
	length, err := y.GetArraySize()
	if err != nil {
		return map[string][]string{}, err
	}

	// For each element, extract the participants
	participants := map[string][]string{}
	for i := 0; i < length; i++ {
		element := y.GetIndex(i)

		// Obtain the participant name from the element
		nameElement := element.Get("name")
		if !nameElement.IsFound() {
			return map[string][]string{}, fmt.Errorf("Element %d missing item 'name'", i)
		}
		name, _ := nameElement.String()

		// Obtain the invalid list for this participant from the element
		invalidElement := element.Get("invalid")

		// If no invalid names were given for this element, add an empty array.
		if !invalidElement.IsFound() {
			participants[name] = []string{}
			continue
		}

		// 'invalid' must be an array
		if !invalidElement.IsArray() {
			return map[string][]string{}, fmt.Errorf("'invalid' at element %d is not an array", i)
		}

		// Collect the data in this array into memory
		arrLength, _ := invalidElement.GetArraySize()
		invalidPeople := []string{}
		for j := 0; j < arrLength; j++ {
			person, _ := invalidElement.GetIndex(j).String()
			invalidPeople = append(invalidPeople, person)
		}

		// Last, add this array to the mapping
		participants[name] = invalidPeople
	}

	return participants, nil
}

// From the givers list and the mapping of invalid users
// determine who will be the recipients.
func determineRecipients(givers []string, invalidMapping map[string][]string) ([]string, error) {
	// Prime the random number generator
	rand.Seed(time.Now().UnixNano())

	// Initiallize with the given list
	recipients := make([]string, len(givers))
	copy(recipients, givers)

	// Iterate until we get a solution that works
	// Well, iterate for up to 1 billion times
	// This is a "brute force" solution - really, not smart
	for i := 0; i < 1000000000; i++ {
		// Shuffle the recipeients.
		rand.Shuffle(len(recipients), func(i, j int) {
			recipients[i], recipients[j] = recipients[j], recipients[i]
		})

		// If any pairing is self or invalid, reject this selection.
		invalidSolution := false
		for i, name := range givers {
			// Self pairing
			if name == recipients[i] {
				invalidSolution = true
				break
			}
			// Invalid pairing
			ok := true
			for _, invalidName := range invalidMapping[recipients[i]] {
				if name == invalidName {
					ok = false
				}
			}
			if !ok {
				invalidSolution = true
				break
			}
		}
		if invalidSolution {
			continue
		}

		// If we got here, then this is a valid solution!
		return recipients, nil
	}

	return []string{}, fmt.Errorf("no solutions found after 1,000,000,000 tries")
}

func main() {
	// Read the participants into memory
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "You must pass one and only one participants file!")
		os.Exit(1)
	}
	participantsFile := os.Args[1]

	// Read file ino memory
	participantsContents, err := ioutil.ReadFile(participantsFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read YAML!", err)
		os.Exit(1)
	}

	// Load file into memory
	participantsData, err := simpleyaml.NewYaml(participantsContents)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse YAML!", err)
		os.Exit(1)
	}

	// From the files, collect the list of participants.
	participantsMap, err := readYaml(participantsData)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Collect the particpants into an individual array
	givers := []string{}
	for name := range participantsMap {
		givers = append(givers, name)
	}

	// Determine the recipients
	recipients, err := determineRecipients(givers, participantsMap)
	if err != nil {
		// Oh no!
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Huzzah! Write to screen and quit.
	maxLen := 0
	for _, name := range givers {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}
	fstring := fmt.Sprintf("%%-%ds ==> %%s\n", maxLen)
	for i, name := range givers {
		fmt.Printf(fstring, name, recipients[i])
	}
	os.Exit(0)
}
