package main // パッケージ宣言

import "fmt" // fmt パッケージ読み込み
// fmt.Println("Hello") → 改行付きで表示
// fmt.Print("Hello") → 改行なしで表示
// fmt.Printf("名前: %s 年齢: %d\n", name, age) → 書式を指定して表示
//
// %s = string
// %d = int
// %f = float
// %v = 値を自動的に表示
// \n = 改行
//
// fmt.Sprintf("名前: %s", name) → 表示せず、整形したstringを返す
// fmt.Errorf("エラー: %s", message) → error型の値を作る
// fmt.Scan(&name) → 標準入力から値を受け取る
///////////////////////////////////////////////////

type User struct {
	Name string
	Age  int
	Job  string
}

///////////////////////////////////////////////////
func main() {
	fmt.Println("Hello, Go!")

	name()

	greet("Kotaro")
	greet("ABC")

	result := add(10, 20)
	fmt.Println(result)

	checkAge(30)
	checkAge(18)

	repeatSum()

	scopeTest()

	printNames()

	printProfile()

	printUser() 

	user := User{
	Name: "Kotaro",
	Age:  27,
	Job:  "Designer",
}
	user2 := User{
	Name: "Yamada",
	Age:  33,
	Job:  "Engineer",
}
	user.Introduce()
	user2.Introduce()

	userJob := user.GetJob()
	fmt.Println("取得した職業:", userJob)

	userName := user.GetName()
	fmt.Println("取得した名前:", userName)


	fmt.Println("変更前:", user.Job)
	user.ChangeJob("Engineer")
	fmt.Println("変更後:", user.Job)

	pointer := &user
	fmt.Println(pointer.Name)
	fmt.Println(pointer.Job)
	user.ChangeJob("Director") //→ 少し遠回りだが、変更ルールをまとめられる
	//pointer.Job = "Director" → 速い・単純
	fmt.Println(user.Job)

	result, err := divide(10, 2)


	if err != nil {
	fmt.Println("エラー: ", err)
} else {
	fmt.Println("結果: ", result)
}

	result2, err2 := divide(10, 0)
	if err2 != nil {
	fmt.Println("エラー: ", err2)
} else {
	fmt.Println("結果: ", result2)
}
	

	err3 := checkNumber(-10)

	if err3 != nil {
	fmt.Println("エラー:", err3)
	} else {
	fmt.Println("正常です")
	}

///////////////////////////////////////////////////////////////////////////
}
///////////////////////////////////////////////////////////////////////////


func name() {
	name := "Kotaro"
	age := 27

	fmt.Println(name)
	fmt.Println(age)

}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a int, b int) int {
	return a + b
}

func checkAge(age int) {
	if age >= 65 {
		fmt.Println("シニアです")
	} else if age >= 20 {
		fmt.Println("成人です")
	} else {
		fmt.Println("未成年です")
	}
}

func repeatSum() {
	total := 0

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		total += i
	}

	fmt.Println("合計:", total)
}

func scopeTest() {
name := "Kotaro"
total := 0

for i := 1; i <= 3; i++ {
	fmt.Println(i)
	total += i
	}
fmt.Println(name)
fmt.Println("合計:", total)
	
}

func printNames() {
	names := []string{"Kotaro", "ABC", "Taro"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	
names = append(names, "XYZ")
	
	for _, name := range names {
	fmt.Println(name)
}
}

func printProfile() {
	profile := map[string]string{
		"name": "Kotaro",
		"job":  "Engineer",
	}

	profile["country"] = "Japan"
	profile["job"] = "Designer"

	fmt.Println("名前:", profile["name"])
	fmt.Println("職業:", profile["job"]) // 「Designer」になる？
	fmt.Println("国:", profile["country"])


	email, ok := profile["email"]

if ok {
	fmt.Println(email)
} else {
	fmt.Println("emailは登録されていません")
}
}



	func printUser() {
	user := User{
		Name: "Kotaro",
		Age:  27,
		Job:  "Designer",
	}

fmt.Println("名前:", user.Name)
fmt.Println("年齢:", user.Age)
fmt.Println("職業:", user.Job)
}


func (u User) Introduce() {
	fmt.Println("こんにちは、", u.Name, "です")
	fmt.Println("職業は:", u.Job, "です")
}


func (u User) GetJob() string {
	return u.Job
}

func (u User) GetName() string {
	return u.Name
}

func (u *User) ChangeJob(newJob string) {
	u.Job = newJob
}

func divide(a int, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("0では割れません")
	}
			return a / b, nil
}



func checkNumber(n int) error{
		if n < 0 {
			return fmt.Errorf("マイナスの値です")
	}
			return nil
}
