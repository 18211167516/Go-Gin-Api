package tests

import (
	"fmt"
	"testing"

	_ "github.com/stretchr/testify/assert"
)

func testGetUsers(t *testing.T){

    testcase := []TestCase{
        {
            code:0,
            param:`name=5678999999`,
            errMsg:`查询成功`,
            method:"GET",
            desc:"验证查询成功",
            haveErr:true,
            showBody:true,
            url:"/api/v1/users",
            content_type:"",
            ext1:1,
        },
    }

    for k,v:=range testcase {
        _,_,w := PerformRequest(v.method,v.url,v.content_type,v.param)
        //assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
        fmt.Println()
        fmt.Printf("第%d个测试用例：%s",k+1,v.desc)
        if v.showBody {
            fmt.Printf("接口返回%s",w.Body.String())
            fmt.Println()
        }

       /*  s := struct{
            Error_code int `json:"error_code"`
            Msg  string `json:"msg"`
            Data  interface{} `json:"data"`
        }{}
        err := util.JsonToStruct([]byte(w.Body.String()),&s)
        assert.NoError(t,err)
        assert.Equal(t, v.errMsg,s.Msg,"错误信息不一致")
        assert.Equal(t,v.code,s.Error_code,"错误码不一致") */
    }
}
func TestUserAll(t *testing.T)  {
    t.Run("TestGetUsers", testGetUsers)
    //t.Run("TestAddTag", testAddTag)
    //t.Run("TestEditTag",testEditTag)
    //t.Run("TestDeleteTag",testDeleteTag)
}