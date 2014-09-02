package main

import (
     "fmt"
     "math"
     "sort"
     "image"
     "os"
     _ "image/gif"
     _ "image/jpeg"
     _ "image/png"
     gwc "github.com/jyotiska/go-webcolors"
     "github.com/nfnt/resize"
)

// This method finds the closest color for a given RGB tuple and returns the name of the color in given mode

func FindClosestColor(RequestedColor []int, mode string) string {
     MinColors := make(map[int]string)
     var ColorMap map[string]string
     if mode == "css3" {
          ColorMap = gwc.CSS3NamesToHex
     } else {
          ColorMap = gwc.HTML4NamesToHex
     }

     for name, hexcode := range ColorMap {
          rgb_triplet := gwc.HexToRGB(hexcode)
          rd := math.Pow(float64(rgb_triplet[0] - RequestedColor[0]), float64(2))
          gd := math.Pow(float64(rgb_triplet[1] - RequestedColor[1]), float64(2))
          bd := math.Pow(float64(rgb_triplet[2] - RequestedColor[2]), float64(2))
          MinColors[int(rd + gd + bd)] = name
     }

     keys := make([]int, 0, len(MinColors))
     for key := range MinColors {
          keys = append(keys, key)
     }
     sort.Ints(keys)
     return MinColors[keys[0]]
}

func ReverseMap(m map[string]int) map[int]string {
    n := make(map[int]string)
    for k, v := range m {
        n[v] = k
    }
    return n
}

func main() {
     reader, err := os.Open("images/image1.png")
     if err != nil {
          fmt.Fprintf(os.Stderr, "%v\n", err)
     }

     image, _, err := image.Decode(reader)
     if err != nil {
          fmt.Fprintf(os.Stderr, "%s", err)
     }

     image = resize.Resize(100, 0, image, resize.Bilinear)
     bounds := image.Bounds()

     ColorCounter := make(map[string]int)
     Limit := 3 // Limiting how many colors to be displayed in output
     TotalPixels := bounds.Max.X * bounds.Max.Y

     for i := 0; i <= bounds.Max.X; i++ {
          for j := 0; j <= bounds.Max.Y; j++ {
               pixel := image.At(i, j)
               red, green, blue, _ := pixel.RGBA()
               RGBTuple := []int{int(red/255), int(green/255), int(blue/255)}
               ColorName := FindClosestColor(RGBTuple, "css21")
               _, present := ColorCounter[ColorName]
               if present {
                    ColorCounter[ColorName] += 1
               } else {
                    ColorCounter[ColorName] = 1
               }
          }
     }

     keys := make([]int, 0, len(ColorCounter))
     for _, val := range ColorCounter {
          keys = append(keys, val)
     }
     sort.Ints(keys)

     ReverseColorCounter := ReverseMap(ColorCounter)

     for _, val := range keys[(len(keys) - Limit):] {
          fmt.Printf("%s %.2f%%\n", ReverseColorCounter[val], ((float64(val) / float64(TotalPixels)) * 100))
     }
}
