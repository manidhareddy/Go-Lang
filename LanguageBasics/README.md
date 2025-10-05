In Go
If a variable starts with an uppercase letter, then that variable is accessible outside the package it was declared in (or exported).

If a variable starts with a lowercase letter, then it is only available within the package it is declared in.

var Email string //outside the package

var passWord string // inside the package itself

scope also plays an important role in varaibles
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
var names []string = {"martin","steve","anand","manidhar"}

for i,n := range names{
    fmt.Println(n);
}
