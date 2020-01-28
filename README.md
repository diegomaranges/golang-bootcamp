# Bootcamp Golang

## Part 2-1

In this folder the file use a global Slide to save the information.

Using this method, all functions from the same file can access the information and edit this.

## Part 2-2

In this folder the file also use a Slide to save the information but don't is global, is local for the main function.

The great difference with the previous is:
> When you declare a local Slide or Array, when you need do something with this information (edit, add or remove element) you need pass like a parameter and return this again when finish the function.
> This is because when yo pass a Slide or Array as parameter Golang send a copy (pass by value) and all changes don't apply in the original.

## Part 2-3

In this folder the file use a Map variable in the main function.

The most important difference with a Slice or Array is:
> The map have two values (key and value) and the Slide only have a one value for each element. And when the function pass map variable like a parameter is passed as reference, and all change will do in the other function apply it.

## Part 3

In this folder the file save charge previus information from a file and save when the user exit from the app.

In the case that the file don't exist, when a user start the app is created
