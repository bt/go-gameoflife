package main
import "fmt"

type Grid struct {
  cell []Cell
}

type Cell struct {
  alive bool
}

func main() {
  fmt.Printf("Conway's Game of Life");
}
