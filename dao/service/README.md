# Dao层 Service的编写

## 介绍

Dao层的Service主要是为了对Dao的扩展，让Dao继承你所编写的Service

例如我要写一个model.User的CRUD, 那么我需要写以下方法

```go
type UserService struct {
  orm *gorm.DB
  //如需要Redis的话再加入一行 redis *redis.Client也是可行的
}

func NewService(orm *gorm.DB) *UserService {
  return &UserService{orm}
}

//创建用户
func (service *UserService) CreateUser(user model.User) (err error) {
  return
}

//获取用户
func (service *UserService) FindUserByID(id uint64) (user model.User, err error) {
  return
}

//更新用户
func (service *UserService) UpdateUser(user model.User) (err error) {
  return
}

//删除用户
func (service *UserService) DestroyUserByID(id uint64) (err error) {
  return
}
```

写完之后，让 "basic_app/dao".Dao结构体继承上就可以了, 例子如下:

```go

type Dao struct {
	*user.UserService

	orm   *gorm.DB
	redis *redis.Client
}

func New(orm *gorm.DB, redis *redis.Client) (dao *Dao) {
	dao = &Dao{orm: orm, redis: redis}
	dao.UserService = user.NewService(orm)
	return
}
```
随后让API Service持有Dao对象，API Service就可以使用上面写的方法了