# McNaughton-Yamada-Thompson

## Overview
This is an algorithm to convert a regular expression to an NFA

### INPUT
A regular expression `r` over and alphabet \sigma.

### OUTPUT
An NFA `N` accepting `L(r)`.

## METHOD
1. Parse `r` into it's contituent subexpressions.
2. Consruct the `N` using:
	- basis rules for subexpressions with no operators
	- inductive rules for constructing larger NFA's from the NFA's of the intermediate subexpresssions for a given expression.
