package utils

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
)

func VerifyCert(crt, privateKey string) (bool, string) {

	//获取下一个pem格式证书数据 -----BEGIN CERTIFICATE-----   -----END CERTIFICATE-----
	certDERBlock, _ := pem.Decode([]byte(crt))
	if certDERBlock == nil {
		return false, "证书解析出错，请检查Crt文件!"
	}

	//第一个叶子证书就是我们https中使用的证书
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return false, "证书解析出错，请检查Crt文件!"
	}

	// 解码私钥------BEGIN RSA PRIVATE KEY-----   -----END RSA PRIVATE KEY-----
	keyDERBlock, _ := pem.Decode([]byte(privateKey))
	if keyDERBlock == nil {
		return false, "私钥解析出错，请检查Key文件!"
	}
	var key interface{}
	var errParsePK error
	if keyDERBlock.Type == "RSA PRIVATE KEY" {
		//RSA PKCS1
		key, errParsePK = x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	} else if keyDERBlock.Type == "PRIVATE KEY" {
		//pkcs8格式的私钥解析
		key, errParsePK = x509.ParsePKCS8PrivateKey(keyDERBlock.Bytes)
	}
	if errParsePK != nil {
		return false, "私钥解析出错，请检查Key文件!"
	}
	return key.(*rsa.PrivateKey).PublicKey.Equal(x509Cert.PublicKey), ""
}

func VerifyDomain(crt, domain string) error {
	var cert tls.Certificate
	//获取下一个pem格式证书数据 -----BEGIN CERTIFICATE-----   -----END CERTIFICATE-----
	certDERBlock, restPEMBlock := pem.Decode([]byte(crt))
	if certDERBlock == nil {
		var err error
		return err
	}
	//附加数字证书到返回
	cert.Certificate = append(cert.Certificate, certDERBlock.Bytes)
	//继续解析Certifacate Chan,这里要明白证书链的概念
	certDERBlockChain, _ := pem.Decode(restPEMBlock)
	if certDERBlockChain != nil {
		//追加证书链证书到返回
		cert.Certificate = append(cert.Certificate, certDERBlockChain.Bytes)
	}
	//第一个叶子证书就是我们https中使用的证书
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return err
	}
	return x509Cert.VerifyHostname(domain)
}

func VerifySslFile(crt string) bool {
	crtTrimSpace := strings.TrimSpace(crt)
	fmt.Println(crtTrimSpace)
	fmt.Println(strings.HasPrefix(crtTrimSpace, "-----BEGIN CERTIFICATE-----"))
	fmt.Println(strings.HasSuffix(crtTrimSpace, "-----END CERTIFICATE-----"))
	return strings.HasPrefix(crtTrimSpace, "-----BEGIN CERTIFICATE-----") && strings.HasSuffix(crtTrimSpace, "-----END CERTIFICATE-----")
}
