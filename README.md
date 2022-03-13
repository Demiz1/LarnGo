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

