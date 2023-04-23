package perm

type PermInterface interface {
	CreatePerm(byteSize int32, permPath string) (bool, error)
	Encrty(permPath string, inputStr string) (string, error)
	DoRsaEncrty(permPath string, inputStr string) (string, error)
	Decrty(permPath string, inputStr string) (string, error)
	DoRsaDecrty(permPath string, inputStr string) (string, error)
}

type PermRsa struct{}
