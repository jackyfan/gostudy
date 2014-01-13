package testtempla

import (
	//"fmt"
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
}

func main() {
	/*f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t, err := t.Parse("hello {{.UserName}}!")
	if err != nil {
		fmt.Print("ddddd")
		fmt.Println("err=", err.Error())
	}

	p := Person{UserName: "Jackyfan",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)*/
	t := template.New("fieldname example1")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)
}
