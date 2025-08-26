package main

import "fmt"

type SortStrategy interface{ sort([]int) }

type BubbleSort struct{}

type QuickSort struct{}

func (BubbleSort) sort(arr []int) { fmt.Println("BubbleSort:", arr) }

func (QuickSort) sort(arr []int) { fmt.Println("QuickSort:", arr) }

type Sorter struct{ strategy SortStrategy }

func (s *Sorter) setStrategy(st SortStrategy) { s.strategy = st }

func (s *Sorter) sort(arr []int) { s.strategy.sort(arr) }
