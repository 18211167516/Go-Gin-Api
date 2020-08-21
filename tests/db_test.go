package tests

import (
	"encoding/xml"
	"fmt"
	"testing"

	_ "go-api/app/models"
)

type Books struct {
	XMLName xml.Name `xml:"books"`
	Nums    int      `xml:"nums,attr"`
	Book    []Book   `xml:"book"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Name    string   `xml:"name,attr"`
	Author  string   `xml:"author"`
	Time    string   `xml:"time"`
}

func TestDb(t *testing.T) {
	bs := Books{Nums: 666}
	//通过append添加book数据
	bs.Book = append(bs.Book, Book{Name: "小红", Author: "阿三", Time: "2018年6月3日"})
	bs.Book = append(bs.Book, Book{Name: "小绿", Author: "阿四", Time: "2018年6月5日"})
	//通过MarshalIndent，让xml数据输出好看点
	data, _ := xml.MarshalIndent(&bs, "", "  ")
	fmt.Println(string(data))
}

func TestAb(T *testing.T) {
	fmt.Println("132123")
}

func TestBc(T *testing.T) {
	fmt.Println("bchchdf")
}
