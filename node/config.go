package node

import ()

type Config struct {
	Port       int
	Tls        bool
	CertFile   string
	KeyFile    string
	ServeFiles []string
}
