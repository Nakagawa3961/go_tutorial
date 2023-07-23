package main

import "fmt"

//userの構造体
type User struct {
	Name    string
	Age     int
	Gender  string
	Address string
}

//以下メソッド
//returnがないとエラーが出る
func (u *User) getName() string {
	return u.Name
}

func (u *User) getAge() int {
	return u.Age
}

func (u *User) getGender() string {
	return u.Gender
}

func (u *User) getAddress() string {
	return u.Address
}

func (u *User) getInfo() string {
	// \nは改行を入れること
	info := fmt.Sprintf("名前: %s\n 年齢: %d\n 性別: %s\n 住所: %s",
		u.getName(), u.getAge(), u.getGender(), u.getAddress())
	return info
}

func main() {
	user := User{
		Name:    "山田",
		Age:     5,
		Gender:  "男性",
		Address: " 福岡県北九州市 ",
	}

	fmt.Println(user.getInfo())
}
