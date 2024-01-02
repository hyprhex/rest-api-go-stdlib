package recipes

type MemStore struct {
    list map[string]Recipe
}

func NewMemStore() *MemStore {
    list := make(map[string]Recipe)
    return &MemStore{
        list,
    }
}
