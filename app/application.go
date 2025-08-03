package app

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"go-blockchain/config"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

var App Application

type Secret interface {
	InitCertificate()
	InitAddress()
}

type Application struct {
	Certificate []byte
	privateKey  *rsa.PrivateKey
	Address     string
}

func NewApplication() {
	App.InitCertificate()
	App.InitAddress()
}

func (app *Application) InitCertificate() {
	if _, err := os.Stat(config.AppConfig.SecretDir + "/cert.pem"); os.IsNotExist(err) {
		certificate, privateKey, genErr := generateKeyAndCert(strings.TrimPrefix(config.AppConfig.Host, "https://"),
			config.AppConfig.SecretDir+"/cert.pem",
			config.AppConfig.SecretDir+"/key.pem")
		if genErr != nil {
			log.Fatal("failed to generate certificate: ", genErr)
		}
		App.Certificate = certificate
		App.privateKey = privateKey

	} else {
		certificate, readErr := ioutil.ReadFile(config.AppConfig.SecretDir + "/cert.pem")
		if readErr != nil {
			log.Fatal("failed to read certificate: ", readErr)
		}
		App.Certificate = certificate
		//read private key
		keyData, err := ioutil.ReadFile(config.AppConfig.SecretDir + "/key.pem")
		if err != nil {
			log.Fatal("failed to read private key: ", err)
		}
		keyBlock, _ := pem.Decode(keyData)
		if keyBlock == nil {
			log.Fatal("invalid private key format")
		}
		var privateKey *rsa.PrivateKey
		switch keyBlock.Type {
		case "RSA PRIVATE KEY":
			privateKey, err = x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
			if err != nil {
				log.Fatal("failed to parse RSA private key: ", err)
			}
		case "PRIVATE KEY":
			// Handle PKCS#8 format
			pk, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
			if err != nil {
				log.Fatal("failed to parse PKCS#8 private key: ", err)
			}
			var ok bool
			privateKey, ok = pk.(*rsa.PrivateKey)
			if !ok {
				log.Fatal("PKCS#8 private key is not RSA")
			}
		default:
			log.Fatal("invalid private key format: ", keyBlock.Type)
		}
		App.privateKey = privateKey
	}
}

func (app *Application) InitAddress() {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&app.privateKey.PublicKey)
	if err != nil {
		log.Fatal("failed to marshal public key: ", err)
	}
	hash := sha256.Sum256(pubKeyBytes)
	address := hex.EncodeToString(hash[:20])
	app.Address = address
	log.Println("address: ", address)
}

func generateKeyAndCert(serverName, certFile, keyFile string) ([]byte, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: serverName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{serverName},
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create certificate: %v", err)
	}

	certOut, err := os.Create(certFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open %s for writing: %v", certFile, err)
	}
	defer certOut.Close()
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certDER}); err != nil {
		return nil, nil, fmt.Errorf("failed to write certificate: %v", err)
	}

	keyOut, err := os.Create(keyFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open %s for writing: %v", keyFile, err)
	}
	defer keyOut.Close()
	if err := pem.Encode(keyOut, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}); err != nil {
		return nil, nil, fmt.Errorf("failed to write private key: %v", err)
	}

	return certDER, privateKey, nil
}
