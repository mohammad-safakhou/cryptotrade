package ports

import "context"

type HelloRepository interface {
	Get() string
	Save(string)
}

type RawDataRepository interface {
	DataReceiver(ctx context.Context) error
}

//Access ID
//DDF0F96626FD41BE9B2B588F39ED24F3
//Secret Key
//47564D69BC03066F4E440ED591AC7AB3F48E64188A17664B
