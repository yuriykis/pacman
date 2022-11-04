package main

type Board struct {
	positions []*Position
}

var boardView = [][]string{
	{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
	{"#", " ", "#", " ", " ", "#", " ", " ", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", " ", " ", "#", " ", " ", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", "#", "#", "#", " ", " ", "#", "#", "#", "#", "#", "#", "#", " ", " ", "#", "#", "#"},
	{"#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", "#"},
	{"#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", "#", "#", "#", "#", " ", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", " ", " ", " ", "#", " ", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", " ", " ", " ", "#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", " ", " ", " ", "#", "#", "#", "#", "#", "#", "#", "#", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", " ", "#", "#", "#", "#"},
	{"#", " ", "#", " ", " ", " ", "#", " ", " ", " ", " ", " ", " ", "#", " ", " ", " ", "#", " ", "#"},
	{"#", " ", "#", " ", " ", " ", "#", " ", " ", " ", " ", " ", " ", "#", " ", " ", " ", "#", " ", "#"},
	{"#", " ", "#", " ", "#", "#", "#", "#", "#", " ", " ", " ", " ", "#", " ", " ", " ", "#", " ", "#"},
	{"#", " ", "#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", " ", " ", "#", " ", "#"},
	{"#", " ", "#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#"},
	{"#", " ", "#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", "#"},
	{"#", " ", "#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", " ", " ", "#", " ", "#"},
	{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
}
