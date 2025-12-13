In Go
- If a variable starts with an uppercase letter, then that variable is accessible outside the package it was declared in (or exported).
- Then it's exported from the package and it's globally visible


'''variable.go
package request
import (
    //all required imports
)
var(
    var Error string // visible outside of this package also
    var timeTaken int // visible inside this package only
)
'''

If a variable starts with a lowercase letter, then it is only available within the package it is declared in.

var Email string //outside the package

var passWord string // inside the package itself

scope also plays an important role in varaibles

var names []string = {"martin","steve","anand","manidhar"}

for i,n := range names{
    fmt.Println(n);
}

## Important points :-

- Variable declaration
    - var foo int
    - var foo int = 42
    - foo := 46
- Can't redeclare variables, but can shadow them.
- All variables must be used.
- Visibility
    - lower case first letter for package scope.
    - upper case first letter to export.
    - no private scope.

- Naming Conventions
    - Pascal or camelCase
        - Capitialize acronyms(HTTP,URL).
    - As shoot as reasonable
        - longer names for longer lives.

- Type conversions
    - destinationType(variable).
    - use ** strconv ** package for strings.

## Special types:-
- bool
- error
- pointers

## Initialization
Go initializes all variables to "zero" by default if you don't:
    - All numerical types get 0 (float0.0, complex 0i).
    - bool gets false
    - string gets "" (the empty string, lenght 0).
    - Everything else gets nil:
        - pointers
        - slices
        - maps
        - channels
        - functions (function variables)
        - interfaces
    - For aggregate types, all members get their "zero" values.

Fscanln - which will scan a line of text ending with new line