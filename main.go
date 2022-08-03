package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

// TODO エラー処理
// TODO ハードコーディングしてる部分を設定ファイルでできるようにする
// FIXME ファイル末尾じゃないとファイルがごちゃる
// FIXME ファイル末尾に改行がないと次のやつが同じ行に出力される
func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("err", err)
	}

	t := time.Now()
	filepath := wd + "/" + fmt.Sprintf("%02d", int(t.Month())) + "/" + fmt.Sprintf("%02d", t.Day()) + ".md"
	//filepath := wd + "/" + "test.txt"
	_, err = os.Stat(filepath)

	//TODO ファイル存在確認
	fmt.Println(filepath)
	if err != nil {
		log.Fatal("err", err)
	}

	seekByte := seekTagetHeader(filepath, `^#+ 行動ログ`)
	fmt.Println(int64(seekByte))
	hhmm := "[" + fmt.Sprintf("%02d", t.Hour()) + ":" + fmt.Sprintf("%02d", t.Minute()) + "]"
	//TODO os.Args[1]の存在確認
	writeSeekPoint(filepath, seekByte, hhmm+" "+os.Args[1])
}

func seekTagetHeader(filepath string, header string) (seekByte int) {
	// TODO 目標見出し未発見時のエラーを返す
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatal("err", err)
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	isTargetHeading := false
	seekByte = 0

	for sc.Scan() {
		lineTx := sc.Text()
		if !isTargetHeading {
			// バイト数記録
			seekByte += len(sc.Bytes()) + 1
			if regexp.MustCompile(header).Match([]byte(lineTx)) {
				// 1.目標見出しの発見
				isTargetHeading = true
				fmt.Println(lineTx)
			}
		} else {
			if !regexp.MustCompile(`^#(.+)`).Match([]byte(lineTx)) {
				// バイト数記録
				seekByte += len(sc.Bytes()) + 1
			} else {
				//2.目標見出し内容の最後列までシークしておしまい
				fmt.Println(lineTx)
				break
			}
		}
	}
	return seekByte
}

func writeSeekPoint(filepath string, seekByte int, content string) {
	fp, err := os.OpenFile(filepath, os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("err", err)
	}
	defer fp.Close()

	fp.Seek(int64(seekByte), 0)
	//file.SeekはO_APPENDで開かれた時の動作が未定義
	fmt.Fprintln(fp, content)
	/*
		w := bufio.NewWriter(fp)

		_, err = w.WriteString(content)
		if err != nil {
			log.Fatal("err", err)
		}
	*/
}
