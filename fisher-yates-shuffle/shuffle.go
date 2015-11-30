package main

import(
  "fmt"
  "math/rand"
  "time"
)

func shuffle(list *[]string) {
  currentList := *list
  rand.Seed(time.Now().Unix())

  for remaining := len(currentList); remaining > 0; remaining-- {

    random := rand.Intn(remaining)
    currentList[random], currentList[remaining - 1] = currentList[remaining - 1], currentList[random]
  }
}

func main(){
  list := []string{"a", "b", "c", "d", "e", "f", "g"}
  shuffle(&list)
  fmt.Println(list)
}
