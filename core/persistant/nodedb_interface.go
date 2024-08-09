package persistant

type NodeDBInterface interface {
	Save(ip string) error
	GetAll() ([]string, error)
	Delete(ip string) error
}
