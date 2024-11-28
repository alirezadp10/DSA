package huffman

import "testing"

func TestEncoder(t *testing.T) {
    res := encode("ABBCCCDDDD")
    assertion := "10010110100011111111"
    if res != assertion {
        t.Errorf("got %s, want %s", res, assertion)
    }
}

func TestDecoder(t *testing.T) {
    res := decode("10010110100011111111", map[string]byte{"100": 65, "101": 66, "0": 67, "11": 68})
    assertion := "ABBCCCDDDD"
    if res != assertion {
        t.Errorf("got %s, want %s", res, assertion)
    }
}
