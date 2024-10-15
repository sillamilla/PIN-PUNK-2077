# Cyberpunk 2077 Matrix Password Finder

## Introduction

The assignment was poorly described in the message, and I tried to do my best to understand what was needed. 

My implementation is based on the mechanics of the game Cyberpunk 2077. However, my code does not search for the best way to find the password for the matrix; it just looks for the first possible one. Finding the best solution would take a couple more days.

## How It Works

1. **Hack_Service**  
   The endpoint `/getSequence` generates a keySequence of 7 elements for hacking and a 5x5 matrix.
   
   The endpoint `/hack` outputs `success` or `fail`.

3. **MetrixSequence_Service**  
   Receives the data and generates a key that fits the matrix, with a length between 1 and 7 elements, having a status from 0 to 3 depending on the key length.
   
   It then saves the data in the database in the following JSON format:

   ```json
   "KEY": {
       "created": "DATE",
       "result_of_hack": "EXAMPLE"
   }


The key must have at least 1 element for saving in DB (to avoid creating entities in the database with an empty key).

**I've also added the `/getAll` endpoint to retrieve all records from the database.

## Matrix Navigation

We start at the initial point (x0, y0) and move horizontally to search for at least 1 element that we have in our list of possible keys values (`keySequence`).

Matrix Example:
  ```json 
{0, 1, 0, 0, 0},   Possible Keys: []int{1,2,3,4,5,6,7}
{0, 0, 7, 0, 6},   The final key will be "1234567".
{0, 2, 3, 0, 0}, 
{0, 0, 0, 0, 0}, 
{0, 0, 4, 0, 5}
```

We always move horizontally, then vertically, and repeat.

The horizontal function calls the vertical one (or the vertical one in reverse, depending on whether we hit the end of the matrix to go back horizontally/vertically).
And then the vertical function calls the horizontal one in the same way.

When we find a value that fits our key, we store it in `HackedKey` and change that position in the matrix to -1 (an impossible value) to avoid using it again. 

After that, we update our coordinates based on +1 or -1 (using `negativeX`/`negativeY`) to continue from the next position.

## Additional Information
You can use this link to use my Postman requests: https://www.postman.com/aerospace-astronomer-79003843/pin-punk-2077/collection/2eupdny/pin-up

I really wanted to use gRPC, but passing the matrix through replicated types or other ways was too painful to look at. So, I used HTTP instead.

There is one test for the CountMatches function, which checks the functionality of the other four matrix navigation functions in parallel.

I hate matrix tasks. Have a great day! 😉
