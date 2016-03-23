package main

import (
	"testing"
	"image"
	"image/color"
	"math"
	"strconv"
)
func testClr( t *testing.T, desc string, expected,actual uint32 ) {
	if ( expected != actual ){
		t.Error(desc," expected ",expected," got ",actual)
	}
}
func TestResizeImage( t *testing.T ) {
	input:= image.NewRGBA( image.Rect(0,0,2,2))
	output:= image.NewRGBA (image.Rect(0,0,1,1))
	test := [][]uint8{
		{ 100,100,100,100,100},
		{ 100,100,0,0,50},
		{ 0,0,0,0,0},
	}
	for i:=0; i < len(test); i++ {
		input.Set(0, 0, color.RGBA{test[i][0], test[i][0], test[i][0], test[i][0]})
		input.Set(0, 1, color.RGBA{test[i][1], test[i][1], test[i][1], test[i][1]})
		input.Set(1, 0, color.RGBA{test[i][2], test[i][2], test[i][2], test[i][2]})
		input.Set(1, 1, color.RGBA{test[i][3], test[i][3], test[i][3], test[i][3]})

		ResizeImageRGBA(input,output)
		clr:= output.At(0,0)
		r,g,b,a := clr.RGBA()
		expected := uint32(test[i][4])<<8
		testClr( t,"R"+strconv.Itoa(i),expected, r )
		testClr( t,"G"+strconv.Itoa(i),expected, g )
		testClr( t,"B"+strconv.Itoa(i),expected, b )
		testClr( t,"A"+strconv.Itoa(i),expected, a )
	}

}

func TestCalcEdge( t *testing.T ){
	test:= [][]float32 {
		{0,0,10,1},
		{1.1,1.1,10,.9},
		{1.9,1.9,10,.1},
		{9,1.9,10,1},
		{9.1,1.9,10,1},
		{9,1.9,9.6,.6},
		{9.4,1.9,9.6,.6},
	}
	for i:=0; i < len(test); i++{
		f:= CalcEdgeFraction(test[i][0],test[i][1],test[i][2])
		if math.Abs(float64(f-test[i][3]))>0.0001{
			t.Error("CaclEdge fail, expected: ", test[i][3]," got ",f)
		}

	}


}
func TestMain( t *testing.T ){

	input:= image.NewRGBA( image.Rect(0,0,2,2))
	test := [][]uint8{
		{ 100,100,100,100,100},
		{ 100,100,0,0,50},
		{ 0,0,0,0,0},
		{ 0,45,55,100,50},
		{ 0,0,100,100,100},
		{ 0,0,100,100,50},
		{ 100,100,100,100,100 },
		{ 0,0,100,100,66 },
	}
	sample := [][]float32{
		{0,0,1,1},
		{0,0,1,1},
		{0,0,1,1},
		{0,0,1,1},
		{0.5,0,1,1},
		{0,0.5,1,1},
		{0.25,0,1,1},
		{0.25,0,1,1},
	}
	for i:=0; i < len(test); i++ {

		input.Set(0, 0, color.RGBA{test[i][0], test[i][0], test[i][0], test[i][0]})
		input.Set(0, 1, color.RGBA{test[i][1], test[i][1], test[i][1], test[i][1]})
		input.Set(1, 0, color.RGBA{test[i][2], test[i][2], test[i][2], test[i][2]})
		input.Set(1, 1, color.RGBA{test[i][3], test[i][3], test[i][3], test[i][3]})
		output := SampleRGBAArea(input, sample[i][0],sample[i][1],sample[i][2],sample[i][3] )
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


