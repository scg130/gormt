package gorm_model

var text = `package model

%s

type %s struct{
%s
}

func (m *%s) TableName() string {
	return "%s"
}
`
