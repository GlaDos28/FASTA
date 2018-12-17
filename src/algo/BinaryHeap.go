package algo

/* --- */

// Basic Binary heap. Is a tree, every node is lesser (by score) than its children.
// Stored as an array, children of i-th node have indices (2i) and (2i + 1).
// Thus root heap element has index 0 and is minimal element of entire heap.
type BinaryHeap struct {
    data []*FastaResultEntry
    size int
}

/* --- */

// Constructs new heap of given size. Initially each node is dummy and filled with entry of score -1,
// i.e. score of real element can not be lesser than zero.
func NewHeapOfSize(size int) *BinaryHeap {
    data := make([]*FastaResultEntry, size)

    for i := 0; i < size; i += 1 {
        data[i] = &FastaResultEntry{ Score: -1 }
    }

    return &BinaryHeap{
        data: data,
        size: size,
    }
}

// Tries to update heap elements, whether given entry
// has score greater than minimal score of heap (i.e. than score of root element heap[0]).
func (heap *BinaryHeap) Update(entry *FastaResultEntry) {
    if entry.Score > heap.data[0].Score {
        heap.data[0] = entry
        heap.dig(0)
    }
}

// 'Digs' element of given index. 'Digging' means correcting heap structure,
// i.e. i-th element can be greater than one of its child, and if it is so,
// it must be swapped with minimal of its children, to correct heap order.
// Procedure is recursive.
func (heap *BinaryHeap) dig(i int) {
    if 2 * i >= heap.size {
        return
    }

    minInd := 2 * i

    if 2 * i + 1 < heap.size && heap.data[2 * i + 1].Score < heap.data[2 * i].Score {
        minInd = 2 * i + 1
    }

    if heap.data[i].Score > heap.data[minInd].Score {
        heap.data[i], heap.data[minInd] = heap.data[minInd], heap.data[i]
        heap.dig(minInd)
    }
}

// Returns root element and removes it from heap.
// Procedure swaps root element with last element and then shrinks heap size,
// finally digging new root element to correct heap order.
func (heap *BinaryHeap) extractRoot() *FastaResultEntry {
    root := heap.data[0]

    heap.data[0], heap.data[heap.size - 1] = heap.data[heap.size - 1], heap.data[0]
    heap.size -= 1

    heap.dig(0)

    return root
}

// Extracts heap elements in descending order.
// Note that this operation destroys heap order completely, so
// it should be called once and heap can not be used after.
func (heap *BinaryHeap) ExtractSorted() []FastaResultEntry {
    // If element score is equal to -1, dummy element is found
    // and it should not be taken into resulting array.
    for heap.size > 0 && heap.data[0].Score == -1 {
        heap.extractRoot()
    }

    result := make([]FastaResultEntry, heap.size)

    for i := heap.size - 1; heap.size > 0; i -= 1 {
        result[i] = *heap.extractRoot()
    }

    return result
}
