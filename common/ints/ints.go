package ints

import (
	"math"

	"github.com/thomaso-mirodin/intmath/intgr"
)

func Min(it Iter) int {
	if !it.HasNext() {
		return 0
	}
	min := math.MaxInt
	for {
		if !it.HasNext() {
			break
		}
		min = intgr.Min(it.Next(), min)
	}

	return min
}

func Max(it Iter) int {
	if !it.HasNext() {
		return 0
	}
	max := math.MinInt
	for {
		if !it.HasNext() {
			break
		}
		max = intgr.Max(it.Next(), max)
	}
	return max
}
func ArrayMin(ints []int) int {
	min := math.MaxInt
	for _, i := range ints {
		min = intgr.Min(i, min)
	}

	return min
}

func ArrayMax(ints []int) int {
	max := math.MinInt
	for _, i := range ints {
		max = intgr.Max(i, max)
	}
	return max
}

func Range(from, to int) []int {
	var res []int
	if from <= to {
		for i := from; i <= to; i++ {
			res = append(res, i)
		}
	}
	return res
}

func Sum(it Iter) int {
	var res int
	for {
		if !it.HasNext() {
			break
		}
		res += it.Next()
	}
	return res
}

func Map(it Iter, fn func(i int) int) Iter {
	res := []int{}
	for {
		if !it.HasNext() {
			break
		}
		res = append(res, fn(it.Next()))
	}
	return FromArray(&res)
}

func Reduce(it Iter, base int, fn func(a, b int) int) int {
	res := base
	for {
		if !it.HasNext() {
			break
		}
		res = fn(res, it.Next())
	}
	return res
}

type Iter interface {
	Next() int
	HasNext() bool
	Vec() Vec
	Sum() int
	Map(fn func(i int) int) Iter
	Reduce(base int, fn func(a, b int) int) int
	Min() int
	Max() int
}

type Vec interface {
	Array() []int
	Iter() Iter
}

type arrBackedVec struct {
	arr *[]int
}

type arrBackedIter struct {
	arr *[]int
	cur int
}

func (it *arrBackedIter) Next() int {
	res := (*it.arr)[it.cur]
	it.cur++
	return res
}

func (it *arrBackedIter) HasNext() bool {
	return it.cur < len(*it.arr)
}

func (it *arrBackedIter) Sum() int {
	return Sum(it)
}

func (it *arrBackedIter) Min() int {
	return Min(it)
}

func (it *arrBackedIter) Max() int {
	return Max(it)
}

func (it *arrBackedIter) Map(fn func(i int) int) Iter {
	return Map(it, fn)
}

func (it *arrBackedIter) Reduce(base int, fn func(a, b int) int) int {
	return Reduce(it, base, fn)
}

func (it *arrBackedIter) Vec() Vec {
	return &arrBackedVec{it.arr}
}

func (it *arrBackedVec) Iter() Iter {
	return &arrBackedIter{arr: it.arr, cur: 0}
}

func (it *arrBackedVec) Array() []int {
	return *it.arr
}

func FromArray(is *[]int) Iter {
	return &arrBackedIter{arr: is}
}

func ArrayVec(is *[]int) Vec {
	return &arrBackedVec{arr: is}
}

type rangeBackedIter struct {
	from, to int
	cur      int
}

type rangeBackedVec struct {
	from, to int
}

func (it *rangeBackedIter) Next() int {
	res := it.cur
	it.cur++
	return res
}

func (it *rangeBackedIter) HasNext() bool {
	return it.cur <= it.to
}

func (it *rangeBackedIter) Sum() int {
	return Sum(it)
}

func (it *rangeBackedIter) Map(fn func(i int) int) Iter {
	return Map(it, fn)
}

func (it *rangeBackedIter) Reduce(base int, fn func(a, b int) int) int {
	return Reduce(it, base, fn)
}

func FromRange(from, to int) Iter {
	return &rangeBackedIter{from: from, to: to, cur: from}
}

func (it *rangeBackedIter) Min() int {
	return it.from
}

func (it *rangeBackedIter) Max() int {
	return it.to
}

func (it *rangeBackedIter) Vec() Vec {
	return &rangeBackedVec{from: it.from, to: it.to}
}

func (it *rangeBackedVec) Iter() Iter {
	return FromRange(it.from, it.to)
}

func (it *rangeBackedVec) Array() []int {
	res := []int{}
	for i := it.from; i < it.to; i++ {
		res = append(res, i)
	}
	return res
}
