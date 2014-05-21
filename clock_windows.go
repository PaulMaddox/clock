// +build windows

// Package clock is a portable library for setting the system clock.
// Tested on Linux/Darwin/Windows
//
//  package main
//
//  import (
//		"time"
//		"github.com/PaulMaddox/clock"
//  )
//
//  func main() {
//
//		// Get the time
//		t, err := clock.Get()
//		if err != nil {
//			panic(err)
//		}
//
//		// Set the time
//		t2 := t.Add(12 * time.Hour)
//		if err := clock.Set(t2); err != nil {
//			panic(err)
//		}
//  }
package clock

import (
	"errors"
	"time"

	"github.com/AllenDang/w32"
)

// Get retrieves the current system time, either via syscall.Gettimeofday on Linux/Darwin
// or via kernel32.GetSystemTime() on Windows. It then parses the result into a
// standard time.Time struct.
func Get() (*time.Time, error) {

	// Gets the system time from the kernel32 API
	st := w32.GetSystemTime()

	// Convert the SYSTEMTIME to time.Time
	t := time.Date(
		int(st.Year),
		time.Month(st.Month),
		int(st.Day),
		int(st.Hour),
		int(st.Minute),
		int(st.Second),
		0, time.UTC)

	return &t, nil

}

// Set sets the current system time, either via syscall.Settimeofday on Linux/Darwin
// or via kernel32.SetSystemtime() on Windows.
func Set(input time.Time) error {

	st := &w32.SYSTEMTIME{
		Year:         uint16(input.Year()),
		Month:        uint16(input.Month()),
		DayOfWeek:    uint16(input.Weekday()),
		Day:          uint16(input.Day()),
		Hour:         uint16(input.Hour()),
		Minute:       uint16(input.Minute()),
		Second:       uint16(input.Second()),
		Milliseconds: 0,
	}

	if success := w32.SetSystemTime(st); success != true {
		return errors.New("unable to set system time")
	}

	return nil

}
