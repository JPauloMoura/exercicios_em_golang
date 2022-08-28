package main

import (
	"bufio"
	"bytes"
	"encoding/ascii85"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

func main() {
	// verificarPdf("Arquivo com senha -->", "protegido.pdf")
	// verificarPdf("Arquivo com senha -->", "file-protegido.pdf")
	// verificarPdf("Arquivo sem senha -->", "file.pdf")
	verificarPdf3("Arquivo sem senha -->", "protegido.pdf")
}

func verificarPdf(file string, path string) {
	pdf, err := os.Open(path)
	checarErro(err)
	defer pdf.Close()

	linha := bufio.NewScanner(pdf)

	var eProtegido bool
	for linha.Scan() {
		eProtegido = strings.Contains(linha.Text(), "/Encrypt")

		if eProtegido {
			fmt.Println(file, " protegido:", eProtegido)
			return
		}
	}

	fmt.Println(file, " protegido:", eProtegido)
}

func verificarPdf2(file string, path string) {
	data, err := ioutil.ReadFile(path)
	checarErro(err)
	// fmt.Println(string(data))

	e := strings.Contains(string(data), "/CryptAlgorithm")
	r := strings.Contains(string(data), "/R ")
	u := strings.Contains(string(data), "/U ")

	fmt.Println("/CryptAlgorithm", e)
	fmt.Println("/R", r)
	fmt.Println("/U", u)
}
func verificarPdf3(file string, path string) {
	data, err := os.Open(path)
	checarErro(err)
	defer data.Close()
	io := ascii85.NewDecoder(data)

	linha := bufio.NewScanner(io)

	// var eProtegido bool
	for linha.Scan() {
		// eProtegido = strings.Contains(linha.Text(), "/Encrypt")
		fmt.Println(linha.Text())
		// if eProtegido {
		// 	fmt.Println(file, " protegido:", eProtegido)
		// 	return
		// }
	}

	// fmt.Println(file, " protegido:", eProtegido)
}

func verificarPdf4(file string, path string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("pdf", path)
	checarErro(err)

	f, err := os.Open(path)
	checarErro(err)

	_, err = io.Copy(fw, f)
	checarErro(err)

}

func checarErro(e error) {
	if e != nil {
		panic(e)
	}
}

// tempo de desenvolvimento: 3hrs
