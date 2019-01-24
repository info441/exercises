# testing

I would argue that the one thing that separates professional from amateur developers the most is a commitment 
to writing automated tests for their features. Here you'll get exposure to writing unit test in Go which you will 
apply to your upcoming assignment.

Refer to this [tutorial](https://drstearns.github.io/tutorials/testing/) if you're stuck.

## Guidelines
1. Start with the `contact/` exercise. There's a subtle bug that you need to fix inside `contact.go`. Run the
test script (`contact_test.go`) in order to see how the function is behaving.
2. With the `mergesort/` exercise, you'll practice adding test cases to `mergesort_test.go`. Add several examples of 
input and expected output to robustly test the given implementation of `mergesort`.
3. In the `reverse/` exercise, you'll work primarily with implementing `Reverse()` in `reverse.go`. Make sure your implementation
passes all of the test cases!
4. The last exercise, `handlers/` will allow you to practice test-driven development. Here you'll get to work with both 
`handlers.go` and `handlers_test.go` in order to create a robust handler that covers additional request that you're required to 
add to `handlers_test.go`.