# Go: Using reflection to create a generic lookup method

## Starting the journey in GoLang

As a newbie GoLang developer I had to deal with a set of new statements to be understood and digested. I also had to face a new way of thinking that shifted me from what I am used to in JS / TS / NodeJS.\
One aspect which was really tough at the beginning was `interfaces`. In my developer experience, I had already encountered interfaces in Java and I didn't like them too much. This wasn't a good way to start with them in Go.\
Then I realized how powerful and useful they could be while creating a whole Go application or library (and they are probably too in Java, no offense).

The other thing that gave me kind of bad time was the lack of Generics, which luckily has been recently filled in [Go 1.18](https://tip.golang.org/doc/go1.18#generics).
The latter was even more annoying when I dealt with a Go project by one of my clients, which is based on Go 1.16 and cannot be upgraded because of some dependency requirements.

Go is sometimes very verbose, and this can be either good or bad, depending on the situation. For example in that codebase I found a package named `arrays` which was taking care of looking up a specific element within a given array/slice.
Why a package? Because they could not find a way to achieve the `lookup in array` methods without creating each for any `type` that is needed. So any type that needs this lookup method, either built-in (string, int, int64) or custom (struct), should have its own method doing pretty much the same as all the others do.

When I see too many repetitions and copy/paste of code I smell a rat and I like to remove all the repetitions as soon as possible. Using code repetitions is not only evil, but when you find a bug in a part of the code they may have been repeated tons of times.\
What I wanted is a generic method that can lookup for a value in a slice only if they both hold the same `type`.

## Hands-on

Creating a generic method was pretty simple as changing the parameters from a given type

```go
func ContainsInt(items []int, value int) bool { ... }
```

to the most generic one in Go: interface{} also known as `any interface`.

```go
func containsAny(items interface{}, value interface{}) bool { ... }
```

But for the body of the method it is another story, even if its original version was pretty trivial:

```go
func ContainsInt(items []int, value int) bool {
    for _, item := range items {
    if item == value {
           return true
       }
    }
    return false
}
```

The best way to deal with this is the `reflect` package which holds all the methods to reflect, meant as analyzing at its deepest level, a data structure.

The first problem is that using an `any interface` type would lose the explicit type check at compile time, which is one of the best things when writing Go code (but I'll tell you more on that later). This is even worse because the expected parameters should be: a slice and a "same-type" single value.

First thing to do is checking that the first parameter holds effectively a slice data structure. This can be achieved with the `Kind` method of the `reflect` package.

```go
if itemsValue.Kind() != reflect.Slice { return false }
```

Once that we have been assured that the first type is effectively a slice, comparing each of its values with the given value is relatively easy. Indeed the `DeepEqual` method allows to compare two values and as the [official documentation](https://pkg.go.dev/reflect#DeepEqual) states:

> Two values of type Interface are deeply equal if they hold deeply equal concrete values.

So for each item in the slice the `Interface` method is called, which returns the current value as an `interface{}` and that can be deeply compared with the given value, which is already of type `interface{}`

```go
for i := 0; i < itemsValue.Len(); i++ {
       if reflect.DeepEqual(itemsValue.Index(i).Interface(), value) {
   return true
       }
    }
```

Moreover we don't even need `type` checking, because the previous method will yield implicitly `false` if the compared values are not of the same type. Again the official documentation:

> Values of distinct types are never deeply equal.

Here is the complete method statement, with a couple of debugging logs:

```go
func containsAny(items interface{}, value interface{}) bool {
    itemsValue := reflect.ValueOf(items)
    if itemsValue.Kind() != reflect.Slice {
       return false
    }

    for i := 0; i < itemsValue.Len(); i++ {
       if reflect.DeepEqual(itemsValue.Index(i).Interface(), value) { // value, short for -> reflect.Value(value).Interface()
           fmt.Printf("%s and %s \n", itemsValue.Index(i).Type().String(), reflect.ValueOf(value).Type().String())
           fmt.Printf("%v and %v \n", itemsValue.Index(i), reflect.ValueOf(value))
           return true
       }
    }
    return false
}
```

## Too much generic

As I wrote earlier now that everything is so generic I lost some benefits of type checking in Go. There will be no way, except at runtime, to know if the method has been called using the correct types, since interface{} can hold anything at compile time.

This led me to reuse the previous method declarations with explicit types from the original package but I replaced their body with a simple call to the generic method.

Several benefits come from that:

* exposed methods have a specific type declaration that is checked at compile time;
* generic method can be unexposed by the package since any call to it will happen within the package;
* there is only a source of truth (or of failure) which is the generic method itself;
* the package can be extended to an unlimited number of types only by creating a new public method declaration with the new given types, and it will need to call the sole generic method in it.

You can find the full code with a bunch of tests [here](https://github.com/sw360cab/go-reflect-compare)

## Conclusions

After starting learning Go almost one year ago and joining my first real and `in-production` project, all based on a GoLang codebase, I am satisfied with having faced these kinds of challenges and having achieved the next level skills.\
Probably Go _gurus_ will disagree with the solution I came up with, but I will be happy to read their wise advice. I know that for sure this was more an intellectual exercise, since with Go 1.18 all of that may be overtaken by built-in `Generics`.
