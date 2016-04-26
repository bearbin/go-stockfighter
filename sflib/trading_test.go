package sflib

import (
	"testing"
)

// TODO: Replace these tests with proper unit tests, and

func TestHeartbeat(t *testing.T) {
	// New client. We don't need an API key for this test.
	cli := NewClient("")
	err := cli.Heartbeat()
	if err != nil {
		t.Error("heartbeat failed - is the API down?")
		return
	}
}

func TestVenueHeartbeat(t *testing.T) {
	// New client. We don't need an API key for this test.
	cli := NewClient("")
	testVenue := "TESTEX"
	response, err := cli.VenueHeartbeat(testVenue)
	if err != nil {
		t.Error("venue heartbeat failed,", err)
		return
	}
	if response.Venue != testVenue {
		t.Error("venue heartbeat gave wrong venue. expected ", testVenue, " got ", response.Venue)
		return
	}
}

func TestVenueStocks(t *testing.T) {
	// New client. We don't need an API key for this test.
	cli := NewClient("")
	testVenue := "TESTEX"
	response, err := cli.VenueStocks(testVenue)
	if err != nil {
		t.Error("venue stocks failed,", err)
		return
	}
	// The test venue should contain only one stock.
	if len(response.Symbols) != 1 {
		t.Error("venue stocks gave wrong number of stocks. expected 1, got ", len(response.Symbols))
		return
	}
	// The test stock should have the symbol "FOOBAR".
	if response.Symbols[0].Symbol != "FOOBAR" {
		t.Error("venue stocks gave wrong symbol for test stock. expected 'FOOBAR', got '", response.Symbols[0].Symbol, "'")
		return
	}
	// The test stock should have the name "Foreign Owned Occluded Bridge
	// Architecture Resources"
	if response.Symbols[0].Name != "Foreign Owned Occluded Bridge Architecture Resources" {
		t.Error("venue stocks gave wrong name for test stock. expected 'Foreign Owned Occluded Bridge Architecture Resources', got '", response.Symbols[0].Name, "'")
		return
	}
}
