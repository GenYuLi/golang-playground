package util

type Color uint8

const (
	Red   Color = iota
	Green       = 1 << iota
	Blue        = 1 << iota
)
