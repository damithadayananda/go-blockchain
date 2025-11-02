package node

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"go-blockchain/app"
	"go-blockchain/config"
	"go-blockchain/controller/request"
	"go-blockchain/core/persistant"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
	"go-blockchain/util"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var NodeRef *Node

type Node struct {
	database persistant.NodeDBInterface
}

func NewNode(database persistant.NodeDBInterface) {
	NodeRef = &Node{
		database: database,
	}
	// Adding known nodes from config
	for _, node := range config.AppConfig.KnownNodes {
		cert, err := ioutil.ReadFile(node.CertificatePath)
		if err != nil {
			app.Logger.Error.Log("Error Reading Certificate: %v", err)
			return
		}
		database.Save(domain.Node{
			Ip:          node.Ip,
			Certificate: cert,
			Address:     node.Address,
		})
	}
	// distribute details to the cluster
	NodeRef.distributeNodeExistence()
}

func (n *Node) SaveNode(node domain.Node) error {
	//avoid duplicate node details
	data, err := n.database.GetAll()
	if err != nil {
		return err
	}
	for _, v := range data {
		if v.Ip == node.Ip {
			return nil
		}
	}
	//saving new node
	return n.database.Save(node)
}

func (n *Node) GetNodes() ([]domain.Node, error) {
	return n.database.GetAll()
}

func (n *Node) RemoveNode(url string) error {
	return n.database.Delete(url)
}

// implementing gossip protocol
// once node is initiated it will inform it's existence to known nodes
// these known nodes will inform their known nodes about newly join node
func (n *Node) distributeNodeExistence() {
	myHost := config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port)
	certificate := app.App.Certificate
	informedNodes, err := n.GetNodes()
	if err != nil {
		return
	}
	informedNodes = append(informedNodes, domain.Node{
		Ip:          myHost,
		Certificate: certificate,
	})
	for _, node := range config.AppConfig.KnownNodes {
		body, _ := json.Marshal(request.AddNodeRequest{
			Url:           config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port),
			InformedNodes: informedNodes,
			Certificate:   certificate,
			Address:       app.App.Address,
		})
		req, _ := http.NewRequest(http.MethodPost, node.Ip+"/node/add", bytes.NewReader(body))
		client := util.GeHttpsClient(node.Certificate)
		client.Timeout = time.Second * time.Duration(config.AppConfig.NodeDistributionTimeOut)
		res, err := client.Do(req)
		var data []byte
		if res != nil {
			data, _ = io.ReadAll(res.Body)
			res.Body.Close()
		}
		app.Logger.Info.Log(fmt.Sprintf("Request to: %s, Response: %v, Error: %v", node, string(data), err))
	}
}

func (n *Node) ValidateNewNode(validate domain.NodeValidate) (transaction.Transaction, error) {
	// do validation logic
	if err := n.Validate(validate); err != nil {
		app.Logger.Error.Log(fmt.Sprintf("Node Validation Failed: %v", err))
		return transaction.Transaction{}, err
	}
	// update node status
	n.database.UpdateNodeStatus(validate.Address, domain.ACTIVE)
	// add transaction to mempool
	txn := transaction.NewTransaction(transaction.Transaction{
		Sender:   app.App.Address,
		Receiver: "Address",
		Amount:   0,
		Fee:      0,
		Data: map[string]interface{}{
			"Address": validate.Address,
			"Status":  domain.ACTIVE.String(),
		},
	})
	// add txn
	//n.txnService.AddTransaction(txn)
	return txn, nil
}

func (n *Node) Validate(validate domain.NodeValidate) error {
	//hash validation
	node, _ := n.database.GetNode(validate.Address)
	var err error
	err = certificateValidity(node.Certificate)
	if err != nil {
		return err
	}
	err = certificateSignatureValidity(node.Certificate, validate.Address)
	if err != nil {
		return err
	}
	return nil
}

func certificateValidity(certPEM []byte) error {
	// Decode the PEM-encoded certificate
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return errors.New("failed to decode PEM block containing certificate")
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse certificate: %v", err)
	}

	// Get current time
	currentTime := time.Now()

	// Check validity period
	if currentTime.Before(cert.NotBefore) {
		return fmt.Errorf("certificate not valid yet, valid from: %v", cert.NotBefore)
	}
	if currentTime.After(cert.NotAfter) {
		return fmt.Errorf("certificate has expired, valid until: %v", cert.NotAfter)
	}

	app.Logger.Error.Log("Certificate is valid from %v to %v", cert.NotBefore, cert.NotAfter)
	return nil
}

func certificateSignatureValidity(certPEM []byte, address string) error {
	// Decode the PEM-encoded certificate
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return errors.New("failed to decode PEM block containing certificate")
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse certificate: %v", err)
	}

	// Extract the public key (assuming RSA)
	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("certificate does not contain an RSA public key")
	}

	// Derive address from the public key (matching InitAddress)
	pubKeyDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %v", err)
	}
	hash := sha256.Sum256(pubKeyDER)
	derivedAddress := hex.EncodeToString(hash[:20])

	// Compare derived address with provided address
	if derivedAddress != address {
		return fmt.Errorf("derived address (%s) does not match provided address (%s)", derivedAddress, address)
	}

	app.Logger.Error.Log("Certificate signature and address verified successfully for issuer: %s", cert.Issuer)
	return nil

}
