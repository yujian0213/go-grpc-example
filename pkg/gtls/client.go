package gtls

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

type Client struct {
	ServerName string
	CaFile string
	CertFile string
	KeyFile string
}

func (t *Client) GetCredentrialsByCaA() ( credentials.TransportCredentials,error)  {
	cert,err := tls.LoadX509KeyPair(t.CertFile,t.KeyFile)
	if err != nil {
		return nil,err
	}
	certPool := x509.NewCertPool()
	ca,err := ioutil.ReadFile(t.CaFile)
	if err != nil {
		return nil,err
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil,err
	}
	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName : t.ServerName,
		RootCAs: certPool,
	})
	return c,nil
}
func (t *Client) GetTLSCredentials() (credentials.TransportCredentials,error)  {
	c,err := credentials.NewClientTLSFromFile(t.CertFile,t.ServerName)
	if err != nil {
		return nil,err
	}
	return c,err
}
