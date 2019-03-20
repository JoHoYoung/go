#### 변수
* 변수는 var 를 통해 선언한다. var 키워드 뒤에 변수명을 적고 그 뒤에 변수타입을 적는다
```
var a int
```
* 변수를 선언하면서 초기값을 지정하지 않으면 Go는 Zero  Value를 기본적으로 할당한다.
```
var number int
var str string
var exist bool

fmt.Println(number)     // 0
fmt.Println(str)        // ""
fmt.Println(exists)     // false
```
 * 상수는 const로 선언. 여러개의 상수를 묶어서 선언할 수 있으며 iota라는 identifier를 사용할 수 있다.
```
const (
    Apple = iota   // 0
    Train          // 1
    Hi             // 2
)
```
* 문자열은 (\` \`) 또는 (" ")를 사용하여 표현할 수 있다.
* 데이터 타입 변환은 T(v)와 같이 사용한다. 

#### 연산자
* 연산자는 다른언어들과 같다.
* Go는 포인터 연산자가 있으며 포인터 산술은 제공하지 않는다.

#### 조건문
* if 조건문 시 ' ( ) '는 쓰지 않는다.
* Go에서는 C/C++ 과는 달리 1, 0을 쓸 수 없고 Boolean 식으로 표현되어야 한다.
```
if val := i*2 ; val < max {
    fmt.Println(val)
}

fmt.Println(val) // if식을 벗어나 에러.
```
* Switch문
1. 다른 언어는 switch 키워드 뒤에 변수나 expression 반드시 두지만, Go는 이를 쓰지 않아도 된다. 이 경우 Go는 switch expression을 true로 생각하고 첫번째 case문으로 이동하여 검사한다
2. 다른 언어의 case문은 일반적으로 리터럴 값만을 갖지만, Go는 case문에 복잡한 expression을 가질 수 있다
3. 다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동하지만, Go는 다음 case로 가지 않는다
4. 다른 언어의 switch는 일반적으로 변수의 값을 기준으로 case로 분기하지만, Go는 그 변수의 Type에 따라 case로 분기할 수 있다
```
func grade(score int) {
    switch {
    case score >= 90:
        println("A")
    case score >= 80:
        println("B")
    case score >= 70:
        println("C")
    case score >= 60:
        println("D")
    default:
        println("No Hope")
    }
}   
```
Type switch
```
switch v.(type) {
case int:
    println("int")
case bool:
    println("bool")
case string:
    println("string")
default:
    println("unknown")
}   
```
break문이 없으면 기본적으로 loop를 빠져나오지만(Go 에 한함) 계속 진행하도록 할 수 있다.
```
func check(val int) {
    switch val {
    case 1:
        fmt.Println("1 이하")
        fallthrough
    case 2:
        fmt.Println("2 이하")
        fallthrough
    case 3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default 도달")
    }
}
```

#### 반복문
* 마찬가지로 다른언어와 달리 '( )'생략.
```
for i := 1; i < 10; i++ {
    print(i)
}
```
* 초기값과 증감식을 생략하고 조건식만을 사용할 수 있다. -> while루프 처럼 사용가능.
```
for i < 100{

    print(i)
    i += 2
}
```
* for range문은 컬렉션으로 부터 한 요소씩 가져와 차례로 실행. for range 문은 "for 인덱스,요소값 := range 컬렉션" 같이 for 루프를 구성하는데, range 키워드 다음의 컬렉션으로부터 하나씩 요소를 리턴해서 그 요소의 위치인덱스와 값을 for 키워드 다음의 2개의 변수에 각각 할당한다.
```
names := []string{"Javascript", "C++", "Go", "Swift"}
 
for index, name := range names {
    println(index, name)
}
```
* break, continue, goto 문
* 경우에 따라 for 루프내에서 즉시 빠져나올 필요가 있는데, 이때 break 문을 사용한다. 만약 for 루프의 중간에서 나머지 문장들을 실행하지 않고 for 루프 시작부분으로 바로 가려면 continue문을 사용한다. 그리고 기타 임의의 문장으로 이동하기 위해 goto 문을 사용할 수 있다. goto문은 for 루프와 관련없이 사용될 수 있다.
```
package main
func main() {
    var a = 1
    for a < 15 {
        if a == 5 {
            a += a
            continue // for루프 시작으로
        }
        a++
        if a > 10 {
            break  //루프 빠져나옴
        }
    }
    if a == 11 {
        goto END //goto 사용예
    }
    println(a)
 
END:
    println("End")
}
```
* break문은 보통 단독으로 사용되지만, 경우에 따라 "break 레이블"과 같이 사용하여 지정된 레이블로 이동할 수도 있다. break의 "레이블"은 보통 현재의 for 루프를 바로 위에 적게 되는데, 이러한 "break 레이블"은 현재의 루프를 빠져나와 지정된 레이블로 이동하고, break문의 직속 for 루프 전체의 다음 문장을 실행하게 한다. 아래 예제는 언뜻 보기에 무한루프를 돌 것 같지만, 실제로는 OK를 출력하고 프로그램을 정상 종료한다. 이는 "break L1" 문이 for 루프를 빠져나와 L1 레이블로 이동한 후, break가 있는 현재 for 루프를 건너뛰고 다음 문장인 println() 으로 이동하기 때문이다.
```
package main
 
func main() {
    i := 0
L1:
    for {
     
        if i == 0 {
            break L1
        }
    }
    println("OK")
}
```
#### 함수의 Pass By Reference
* C에서 Call By Value, Call By Reference와 같은 개념이다.
* 함수 정의시 매개변수에 포인터 변수로 선언(' * '), 해당함수 사용시 주소값 전달 (' & ')

#### 함수의 가변인자함수 ( js의 spread개념)
* 고정된 수의 파라미터를 전달하지 않고 다양한 숫자의 파라미터를 전달하고자 할 때
```
func say(msg ...string) {
    for _, s := range msg {
        println(s)
    }
}
```

#### 함수의 리턴값
* Go 에서는 리턴값이 복수 개일 수도 있다.
* Go 언어는 Named Return Parameter 라는 기능을 제공하는데, 이는 리턴되는 값들을 리턴 파라미터들에 할당할 수 있는 기능
```
func sum(nums ...int) (int, int) {
    s := 0      // 합계
    count := 0  // 요소 갯수
    for _, n := range nums {
        s += n
        count++
    }
    return count, s
}

func sum(nums ...int) (count int, total int) {
    for _, n := range nums {
        total += n
    }
    count = len(nums)
    return
}
// parameter에 리턴값 할당. 이런 경우에도 return은 꼭 적어줘야함.
```
#### 익명함수.
```
sum := func(n ...int) int {
    s := 0
    for _, i := range n{
       s += i
    }
    return s
}
```
#### 일급함수
* 함수는 일급함수로서 Go의 기본 타입과 동일하게 취급되며, 다른함수의 파라미터로 전달하거나 리턴값으로도 사용할 수 있다.
```
func main() {
    //변수 add 에 익명함수 할당
    add := func(i int, j int) int {
        return i + j
    }
 
    // add 함수 전달
    r1 := calc(add, 10, 20)
    println(r1)
 
    // 직접 첫번째 파라미터에 익명함수를 정의함
    r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
    println(r2)
 
}
 
func calc(f func(int, int) int, a int, b int) int {
    result := f(a, b)
    return result
}
```
* type문을 사용한 함수 원형 정의
> type 문은 구조체(struct), 인터페이스 등 Custom Type(혹은 User Defined Type)을 정의하기 위해 사용된다. type 문은 또한 함수 원형을 정의하는데 사용될 수 있다. 즉, 위 예제에서 func(x int, y int) int 함수 원형이 코드 상에 계속 반복됨을 볼 수 있는데, 이 경우 type 문을 정의함으로써 해당 함수의 원형을 간단히 표현할 수 있다.
```
// 원형 정의
type calculator func(int, int) int
 
// calculator 원형 사용
func calc(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}

```
#### 클로저 (Js, Swift와 같은개념)
* 실행 컨텍스트가 끝나도 환경을 기억해서 그때의 변수를 쓰는 함수(?) 같은 개념
* 클로저는 독립적인 (자유) 변수를 가리키는 함수이다. 또는, 클로저 안에 정의된 함수는 만들어진 환경을 ‘기억한다’.
```
func nextValue() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```
#### 배열
* 연속적인 메모리 공간에 동일한 타입의 데이타를 순서적으로 저장하는 자료구조이다. Go에서 배열의 첫번째 요소는 0번, 그 다음은 1번, 2번 등으로 인덱스를 매긴다 (Zero based array).
```
func main() {
    var a [3]int  //정수형 3개 요소를 갖는 배열 a 선언
    a[0] = 1
    a[1] = 2
    a[2] = 3
    println(a[1]) // 2 출력
}
```
* 배열 크기는 생략할 수 있다.
```
var a [...]int{1,2,3,4..}    // 인자 갯수만큼 자동할당
```

#### 컬렉션 - Slice
* Go 배열은 크기를 동적으로 증가시키거나 부분 배열을 추출할 수 없다. Slice는 이런 제약점을 넘어 유용항ㄴ 기능을 제공.
* 크기를 미리 지정하지 않을 수 있고 크기를 동적으로 변경, 부분배열을 추출할 수 있다.
```
var a []int
var a2 = make([]int, 5, 10) 
a2 = a2[2:5]    // 2번 인덱스 원소부터 5 - 1번 인덱스 원소까지 부분 Slice
a2.append(a2, 3 ,4 ,5) // a2에 3개 append
```

#### 컬렉션 - Map
* Map은 key : value의 해쉬테이블
* var map[Key타입]value타입 으로 선언.
```
var myamp map[int]string    // 이렇게 선언시 Nil Map이 되어 값을 쓸 수 없다.
mymap = make(map[int]string) 또는
myamp := map[string]string{
"a":"b"
}                               // 이런식으로 리터럴로 초기화해야함.
```
* 새로운 데이터 추가는 간단한 편이다.
```
func main() {
    var m map[int]string
 
    m = make(map[int]string)
    //추가 혹은 갱신
    m[901] = "Apple"
    m[134] = "Grape"
    m[777] = "Tomato"
 
    // 키에 대한 값 읽기
    str := m[134]
    println(str)
 
    noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
    println(noData)
 
    // 삭제
    delete(m, 777)
}
```
* 맵은 map[key]로 값을 가져올때 두개의 값을 리턴한다. 첫번째 키는 value고 두번째 키는 그 키에 해당하는 값이 존재하는지 아닌지를 나타내는 bool 값이다.
* 컬렉션 Map에 for range를 사용하면 Map키와 Map 값 2개를 리턴한다.
```
func main() {
    myMap := map[string]string{
        "A": "Apple",
        "B": "Banana",
        "C": "Charlie",
    }
 
    // for range 문을 사용하여 모든 맵 요소 출력
    // Map은 unordered 이므로 순서는 무작위
    for key, val := range myMap {
        fmt.Println(key, val)
    }
}
```

#### 패키지 import
```
import "fmt"
 
func main(){
 fmt.Println("Hello")
}
```
* 패키지는 함수, 구조체, 메서드 등이 존재하는데, 이름이 대문자로 시작하면 사용할 수있다.
* 소문자는 non-public, 대문자는 public / 구별해서 잘 사용.
* 패키지 실행시 처음으로 호출되는 함수인 init() 함수. init()함수는 패키지가 로드되면서 실행되는 함수로 별도의 호출없이 자동으로 호출.
```
package testlib
 
var pop map[string]string
 
func init() {   // 패키지 로드시 map 초기화
    pop = make(map[string]string)
}
```
* import 호출 시 _ 라는 alias를 지정하면 init()함수만을 호출한다.
```
package main
import _ "other/xlib"
```
* 만약 패키지 이름이 동일하지만, 서로 다른 버젼 혹은 서로 다른 위치에서 로딩하고자 할 때는, 패키지 alias를 사용해서 구분할 수 있다.
```
import (
    mongo "other/mongo/db"
    mysql "other/mysql/db"
)
func main() {
    mondb := mongo.Get()
    mydb := mysql.Get()
    //...
}
```
* 패키지를 만들때는, 폴더를 하나 만들고 안에 .go 파일들을 만들어 구성. 모두 동일핝 패키지명을 가지며 패키지명은 폴더의 이름과 같게한다.
* Go 패키지를 빌드하고 /pkg 폴더에 인스톨하기 위해서 "go install" 명령어를 입력하면 폴더안에서 실행할 수 있다

#### 구조체
* Go의 struct는 필드 데이터만을 가지며 메스드를 갖지 않는다.
* OOP를 고유의 방식으로 지원하며 클래스, 객체, 상속 개념이 없다.
* 매서드는 별도로 분리하여 정의한다
```
type person struct{
    name string
    age int
}
```
* 구조체 객체 생성법
```
# 빈 객체 생성 후, 나중에 필드값을 채워넣는 방법.
me := person{}
me.name = "호영"
me.age = 220

# 초기값과 함께 할당하는 방법
me := person{name: "호영", age:220}

# 내장함수 new()를 사용하는 방법.
# 모든 필드를 zero value로 초기화하고 포인터를 리턴한다.
# 포인터로 필드엑세스시 (.)을 사용하는데 이때 자동으로 Dereference된다.
me := new(person)
me.name = "호영"
me.age = 220

# struct를 함수의 파라미터로 넘긴다면 Pass by value로 객체를 복사해서 전
# 객체를 pass by reference로 전달하려면 struct의 포인터를 전달해야 한다.
```
* 생성자 함수를 사용할 수 있다. 행성자 함수는 struct 타입을 리턴하는 함수로서 함수 본문에서 필요한 필드를 초기화한다.
```
func newPerson() *person{
  p := person{}
  person.name = "KK"
  person.age = "0"
  return &p
}

func main() {
    per := newPerson()
}
```

#### Go 메서드
* Go에서는 struct가 필드만을 가지며 메서드는 별도로 분리되어 정의된다.
* func 키워드와 함수명 사이에 그 함수가 어떤 struct를 위한 메서드인지를 표시한다. receiver라고 한다.
```
func (p person) nextage() int{
    p.age++;
    return p.age
}               // 이것은 Value receiver이며 해당 객체의 값을 변경시키지 않는다. 즉 Call By Value라는 소리.

func (p *person) nextage() int{
    p.age++
}               // 이와같이 pointer receiver로 구현한다면 전달된 객체의 값도 변경된다.
```

#### 인터페이스
* 구조체는 필드들의 집합, 인터페이스는 메서드들의 집합이다.
* 하나의 사용자 정의 타입(구조체)이 인터페이스를 구현하기 위해서는 그 인터페이스가 갖는 모든 메서드들을 구현합면 된다.
```
type animal interface {
    age() int
}
type person struct {
    age int
    name string
}

func newPerson() *person{
    p := person{}
    p.name ="in"
    p.age = 0
    return &p
}

func (p person) age() int{
    retunr p.age
}
```
* 인터페이스를 사용하는 일반적인 예로 함수가 파라미터로 인터페이스를 받는 경우가 있다.
* 함수 파라미터가 interface인 경우 어떤 타입이든 해당 인터페이스의 모든 함수를 구현하기만 했으면, 파라미터로 쓸 수 있따는 것을 의미한다.
```
func getage(animal ...animals){
    for _, a: range animals{
        age = animal.age()
        print(age)
    }
}
```
#### 인터페이스 타입.
* 빈 인터페이스 v interface{} 의 의미는 Go의 모든타입은 적어도 0개의 메서드를 구현하므로. v interface{}는 모든 타입을 나타내ㅡㄴㄴ 의미이다.
* 즉 어떤 타입도 담을 수 있는 컨테이너 같은 개념이다.
* Type assertion : interface타입 x와 어떤타입 T에 대해 x가 nil이 아니며 T타입이라는걸 확인할때 사용.
```
var a interface{} = 1

j := a.(int) // 21
k := a.(string) // error
```

#### 에러처리
* Go는 내장타입으로 error라는 interface타입을 갖는다.
* error 인터페이스는 다음과 같다.
```
type error interface{
    Error() string
}
```
* 즉 특정 구조체에서 Error 메서드를 정의하면 error interface를 구현하는것이 된다.

#### 지연실행 defer
* defer키워드는 특정 문장 혹은 함수를 나중에 실행하게 한다.

#### panic 함수
* panic함수는 함수를 즉시 멈추고 defer함수들을 모두 실행한 후, 즉시리턴. 마지막에는 에러를 내고 종료하게 된다.

#### recover함수
* panic에 의한것을 정상으로 

#### Goroutine
* 런타임이 관리하는 가상적 쓰레드.
* go 키워드를 사용하여 함수를 호출하면 런타임시 새로운 goroutine을 실행.
* 비동기적으로 함수루틴을 실행하므로 여러코드를 동시에 실행하는데 사용.
```
func say(s string) {
    for i := 0; i < 10; i++ {
        fmt.Println(s, "***", i)
    }
}
 
func main() {
    // 함수를 동기적으로 실행
    say("Sync")
 
    // 함수를 비동기적으로 실행
    go say("Async1")
    go say("Async2")
    go say("Async3")
 
    // 3초 대기
    time.Sleep(time.Second * 3)
}
```
#### 여러 Go루틴들이 끝날 때까지 기다리기.
```
import (
    "fmt"
    "sync"
)
 
func main() {
    // WaitGroup 생성. 2개의 Go루틴을 기다림.
    var wait sync.WaitGroup
    wait.Add(2)
 
    // 익명함수를 사용한 goroutine
    go func() {
        defer wait.Done() //끝나면 .Done() 호출
        fmt.Println("Hello")
    }()
 
    // 익명함수에 파라미터 전달
    go func(msg string) {
        defer wait.Done() //끝나면 .Done() 호출
        fmt.Println(msg)
    }("Hi")
 
    wait.Wait() //Go루틴 모두 끝날 때까지 대기
}
```

#### Go 채널
* 채널은 데이터를 주고 받는 통로. 채널은 make()함수를 통해 미리 생성되어야 한다.
* <- 를 통해 데이터를 보내고 받는다.
```
func routine1(ourChannel chan string) {
        // 채널에 값을 넣습니다.
        ourChannel<- "data"
}

func routine2(ourChannel chan string) {
        // 채널로부터 값을 받습니다.
        fmt.Println(<-ourChannel)
        // 출력값 : data
}

func main() {
        // string 채널을 위한 메모리를 할당합니다.
        ourChannel := make(chan string)
        
        go routine1(ourChannel)
        go routine2(ourChannel)
} 
```
* 채널을 닫으면 송신은 불가능 하지만 수신은 가능하다.