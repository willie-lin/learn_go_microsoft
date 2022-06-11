package main

// 若要声明变量，需要使用 var 关键字：

//var firstname string
//var firstname, lastname string
//var age int

// 另一种编写前一个语句的方法是在 var 关键字后面使用括号，就像通过一个专用块来声明变量一样
//var (
//	firstname, lastname string
//	age                 int
//)

// 声明了变量，但有时候，你需要为它们赋予初始值
//var (
//	firstname strin  = "John"
//	lastname  string = "Doe"
//	age       int    = 30
//)

// 如果你决定初始化某个变量，则不需要指定其类型，因为当你使用具体值初始化该变量时，Go 会推断出其类型
//var (
//	firstname = "John"
//	lastname  = "Doe"
//	age       = 30
//)

// 可以在单行中声明和初始化变量。 使用逗号将每个变量名称隔开，并对每个值执行相同的操作（按同一顺序）
//var (
//	lastname, firstname, age = "Doe", "john", 30
//)

//有时，你需要在代码中加入静态值，这称为常量
//const HTTPStatusOk = 200

// 在一个块中声明多个常量
const (
	StatusOk                  = 200
	StatusBadRequesr          = 400
	StatusInternalServerError = 500
)

// 整数数字
var integer8 int8 = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807

func main() {
	//firstname, lastname, age := "John", "Doe", 30

	//fmt.Printf("%s %s is %d years old\n", firstname, lastname, age)
	//fmt.Println(integer8, integer16, integer32, integer64)
	//rune := 'G'
	//fmt.Println(rune)
	//var integer32 int = 2147483647
	//fmt.Println(integer32)
	////var integer uint = -10
	////fmt.Println(integer)
	//fmt.Println(math.MaxFloat32, math.MaxFloat64)
	//number1, _ := strconv.Atoi(os.Args[1])
	//number2, _ := strconv.Atoi(os.Args[2])
	//fmt.Println("Sum:", number1+number2)

	//sum := sum(os.Args[1], os.Args[2])
	//fmt.Println(sum)
	//
	//sum, mul := calc(os.Args[1], os.Args[2])
	//fmt.Println("Sum:", sum)
	//fmt.Println("Mul:", mul)
}
