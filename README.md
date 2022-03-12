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
