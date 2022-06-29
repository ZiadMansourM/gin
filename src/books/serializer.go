package books

import "github.com/ziadmansourm/gin/users"

type BookSerializer struct {
	Id     uint64               `json:"id"`
	Title  string               `json:"title"`
	Author users.UserSerializer `json:"author"`
}
