package main

import "testing"

const (
	exampleOne   = "aaaaa-bbb-z-y-x-123[abxyz]"
	exampleTwo   = "a-b-c-d-e-f-g-h-987[abcde]"
	exampleThree = "not-a-real-room-404[oarel]"
	exampleFour  = "totally-real-room-200[decoy]"
)

func TestSectorIDs(t *testing.T) {
	ids := map[string]int{
		exampleOne:   123,
		exampleTwo:   987,
		exampleThree: 404,
		exampleFour:  200,
	}
	for room, id := range ids {
		if result := SectorID(room); result != id {
			t.Errorf("Room %s should have sector ID %d (but has %d)", room, id, result)
		}
	}
}

func TestIsValid(t *testing.T) {
	valid := map[string]bool{
		exampleOne:   true,
		exampleTwo:   true,
		exampleThree: true,
		exampleFour:  false,
	}
	for room, valid := range valid {
		if result := IsValid(room); result != valid {
			t.Errorf("Room %s validity should be %b, but was %b", room, valid, result)
		}
	}
}

func TestDecrypt(t *testing.T) {
	var (
		encrypted = "qzmt-zixmtkozy-ivhz-343[abc]"
		decrypted = "very encrypted name"
	)
	if result := Decrypt(encrypted); result != decrypted {
		t.Errorf("decrypt(%s) == %s != %s", encrypted, result, decrypted)
	}
}
