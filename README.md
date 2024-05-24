[![PkgGoDev](https://pkg.go.dev/badge/github.com/awmanoj/nikparse)](https://pkg.go.dev/github.com/awmanoj/nikparse)

# nikparse

## Parse & Validate KTP Population Identification Number (NIK).

<img width="565" alt="Screenshot 2024-05-24 at 8 05 24 PM" src="https://github.com/awmanoj/nikparse/assets/1171470/2504ee40-e057-42c3-bc49-40d292aa6e67">

## Usage 

```
import "github.com/awmanoj/nikparse"
```

```
	info, err := nikparse.ParseNIK(nik)
	if err != nil {
		log.Println("err parsing nik", err)
		return
	}

	fmt.Printf("%s|%s-%s-%s|%s|%s|%s|%s|%s\n", nik, info.DateOfBirth, info.MonthOfBirth, info.YearOfBirth, info.Gender,
		info.Province, info.District, info.SubDistrict, info.KodePOS)
```

