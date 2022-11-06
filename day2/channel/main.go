import "fmt"

func main()  {
	var ch chan int 
	fmt.Printf("ch is nul %t \n",ch == nil)
	fmt.Printf("len is nul %d \n",len(ch))

	ch = make(chan int, 10)

	ch <- 1
	close(chan type)

	for v := range ch {
		fmt.Println(v)
	}

}