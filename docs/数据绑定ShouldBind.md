# 数据绑定 ShouldBind

basic_app封装了一份关于bind支持添加默认值的功能。<br />
<br />之前开发gin-vue-admin的时候发现，获取分页信息的时候需要手动验证，而且语法看着比较复杂，其功能意思就是判断是否发送了page和pageSize参数。当时有一个想法，如果没有传递page参数就默认是1就好了，没有传递pageSize就默认25吧。<br />
<br />然后我当时这样写
```go
type PaginationRequest struct {
    Page int `query:"page"`
    PageSize int `query:"page_size"`
}

func IndexXXX(c *gin.Context) {
    var (
        request PaginationRequest
        err 	error
    )

    if err = c.ShouldBindQuery(&request); err != nil {
        return
    }

    if request.Page <= 0 {
    	request.Page = 1
    }

    if request.PageSize <= 0 {
    	request.PageSize = 25
    }

    //...其他代码就不写了
}
```
当时写到这的时候还是觉得有些不妥，因为要多写7行代码。于是我在github上把go-default的代码扒了下来。<br />现在可以通过tag的方式在ShouldBind的时候设置默认值<br />

```go
import (
    "basic_app/library/bind"
)

type PaginationRequest struct {
    Page int `query:"page" default:"1"` //若page == 0 那么默认会拿到1
    PageSize int `query:"page_size" default:"25"`//若page_size == 25 那么默认会拿到25
}

func IndexXXX(c *gin.Context) {
    var (
        request PaginationRequest
        err 	error
    )

    if err = bind.ShouldBindQuery(&request); err != nil {
    	return
    }
}
```

<br />另外这里放一个gin-vue-admin的分页验证片段
```go
func GetXxxList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)//先绑定分页请求page和pageSize参数
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])//再验证参数是否符合
	if PageVerifyErr != nil { //不符合就报错，然后返回错误
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}

    //然后开始业务代码
}
```

