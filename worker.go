package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"github.com/benmanns/goworker"
)

func myFunc(queue string, args ...interface{}) error {
	fmt.Printf("From %s, %v\n", queue, args)
	fmt.Printf("group_id : %s\n", args[0])
	fmt.Printf("id : %s\n", args[1])
	fmt.Printf("url : %s\n", args[2])

	if id, ok := args[1].(string); ok {
		if url, ok := args[2].(string); ok {
			// 動画ファイルダウンロードして {id}.mpegで保存する
			downloadFromUrl(url, id)
		} else {
			fmt.Println("failed read url")
		}
	}

	// TODO：動画からサムネイル生成する
	// ./generator ./{id}.mpeg 2 126 73 4 ./thmbnails/{id}.jpg

	return nil
}

func init() {
	goworker.Register("MyClass", myFunc)
}

func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}

func downloadFromUrl(url, id string) {
	fmt.Printf("Downloading %s to %s\n", url, id)
	fileName := fmt.Sprintf("%s.%s", id, "mpeg")

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error while creating %s : %s\n", url, err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error while downloading $s : %s\n", url, err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Printf("Error while Copy $s : %s\n", url, err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

/*
// run
go run main.go -queues=default

// queue
redis-cli  RPUSH resque:queue:default '{"class":"MyClass","args":["a123456","z99", "http://banner-mtb.dspcdn.com/mtbimg/video/bb5/bb59adddc40801a2f9fa10f2116d4185c56a0213"]}'
 */