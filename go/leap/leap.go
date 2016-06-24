package leap

const testVersion = 2

// IsLeapYear reports if the received year is a leap year in the Gregorian calendar.
func IsLeapYear(y int) bool {
	return (y%4 == 0) && ((y%100 != 0) || (y%400 == 0))
}
