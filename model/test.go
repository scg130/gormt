package model

type Test struct{
	Id int	`json:"id" db:"id"`
	TestTxt string	`json:"test_txt" db:"test_txt"`
	Title string	`json:"title" db:"title"`
}

func (*Test) TableName() string {
	return "test"
}
