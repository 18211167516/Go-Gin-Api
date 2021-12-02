> [!NOTE]

* 核心代码位于在`core/session`目录下
* 使用包`github.com/18211167516/sessions` 是我fork的`github.com/gorilla/sessions`
* 建议各位在使用别人的包尽量要fork下来，方便调试和优化

### 结构

```
    ├─middleware  	    （中间件）
    |  StartSession.go     （开启session中间件）
    ├─stores            （多存储实现方式、可自定义）
    |  cookieStore.go   （cookie存储）
    |  redisStore.go    （redis存储）
    |  store.go    		（定义的store接口，自定义均要实现该接口）
    ├─session.go  	    （session实现）
```

### 使用

```
package routes

import (
	"go-api/app/middleware"
	"go-api/app/response"
	"go-api/core/session"
	coremiddleware "go-api/core/session/middleware"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

func testRoute(r *gin.Engine) {

	test := r.Group("/test", middleware.DefaultLog(), middleware.Recovery(), coremiddleware.StartSession())
	{

		test.GET("/panic", func(c *gin.Context) {
			box := response.SysLoginUserResponse{
				ID:       "1",
				Name:     "白",
				RealName: "Bai",
				Type:     2,
				Password: "string",
			}
			s := session.Default(c)
			count := s.Get("user")
			s.Set("user", tool.StructToJson(box))
			s.Save()

			tool.JSONP(c, 0, "层高", count)
		})

	}

}

```

### 支持方法

> [!NOTE] 实例化sesssion

```
func NewSession(c *gin.Context, store stores.Store) *session {
}
```

> [!NOTE] 获取sesssion

```
func Default(c *gin.Context) Session {
	return c.MustGet(global.VP.GetString("session.cookie")).(Session)
}
```

> [!NOTE] 获取值

```
func (s *session) Get(key interface{}) interface{} {}
```
> [!NOTE] 设置值

```
func (s *session) Set(key interface{}, val interface{}) {}
```
> [!NOTE] 删除值

```
func (s *session) Delete(key interface{}) {}
```

> [!NOTE] 清空session

```
func (s *session) Clear() {}
```

> [!NOTE] 保存值

```
func (s *session) Save() error {}
```

### 自定义store

> 需要实现 接口
```
type Store interface {
	// Get should return a cached session.
	Get(r *http.Request, name string) (*Session, error)

	// New should create and return a new session.
	//
	// Note that New should never return a nil session, even in the case of
	// an error if using the Registry infrastructure to cache the session.
	New(r *http.Request, name string) (*Session, error)

	// Save should persist session to the underlying store implementation.
	Save(r *http.Request, w http.ResponseWriter, s *Session) error
}
```