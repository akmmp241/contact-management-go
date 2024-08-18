package domains

type User struct {
	Id       int
	Username string
	Password string
	Name     string
	Token    string
}

type Contact struct {
	Id        int
	UserId    int
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type Address struct {
	Id        int
	ContactId int
	Street    string
	City      string
	Province  string
	Country   string
	Postcode  string
}
