package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:servers`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:server`
	Description string   `xml:",innerxml"`
}

type server struct {
	//XMLName    xml.Name `xml:servers`
	ServerName string `xml:serverName`
	ServerIP   string `xml:serverIP`
}

func main() {
	//file, err := os.Open("D:/Development/Go_Repo/go/src/go-web-practice/chap07/server.xml")
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//defer file.Close()
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//v := Recurlyservers{}
	//err = xml.Unmarshal(data, &v)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//
	//fmt.Println(v)

	v := &Recurlyservers{Version: "1"}
	v.Svs = append(append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"}), server{"BeijingVPN", "127.0.0.1"})
	output, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
