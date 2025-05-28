package mockerybug

type itf interface {
	Use(*Struct)
}

type Struct struct {
}
