package rt

import "fmt"

type Server struct {
	Protocol, Host string
	Port           int
}

func (s *Server) URL(endpoint string) string {
	return fmt.Sprintf("%s://%s:%d%s", s.Protocol, s.Host, s.Port, endpoint)
}
