# Dao层面 Service的编写

<a name="kfAY2"></a>
### Dao 是什么？
dao的全称是Data Access Object，数据访问对象。负责处理数据层面的业务, api需要持有dao对象才可以进行数据库操作。<br />
<br />在basic_app项目中会有一个dao目录，这是dao目录内的项目结构<br />![image.png](https://cdn.nlark.com/yuque/0/2020/png/438583/1597287437291-8193ecba-a10e-4209-8fa5-af48e2521da7.png#align=left&display=inline&height=154&margin=%5Bobject%20Object%5D&name=image.png&originHeight=308&originWidth=588&size=21722&status=done&style=none&width=294)<br />model目录主要放置与MySQL相关的表的数据映射结构体，model自身不允许import项目中的api和dao的包，这样容易引起环包错误。<br />
<br />service是业务核心所在，所有关于数据库的相关业务层操作都需要在此实现，主要实现步骤和API基本一致<br />
<br />事例
```go
package service

type UserService struct {
	orm *gorm.DB
}

func NewService(orm *gorm.DB) *UserService {
	return &UserService{orm: orm}
}

//复杂型查询表单
type FindUsersForm struct {
    Name string //准确寻找用户名
    Query string //模糊查询 用户名, 邮箱，手机号
    Page int //分页
    PageSize int //分页数量

    CreatedAt *time.Time
    IsDeleted bool
}

func (service *UserService) FindUsers(form FindUsersForm) (users []model.Users, err error) {
    //...相关业务操作
    return
}

func (service *UserService) FindUserByID(id uint64) (user model.User, err error) {
	return
}

func (service *UserService) FindUserByName(name string) (user model.User, err error) {
    return
}

```

<br />另外还有一个cache目录，其主要目的是为了开发缓存接口，个人想法是将sql和cache操作隔开，在dao/service中，通过gorm来操作数据库，通过dao/cache包来操作缓存，实现业务隔离。其次，在其他api上使用也更加方便，api可以直接持有cache对象而非需要先持有dao对象<br />
<br />最后，因为api接口普遍是持有dao对象的，但是我们编写的UserService一般情况很少会让api持有（但也不是不行，只要你的userService没有和api形成环包错误，你也可以让api持有userService），所以我们需要让Dao继承UserService的所有方法，如下:<br />![image.png](https://cdn.nlark.com/yuque/0/2020/png/438583/1597289197519-4d4983ec-04ee-451e-b5e4-bd49f8d2a3e0.png#align=left&display=inline&height=441&margin=%5Bobject%20Object%5D&name=image.png&originHeight=882&originWidth=1068&size=93852&status=done&style=none&width=534)<br />

