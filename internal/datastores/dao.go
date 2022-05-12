package datastores

type DAO struct {
	Repo  Repo
	Cache Cache
}

func NewDAO() (*DAO, error) {
	rdb, err := NewRedis()
	if err != nil {
		return nil, err
	}

	db, err := NewMongo()
	if err != nil {
		return nil, err
	}
	return &DAO{
		Repo:  NewRepo(db),
		Cache: NewCache(rdb),
	}, nil
}
