// tinygo build -o main.wasm -target wasm ./main.go
// https://wasmbyexample.dev/examples/reading-and-writing-graphics/reading-and-writing-graphics.go.en-us.html
package main

const CHECKERBOARD_SIZE int = 20;

const BUFFER_SIZE int = CHECKERBOARD_SIZE * CHECKERBOARD_SIZE * 4;
var graphicsBuffer [BUFFER_SIZE]uint8;

func main() {}

//export getGraphicsBufferPointer
func getGraphicsBufferPointer() *[BUFFER_SIZE]uint8 {
  return &graphicsBuffer
}

//export getGraphicsBufferSize
func getGraphicsBufferSize() int {
  return BUFFER_SIZE;
}

//export generateCheckerBoard
func generateCheckerBoard(
  darkValueRed uint8,
  darkValueGreen uint8,
  darkValueBlue uint8,
  lightValueRed uint8,
  lightValueGreen uint8,
  lightValueBlue uint8,
) {
  for y := 0; y < CHECKERBOARD_SIZE; y++ {
    for x := 0; x < CHECKERBOARD_SIZE; x++ {
      isDarkSquare := true;

      if y % 2 == 0 {
        isDarkSquare = false;
      }

      if x % 2 == 0 {
        isDarkSquare = !isDarkSquare;
      }

      squareValueRed := darkValueRed;
      squareValueGreen := darkValueGreen;
      squareValueBlue := darkValueBlue;
      if !isDarkSquare {
      squareValueRed = lightValueRed;
      squareValueGreen = lightValueGreen;
      squareValueBlue = lightValueBlue;
      }

      squareNumber := (y * CHECKERBOARD_SIZE) + x;
      squareRgbaIndex := squareNumber * 4;

      graphicsBuffer[squareRgbaIndex + 0] = squareValueRed;
      graphicsBuffer[squareRgbaIndex + 1] = squareValueGreen;
      graphicsBuffer[squareRgbaIndex + 2] = squareValueBlue;
      graphicsBuffer[squareRgbaIndex + 3] = 255;
    }
  }
}
