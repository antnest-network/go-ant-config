package config

var DefaultQueenAddresses = []string{
	"/ip4/122.9.61.5/tcp/4001/p2p/12D3KooWJHoX2Xf1tmgWaK3insPaKYGntYxTDmWj8489Dxx72Ymx",
}

type Ant struct {
	Chain          Chain
	QueenAddresses []string
}

type Chain struct {
	Endpoint       string
	LockerContract string
}
