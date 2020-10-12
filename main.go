package main

import (
	"encoding/json"
	"fmt"
	"flag"
	"strings"
	"io/ioutil"
	"net/http"
)

type CrtShResult struct {
	Name string `json:"name_value"`
}

func fetchCrtSh(domain string) ([]string, error) {
	var results []CrtShResult

	resp, err := http.Get(
		fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain),
	)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	output := make([]string, 0)

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &results); err != nil {
		return []string{}, err
	}

	for _, res := range results {
		output = append(output, res.Name)
	}
	return output, nil
}


func cleanDomain(d string) string {
	d = strings.ToLower(d)

	if len(d) < 2 {
		return d
	}


	if d[0] == '.' {
		d = d[1:]
	}

	return d
}

func main() {

	flag.Parse()

	input := flag.Arg(0)

	output, _:= fetchCrtSh(input)

	printed := make(map[string]bool)

	domains := []string {}

	for _, o := range output {
		domain := strings.Split(o,  "\n")
		for _, d := range(domain) {
			domains = append(domains, d)
		}
	}

	for _, d := range domains {
		if strings.ContainsAny(d, "@") {
			continue
		}

		if _, ok := printed[d]; ok {
			continue
		}

		printed[d] = true

		d = strings.ToLower(d)

		if len(d) < 2 {
			fmt.Println(d)
		}

		if d[0] == '*' || d[0] == '%' {
			d = d[1:]
		}

		if d[0] == '.' {
			d = d[1:]
		}
		fmt.Println(d)
	}



}
