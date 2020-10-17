package http

import ( 
    "net/http" 
    "html/template"
    "log"
    json "encoding/json"
    "math"
)


func TotalPage(nums, prepage int) int {
    return int(math.Ceil(float64(nums) / float64(prepage)))
}

type ApiData struct {
    //［结构体变量名 ｜ 变量类型 ｜ json 数据 对应字段名]
    ErrCode int         `json:"errcode"` //接口响应状态码
    Msg     string      `json:"msg"`     //接口响应信息
    Data    interface{} `json:"data"`
}

func OutJson(w http.ResponseWriter, errorCode int, msg string, data interface{}) error {
    apiData := &ApiData{errorCode, msg, data}
    err := json.NewEncoder(w).Encode(apiData)
    if err != nil {
        log.Println(err)
    }

    return err
}

/* 
参数一 templFile := "login.gtpl"
    type Data struct{ Title string }
参数二 data:= Data{Title:"夕阳西下"}
*/
func OutHtml(w http.ResponseWriter, templFile string, data interface{}) {
    t, err := template.ParseFiles(templFile)
    checkErr(err)
    
    err = t.Execute(w, data)
    checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        log.Println(err)
    }
}