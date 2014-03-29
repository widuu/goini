goini	
========
[![Build Status](https://drone.io/github.com/widuu/goini/status.png)](https://drone.io/github.com/widuu/goini/latest)

**[The official website](http://www.widuu.com)**
##描述

使用goini更简单的读取go的ini配置文件以及根据特定格式的各种配置文件。

##安装方法

	gp get github.com/widuu/goini

##使用方法

>ini配置文件格式样列

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

>初始化

	conf := goini.SetConfig("./conf/conf.ini") //goini.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置

>获取单个配置信息

	username := conf.GetValue("database", "username") //database是你的[section]，username是你要获取值的key名称
	fmt.Println(username) //root

>删除一个配置信息

	conf.DeleteValue("database", "username")	//username 是你删除的key
	username = conf.GetValue("database", "username")
	if len(username) == 0 {
		fmt.Println("username is not exists") //this stdout username is not exists
	}

>添加一个配置信息

	conf.SetValue("database", "username", "widuu")
	username = conf.GetValue("database", "username")
	fmt.Println(username) //widuu 添加配置信息如果存在[section]则添加或者修改对应的值，如果不存在则添加section

>获取所有配置信息

	conf.ReadList() //返回[]map[string]map[string]string的格式 即setion=>key->value

---

goini
========


##About

使用goini更简单的读取go的ini配置文件以及根据特定格式的各种配置文件。

##install 

	gp get github.com/widuu/goini

##use example

>conf.ini

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

>initialize

	conf := goini.SetConfig("./conf/conf.ini") //goini.SetConfig(filepath) filepath = directory+file

>To obtain a single configuration information

	username := conf.GetValue("database", "username") //username is your key you want get the value
	fmt.Println(username) //root

>To delete a configuration information

	conf.DeleteValue("database", "username")	//username is your delete the key
	username = conf.GetValue("database", "username")
	if len(username) == 0 {
		fmt.Println("username is not exists") //this stdout username is not exists
	}

>Add a configuration information

	conf.SetValue("database", "username", "widuu")
	username = conf.GetValue("database", "username")
	fmt.Println(username) //widuu Adding/section configuration information if there is to add or modify the value of the corresponding, if there is no add section

>Get all the configuration information

	conf.ReadList() //return []map[string]map[string]string  example:setion=>key->value



