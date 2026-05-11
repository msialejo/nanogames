package main

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Rect struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	W float32 `json:"w"`
	H float32 `json:"h"`
}
