package main

import (
	"flag"
	"io"
	"log"
	"net"
)

/*
	一个简单的时间获取客户程序 - 1.2 第 6 页, 图 1-5
*/

var ip = flag.String("ip", "129.6.15.28", "ipv4 or ipv6 address") // NIST Internet Time Severs: https://tf.nist.gov/tf-cgi/servers.cgi

func main() {
	flag.Parse()

	// 这一步相当于实现了 c 语言前 17 行的内容, 涉及 c 语言的关键步骤;
	// Go 标准库已经封装好了底层 socket 连接的所有操作, 具体可以浏览源码
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: net.ParseIP(*ip), Port: 13})
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close() // 简单处理了, 实际应该处理 Close 返回的 error, 其他地方皆如此这般

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf) // 接 c 语言实现的第 19 行, 只不过我们的实现没有 MAXLINE 的限制

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			return
		}

		log.Print(string(buf[:n]))
	}

}
