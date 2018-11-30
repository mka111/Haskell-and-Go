package main

import(
  "image"
  "image/color"
  "math"
  "os"
  "image/png"
  "image/draw"
//  "fmt"
)

func DrawCircle(outerRadius, innerRadius int, outputFile string){
  rectImage := image.NewRGBA(image.Rect(0, 0, 200, 200))
  blackclr := color.RGBA{0,0,0,255}
  whiteclr := color.RGBA{255,255,255,255}
  draw.Draw(rectImage, rectImage.Bounds(), &image.Uniform{whiteclr}, image.ZP, draw.Src)

  //check for all possible points in the inner and outer circle, where distance
  //between two points is equal to outerRadius - innerRadius, set that equal to whiteclr

  xmin := 100-outerRadius
  xmax := 100+outerRadius

  ymin := 100-outerRadius
  ymax := 100+outerRadius

//  radius := (outerRadius - innerRadius)

  x := xmin
  //y:=xmin


  for x <= xmax{
    y := ymin
    for y <= ymax {
      y1 := y-100
      x1 := x-100
      r1 := int(math.Sqrt(float64(float64(x1*x1) + float64(y1*y1))))
      if(r1>= innerRadius && r1 <= outerRadius){
        rectImage.Set(x,y,blackclr)
      }
    //  }
      y = y+1
    }
    x = x+1
  }





  f, _ := os.Create(outputFile)
  png.Encode(f, rectImage)

}

//func main(){
//  DrawCircle(40,20,"out.png")
//}
