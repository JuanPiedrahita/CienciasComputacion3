//Presentado Por:
//Juan Sebastian Piedrahita 20141020036
//Juan David  Cubillos 20141020050
//Edison Peñuela 20141020018
//Grupo 81 
//Ciencias de la Computacion 3
package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
	"regexp"
)



func esTipo(x string) string{
	vn:=regexp.MustCompile("^([0-9])+$")
	if vn.MatchString(x) {return "vn"}
	ol:=regexp.MustCompile("^['|'|'&']$")
	if ol.MatchString(x) {return "ol"}
	oc:=regexp.MustCompile("^((([<|>]?)([=]?))|(!{1}={1}))$")
	if oc.MatchString(x) {return "oc"}
	vl:=regexp.MustCompile("^(false|true)$")
	if vl.MatchString(x) {return "vl"}
	return "No se reconoce el token"
}

func main(){
	reader:=bufio.NewReader(os.Stdin)
	x, _ := reader.ReadString('\n')
	x = x[0:len(x)-1]
	s := strings.Split(x, " " )
	for i:=0;i<len(s);i++{
		fmt.Println(esTipo(s[i]),s[i])
	}
}


