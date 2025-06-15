package bst

type User struct {
	id   int64
	name string
}

func Less(a, b User) bool {
	return a.id < b.id
}

type UserBst BST[User]

func NewUserBst() *UserBst {
	return &UserBst{
		less: Less,
	}
}
