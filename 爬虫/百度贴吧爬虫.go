package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func SpiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:")
		return
	}
	//fmt.Println("result=", result)

	f, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	f.WriteString(result)
	f.Close()

	strconv.F

	page <- i
}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read page done!")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func working2(start, end int) {
	fmt.Printf("crawling %d page to %d...\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("%d page crawl done...", <-page)
	}
}

func working(start, end int) {
	fmt.Printf("crawling %d page to %d...\n", start, end)

	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err:")
			continue
		}
		//fmt.Println("result=", result)

		f, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
		if err != nil {
			fmt.Println("os.Create err:", err)
			continue
		}
		f.WriteString(result)
		f.Close()
	}

}

func main() {

	var start, end int
	fmt.Print("start index(>= 1):")
	fmt.Scan(&start)
	fmt.Printf("end index(>=%d):", start)
	fmt.Scan(&end)

	working(start, end)
	working2(start, end)
}
