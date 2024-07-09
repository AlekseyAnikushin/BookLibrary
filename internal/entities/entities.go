package entities

type Author struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Biography string `json:"biography"`
	Birthdate string `json:"birthdate"`
}

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	AuthorId int    `json:"author_id"`
	Year     uint16 `json:"year"`
	Isbn     string `json:"isbn"`
}

type BookAndAuthor struct {
	Book   `json:"book"`
	Author `json:"author"`
}
