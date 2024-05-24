package nikparse

import (
	"testing"
)

func TestParseNIK(t *testing.T) {
	tests := []struct {
		name      string
		nik       string
		expected  *NIKInfo
		expectErr bool
	}{
		{
			name: "Valid NIK",
			nik:  "3201010201980001",
			expected: &NIKInfo{
				Province:     "JAWA BARAT",
				District:     "KAB. BOGOR",
				SubDistrict:  "CIBINONG",
				DateOfBirth:  "02",
				MonthOfBirth: "01",
				YearOfBirth:  "1998",
				KodePOS:      "43271",
				Gender:       "LAKI-LAKI",
			},
			expectErr: false,
		},
		{
			name:      "Invalid NIK length",
			nik:       "320101020198000",
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "Invalid NIK non-numeric",
			nik:       "32010X0201980001",
			expected:  nil,
			expectErr: true,
		},
		{
			name: "Invalid birth date",
			nik:  "3201013202980001", // Invalid date
			expected:  nil,
			expectErr: true,
		},
		{
			name: "Edge case: leap year",
			nik:  "3201012902960001", // 29 Feb 1996
			expected: &NIKInfo{
				Province:     "JAWA BARAT",
				District:     "KAB. BOGOR",
				SubDistrict:  "CIBINONG",
				DateOfBirth:  "29",
				MonthOfBirth: "02",
				YearOfBirth:  "1996",
				KodePOS:      "43271",
				Gender:       "LAKI-LAKI",
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseNIK(tt.nik)
			if (err != nil) != tt.expectErr {
				t.Errorf("ParseNIK() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if tt.expected != nil && result != nil {
				if *result != *tt.expected {
					t.Errorf("ParseNIK() = %+v, expected %+v", result, tt.expected)
				}
			} else if result != tt.expected {
				t.Errorf("ParseNIK() = %+v, expected %+v", result, tt.expected)
			}
		})
	}
}
