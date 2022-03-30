package pkg

func Max (ch chan int64, nums []int64) {

	var max int64

	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	ch <- max
}

func Divide(input []int64, res int) (max int64) {

	var maxx []int64
	var maxs []int64
	var left int
	right := len(input)/res

	for i := 1; i <= res; i++ {
		maxs = input[left:right]

		ch := make (chan int64)

		go Max(ch, maxs)

		maxx = append(maxx, <-ch)

		left = right

		right = left + (len(input)/res)
	}

	ch1 := make(chan int64)

	go Max(ch1, maxx)

	max = <-ch1

	return max
}







// 	for i := 1; i < res; i++{

// 		ch1 := make(chan int64)

// 		go Max(ch1, input[i*limit:(i+int64(1))*limit])

// 		num := <- ch1

// 		maxes = append(maxes, num)
// 	}

// 	for _, i := range maxes{
// 		if i > max{
// 			max = i
// 		}
// 	}
// 	return
// }

	// ch1 := make(chan int64)
	// ch2 := make(chan int64)

	// go Max(ch1, input[:len(input)/2])
	// go Max(ch2, input[len(input)/2:])

	// left, right := <-ch1, <-ch2

	// if left > right {
	// 	max = left
	// } else {
	// 	max = right
	// }

	// return



//     for i := 0; i < num; i++ {
//     if i == 0 {
//       go Max(ch1, input[:x])
//       go Max(ch2, input[x:y])
      
//       left, right := <-ch1, <-ch2

//       if left > right {
//         max = left
//       } else {
//         max = right
//       }

//       x = x + n
//       y = y + n
//     }

//     go Max(ch1, input[x:x])
//     go Max(ch2, input[x:y])

//     left, right := <-ch1, <-ch2

//     if right > max {
//       max = right
//     } 
//     if left > max {
//       max = left
//     }

//     x = x + n
//     y = y + n

//     fmt.Println("i")

//     if i == num {
//       break
//     }
//   }


//   return
// }


func Fill(value int64) []int64 {

	var number int64 = 1000

	result := make([]int64, number)

	// result[len(result) - 1] = value

	var i int64 = 1

	for ;i <= number; i++ {
		result[i-1] = i
	}

	return result
}