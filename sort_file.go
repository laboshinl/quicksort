package main

import(
 "fmt"
 "math/rand"
 "bufio"
 "strings"
 "os"
 "io"
)

func swap(arr []byte, i int, j int){
  arr[i],arr[j]=arr[j],arr[i]
}

func mysort(arr []byte,low int, high int){
  if low < high {
    p := partition(arr, low, high)
    mysort(arr, p + 1, high)
    mysort(arr, low, p - 1)
  }
}

func partition(arr []byte , low int, high int) int {
  p := arr[high]
  i := low - 1
  for  j := low; j < high + 1; j ++ {
    if arr[j] <= p {
      if i < j {
        swap(arr, i + 1, j)
      }
      i ++
    }
  }
  return i
}

func main(){
 array := make([]byte, 10)
 rand.Read(array)
 f, _ := os.Create("to_sort.txt")
 w := bufio.NewWriter(f)
 fmt.Fprint(w, strings.Trim(fmt.Sprint(array), "[]"))
 w.Flush()
 var number byte
 for{
   _, err := fmt.Fscanf(f, "%d ", &number)
   if err != nil {
     if err == io.EOF {
       break
     }
     fmt.Println(err)
     os.Exit(1)
   }
   array = append(array, number)
 }
 defer f.Close()
 mysort(array, 0, len(array) - 1)
 fmt.Println(array)
}
