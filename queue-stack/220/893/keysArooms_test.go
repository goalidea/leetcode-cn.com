package keysArooms

import "testing"

func TestKeysAndRooms(t *testing.T) {
	rooms := [][]int{
		{1},
		{2},
		{3},
		{},
	}

	r := canVisitAllRooms(rooms)

	if !r {
		t.Error("visit room error")
	}
}
