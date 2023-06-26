package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide the path to the text file as a command-line argument.")
		return
	}
	fmt.Println("MADE BY TREPHY")

	fmt.Println("Checking for nitro codes in the text file...")

	filePath := os.Args[1]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Iterating over the list of strings from the file:")

	var validLines []string
	var validCount, invalidCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nitro := scanner.Text()
		url := fmt.Sprintf("https://discordapp.com/api/v9/entitlements/gift-codes/%s?with_application=false&with_subscription_plan=true", nitro)

		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		if response.StatusCode == http.StatusOK {
			fmt.Printf("Valid | %s\n", nitro)
			validLines = append(validLines, nitro)
			validCount++
			// Write the nitro code to the file
			file, err := os.OpenFile("Nitro Codes.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			file.WriteString(nitro + "\n")
		} else {
			fmt.Printf("Invalid | %s\n", nitro)
			invalidCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Valid lines:")
	for _, line := range validLines {
		fmt.Println(line)
	}

	fmt.Printf("Valid lines: %d\n", validCount)
	fmt.Printf("Invalid lines: %d\n", invalidCount)

}
