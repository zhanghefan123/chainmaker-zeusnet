package global_var

import "chainmaker.org/chainmaker/protocol/v2"

type SelfDefinedDataStructure struct {
	ConnectedPeerListBefore map[string]bool
	ConnectedPeerList       map[string]bool
	ReaderResponseChannel   chan map[string]bool
	WantToReadChannel       chan int
	Net                     protocol.Net
}

var GlobalStructure = SelfDefinedDataStructure{
	ConnectedPeerListBefore: map[string]bool{},
	ConnectedPeerList:       map[string]bool{},
	ReaderResponseChannel:   make(chan map[string]bool, 1),
	WantToReadChannel:       make(chan int, 1),
	Net:                     nil,
}
