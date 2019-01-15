/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 11:42
 * Description：
*/
package document

import (
	"baliance.com/gooxml/document"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//获取当前目录
func exePath() (string, error) {
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}
		err = fmt.Errorf("%s is directory", p)
	}
	if filepath.Ext(p) == "" {
		p += ".exe"
		fi, err := os.Stat(p)
		if err == nil {
			if !fi.Mode().IsDir() {
				return p, nil
			}
			err = fmt.Errorf("%s is directory", p)
		}
	}
	return "", err
}
func Parse() {
	path, _ := exePath()
	log.Println("path:", path)
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("path2:", os.Args)
	doc, err := document.Open(`D:\project\go\src\qe-server\build\ml.docx`)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	for i, para := range doc.Paragraphs() {
		//run为每个段落相同格式的文字组成的片段
		fmt.Println("-----------第", i, "段-------------")
		for j, run := range para.Runs() {
			fmt.Print("\t-----------第", j, "格式片段-------------")
			fmt.Print(run.Text())

		}
		fmt.Println()
	}
}
