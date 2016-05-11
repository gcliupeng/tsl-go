package main
import(
	"crypto/tls"
	"fmt"
	"net"
)
func handle(conn net.Conn) {
	conn.Write([]byte("abcdefg"))
}
func main() {
	cfg := new(tls.Config)
	if cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem"); err == nil {
		cfg.Certificates = append(cfg.Certificates, cert)
	} else {
		fmt.Printf("%v",err)
		return
	}
	listen,_:=tls.Listen("tcp",":8082",cfg)
	for {
		conn,_:=listen.Accept()
		go handle(conn)
	}
}