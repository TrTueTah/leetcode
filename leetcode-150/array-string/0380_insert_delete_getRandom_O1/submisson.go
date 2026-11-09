package insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	list  []int
	store map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		list:  make([]int, 0),
		store: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.store[val]; exist {
		return false
	}

	this.list = append(this.list, val)
	index := len(this.list) - 1
	this.store[val] = index
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exist := this.store[val]

	if !exist {
		return false
	}

	last := this.list[len(this.list)-1]
	this.list[index] = last
	this.store[last] = index
	this.list = this.list[:len(this.list)-1]
	delete(this.store, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.list))
	return this.list[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
