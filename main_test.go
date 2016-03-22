package main

import (
	"testing"
	"image"
	"image/color"
)


func TestMain( t *testing.T ){
	input:= image.NewRGBA( image.Rect(0,0,2,2))
	test := [][]uint8{
		{ 100,100,100,100,100},
		{ 100,100,0,0,50},
		{ 0,0,0,0,0},
		{ 0,45,55,100,50},
	}
	for i:=0; i < len(test); i++ {

		input.Set(0, 0, color.RGBA{test[i][0], test[i][0], test[i][0], test[i][0]})
		input.Set(0, 1, color.RGBA{test[i][1], test[i][1], test[i][1], test[i][1]})
		input.Set(1, 0, color.RGBA{test[i][2], test[i][2], test[i][2], test[i][2]})
		input.Set(1, 1, color.RGBA{test[i][3], test[i][3], test[i][3], test[i][3]})
		output := SampleRGBAArea(input, 0, 0, 1, 1)
		expected := color.RGBA{test[i][4], test[i][4], test[i][4], test[i][4]}

		if ( output.R != expected.R ) {
			t.Error("num: ",i,"Red Expected: ", expected.R, " got ", output.R)
		}
		if ( output.G != expected.G ) {
			t.Error("num: ",i,"Green Expected: ", expected.G, " got ", output.G)
		}
		if ( output.B != expected.B ) {
			t.Error("num: ",i,"Green Expected: ", expected.B, " got ", output.B)
		}
		if ( output.A != expected.A ) {
			t.Error("num: ",i,"Green Expected: ", expected.A, " got ", output.A)
		}
	}
}


