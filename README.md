# Douban Top 250

并发快速爬取豆瓣 top250 影片，并按影片顺序存储到数组中。

所有爬取结束，按顺序打印出 top250 影片，打印格式为 markdown 格式

## 使用

```sh
git clone https://github.com/songhuangcn/douban-top-250.git
cd douban-top-250 && make setup && make run
```

运行后，你的终端会展示以下类似数据：
```md
1. [肖申克的救赎](https://movie.douban.com/subject/1292052/)
2. [霸王别姬](https://movie.douban.com/subject/1291546/)
3. [阿甘正传](https://movie.douban.com/subject/1292720/)
4. [泰坦尼克号](https://movie.douban.com/subject/1292722/)
5. [千与千寻](https://movie.douban.com/subject/1291561/)
...
```
