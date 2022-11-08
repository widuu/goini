package goini

import (
	"fmt"
	"log"
	"testing"
)

func TestUnit(t *testing.T) {
	configPath := "/Users/zen/Github/goini/conf/conf.ini"
	conf := &Config{
		filepath: configPath,
	}
	// 只要变量在下文中没有被调用,错误处理绝对不写成if err!=nil,天王老子都不能改变,我说的
	if username, err := conf.GetValue("database", "username"); err != nil {
		fmt.Println(err)
	} else {
		log.Printf("value is %v\n", username)
	}
}
func TestAnother(t *testing.T) {
	configPath := "/Users/zen/Github/goini/conf/conf.ini"
	conf := SetConfig(configPath)
	if ret, err := conf.GetValue("database", "username"); err != nil {
		fmt.Println(err)
	} else {
		t.Log(ret)
	}
}
func TestAll(t *testing.T) {
	conf := SetConfig("./conf/conf.ini")
	if username, err := conf.GetValue("database", "username"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(username) //root
	}

	conf.DeleteValue("database", "username")
	if username, err := conf.GetValue("database", "username"); err != nil {
		fmt.Println(err)
	} else {
		if len(username) == 0 {
			fmt.Println("username is not exists") //this stdout username is not exists
		}
		conf.SetValue("database", "username", "zen")
		if username, err = conf.GetValue("database", "username"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(username) //zen
		}

	}

	data := conf.ReadList()
	fmt.Println(data)
}
