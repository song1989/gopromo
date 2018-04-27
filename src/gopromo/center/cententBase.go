package center

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"net"
)

func GetFactory(host string, port string) (thrift.TTransport, *thrift.TBinaryProtocolFactory) {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalln("center tSocket error:", err)
	}

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	transport, err := transportFactory.GetTransport(tSocket)
	if err != nil {
		log.Fatalln("center ransport error:", err)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	return transport, protocolFactory
}
