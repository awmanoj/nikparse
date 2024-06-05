package main

import (
	"fmt"
	//"log"
	"os"
	"github.com/awmanoj/nikparse"
)

func main() {
	nik := os.Args[1]

	// skip geo data validation
	nikparse.DoNotValidateGeoData = true
	
	info, err := nikparse.ParseNIK(nik)
	if err != nil {
		if !nikparse.DoNotValidateGeoData {
			fmt.Println(",,,,,,")
			//log.Println("err parsing nik", err. nik)
		} else {
			fmt.Println(",")
			//log.Println("err parsing nik", err. nik)
		}
		return
	}

	if !nikparse.DoNotValidateGeoData {
		fmt.Printf("%s|%s-%s-%s|%s|%s|%s|%s|%s\n", nik, info.DateOfBirth, info.MonthOfBirth, info.YearOfBirth, info.Gender,
			info.Province, info.District, info.SubDistrict, info.KodePOS)
	} else {
		fmt.Printf("%s,%s-%s-%s\n", info.Gender, info.DateOfBirth, info.MonthOfBirth, info.YearOfBirth)
	}
}
