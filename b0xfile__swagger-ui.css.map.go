// Code generaTed by fileb0x at "2019-07-04 12:13:14.740673 +0800 CST m=+0.825015659" from config file "b0x.yaml" DO NOT EDIT.
// modified(2019-07-04 12:08:17.131302308 +0800 CST)
// original path: swagger-ui/dist/swagger-ui.css.map

package atreugoswagger

import (
	"os"
)

// FileSwaggerUICSSMap is "/swagger-ui.css.map"
var FileSwaggerUICSSMap = []byte("\x7b\x22\x76\x65\x72\x73\x69\x6f\x6e\x22\x3a\x33\x2c\x22\x73\x6f\x75\x72\x63\x65\x73\x22\x3a\x5b\x5d\x2c\x22\x6e\x61\x6d\x65\x73\x22\x3a\x5b\x5d\x2c\x22\x6d\x61\x70\x70\x69\x6e\x67\x73\x22\x3a\x22\x22\x2c\x22\x66\x69\x6c\x65\x22\x3a\x22\x73\x77\x61\x67\x67\x65\x72\x2d\x75\x69\x2e\x63\x73\x73\x22\x2c\x22\x73\x6f\x75\x72\x63\x65\x52\x6f\x6f\x74\x22\x3a\x22\x22\x7d")

func init() {

	f, err := FS.OpenFile(CTX, "/swagger-ui.css.map", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileSwaggerUICSSMap)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
