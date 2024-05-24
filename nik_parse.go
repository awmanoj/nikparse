/*******************************************************
 * Parse & Validate KTP Population Identification Number (NIK).
 * Implemented by @awmanoj 2024 
 * Inspired and ported from original code in Javascript by @bachors 2018 | https://github.com/bachors/nik_parse.js
 *******************************************************/
package nikparse

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const kProvinceCodeKey = "provinsi"
const kDistrictCodeKey = "kabkot"
const kSubDistrictCodeKey = "kecamatan"

type NIKInfo struct {
	Province     string // provinsi
	District     string // kabkot
	SubDistrict  string // kecamatan
	DateOfBirth  string
	MonthOfBirth string
	YearOfBirth  string
	KodePOS      string
	Gender       string
}

func ParseNIK(nik string) (*NIKInfo, error) {
	// validate length
	if len(nik) != 16 {
		return nil, fmt.Errorf("err invalid NIK length")
	}

	// province code
	provinceCode := nik[0:2]

	// district code
	districtCode := nik[0:4]

	// subdistrict code
	subDistrictCode := nik[0:6]

	// Validate geo data presence
	if !HasGeoData(provinceCode, districtCode, subDistrictCode) {
		return nil, fmt.Errorf("err province or district or subdistrict not found: %s | %s | %s",
			provinceCode, districtCode, subDistrictCode)
	}

	// extract names
	provinceName := GetProvince(provinceCode)
	districtName := GetDistrict(districtCode)
	subDistrictPOSName := GetSubDistrict(subDistrictCode)

	// separate and extract subdistrict data and POS
	splits := strings.Split(strings.ToUpper(subDistrictPOSName), " -- ")
	subDistrictName := splits[0]
	kodePOS := splits[1]

	// current Year
	currentYear := time.Now().Year()
	currentYearLastTwoDigits := currentYear % 100

	yearOfBirthRaw := nik[10:12] // tahun NIK
	monthOfBirth := nik[8:10]    // bulan NIK
	dateOfBirthRaw := nik[6:8]   // tanggal NIK

	dateOfBirthInt, err := strconv.Atoi(dateOfBirthRaw)
	if err != nil {
		return nil, fmt.Errorf("err parsing dateOfBirth")
	}

	yearOfBirthInt, err := strconv.Atoi(yearOfBirthRaw)
	if err != nil {
		return nil, fmt.Errorf("err parsing yearOfBirthRaw")
	}

	// gender
	gender := "LAKI-LAKI"
	if dateOfBirthInt > 40 {
		gender = "PEREMPUAN"
	}

	// date of birth
	var dateOfBirth string = fmt.Sprintf("%02d", dateOfBirthInt)
	if dateOfBirthInt > 40 {
		dateOfBirth = fmt.Sprintf("%02d", dateOfBirthInt-40)
	}

	// tahun lahir
	yearPrefix := "19"
	if yearOfBirthInt < currentYearLastTwoDigits {
		yearPrefix = "20"
	}
	yearOfBirth := fmt.Sprintf("%s%02d", yearPrefix, yearOfBirthInt)

	dobStr := fmt.Sprintf("%s-%s-%s", dateOfBirth, monthOfBirth, yearOfBirth)
	if !isValidDate(dobStr, "02-01-2006") {
		return nil, fmt.Errorf("err invalid date of birth: %s", dobStr)
	}

	return &NIKInfo{
		Province:     provinceName,
		District:     districtName,
		SubDistrict:  subDistrictName,
		DateOfBirth:  fmt.Sprintf("%s", dateOfBirth),
		MonthOfBirth: monthOfBirth,
		YearOfBirth:  yearOfBirth,
		KodePOS:      kodePOS,
		Gender:       gender,
	}, nil
}

func HasGeoData(provinceCode, districtCode, subDistrictCode string) bool {
	_, ok1 := GeoData[kProvinceCodeKey][provinceCode]
	_, ok2 := GeoData[kDistrictCodeKey][districtCode]
	_, ok3 := GeoData[kSubDistrictCodeKey][subDistrictCode]
	return ok1 && ok2 && ok3
}

func GetProvince(provinceCode string) string {
	return GeoData[kProvinceCodeKey][provinceCode]
}

func GetDistrict(districtCode string) string {
	return GeoData[kDistrictCodeKey][districtCode]
}

func GetSubDistrict(subDistrictCode string) string {
	return GeoData[kSubDistrictCodeKey][subDistrictCode]
}

func isValidDate(dateStr, layout string) bool {
	_, err := time.Parse(layout, dateStr)
	return err == nil
}
