[![PkgGoDev](https://pkg.go.dev/badge/github.com/awmanoj/nikparse)](https://pkg.go.dev/github.com/awmanoj/nikparse)

# nikparse

## Parse & Validate KTP Population Identification Number (NIK).

<img width="565" alt="Screenshot 2024-05-24 at 8 05 24 PM" src="https://github.com/awmanoj/nikparse/assets/1171470/2504ee40-e057-42c3-bc49-40d292aa6e67">

A NIK is a 16-digit string that encodes the holder's province, district, sub-district, date of birth, and gender. `ParseNIK` decodes those fields and validates them. The geographic lookup table is embedded in the package, so there is no external data dependency at runtime.

## Install

```
go get github.com/awmanoj/nikparse
```

## Usage

```go
import "github.com/awmanoj/nikparse"
```

```go
	info, err := nikparse.ParseNIK(nik)
	if err != nil {
		log.Println("err parsing nik", err)
		return
	}

	fmt.Printf("%s|%s-%s-%s|%s|%s|%s|%s|%s\n", nik, info.DateOfBirth, info.MonthOfBirth, info.YearOfBirth, info.Gender,
		info.Province, info.District, info.SubDistrict, info.KodePOS)
```

`ParseNIK` returns an `*NIKInfo`:

| Field | Description |
| --- | --- |
| `Valid` | true when parsing succeeded |
| `SkippedGeoDataValidation` | true when geo-data validation was skipped (see below) |
| `Province`, `District`, `SubDistrict` | decoded geographic names (empty when geo validation is skipped) |
| `DateOfBirth`, `MonthOfBirth`, `YearOfBirth` | date of birth (day, month, 4-digit year) |
| `KodePOS` | postal code (empty when geo validation is skipped) |
| `Gender` | `LAKI-LAKI` (male) or `PEREMPUAN` (female) |

### Note on the year

The NIK stores only a two-digit year, which is inherently ambiguous. The century is inferred relative to the current year: a two-digit year at or before the current year is treated as `20xx`, otherwise `19xx`. The result therefore depends on when the code runs.

### Skipping geo-data validation

When you only need the date of birth and gender — for example when the geographic digits come from unreliable OCR — set the package-level flag to bypass the geo lookup and its validation:

```go
	nikparse.DoNotValidateGeoData = true
	info, err := nikparse.ParseNIK(nik) // Province/District/SubDistrict/KodePOS are left empty
```

This is a package-global flag and is not safe to toggle concurrently from multiple goroutines.

## Example CLI

```
go run ./cmd/nikparse-example 3201010201980001
```

## Testing

```
go test ./...
```
