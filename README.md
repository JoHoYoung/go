#### 변수
* 변수는 var 를 통해 선언한다. var 키워드 뒤에 변수명을 적고 그 뒤에 변수타입을 적는다
```
var a int
```
* 변수를 선언하면서 초기값을 지정하지 않으면 Go는 Zero Value를 기본적으로 할당한다.
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