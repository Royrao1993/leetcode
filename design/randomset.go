package design

import "math/rand"

type RandomizedSet struct {
	Map  map[int]int
	List []int
}

func NewRandomizedSet() *RandomizedSet {
	return &RandomizedSet{map[int]int{}, []int{}}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.Map[val]; ok {
		return false
	} else {
		// insert
		s.Map[val] = len(s.List)
		s.List = append(s.List, val)
		return true
	}
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.Map[val]; ok {
		// take care of that if the value to be removed is in the middle of the list,
		// move the value to the end of the list first in case of the mismatch of index
		idx := s.Map[val]
		last := len(s.List) - 1
		if idx < last {
			lastVal := s.List[last]
			s.Map[lastVal], s.Map[val] = s.Map[val], s.Map[lastVal]
			s.List[idx], s.List[last] = s.List[last], s.List[idx]
		}
		delete(s.Map, val)
		s.List = s.List[:last]
		return true
	} else {
		return false
	}
}

func (s *RandomizedSet) GetRandom() int {
	return s.List[rand.Intn(len(s.List))]
}
