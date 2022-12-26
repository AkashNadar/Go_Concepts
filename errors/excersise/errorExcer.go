package errorparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, min, sec int
}

type TimeParseError struct {
	msg   string
	input string
}

func (e *TimeParseError) Error() string {
	return fmt.Sprintf("msg :%v, input :%v", e.msg, e.input)
}

func ParseTime(input string) (Time, error) {
	compo := strings.Split(input, ":")
	if len(compo) != 3 {
		return Time{}, &TimeParseError{"Length is smaller ", input}
	} else {
		hours, err := strconv.Atoi(compo[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hours :%v", hours), input}
		}
		min, err := strconv.Atoi(compo[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing min :%v", min), input}
		}
		sec, err := strconv.Atoi(compo[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing sec :%v", sec), input}
		}

		if hours > 23 || hours < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("Hours out of range %v", hours), input}
		}
		if min > 59 || min < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("min out of range %v", min), input}
		}
		if sec > 59 || sec < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("Hours out of range %v", sec), input}
		}

		return Time{hours, min, sec}, nil
	}

}
