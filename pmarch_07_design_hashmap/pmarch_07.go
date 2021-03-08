package pmarch_07

type Pair struct {
	Key   int
	Value int
}

type MyHashMap struct {
	m []Pair
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	for i, p := range m.m {
		if p.Key == key {
			m.m[i] = Pair{key, value}
			return
		}
	}
	m.m = append(m.m, Pair{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	for _, p := range m.m {
		if p.Key == key {
			return p.Value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	for i, p := range m.m {
		if p.Key == key {
			m.m = append(m.m[0:i], m.m[i+1:]...)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
