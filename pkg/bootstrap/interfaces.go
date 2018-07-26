package bootstrap

type Tokens interface {
	GetBootstrapToken() (string, error)
}
