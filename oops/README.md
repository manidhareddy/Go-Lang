# Object-Oriented Programming

- abstraction
- encapsulation
- polymorphism
- inheritance

## Abstraction

Abstraction: decoupling behavior from the implementation details.

## Encapsulation

Enacapsulation: hiding implementation details from misuse.

It's hard to  maintain an abstraction if the details are exposed.

Encapsulation usually means controlling the visibility of names.
("private" variables)

## Polymorphism

Polymorphism literally means "many shapes" - multiple types behind a single interface.

Three main types

- ad-hoc - function/opertaor overloading
- parametric - generic programming
- subtype - inheritance
- protocol-oriented - interface type 

## Inheritance

Inheritance has conflicting meanings
- substitution (subtype) polymorphism.
- structural sharing of implementation details.

# OO in Go

Go offer four main supports for object oriented prog
- encapsulation using the package for visibility control.
- abstraction & ploymorphism using interface types
- enhanced composition to provide structure sharing

Go does not offer inheritance or substituability based on types.

Substitutability is based only on interfaces: purely a function of abstract behavior.

- Go allows defining methods to amy user-defined type.

- Go allows any object to implement the methods of an interface.




