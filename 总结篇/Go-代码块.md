# 1.go代码片段

## 	1.go实现分页

```go
				//go实现分页
				func PageEdior(pageCount float64,pageIndex int)([]int){
					//判断显示哪些页码
					var pages []int
				
					if pageCount <= 5{
						pages = make([]int,int(pageCount))
						i := 1
						for pageCount > 0{
							pages[i-1] = i
							pageCount -= 1
							i += 1
						}
					}else  if pageIndex <= 3{
						pages = make([]int,5)
						i := 1
						//当前页码等于
						var temp = 5
						for temp > 0{
							pages[i-1] = i
							temp -= 1
							i += 1
						}
					}else if pageIndex >= int(pageCount) - 2{
						pages = make([]int,5)
						//给后三页赋值
						temp := 5
						i := 1
						for temp >0{
							pages[i-1] = int(pageCount) - temp + 1
							temp -= 1
							i +=1
						}
					}else {
						pages = make([]int,5)
						temp := 2
						i := 1
						for temp > -3{
							pages[i-1] = pageIndex - temp
							temp -= 1
							i += 1
						}
						beego.Info(pages)
					}
					return pages
				}


```

## 	 