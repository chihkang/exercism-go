package leap

const testVersion = 3

// IsLeapYear is a function that takes in an int and returns a bool if the input year is a leap year
func IsLeapYear(year int) bool {
	var res = false
	if year%4 == 0 {
		if year%100 != 0 {
			res = true
		} else {
			if year%400 == 0 {
				res = true
			}
		}
	}
	return res
}
