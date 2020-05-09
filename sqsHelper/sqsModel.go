package sqsHelper

type SqsModel struct {
	CallUUID   string
	ChunkName  string
	ChunkIndex int
	ChunkCount int
}

const (
	pcapMsgPrefix = "Pcap processor Chunk info-"
)
