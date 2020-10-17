package http

import (
    "log" 
    "context"
    "net/http"  
    service "kfd/internal/logic/service" 
)

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var (
            token string
        )

        token = r.Header.Get("token")
        
        user, _ := service.AuthService.VerifyToken(context.TODO(), token)
        log.Println("打印 秘钥进行校验 参数：",user)
        if user == nil || user.UserId == 0 {
            OutJson(w, 1, "midwares token error ", nil)
            return
        } else {
            //想把user放到 w 或者 r 里面 方便后续handler能够获取
            //后面改为t.Context() https://www.cnblogs.com/yjf512/p/10399190.html
            //r.Header.Set("User-Agent", user.Account) 
            ctx := context.WithValue(r.Context(), "account", user.Account)
            ctx2 := context.WithValue(ctx, "user_id", user.UserId)
 
            next.ServeHTTP(w, r.WithContext(ctx2))
        }
    })
}
