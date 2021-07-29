# GO Projects and Other Codes
This repository contains some of my works along my journey of learning Golang

## Credit Card Validator
To run this project, 
1. Download the folder
2. Go to `test` directory
3. Run the program using:
```sh
go run . input1 input2
```
The `input1` and `input2` are the command line arguments that generates 4 random numbers for each argument with their values as its prefix. They are then validated against the algorithm and the relevant output is displayed. For example,
```sh
go run . 582134 18009
```
![Output](https://user-images.githubusercontent.com/40364058/127534580-55c9745d-0840-45ac-89bf-2b9d2c55c33a.png)

The algorithm used to generate random numbers for credit cards is taken from this answer on [Stack Overflow]

[//]: #
   [Stack Overflow]: <https://stackoverflow.com/a/31832326>
