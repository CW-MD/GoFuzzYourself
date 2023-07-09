package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file")
		os.Exit(1)
	}
	lines := readFile()
	//getTest()
	getRequest(lines)

}

func readFile() []string {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Printf("Something went wrong: %s\n", err)
		os.Exit(1)
	}

	// content, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Printf("Error reading file: %s\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("File Contents: \n%s\n", content)
	// getRequest(string(content))

	return lines
}

func getRequest(urlArr []string) []string {

	responsesBodies := []string{}

	for i, url := range urlArr {
		err := func() error {
			resp, err := http.Get(url)
			if err != nil {
				//todo fix
				fmt.Printf("Error making a request to: %s\n", url)
				return err
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Something went wrong, %s\n", err)
				return err
			}
			stringBody := string(body)
			responsesBodies = append(responsesBodies, stringBody)
			fmt.Printf("%s : %s/n", url, stringBody)
			_ = i
			return nil
		}()
		if err != nil {
			fmt.Printf("Error processing url %s: %v", url, err)
			continue
		}
	}

	return responsesBodies
}

func getTest() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err0r %s", err)
	}
	fmt.Printf("%s", string(body))
}
