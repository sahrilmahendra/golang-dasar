package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`s([a-z][a-z][a-z][a-z])l`)

	fmt.Println(regex.MatchString("sahril"))
	fmt.Println(regex.MatchString("sahryl"))
	fmt.Println(regex.MatchString("saHril"))
	fmt.Println(regex.MatchString("fahrul"))

	fmt.Println(regex.FindAllString("sahril sahrul sahrol fajrul past", 10))
}
