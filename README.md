# 授权中心系统设计SSO(单点登录)

```sequence
title:单点登录系统设计


participant 站点1 as web1
participant 站点2 as web2
participant 认证中心 as sso

web1->web1:本地加载用户信息
web1->sso:加载用户信息
sso->sso:授权信息写入根域
sso-->web1:获得用户信息
web2->web2:获取用户信息\n流畅使用

```


```sequence
title:前端校验

participant 前端 as front
participant 认证中心 as sso

front->front:加载前端资源
front->front:初始化检查cookie
front->sso:加载用户信息
sso->sso:授权信息写入根域
sso-->front:获得用户信息

```


```sequence
title:后端校验

participant 后端 as end
participant 认证中心 as sso

end->end:初始化检查cookie
end->sso:加载用户信息
sso->sso:授权信息写入根域
sso-->end:获得用户信息

```

* 考虑用户安全性，可以使用AES等

# 数据库设计
https://blog.csdn.net/hzw2312/article/details/54612962
## 用户信息表 User

## 用户组表  UserGroup (类似于角色)
用于批量配置用户权限
（人的集合、和其它集团系统的单位信息表，存在共同相同）
权限的集合

## 权限表 Resource



# API设计
## 参考Spring Security
```
  <security:authorize ifAnyGranted="ROLE_ADMIN,ROLE_ADD_FILM">
    <a href="<%=basePath %>manager/insertFilm.jsp">添加影片信息</a><br />
    </security:authorize>
```
范例里的ROLE,对应数据库的的UserGroup
* 按钮级权限控制: has_permission ,has_role 的验证

## Service的API
* AllRoleData
* HasPermission
* HasRole
