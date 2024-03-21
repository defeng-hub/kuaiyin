package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	AMap    AMap // 这里配置对应配置文件的结构即可
	Ftpsync Ftpsync
}

type AMap struct {
	Key string
}

type Ftpsync struct {
	Side      string
	Host      string
	Ftpuser   string
	Password  string
	Localdir  string
	Backupdir string
	Watchdir  []string
}
