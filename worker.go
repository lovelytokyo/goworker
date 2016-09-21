package main

import (
	"os"
	"io"
	"fmt"
	"os/exec"
	"net/http"

	"github.com/benmanns/goworker"
)

func myFunc(queue string, args ...interface{}) error {
	if id, ok := args[1].(string); ok {
		if url, ok := args[2].(string); ok {
			// 動画ファイルダウンロードして {id}.mpegで保存する
			downloaded := downloadFromUrl(url, id)
			if downloaded == false {
				fmt.Println("failed download video")
				return nil
			}

			// ./generator <video> <interval> <width> <height> <columns> <output>
			cmd := fmt.Sprintf("generator ./video/%s.mpeg 2 126 73 4 ./img/%s.jpg", id, id)
			fmt.Printf("cmd: %s\n", cmd)

			_, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				fmt.Printf("execute generator err : %s", err)
				return nil
			}

		} else {
			fmt.Println("failed read url")
		}
	}

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

func downloadFromUrl(url, id string) bool {

	fmt.Printf("Downloading %s to %s\n", url, id)
	fileName := fmt.Sprintf("./video/%s.%s", id, "mpeg")

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error while creating %s : %s\n", url, err)
		return false
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error while downloading $s : %s\n", url, err)
		return false
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Printf("Error while Copy $s : %s\n", url, err)
		return false
	}

	fmt.Println(n, "bytes downloaded.")
	return true
}

/*
// run
go run main.go -queues=default

// queue
redis-cli  RPUSH resque:queue:default '{"class":"MyClass","args":["a123456","z99", "http://banner-mtb.dspcdn.com/mtbimg/video/bb5/bb59adddc40801a2f9fa10f2116d4185c56a0213"]}'
 */