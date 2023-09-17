package log

import (
	"io"
	stlog "log"
	"net/http"
	"os"
)

//log服务要干什么（自定义日志文件，注册处理函数（将请求的body写入日志））

var log *stlog.Logger

type filelog string

func (fl filelog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// filelog类型实现了io.Writer接口，意味着可以往filelog里写东西
// 将destination转为filelog类型，意味着可以往destination里写东西
func Run(destination string) {
	log = stlog.New(filelog(destination), "[go] - ", stlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
