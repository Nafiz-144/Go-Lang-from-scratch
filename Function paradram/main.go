package main

/*go
func add(num1 int , num2 int){ //parameter
    c:= num1 + num2
    fmt.Println(c)
}
func main(){
    add(2,4) // argument
}
```



argument -> আগে
parameter -> পরে।

First Order Function (first order function সেগুলোই যেগুলো parameter হিসেবে string , int, array type এর নিবে )
Standard Function or named function
anonomyous Function (যে ফাংশনের নাম থাকে না )
Immidiately invoke function Expression (IIFE) (যে anonomyous function কে সাথে সাথে invoke করতে হয় )
Higher Order Function(first class function) যে function কোন একটা function কে receive করে as parameter বা return করে বা দুইটাই করে একসাথে তাঁকে বলা হয় Higher Order Function
callback function (Higher Order function এ আমরা কোন একটা function কে pass করা হয় সেটা হলো callback function)
Discrete Mathematic এর একটা concept আছে logic (First Order Logic , Higher Order Logic) যার থেকে inspired হয়ে বানানো হয়েছে Functional Paradigm language .

আবার এই functional paradigm থেকে inspired হয়েছে Go programming language .

LOGIC
Object (People , animal , car etc)
Property (Color , student ) -> বৈশিষ্ট
Relation
লজিক এ আমরা এই তিন টা জিনিসের উপর depend করে লজিক বানায়।


First Order Logic
RULE : All Customers must pay their pizza bills
       All student must wear their uniforms All customers must pay tips to the waiter .
Higher Order Logic
     Any rule that applies to all customers must also apply to tutul


First Order Logic যা Object , Property , Relation নিয়ে যে rule বানানো হয় সেটা হলো First Order Logic
Higher Order Logic যা ওই rule এর উপর depend করে আবার rule create করে।



Higher Order function- Any One Of the following 3
1.Parameter -> Function
2. function return
3. both


Argument হিসেবে Function পাঠানো হয়েছে।

```go
package main
import "fmt"

func processOperation(a int , b int , op func (p int , q int)){
    op(a,b)
}

func add(x int , y int){
    z:= x+y
    fmt.Println(z)
}
func main(){
    processOperation(2,5,add)
}
```

```go
Function return
package main
import "fmt"

func call() func (x int , y int){
    return add
}

func add(x int , y int){
    z:= x+y ;
    fmt.Println(z)
}

func main(){
    sum:= call()
    sum(2,3)
}
```




GO Function এর দুইটা phase থাকে
1. Compiler phase (Compile Time)
2.  Execution phase (Runtime )
Compiler phase এ সব Read only type এর জিনিস গুলো প্রবেশ করবে আর function গুলো প্রবেশ করবে।



Go প্রথম check করবে init function আছে কি না যদি থাকে তাহলে তা প্রথম automatice execute হবে।
আর তা stack এ এই ফাংশনের stack frame create হবে আর কাজ শেষ হলে তা pop হয়ে যাবে।

তারপর main function execute হবে আর stack এ stack frame create করবে। এভাবে কলতে থাকবে।



Higher Order function কে first class function কেন বলে?
আমরা দেখেছি variable এর মধ্যে আমরা string , int , bool এগুলো assign করতে পারি আর এগুলো হলো first class citizen

আমরা আরো দেখেছি variable এর মধ্যে function , anonomyous function ও assign করতে পারি javascript এর মতো যা আমরা c programming এর মধ্যে পারি না ।

আমরা যেহেতু function কে first order citizen এর মতো assign করতে পারি তাই higher order function সেগুলোকে বলা হয় first class function .*/