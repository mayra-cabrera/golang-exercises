package leap

// TestVersion refers to testing version number.
const TestVersion = 1

// IsLeapYear reports if a year is a leap year.
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
