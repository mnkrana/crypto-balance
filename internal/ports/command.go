package ports

type Command interface {
	ExecuteRequest(string, any) (string, error)
}
