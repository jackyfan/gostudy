package testos

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Testos() {
	//os.Setenv("WEB", "http://coolshell.cn") //设置环境变量
	//fmt.Println(os.Getenv("WEB"))               //读出来

	for _, env := range os.Environ() { //穷举环境变量
		e := strings.Split(env, "=")
		fmt.Println(e[0], "=", e[1])
	}
}
func Testexec() {
	cmd := exec.Command("ipconfig")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Command Error!", err.Error())
		return
	}
	fmt.Println(string(out))
}
