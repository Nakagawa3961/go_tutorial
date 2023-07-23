package main

import "fmt"

// ユーザー情報を表す構造体
type User struct {
	Name    string
	Age     int
	Gender  string
	Address string
}

func main() {
	// ユーザーリストを作成
	//var 変数 [] 型
	var userList []User

	// ユーザーを追加
	user1 := User{Name: "半田隆太", Age: 26, Gender: "男性", Address: "福岡県北九州市"}
	user2 := User{Name: "半田隆太１", Age: 26, Gender: "男性", Address: "福岡県北九州市"}
	user3 := User{Name: "半田隆太２", Age: 26, Gender: "男性", Address: "福岡県北九州市"}

	userList = append(userList, user1, user2, user3)

	// ユーザーリストを表示
	//fmt.printlnで空行を追加
	//printlnは改行あり
	// \nで改行
	for _, user := range userList {
		fmt.Printf("名前:%s\n 年齢:%d\n 性別:%s\n 住所:%s\n", user.Name, user.Age, user.Gender, user.Address)
		fmt.Println()
	}
}
