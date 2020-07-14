package model

// Movie - структура в котором описывается имя фильма и все сопутствующее
type Movie struct {
	ID        int
	NameMovie string
	YearMovie string
	Contry    string
}

// Переменные с большой буквы, чтоб их было видно за пределами пакета,
// в идеале писать защищенные (с маленькой) и функции для доступа к ним
