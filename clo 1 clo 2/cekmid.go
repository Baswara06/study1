package main
import "fmt"

func main() {
    var dragma, attic, mine, oboli int
    
    fmt.Scan(&dragma)
    attic = dragma / 60
    mine = dragma / 100
    oboli = dragma / 6
     
    fmt.Print(attic, mine, oboli)
}