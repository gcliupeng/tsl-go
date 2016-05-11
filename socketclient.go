package main
import(
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"io/ioutil"
	//"net"
)
func main() {
	CA_Pool := x509.NewCertPool()
	severCert, err := ioutil.ReadFile("./cert.pem")
	if err != nil {
    	log.Fatal("Could not load server certificate!")
	}
	CA_Pool.AppendCertsFromPEM(severCert)

	//config := tls.Config{RootCAs: CA_Pool}	
	config:=tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "10.10.33.10:8082", &config)
	if err != nil {
    	log.Fatalf("client: dial: %s", err)
	}
	buf:=[100]byte{}
	conn.Read(buf[:])
	fmt.Printf(string(buf[:]))
}