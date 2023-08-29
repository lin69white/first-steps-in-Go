package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person map[string]string

func RemoveSpecialReturn(s string) string {
	var newline = "\n"
	var cabret = "\r"

	if strings.HasSuffix(s, newline) {
		s = strings.TrimSuffix(s, newline)
	}
	if strings.HasSuffix(s, cabret) {
		s = strings.TrimSuffix(s, cabret)
	}

	return s
}

func inputPersonData() Person {
	person := make(Person)
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("請輸入名字:")
		if name, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\n錯誤: ", err)
		} else {
			person["name"] = RemoveSpecialReturn(name)
			break
		}
	}

	for {
		fmt.Print("請輸入地址:")
		if address, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\n錯誤: ", err)
		} else {
			person["address"] = RemoveSpecialReturn(address)
			break
		}
	}

	return person
}

func marshJSON(person Person) {
	if data, err := json.Marshal(person); err != nil {
		fmt.Println("\n序列化為JSON時出錯誤: ", err)
	} else {
		fmt.Println("\nJSON:", string(data))
	}
}

func main() {
	person := inputPersonData()

	marshJSON(person)
}
