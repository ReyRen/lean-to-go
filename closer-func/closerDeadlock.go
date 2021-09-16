package main

/*func main() {
	cs := make([](chan int), 10)
	for i := 0; i < len(cs); i++ {
		cs[i] = make(chan int)
	}

	for i := range cs {
		go func(index int) {
			cs[index] <- index
		}(i)
	}

	for i := 0; i < len(cs); i++ {
		t := <-cs[i]        //读取值的时候，可能会出现一只阻塞的情况
		fmt.Println(t)
	}
}*/
