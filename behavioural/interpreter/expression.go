package main

type Expression interface {
	Interpret() int
}