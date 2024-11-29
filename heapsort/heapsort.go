package heapsort

type Heapsort struct {
    heap []int
}

func (h *Heapsort) push(num int) {
    h.heap = append(h.heap, num)
    h.bubbleUp(len(h.heap) - 1)
}

func (h *Heapsort) bubbleUp(index int) {
    parentIndex := (index - 1) / 2
    parentValue := h.heap[parentIndex]
    if h.heap[index] < parentValue {
        h.heap[index], h.heap[parentIndex] = h.heap[parentIndex], h.heap[index]
        h.bubbleUp(parentIndex)
    }
}

func (h *Heapsort) pop() int {
    res := h.heap[0]
    h.heap[0] = h.heap[len(h.heap)-1]
    h.heap = h.heap[:len(h.heap)-1]
    h.bubbleDown(0)
    return res
}

func (h *Heapsort) bubbleDown(index int) {
    leftChildIndex := (index * 2) + 1
    rightChildIndex := (index * 2) + 2
    smallest := index
    heapSize := len(h.heap)
    if leftChildIndex < heapSize && h.heap[leftChildIndex] < h.heap[smallest] {
        smallest = leftChildIndex
    }
    if rightChildIndex < heapSize && h.heap[rightChildIndex] < h.heap[smallest] {
        smallest = rightChildIndex
    }
    if smallest != index {
        h.heap[smallest], h.heap[index] = h.heap[index], h.heap[smallest]
        h.bubbleDown(smallest)
    }
}
