package inmemorydb

type InMemoryNode struct {
	ips []string
}

func NewInMemoryNode() *InMemoryNode {
	return &InMemoryNode{}
}

func (n *InMemoryNode) Save(ip string) error {
	n.ips = append(n.ips, ip)
	return nil
}
func (n *InMemoryNode) Delete(ip string) error {
	for k, v := range n.ips {
		if v == ip {
			n.ips = append(n.ips[:k], n.ips[k+1:]...)
		}
	}
	return nil
}
func (n *InMemoryNode) GetAll() ([]string, error) {
	return n.ips, nil
}
