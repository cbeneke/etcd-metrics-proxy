package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Proxy struct {
	client    *http.Client
	proxyIp   string
	proxyPort string
	proxyPath string
}

func (p *Proxy) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error

	resp, err = p.client.Get("https://" + p.proxyIp + ":" + p.proxyPort + p.proxyPath)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	for k, v := range resp.Header {
		wr.Header().Set(k, v[0])
	}
	wr.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(wr, resp.Body)
	_ = resp.Body.Close()
}

func createTlsConfig(certFile string, keyFile string, caFile string) (error, *tls.Config) {
	// Load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err, nil
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return err, nil
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()

	return nil, tlsConfig
}

func main() {
	bindIp := flag.String("bindIp", "", "IP address to bind to")
	proxyIp := flag.String("proxyIp", "", "IP address to proxy to")
	bindPort := flag.String("bindPort", "", "port to bind to")
	proxyPort := flag.String("proxyPort", "2379", "port to proxy to")
	proxyPath := flag.String("proxyPath", "/metrics", "path to proxy to")
	certFile := flag.String("certFile", "", "path to client cert file")
	keyFile := flag.String("keyFile", "", "path to client key file")
	caFile := flag.String("caFile", "", "path to client ca file")

	flag.Parse()

	err, tlsConfig := createTlsConfig(*certFile, *keyFile, *caFile)

	if err != nil {
		log.Fatal("createTlsConfig: ", err.Error())
		return
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	proxy := &Proxy{
		client:    client,
		proxyIp:   *proxyIp,
		proxyPort: *proxyPort,
		proxyPath: *proxyPath,
	}

	err = http.ListenAndServe(*bindIp+":"+*bindPort, proxy)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
