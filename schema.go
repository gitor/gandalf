package gandalf

type user struct {
	Name  string `bson:"_id"`
	Keys  []string
}

type repository struct {
	Name  string `bson:"_id"`
	Users []string
}
