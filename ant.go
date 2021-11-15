package config

var DefaultQueenAddresses = []string{
	"/ip4/13.57.3.254/tcp/4001/p2p/12D3KooWSHbYhw3MeVHGQU2nYUprREDkgiYqNbzpcvQdjiPRuLew",
	"/ip4/3.11.8.114/tcp/4001/p2p/12D3KooWBRevRigHtU5x7WEpoLTFb7fsSWNnAk7LY7RM6Brn4dBW",
	"/ip4/16.163.58.59/tcp/4001/p2p/12D3KooWELoezRGrWa5LLvzb1x7VekJ5exrz5SdSUUyGquVmxZ5f",
}

type Ant struct {
	Chain          Chain
	QueenAddresses []string
}

type Chain struct {
	Endpoint       string
	LockerContract string
}
