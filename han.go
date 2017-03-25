package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Text2 struct {
	Hostname string
	Echos    string
}

type Text4 struct {
	Hostname string
	Echos    string
	LinkTxt  string
}

type Text2Slice struct {
	SliceTxt []Text2
}

type Text3Slice struct {
	SliceTxt []Text2
}

type Text4Slice struct {
	SliceTxt []Text4
}

type Topological struct {
	SliceTxt []Text4
}

func (m *Text2Slice) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	text1, err := json.Marshal(m.SliceTxt)
	if err != nil {
	}
	w.Write(text1)
}

func (m *Text3Slice) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	text1, err := json.Marshal(m.SliceTxt)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(text1)
}

func (m *Text4Slice) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	text1, err := json.Marshal(m.SliceTxt)
	if err != nil {
	}
	w.Write(text1)
}

func (m *Topological) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("topological.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	w.Write(data)
}

type Static struct{}

func (m *Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("." + r.URL.Path)
	fmt.Println(r.URL.Path)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}

func main() {
	mux := http.NewServeMux()

	st2txt1 := &Text2{
		"HAN-HAILIN-LB1",
		"the machine ip \n  is 1.1.1.1",
	}

	st2txt2 := &Text2{
		"HAN-HAILIN-LB2",
		"the machine ip is 1.1.1.2",
	}

	st2 := &Text2Slice{}
	st2.SliceTxt = append(st2.SliceTxt, *st2txt1, *st2txt2)
	mux.Handle("/st2", st2)

	//st3 Handler
	st3txt1 := &Text2{
		"HAN-HAILIN-WEB1",
		"the machine ip is 2.1.1.1",
	}

	st3txt2 := &Text2{
		"HAN-HAILIN-WEB2",
		"the machine ip is 2.1.1.2",
	}

	st3txt3 := &Text2{
		"HAN-HAILIN-WEB3",
		"the machine ip is 2.1.1.3",
	}

	st3txt4 := &Text2{
		"HAN-HAILIN-WEB4",
		"the machine ip is 2.1.1.4",
	}

	st3 := &Text3Slice{}
	st3.SliceTxt = append(st3.SliceTxt, *st3txt1, *st3txt2, *st3txt3, *st3txt4)
	mux.Handle("/st3", st3)

	//st4 Handler
	st4txt1 := &Text4{
		"HAN-HAILIN-DB1",
		"数据库的IP地址是：3.1.1.1，请注意，这是一个写库，请不要从这个数据库中读取信息",
		"just write",
	}

	st4txt2 := &Text4{
		"HAN-HAILIN-DB2",
		"the machine ip is 3.1.1.2 the machine ip is 3.1.1.2 the machine ip is 3.1.1.2 the machine ip is 3.1.1.2 the machine ip is 3.1.1.2the machine ip is 3.1.1.2",
		"only read",
	}

	st4 := &Text4Slice{}
	st4.SliceTxt = append(st4.SliceTxt, *st4txt1, *st4txt2)
	mux.Handle("/st4", st4)

	Topological := &Topological{}
	mux.Handle("/topological", Topological)

	mux.Handle("/static/", &Static{})

	http.ListenAndServe(":3000", mux)

}
