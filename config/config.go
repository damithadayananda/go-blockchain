package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var AppConfig config

type config struct {
	InGenesis                bool         `yaml:"inGenesis"`
	GenesisTransactions      string       `yaml:"genesisTransactions"`
	Complexity               int8         `yaml:"complexity"`
	MaxTransactionsPerBlock  int          `yaml:"maxTransactionsPerBlock"`
	Host                     string       `yaml:"host"`
	Port                     int          `yaml:"port"`
	NodeDistributionTimeOut  int          `yaml:"nodeDistributionTimeOut"`
	BlockDistributionTimeOut int          `yaml:"blockDistributionTimeout"`
	BlockInquiryTimeOut      int          `yaml:"blockInquiryTimeOut"`
	KnownNodes               []*KnownNode `yaml:"knownNodes"`
	SecretDir                string       `yaml:"secretDir"`
}

type KnownNode struct {
	Ip              string `yaml:"ip"`
	CertificatePath string `yaml:"certificatePath"`
	Certificate     []byte `yaml:"certificate"`
	Address         string `yaml:"address"`
}

func InitConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	b, e := os.ReadFile(configPath)
	if e != nil {
		panic(configPath)
	}
	if err := yaml.Unmarshal(b, &AppConfig); err != nil {
		panic(err)
	}
	AppConfig.parseCertificates()
}

func (p *config) parseCertificates() {
	for _, node := range AppConfig.KnownNodes {
		b, err := ioutil.ReadFile(node.CertificatePath)
		if err != nil {
			log.Fatal(fmt.Sprintf("Read certificate file failed for known hosts: %v", err))
		}
		node.Certificate = b
	}
}
