# hmac

In this exercise, you will practice working with `HMAC` signatures as described in this [tutorial](https://drstearns.github.io/tutorials/sessions/#secsessiontokens).
This will help you prepare for `Assignment 3: Tracking Sessions`.

## Context
What you are building in this exercise is a command line function that allows you to specify certain arguments in order to 
create a `HMAC` signature that you can verify.
For example running:
```bash
hmac sign secret < test.txt
```
- `hmac`: Executable being ran
- `sign`: String that represents a function to call inside your executable
- `secret`: A random string that represents your hashing key. This can literally be anything
- `<`: Redirect command. See more [here](https://www.digitalocean.com/community/tutorials/an-introduction-to-linux-i-o-redirection)
- `text.txt`: Text file that we use to base our signature on

Will generate a `HMAC` signature to `stdout` (your terminal)
```bash
AHYFAcm0TnHvpWdQoyeWdeHgy-t54nK-4u8xsK2_cTg=
```
With the provided signature, you can now verify it as shown below:
```bash
hmac verify secret AHYFAcm0TnHvpWdQoyeWdeHgy-t54nK-4u8xsK2_cTg= < text.txt
```
Which will result in an output like the following:
```bash
Signature is Valid
```
Or 
```bash
Signature is Invalid
```
## Instructions
In order to finish this exercise, you will need to add code to the `hmac.go` file. 
Instructions are provided for each function. 
`hmac_test.go` is provided for testing purposes.

General Tips:
- Work on `Sign` first then `Verify`, run the test along the way in order to 
incrementally build up your functions.
- Remember that you will need to run `go install` every time you make changes to the files in order 
to update your executable.


