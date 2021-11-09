package config

var DefaultQueenAddresses = []string{
	"/ip4/13.57.3.254/tcp/4001/p2p/12D3KooWSHbYhw3MeVHGQU2nYUprREDkgiYqNbzpcvQdjiPRuLew",
}

type Ant struct {
	Chain          Chain
	QueenAddresses []string
}

type Chain struct {
	Endpoint       string
	LockerContract string
}
