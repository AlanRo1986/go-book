package util

import (
	"book/conf"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	shortForm = "2019-05-01 00:00:00"
	PUT_JSON  = 0
	PUT_HTML  = 1
)

func TimeI() int64 {
	return time.Now().Unix()
}

func TimeFormat(ti int) string {
	if ti < 86400 {
		return ""
	}
	t := time.Unix(int64(ti), 0)
	//temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	//str := temp.Format(shortForm)
	str := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return str
}

func Md5(text string) string {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}

func Output(w http.ResponseWriter, data string, t int) {
	if PUT_JSON == t {
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
	} else if PUT_HTML == t {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	w.Header().Set("anthor", "by alan")
	_, err := fmt.Fprint(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ShowError(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	ShowError(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func GetViewTpl(name string) string {
	filePath := conf.ROOT + "view/" + name
	fmt.Println(filePath)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 066)
	ShowError(err)

	defer file.Close()

	reader := bufio.NewReader(file)
	var str string;
	for {
		con, err := reader.ReadString('\n')
		str += con
		if err == io.EOF {
			break
		}
	}
	return str;
}

func Sum(vals ...interface{}) int {
	total := 0
	for _ = range vals {
		total++
	}
	return total
}
