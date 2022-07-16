# AlgebraCalculator

## General usage
Typ your formulas in input.txt and run main.go
The calculated formulas will be printed in the console.

## Formula syntax
A formula is identified by an "=". 
Before the "=" must be a letter or word representing the name of the term.  
Example: a = 1

A term can have internal variables. These must be set after the name.  
Example: a\<t\> = 2 * t

Every name, operator or any other symbol must be separated by spaces, or ">" and "<" symbols.
The "<" and ">" don't change any anything in the formula and can just be used to improve readability.  
Example: a\<t\> = 2 * t the same as: a t = 2 * t

After the "=" comes the term of the formula.

## Term syntax

### Numbers
A Number can be an integer or a decimal number. 
Decimla numbers must be defined with an "."

### Vectors 
An vector can have an infinite amount of dimensions. 
To create an vector use the "," opperator.
It merges the number before and the number after it to an vector.
The "," operator can the chained to create n-dimensional vectors.

Example: a = 1 , 8 , 5  
"a" is an 3 dimesional vector with the fist vaule being 1 the second being 8 and the third being 5.

### Operators 
An operator is merges the vector before and after it acording to its function.  
List of Opperators:  
* "+" -> addition
* "-" -> subtraction
* "*" -> multiplication
* "/" -> division
* "pow" -> power
* "dot" -> dot product
* "dist" -> distance between two vectors

### Math functions
A math function processes the following vector according to its purpose.  
List of math functions:
* "sqrt" -> square root
* "degree" -> radians to degree
* "radians" -> degree to radians
* "sin" -> sine in radians
* "cos" -> cosine in radians
* "tan" -> tangent in radians
* "len" -> the length of an vectors

### Inserting other terms
A previously created formula can be inserted in a term. 
Just typ the name of the term it will be pasted in. 
If the formula has internal variables you will need to specify how they should be replaced.
You can do this by definig the replacements after the term.
When you whis to keep the internal variables as they are just replace them with the same name.  
Example:  
a\<t\> = 2 * t  
b = a\<1\> -> b = 2  
c\<t\> = a\<t\> * 2 -> c\<t\> = 4 * t

### Sub operations
If you want to get a specific dimention of a vector or variabele just type "\<name\>.\<dimensions\>"  
Exampel:  
a = 3 , 5 , 7  
b = a.1 -> b = 5  

### Term functions
Term functions are operations that can only be applied on other terms.
List of term functions:
* "gauss" -> tries to solve the interal varibles of the term with gaussian elimination.
