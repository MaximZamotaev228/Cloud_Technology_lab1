package main

import (
 "fmt"
 "sort"
)

// Функция сортировки
func Sorting(nums []int, sortEvenAsc bool, sortOddAsc bool) {
 // Разделение на четные и нечетные числа
 var even, odd []int
 for _, num := range nums {
  if num%2 == 0 {
   even = append(even, num) // Четные числа
  } else {
   odd = append(odd, num) // Нечетные числа
  }
 }

 // Сортировка четных чисел
 if sortEvenAsc {
  sort.Ints(even) // По возрастанию
 } else {
  sort.Sort(sort.Reverse(sort.IntSlice(even))) // По убыванию
 }

 // Сортировка нечетных чисел
 if sortOddAsc {
  sort.Ints(odd) // По возрастанию
 } else {
  sort.Sort(sort.Reverse(sort.IntSlice(odd))) // По убыванию
 }

 // Вывод результата: сначала четные, потом нечетные
 fmt.Println(append(even, odd...))
}

func main() {
 // Пример массива чисел
 nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

 // Параметры сортировки
 sortEvenAsc := false // Четные числа по убыванию
 sortOddAsc := true   // Нечетные числа по возрастанию

 // Вызов функции сортировки
 Sorting(nums, sortEvenAsc, sortOddAsc)
}
 
