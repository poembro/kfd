package util

import ( 
    "math" 
)

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(currpage, prepage int, nums int64) map[string]interface{} {
 
    var firstpage int //前一页地址
    var lastpage int  //后一页地址
    //根据nums总数，和prepage每页数量 生成分页总数
    totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //currpage总数
    if currpage > totalpages {
        currpage = totalpages
    }
    if currpage <= 0 {
        currpage = 1
    }
    var pages []int
    switch {
    case currpage >= totalpages - 5 && totalpages > 5: //最后5页
        start := totalpages - 5 + 1
        firstpage = currpage - 1
        lastpage = int(math.Min(float64(totalpages), float64(currpage+1)))
        pages = make([]int, 5)
        for i, _ := range pages {
            pages[i] = start + i
        }
    case currpage >= 3 && totalpages > 5:
        start := currpage - 3 + 1
        pages = make([]int, 5)
        firstpage = currpage - 3
        for i, _ := range pages {
            pages[i] = start + i
        }
        firstpage = currpage - 1
        lastpage = currpage + 1
    default:
        pages = make([]int, int(math.Min(5, float64(totalpages))))
        for i, _ := range pages {
            pages[i] = i + 1
        }
        firstpage = int(math.Max(float64(1), float64(currpage-1)))
        lastpage = currpage + 1
        //fmt.Println(pages)
    }
    paginatorMap := make(map[string]interface{})
    paginatorMap["pages"] = pages
    paginatorMap["totalpages"] = totalpages
    paginatorMap["firstpage"] = firstpage
    paginatorMap["lastpage"] = lastpage
    paginatorMap["currpage"] = currpage
    return paginatorMap
}