goini
========

**[The official website](https://zhangyiming748.github.io/)**
##描述

使用goini更简单的读取go的ini配置文件以及根据特定格式的各种配置文件。

##安装方法

```shell
$ go get -u github.com/zhangyiming748/goini
```

##使用方法

+ ini配置文件格式样列

```dosini
[database]
username = root
password = password
hostname = localhost

[admin]
username = root
password = password

[nihao]
username = root
password = password
```

+ 初始化

```go
conf := goini.SetConfig("./conf/conf.ini") //goini.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置
```

+ 获取单个配置信息

```go
username := conf.GetValue("database", "username") //database是你的[section]，username是你要获取值的key名称
fmt.Println(username) //root
```

+ 删除一个配置信息

```go
conf.DeleteValue("database", "username")//username 是你删除的key
username = conf.GetValue("database", "username")
if len(username) == 0 {
fmt.Println("username is not exists") //this stdout username is not exists
}
```

+ 添加一个配置信息

```go
conf.SetValue("database", "username", "widuu")
username = conf.GetValue("database", "username")
fmt.Println(username) //widuu 添加配置信息如果存在[section]则添加或者修改对应的值，如果不存在则添加section
```

+ 获取所有配置信息

```go
conf.ReadList() //返回[]map[string]map[string]string的格式 即setion=>key->value
```
