### URL与Method命名规范

> 我在创建一个API的时候会习惯在文件名前面加(index, find, update, destroy, create)等关键字，这很方便在后期维护的时候在方便找自己想要修改的api代码文件。


> URL一律使用下划线风格，严禁驼峰样式

### 例子: 设计枚举API功能(enumeration)

#### 若是寻找多个对象的API, 则URL我会命名[GET]/enumerations (加复数形式简便直观)
#### 若是寻找单个对象的API, 则URL我会命名[GET]/enumerations/:id (后面穿插一个:id，在gin代码中通过c.Param(":id")即可拿到id值)
#### 若是创建单个或多个对象的API, 则URL我会命名[POST]/enumerations (然后通过JSON的格式传递数据)
#### 若是删除或软删一个对象的API, 则URL我会命名[DELETE]/enumerations/:id (然后通过JSON的格式传递数据)

#### 若是删除或软删一个对象的API, 则URL我会命名[DELETE]/enumerations (然后通过JSON的格式传递数据，并且通过json传递ids批量删除指定的对象)
#### 若是更新一个对象的API, 则URL我会命名[PUT]/enumerations/:id
#### 若是更新一个对象的API, 则URL我会命名[PUT]/enumerations
> *Query只有在GET, DELETE的请求方式才会采用, PUT, POST一律采用JSON传输数据

### API编写规范
basic_app这套基本框架参考了微服务的一些标准思想，宗旨是为了规范化的开发API,将一个API系列看作是一个模块，模块自有一个Service且持有Dao数据对象。在router上再new一个Service并把API一律注册上

Dao是一个数据库服务聚合对象，一切与数据库关联(mysql, redis, queue)的业务都需要持有Dao对象

module(basic_app/dao/service)是各自模块的数据库方法，Dao对象只是聚合了service包下的所有对象

```go
//定义API服务, 例如我要在/api/v1/manager(v1版本的后台管理API)中添加用户管理
//则创建一个文件夹(users)，在go中包名就是users，然后创建一个Service结构体，可以完全照抄其他的Service怎么写的

type Service struct {
  dao *dao.Dao //数据库对象
}

type IndexUserRequest struct {
  Page     string `query:"page"`//页数
  PageSize string `query:"page_size"` //一页的尺寸
  Query    string `query:"query"`//模糊查询用户名，手机号，邮箱
}

//若是寻找多个对象的API，请把方法写成
func (service *Service) IndexUsers(c *gin.Context) {
  //定义request变量和err变量，不用在意没有用到会怎么样，栈内存回收很快的
  var (
    request IndexUserRequest
    err error
  )

  if err = c.ShouldBindQuery(&request); err != nil {
    result.ParamError.Render(c) //他在里面会检查你的语言(zh简体，tw繁体, en英语). 详情参考library/result包的代码
    return
  }

  ...
}

//若是寻找单个对象的API，请把方法写成
func (service *Service) FindUser(c *gin.Context) {
  id := c.Param("id")
  ...
}

//若是更新单个对象的API，请把方法写成
func (service *Service) UpdateUser(c *gin.Context) {
  id := c.Param("id")
  ...
}

//若是更新多个对象的API，请把方法写成
func (service *Service) UpdateUsers(c *gin.Context) {
  id := c.Param("ids")
  ...
}

//若是删除单个对象的API，请把方法写成
func (service *Service) DestroyUser(c *gin.Context) {
  id := c.Param("id")
  ...
}

//若是删除多个对象的API，请把方法写成
func (service *Service) DestroyUsers(c *gin.Context) {
  id := c.Param("ids")
  ...
}
```

### 规范来源

命名规范来源Ruby On Rails框架，目前所有MVC框架的祖师爷，业界标杆，你现在用的任何MVC框架都有Rails的影子。
代码规范来源 B站微服务架构思路