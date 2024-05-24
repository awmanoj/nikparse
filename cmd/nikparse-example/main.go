package main

import (
	"fmt"
	"log"
	"os"
	"github.com/awmanoj/nikparse"
)

func main() {
	nik := os.Args[1]
	info, err := nikparse.ParseNIK(nik)
	if err != nil {
		log.Println("err parsing nik", err)
		return
	}

	fmt.Printf("%s|%s-%s-%s|%s|%s|%s|%s|%s\n", nik, info.DateOfBirth, info.MonthOfBirth, info.YearOfBirth, info.Gender,
		info.Province, info.District, info.SubDistrict, info.KodePOS)
}
