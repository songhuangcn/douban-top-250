package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

const (
	BASE_URL  = "https://movie.douban.com/top250"
	MAX_DEPTH = 10 // 最高并发和最大深度都是 10，因为总共只有 250(总数) / 25(每页数量) = 10 页
)

var (
	movies = make([]*Movie, 250)
)

type Movie struct {
	name string
	link string
}

func main() {
	c := colly.NewCollector(
		colly.MaxDepth(MAX_DEPTH),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: MAX_DEPTH})

	// 访问下一页
	c.OnHTML(".paginator > .next > a[href]", func(e *colly.HTMLElement) {
		nextPageURL := e.Request.AbsoluteURL(e.Attr("href"))
		e.Request.Visit(nextPageURL)
	})

	// 存储当前页面的影片
	c.OnHTML("#content li", func(e *colly.HTMLElement) {
		index, err := strconv.Atoi(e.ChildText(".pic em"))
		if err != nil {
			fmt.Println(err)
			return
		}
		// 名称使用图片的标识，因为真实名称元素含有很多其他语言名称，这里只想要简体中文版名称
		name := e.ChildAttr(".pic img", "alt")
		link := e.ChildAttr(".info a", "href")

		movies[index-1] = &Movie{name, link}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Request URL: %s, error: %s\n", r.Request.URL, err)
	})

	c.Visit(BASE_URL)

	// 等待所有爬取结束
	c.Wait()

	printResult()
}

func printResult() {
	for i, movie := range movies {
		fmt.Printf("%d. [%s](%s)\n", i+1, movie.name, movie.link)
	}
}
