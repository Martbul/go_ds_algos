package main

type MinHeap struct {
	length int //used for insertion and deletion
	data   []int
}

func (mh *MinHeap) inseart(value int) {
	mh.data[mh.length] = value
	mh.heapifyUp(mh.length)
	mh.length++
}

func (mh *MinHeap) delete() int {
	if mh.length == 0 {
		return -1
	}

	out := mh.data[0]
	if mh.length == 1 {
		mh.data = []int{}
		return out
	}

	mh.length--
	mh.data[0] = mh.data[mh.length]
	mh.heapifyDown(0)
	return out

}
func (mh *MinHeap) heapifyDown(index int) {
	lIndx := leftChild(index)
	rIndx := rightChild(index)

	if lIndx >= mh.length || lIndx >= mh.length {
		return
	}

	lV := mh.data[lIndx]
	rV := mh.data[rIndx]
	v := mh.data[index]
	if lV > rV && v > rV {
		mh.data[index] = rV
		mh.data[rIndx] = v
		mh.heapifyDown(rIndx)
	} else if rV > lV && v > lV {
		mh.data[index] = lV
		mh.data[lIndx] = v
		mh.heapifyDown(lIndx)
	}

}
func (mh *MinHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}

	p := parent(index)
	paretV := mh.data[p]
	v := mh.data[index]

	if paretV > v {
		mh.data[index] = paretV
		mh.data[p] = v
		mh.heapifyUp(p)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}
