# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`nikparse` is a Go library (`github.com/awmanoj/nikparse`) that parses and validates Indonesian KTP National Identification Numbers (NIK). It is ported from the original JavaScript implementation by @bachors. A NIK is a 16-digit string that encodes geographic origin, date of birth, and gender.

## Commands

```bash
go test ./...                       # run all tests
go test -run TestParseNIK           # run a single test
go build ./...                      # build everything
go build -o cmd/nikparse-example/nikparse-example ./cmd/nikparse-example   # build the example CLI
./cmd/nikparse-example/nikparse-example 3201010201980001                   # run CLI on one NIK
./process.sh                        # batch-parse NIKs from GENDER-DOB-FROM-OCR.csv into RESULT.CSV
```

## Architecture

Three source files make up the library (package `nikparse`, at repo root):

- **nik_parse.go** — `ParseNIK(nik string) (*NIKInfo, error)` is the single entry point. It slices the 16-digit string by byte position to extract fields:
  - `[0:2]` province code, `[0:4]` district code, `[0:6]` sub-district code
  - `[6:8]` day of birth — **if > 40, gender is PEREMPUAN (female) and 40 is subtracted** to get the real day; otherwise LAKI-LAKI (male)
  - `[8:10]` month, `[10:12]` two-digit year
  - The century is inferred by comparing the two-digit year against the current year's last two digits (`< current → 20xx`, else `19xx`), so output depends on `time.Now()`.
- **geodata.go** — embeds `assets/data.txt` via `//go:embed` and unmarshals it at `init()` into the package-global `GeoData` (`map[string]map[string]string`). Top-level keys are `provinsi`, `kabkot`, `kecamatan`; each maps a numeric code prefix to a name. For `kecamatan`, the value packs both the sub-district name and postal code as `"NAME -- KODEPOS"`, split on `" -- "` in `ParseNIK`.
- **assets/data.txt** — the embedded JSON geo-code lookup table. It is compiled into the binary; there is no external data dependency at runtime.

## The `DoNotValidateGeoData` flag

`nikparse.DoNotValidateGeoData` is a package-global bool (default `false`). When `true`, `ParseNIK` skips both the geo-data presence check and name/postal-code lookups, returning only DOB and gender. This exists for OCR-derived data where only DOB/gender are wanted and the geo codes may be unreliable. It is a global mutable flag (a deliberate hack), so setting it affects all callers — see `cmd/nikparse-example/main.go`, which sets it and switches output format based on its value.

## Notes

- Tests in `nik_parse_test.go` assert exact geo names/postal codes, so they implicitly depend on the contents of `assets/data.txt`.
- `pc.tsv`, `GENDER-DOB-FROM-OCR.csv`, and `process.sh` are ad-hoc batch-processing artifacts, not part of the library.
