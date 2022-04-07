# Balad Technical interview: go-redis
During my interviews with Balad, i was tasked to implement a simple redis like in memory caching system using golang.

## Design
I have tried to make the design as minimalistic and easy to understand possible which makes the code more maintainable.
This software has 3 entities:
- Database: which is an instance of a database holding all the data and behaviour related to that database 
- Container: which is dependency injection tool, for storing the state of the program.
- Parser: which is in charge of reading inputs from the cli and delegating to the sufficient methods of the 
- Container or Database structs.

## Setup
This project is written using go-lang, and there was no need for any additional dependencies, so it must work out of 
the box!

