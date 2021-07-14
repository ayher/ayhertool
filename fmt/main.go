package fmt

import (
	"encoding/json"
	realfmt "fmt"
	"runtime"
	"strings"
)

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = realfmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return realfmt.Sprintf(msg, v...)
}

func Println(fp interface{}, vp ...interface{})  {
	var f interface{}
	var v []interface{}
	if b,ok:=fp.(string);ok{
		f=b
	}else if e,ok:=fp.(error);ok{
		f=e.Error()
	}else{
		bb,_:=json.Marshal(fp)
		f=string(bb)
	}


	for _, item:= range vp{
		if b,ok:=item.(string);ok{
			v=append(v,b)
		}else{
			bb,_:=json.Marshal(item)
			v=append(v,string(bb))
		}
	}

	pc,_,line,_ := runtime.Caller(1)
	ff := runtime.FuncForPC(pc)
	me:=realfmt.Sprintf("[%s:%d]",ff.Name(),line)
	realfmt.Printf("\033[1;34;8m%s\033[0m",me)
	realfmt.Printf("\033[1;32;32m %s\033[0m\n",formatLog(f, v...))
}

func Error(f interface{}, v ...interface{})  {
	pc,_,line,_ := runtime.Caller(1)
	ff := runtime.FuncForPC(pc)
	me:=realfmt.Sprintf("[%s:%d]",ff.Name(),line)
	realfmt.Printf("\033[1;34;8m%s\033[0m",me)
	realfmt.Printf("\033[1;31;31m %s\033[0m\n",formatLog(f, v...))
}

func Debug(f interface{}, v ...interface{})  {
	pc,_,line,_ := runtime.Caller(1)
	ff := runtime.FuncForPC(pc)

	me:=realfmt.Sprintf("[%s:%d]",ff.Name(),line)
	realfmt.Printf("\033[1;34;8m%s\033[0m",me)
	realfmt.Printf("\033[1;31;33m %s\033[0m\n",formatLog(f, v...))
}

func Printf(format string, a ...interface{}) (n int, err error)  {
	return realfmt.Printf( format, a...)
}

func Print( a ...interface{}) (n int, err error)  {
	return realfmt.Print(a...)
}

func Sprintf(f interface{}, v ...interface{}) string {
	return realfmt.Sprintf(f.(string),v...)
}

func Errorf(s string,f ...interface{}) (error) {
	return realfmt.Errorf(s,f)
}

func Scanln(f interface{}) (int,error) {
	return realfmt.Scanln(f)
}

func Scanf(f string, a ...interface{}) (int,error) {
	return realfmt.Scanf(f,a...)
}

func Scan(a ...interface{}) (int,error) {
	return realfmt.Scan(a...)
}