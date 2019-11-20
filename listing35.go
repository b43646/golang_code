package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	// 声明字节缓存类型变量 b
	var b bytes.Buffer
	// 将字符串写入变量b
	b.Write([]byte("Hello"))
	// 追加字符串到变量b
	_, _ = fmt.Fprintf(&b, "World!")
	// 将buffer的内容写入到标准输出
	_, _ = io.Copy(os.Stdout, &b)
}
