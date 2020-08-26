# API 代码规范

回到上一篇的上下文，现在我需要开发用户相关的api，<br />首先这个api他是面向管理用户的，那么就是复数形式的user，即users
```go
package users

import "basic_app/dao"

type Service struct {
	dao *dao.Dao
}

func NewService() *Service {
	return &Service{dao: dao.Current}
}

```
首先我们需要先写如上所示的Service结构体，之所以要写结构体，我的想法是希望放与该apis上下文所相关的其他东西到Service结构体内，就比如我有一个图片需要上传到oss，那么我就需要把他放Service内，这样service下的其他api方法都能获取到而且还不用再次import，另外我如果需要websocket功能，我只需要继承"basic_app/library/websocket".WebSocketService相应的扩展<br />
<br />我个人的建议是一个api单独放一个文件，并在文件前加action_名称，例如

| Action | 文件名称 | 函数名称 |
| --- | --- | --- |
| index | index_users.go | IndexUsers |
| find | find_user.go | FindUser |
| update | update_user.go | UpdateUser |
| destroy | destroy_user.go | DestroyUser |
| create | create_user.go | CreateUser |


<br />完整例子
```go
package users

type IndexUsersRequest struct {
    Page  	 int	`query:"page"`
    PageSize int	`query:"page_size"`
    Query 	 string `query:"query"` //模糊查询用户
}

func (service *Service) IndexUsers(c *gin.Context) {
    var (
    	request IndexUsersRequest
        err		error
    )

    if err = c.ShouldBindQuery(&request); err != nil {
        //参数出错了
        return
    }
}
```

