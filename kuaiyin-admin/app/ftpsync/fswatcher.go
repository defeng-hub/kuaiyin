package ftpsync

import (
	ext "go-admin/config"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-admin-team/go-admin-core/sdk/config"
)

func FtpFswatch() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	timer := make(map[string]*time.Timer, 0)

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Print("event:", event.Op)
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					// 重置计时器
					if _, ok := timer[event.Name]; ok {
						timer[event.Name].Stop()
					}
					// 启动一个计时器，在一定时间后处理事件
					timer[event.Name] = time.AfterFunc(5*time.Second, func() {
						//非内网环境将文件上传到ftp
						if config.ApplicationConfig.Mode != "lan" {
							log.Print("Ftp Upload file:", event.Name)
							err = FtpClientConn.Upload(event.Name, event.Name)
							if err != nil {
								log.Println(err)
							}
						} else {
							//内网模式将文件上传到项目目录下。
							if strings.Contains(event.Name, "static") {
								fname := strings.Split(event.Name, "static")[1]

								err = os.Rename(event.Name, "/opt/"+config.ApplicationConfig.Name+"/sys/static"+fname)
								if err != nil {
									log.Println(err)
								}
							} else if strings.Contains(event.Name, "logs") {
								fname := strings.Split(event.Name, "logs")[1]
								err = os.Rename(event.Name, "/opt/"+config.ApplicationConfig.Name+"/sys/logs"+fname)
								if err != nil {
									log.Println(err)
								}
							} else {
								log.Print("Invailid file:", event.Name)
							}
						}
						// 删除计时器
						delete(timer, event.Name)
					})

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	//err = watcher.Add("/Users/xz/go/src/insurance_sys/static/upload")
	//判断Watchdir数组是否为空
	if len(ext.ExtConfig.Ftpsync.Watchdir) != 0 {
		for _, v := range ext.ExtConfig.Ftpsync.Watchdir {
			err = watcher.Add(v)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	// Block main goroutine forever.
	<-make(chan struct{})
}
