package models

type User struct {
	ID       *int64 `gorm:"primarykey"`
	Username string  `gorm:"unique"`
	Password string
	Pfp      string
}

type Post struct {
	ID      *uint64 `gorm: "primarykey"`
	Content string
	Author  int64
	Under   *int64
}

type Like struct {
	ID      *int64 `gorm: "primarykey"`
	Author  int64
	Under   int64
	Matcher string `gorm:"unique"`
}
