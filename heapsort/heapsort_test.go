package heapsort

import (
    "testing"
)

func TestHeapsort(t *testing.T) {
    k := []int{4, 6, 1, 3, 2}
    h := Heapsort{}
    for _, v := range k {
        h.push(v)
    }
    assertion := []int{1, 2, 3, 4, 6}
    for i := range k {
        v := h.pop()
        if v != assertion[i] {
            t.Errorf("error! %v!=%v", v, assertion[i])
        }
    }
}
