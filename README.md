# LarnGo
A repository with simple projects ive made for me to learn go

Focus will be on the use of mathematical functionalitys with [gonum](https://www.gonum.org/)

Projects to make, in order to learn go
- [ ] Graph the normal gaussian distribution 
- [ ] Graph several common distributions on top of each other
- [ ] Create some matrices and multiply them, print the answers in a nice way
- [ ] Create the vanilla Euler forward function
- [ ] Define a diferential equation (a simple system)
- [ ] Create 3 different system matrices, evaluate them for a cuple of loops and plot their states in a nice way
- [ ] Create A system with inputs ( static input array)
- [ ] Create A system with inputs (let a user decide input)

<hr>
## Go tour
First, lets run though the go tour and learn some basics.

**Simple variables**
Variables can be defined in a nubmer of ways.
WHY DOES IT COMPLAIN WHEN I HAVE NOT USED A DECLARED VARIABLE? 
    HOW BIG OF AN ISSUE CAN IT POSSIBBLY BE?

Simples, go inferes the type
>a := 5

Variable with no value assignment (datatypes have default values tho)
> var NAME TYPE
> var a uint16

The same format also allows for value assignment
> var NAME TYPE = VALUE
> var a uint16 = 5

multiple variables with the same type can be created in the same line
>var NAME1, NAME2, NAME3 TYPE 
> var a,b,c float64

And just like before, we can assign values to theses
>var NAME1, NAME2, NAME3 TYPE = 1,2,3
> var a,b,c float64 = 1,24,5

<hr>
**Loops**
there is only the *for* loop in go

for is implemented just like "normal"...

while is implemented like this
> for CONDITION {STUFF}

A forever loop is even easier than "normal"
> for {STUFF}
<hr>
**IF**
Little differances here... but some nice functionality is added.

firstly, statement look
> if CONDIFTION {STUFF}
no parenthesis needed.

The added functionallity:
we can define variables in the if statement!
if a:= 5 < 10{STUFF}
the variable "a" is valid inside the if-scope

we can have a function which might return an error
>fun(input)(error)
Put it inside the if statement and check if the value is not nil (meaning there was an error, and handle it inside the if statement

if a:=fun(input);a != nil{
    //handle the error
    a contains the error
}


<hr>
**Switches**
note: im not a huge fan of switches, donno why

Seems pretty stright forward.
Only evaluates the first matching switch!
Statement is defined in the following way:
switch VARAIABLE {
    case VALUE:
        STUFF
    case VALUE:
        STUFF
    default:
        STUFF
}

one neat functionallity is that VARAIABLE is optional, meaning we can chain long if-else in a tidy way.

[from tour](https://go.dev/tour/flowcontrol/11)
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}

<hr>
**DEFER**
no idea what this is, is this somehting new and unique to go?
Its like a function wrapper inside a function? the statement is executed after the rest of the function have returned... Sounds very odd

we can have multiple defers in a single function, and they are executed in last-in-first-out manner, after the surounding functino is done.

<hr>
**POINTERS**
i feel rusty on this subject, i hope i will de-rust fast ^^

& gives the momory address
* gives the value the address points to

g := 123
address of g is found with
p := &g

the data type of p is ASTERIX-DATATYPE
g = int -> &g = *int

i remember,in c, we can have a 16-byte number, extract its address and read it like a 8-byte number, donno if we can do it in go tho...

<hr>
**STRUCTS**
Yay, always liked structs :)

type STRUCTNAME struct{
    var1 DATATYPE
    var2 DATATYPE
}

instance := STRUCTNAME{value1,value2,...}

Instance with no initial values 
instance := STRUCTNAME{}

Struct uses the dot (.) notation like normal

We can create a struct instance and only set a subset of the fields with their name

instance := STRUCTNAME{FIELDn : VALUEn, FIELD: VALUE}

<hr>
**ARRAYS**
Arrays seems pretty normal aswell.
Declaration can be done in 3 main ways it seems
- size
- size and some initial values

> var a [5]int

> a := [5]int{1,2}

Its odd, feels like i should be able to use the first notation and also apply values, but guess not.

Arrays can also be displayed nicely with print function.

<hr>
**SLICE**
Wait. A slice is not a dynamic array as im used to it. Its a dynamic window to view elements of an array.
I dont really see the need for this... Must be more to this.

Soooo, a slice only reference another array/section of array. Canging values in the slice changes the arrays value at that location.

But slices have both a capacity and lenght property we can read...
cap = length of underlying array
len = lenght of slice.

slices of no length have value nil

There is a function called make(), which sounds a bit like malloc on the surface.
We can create a slice with make and specify its lenght and capacity in the make arguments.

>`a:= make([]int,5,10)` Creates a slice with capacity 10 and initial lenght of 5

When making the tour, i get a feeling that slices are the prefered way of working with arrayed-data.
Slices can also store slices. So thats also neat. that way, we can create a N-d slice.

Aiit, here it comes. Slices can be extended with the use of [`append()`](https://pkg.go.dev/builtin#append).

<hr>
**RANGE**
Range works like itterate in python. nice in for loops. They return both the index of the array and the value at that position. making stuff neat.

`for INDEX, VALUE := range ARRAY {logic}`

> We can ignore outputs from functions with the use of `_`. its like `~` in matlab.

