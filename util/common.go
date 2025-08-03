package util

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"go-blockchain/core/block"
	"net/http"
)

func calculateHash(block block.Block) []byte {
	sha := sha256.New()
	byt, e := json.Marshal(block)
	if e != nil {
		fmt.Println(e)
	}
	sha.Write(byt)
	return sha.Sum(nil)
}

func GeHttpsClient(cert []byte) *http.Client {
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		panic("failed to append server certificate to cert pool")
	}
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		}}
}
