package main

import (
	"context"
	"flag"
	"fmt"

	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	from := flag.String("from", "/ip4/192.168.1.13/tcp/14001/ipfs/QmXNZRTd54Zvarf4sswVvUUnpb4gPQNAhFViozVgG8uwri", "from address")
	shell := flag.String("shell", "/ip4/127.0.0.1/tcp/5001", "ipfs shell address")
	flag.Parse()

	ma, err := multiaddr.NewMultiaddr(*from)
	if err != nil {
		panic(err)
	}
	pi, err := peer.AddrInfoFromP2pAddr(ma)
	if err != nil {
		panic(err)
	}

	api, err := NewIPFSApi(*shell)
	if err != nil {
		panic(err)
	}
	err = api.Swarm().Connect(context.Background(), *pi)
	if err != nil {
		panic(err)
	}
	fmt.Println("swarm connect")
	infos, err := api.Swarm().Peers(context.Background())
	if err != nil {
		panic(err)
	}
	for _, info := range infos {
		fmt.Println(info.ID(), info.Address())
	}
}

func NewIPFSApi(path string) (*httpapi.HttpApi, error) {
	addr, e := multiaddr.NewMultiaddr(path)
	if e != nil {
		return nil, e
	}
	return httpapi.NewApi(addr)

}
