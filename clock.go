// +build !windows

// Package clock is a portable library for setting the system clock.
// Supports nanosecond resolution for Linux/Darwin, and second resolution on Windows.
// All time values get/set are UTC.
// Tested on Linux/Darwin/Windows.
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
	"syscall"
	"time"
)

// Get retrieves the current system time, either via syscall.Gettimeofday on Linux/Darwin
// or via kernel32.GetSystemTime() on Windows. It then parses the result into a
// standard time.Time struct.
func Get() (*time.Time, error) {

	var tv syscall.Timeval
	if err := syscall.Gettimeofday(&tv); err != nil {
		return nil, err
	}

	var output time.Time
	seconds, nanoseconds := tv.Unix()
	output = time.Unix(seconds, nanoseconds)

	return &output, nil

}

// Set sets the current system time, either via syscall.Settimeofday on Linux/Darwin
// or via kernel32.SetSystemtime() on Windows.
func Set(input time.Time) error {

	tv := syscall.Timeval{
		Sec:  input.Unix(),
		Usec: int32(input.UnixNano()),
	}

	err := syscall.Settimeofday(&tv)
	if err != nil {
		return err
	}

	return nil

}
