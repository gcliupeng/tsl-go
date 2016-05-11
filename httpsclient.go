package main

import ( 
"crypto/tls" 
"crypto/x509" 
"fmt" 
"io/ioutil" 
"net/http" 
)

func main() { 
pool := x509.NewCertPool() 
caCertPath := "cert.pem"

caCrt, err := ioutil.ReadFile(caCertPath) 
if err != nil { 
fmt.Println("ReadFile err:", err) 
return 
} 
pool.AppendCertsFromPEM(caCrt)

tr := &http.Transport{ 
TLSClientConfig: &tls.Config{RootCAs: pool}, 
//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
} 
client := &http.Client{Transport: tr} 
resp, err := client.Get("https://10.10.33.10:8081") 
if err != nil { 
fmt.Println("Get error:", err) 
return 
} 
defer resp.Body.Close() 
body, err := ioutil.ReadAll(resp.Body) 
fmt.Println(string(body)) 
}