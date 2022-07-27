package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	defer recovery()
	var domain string
	var brute bool
	flag.StringVar(&domain, "d", "", "eg: bing.com")
	flag.BoolVar(&brute, "brute", false, "BruteForce Mode")
	flag.Parse()

	if brute == false {
		dom := "https://crt.sh/?q=" + domain
		resp, err := http.Get(dom)
		if err != nil {
			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		sb := string(body)
		replacer := strings.NewReplacer(">", "\n", "<", "\n", "*.", "\n")
		output := replacer.Replace(sb)
		for {
			idxEnd := strings.Index(output, "."+domain)
			left := strings.LastIndex(output[:idxEnd], "\n")
			right := strings.Index(output[idxEnd:], "\n")
			fmt.Print(output[left : idxEnd+right])
			output = strings.Replace(output, output[left:idxEnd+right], "", -1)

		}
	} else {
		file, err := os.Open("subword.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		var brut string
		var brut1 string
		scanner := bufio.NewScanner(file)
		// optionally, resize scanner's capacity for lines over 64K, see next example
		for scanner.Scan() {
			brut = "https://" + scanner.Text() + "." + domain
			brut1 = "http://" + scanner.Text() + "." + domain

			fmt.Println(brut)
			resp, err := http.Get(brut)
			if err != nil {
				fmt.Print("")
			}
			if resp.StatusCode == 200 {
				fmt.Println(brut)
			} else {
				resp, err = http.Get(brut1)
				if err != nil {
					fmt.Print("")
				}
				if resp.StatusCode == 200 {
					fmt.Println(brut1)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

}
func recovery() {
	if r := recover(); r != nil {

	}

}
