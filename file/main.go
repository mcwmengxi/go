package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main(){
	file4()
}

func file(){
	file,err := os.OpenFile("./test.txt",os.O_CREATE | os.O_RDWR,0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	i:=0
	for{
		b := make([]byte, 12)
		i++
		f,e := file.Read(b)
		defer file.Close()
		file.Seek(0,io.SeekEnd) //光标移到最后
		file.Write([]byte("12"))
		// if i<5 {
			// file.Seek(36,io.SeekStart)
			// file.Seek(-3,1)
		// }
		if e!= nil{
			return
		}
		fmt.Println(string(b),f)
	}
}
func file1(){
	file,err := os.OpenFile("./test.txt",os.O_CREATE | os.O_RDWR,0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		// str, e := reader.ReadString('\n')
		str1,isPrefix, e1 := reader.ReadLine()
		if  e1 != nil {
			fmt.Println(err)
			return
		}
		fmt.Println( string(str1),isPrefix)
	}
}
func file2(){
	file,err := os.OpenFile("./test.txt",os.O_CREATE | os.O_RDWR,0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	readAll,_ := ioutil.ReadAll(file)
	readAll1,_ := os.ReadFile("./test.txt")

	fmt.Println( string(readAll),string(readAll1))

}
func file3(){
	dir,_ := os.ReadDir("./")
	for k,v := range dir {
		fmt.Println(v.Info())
		fmt.Println(k,v.IsDir(),v.Name(),v.Type())
	}

	// fmt.Println(dir)
}

func file4(){
	file,err := os.OpenFile("./test.txt",os.O_CREATE | os.O_RDWR,0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	reader := bufio.NewReader(file)
	i:=0
	for{
		i++
		str, e := reader.ReadString('\n')
		writer.WriteString(strconv.Itoa(i)+" "+str)
		if  e != nil {
			fmt.Println(err)
			break
		}
	}
	file.Seek(0,io.SeekStart)
	file1,_ := os.OpenFile("./test_copy.txt",os.O_CREATE | os.O_RDWR,0777)
	// 先复制一份
	io.Copy(file1, file)
	defer writer.Flush()
	defer file.Close()
}