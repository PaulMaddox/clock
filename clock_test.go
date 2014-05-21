package clock

import (
	"log"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {

	current, err := Get()
	if err != nil || current == nil {
		t.Fatalf("Could not obtain current time (%s)", err)
	}

	log.Printf("Got Time: %s", current)

}

func TestSetTime(t *testing.T) {

	current, err := Get()
	if err != nil || current == nil {
		t.Fatalf("Could not obtain current time (%s)", err)
	}

	target := current.Add(1 * time.Hour)

	log.Printf("Setting time to: %d:%d", target.Hour(), target.Minute())

	err = Set(target)
	if err != nil {
		t.Fatalf("Could not set the time (%s)", err)
	}

	modified, err := Get()
	if err != nil || current == nil {
		t.Fatalf("Could not obtain modified time (%s)", err)
	}

	if target.Hour() != modified.Hour() {
		t.Fatalf("Clock was not modified (expecting %d:%d but got %d:%d)", target.Hour(), target.Minute(), modified.Hour(), target.Minute())
	}

	reverted := current.Add(-1 * time.Hour)
	err = Set(reverted)
	if err != nil {
		t.Fatalf("Could not revert time (%s)", err)
	}

}
