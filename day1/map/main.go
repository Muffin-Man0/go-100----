package main
import "fmt"
func testMap()  {
	
}

func main()  {
	var m map[string]int 
	fmt.Println(len(m))
	m = make(map[string]int) //
	m = make(map[string]int, 10)

	m["HAO"] = 123
	// delete(m,"HAO")
	_,ok := m["HAO"]
	
	if ok {
		fmt.Println("有的")
	}
	for key,value := range m{
		m[key]+=1
		fmt.Printf("key %s,value %d\n",key,value)
	}
	fmt.Println(m)
}

/**
123 “123”  hash("123") = 1421242123 %7 都是等于3
123 “123”  hash("124") = 52123421422  %7 都是等于3
*/