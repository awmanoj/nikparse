/*******************************************************
 * Parse & Validate KTP Population Identification Number (NIK).
 * Implemented by @awmanoj 2024 
 * Inspired and ported from original code in Javascript by @bachors 2018 | https://github.com/bachors/nik_parse.js
 *******************************************************/
// Package nikparser provides functions to parse Indonesian National Identification Number (NIK).
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

// This global variable is a UGLY hack for skipping province data in the validity of NIK 
// for the cases where interest is only in gathering DOB and Gender 
var DoNotValidateGeoData = false 

// NIKInfo contains the parsed information from a NIK.
type NIKInfo struct {
	Valid	 	 				bool 
	SkippedGeoDataValidation 	bool 
	Province     				string // provinsi
	District     				string // kabkot
	SubDistrict  				string // kecamatan
	DateOfBirth  				string
	MonthOfBirth 				string
	YearOfBirth  				string
	KodePOS      				string
	Gender       				string
}

// ParseNIK parses an Indonesian National Identification Number (NIK) and returns the parsed information.
//
// The NIK should be a 16-digit string. The function extracts the province, district, sub-district,
// date of birth, gender, and unique code from the NIK.
//
// Example NIK: "3201010201980001"
//
// The date of birth is assumed to be in the format DDMMYY. If the date part is invalid, an error is returned.
//
// Gender is determined by the date of birth; if the day is greater than 40, the gender is female ("PEREMPUAN"),
// otherwise male ("LAKI-LAKI").
//
// Returns an NIKInfo struct containing the parsed information, or an error if the NIK is invalid.
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
	if !DoNotValidateGeoData && !HasGeoData(provinceCode, districtCode, subDistrictCode) {
		return nil, fmt.Errorf("err province or district or subdistrict not found: %s | %s | %s",
			provinceCode, districtCode, subDistrictCode)
	}

	var provinceName, districtName, subDistrictPOSName, subDistrictName, kodePOS string 
	if !DoNotValidateGeoData {
		// extract names
		provinceName = GetProvince(provinceCode)
		districtName = GetDistrict(districtCode)
		subDistrictPOSName = GetSubDistrict(subDistrictCode)

		// separate and extract subdistrict data and POS
		splits := strings.Split(strings.ToUpper(subDistrictPOSName), " -- ")
		subDistrictName = splits[0]
		kodePOS = splits[1]
	}

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
		Valid: 		  				true, 
		SkippedGeoDataValidation: 	DoNotValidateGeoData, 
		Province:     				provinceName,
		District:     				districtName,
		SubDistrict:  				subDistrictName,
		DateOfBirth:  				fmt.Sprintf("%s", dateOfBirth),
		MonthOfBirth: 				monthOfBirth,
		YearOfBirth:  				yearOfBirth,
		KodePOS:      				kodePOS,
		Gender:       				gender,
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
