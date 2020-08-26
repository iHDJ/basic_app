# Restful API规范

<a name="uzCOo"></a>
### API规范
> URL的设定必须是下划线模式的，切勿使用驼峰url

<a name="30843975"></a>
#### 请求方法对照表
例如我要开发一个用户管理模块，model 名称为 User

| Method | URL | Action 行为 | 描述 |
| --- | --- | --- | --- |
| GET | /.../users | index | 获取用户列表并分页 |
| GET | /.../users/:id | find | 获取用户$(id)的信息 |
| POST | /.../users | create | 创建用户 |
| PUT | /.../users/:id | update | 更新用户 |
| PUT | /.../users | update | 批量更新用户，通过json取字段ids得到要更新的users |
| DELETE | /.../users/:id | destroy | 销毁用户 |
| DELETE | /.../users | destroy | 批量销毁用户，通过query取字段ids |

<a name="9706d8a8"></a>
#### ID字段从url带入而非从query或body带入
如上表格，你会发现除了批量操作，都是使用url带入id字段。
<a name="srhka"></a>
#### 参数的带入方式
> 参数必须下划线命名法，不要大小驼峰

GET和DELETE方式的请求，数据普遍放在Query中传递<br />PUT和POST方式的请求，必须使用Body中以JSON格式传递<br />

<a name="0YU2E"></a>
#### 若是开发涉嫌安全的API，尽量使用POST方式
例如 稠江街道综合治理平台 的发送短信功能，使用的是GET 携带 Query ? phone = 15669829323的请求。<br />那么就可以F12双击请求，复制粘贴URL，修改手机号字符串，随意的乱发请求。在没有做请求签证的情况下对方是完全可以耗光我们的短信数量的。
