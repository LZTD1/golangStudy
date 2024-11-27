package localStore

type Store interface {
	Load() error
	Save() error
	Links() []string
	Data() map[string][]int64
	SetLinks([]string)
	SetData(map[string][]int64)
}
