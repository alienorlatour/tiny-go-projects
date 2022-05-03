# Hello vibrating stars

## In this chapter
- Write to the standard output
- Test writing to the standard output
- Write table-driven tests
- Use flags to read command-line parameters
- Use a hash table to hold key-value pairs

## Introduction

`Hello world` in the desired human language



## All space travel begins on the ground

During the setup of your environment, you wrote a minimalistic Hello World. Let's understand it.

```go
package main
```

The file starts with the name of te package, `main`, which tells the compiler that the `main()` function will be found here and that's what should be executed when the program is run.

```go
import "fmt"
```

`fmt` is the standard-library package for formatting and outputting stuff. A very useful function for cheap debugging!

Non-standard libraries are identified by the URL to their repository. More on this later. For the moment, know that any import that does not look like a URL is from the standard library.

Finally, we have the `main()` function itself. It does not take any argument and does not return anything. Simple. Go is a simple language.

```go
func main() {
	fmt.Println("Hello, world!")
}
```

Inside the `fmt` package, `Println` _formats using the default formats for its operands and writes to standard output._

#### Why the long letter, my friend?

The long chapter about scope and visibility:

> Any symbol starting with a capital is exposed to external users of the package. Anything that starts with a lower-case letter stays inside.

This applies to variables, constants, functions, types, etc.

And that's it. Really.

That's why the function `Println()` starts with a capital.

TODO: Add the  

## Let's test!

First, we need a test file `main_internal_test.go` : 
- `main` for the function we test;
- `internal` because we want to access unexposed methods;
- `test` for the testing file.

## Let's raise the standard

To test the standard output, we need to name the test `Example_<function_name>`, here we test `Example_greet` method.

Let's add the call to the method:
```go
main()
```

To assert the expected output `Hello, world` to the obtained standard output line, we will use [Examples](https://pkg.go.dev/testing#hdr-Examples).
The tested function should include a line comment that begins by `Output:`. The text is compared to the standard output of the function.

```go
// Output: Hello, world!
```

The full file looks like below:
```go
package main

func Example_main() {
	main()
	// Output: Hello, world!
}
```

### Let's run the test

To run a test, you may call the test command from go:

```bash
go test 
```

The output lists the launched tests and their results.
```bash
=== RUN   Example_main
--- PASS: Example_main (0.00s)
PASS
```

## Are you polyglot?

Because we want to display `Hello, word` in different languages, passing the desired one is necessary.

### Duck typing

Typing is important. So we know what we are talking about. 

The input language will be a `string` but more precisely a `locale`.

```go
type locale string
```

### Selecting the right language

Now that we have a type, we can pass it as a parameter to the function `greet`.  
The new signature will be as below:
```go
func greet(l locale) string
```

For the first iteration, we can add a `switch` on the `locale` and return the corresponding greeting.
The default value for the moment is just an empty string.

```go
switch l {
	case "en":
		return "Hello, world!"
	case "fr":
		return "Bonjour le monde!"
	default:
		return ""
	}
```

In the main, we need to pass the desired `locale` to `greet` function and print the output.
For example, `"en"` for english.

```go
func main() {
	hello := greet("en")
	fmt.Println(hello)
}
```

### Enrich the test

Now we want to test the `greet` function with the various possible input. 
Since we are testing the returned value of the function, the testing function is called `Test_greet`.  

We will make a call to the `greet` function by passing the desired input language and store the output in a variable, so we can verify it.

```go
msg := greet("en")
```

`Errorf` from the testing package, creates an error with the given message as argument.
We will use it the expected output is different from the obtained one.

The full test looks like below:
```go
func Test_greet(t *testing.T) {
	msg := greet("en")
	if msg != "Hello, world!" {
		t.Errorf("expected: Hello, world!, got: %s", msg)
	}

	msg = greet("fr")
	if msg != "Bonjour le monde!" {
		t.Errorf("expected: Bonjour le monde!, got: %s", msg)
	}
}
```

## [Second iteration] Using a `map` 

Adding entries to a `switch` clause in Go reduces the readability of the code: it increases the size of the function, sometimes beyond screen dimensions.
For this reason, we decided to use a `map`, a very common and useful data structure in Go. A `map` is a set of pairs of distinct keys and their associated values.

Our map will associate a greeting message to every locale as a pair of `{locale, greeting}`.
For this chapter, we will use a global variable - something one shouldn't do in a production project.
```go
// dictionary holds greeting for each supported language
var dictionary = map[locale]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}
```
We now want to use this dictionary instead of the `switch` in the `greet` function.

Accessing an item in our `map` returns a `greeting` - the message associated with the key locale `l` - and a boolean that informs us of whether the key was found

```go
// greet says hello to the world
func greet(l locale) string {
	msg, ok := dictionary[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return msg
}
```





<hr>

On a un hello wolrd, on est contente.
Ecrivons un test
Example permet de tester la sortie standard - lien vers la doc
Pourquoi _internal_test.go
Run the test with the wrong string, make sure it’s red
Correct the test
Check your CI
-- commit - commit conventions are not part of the scope
passer une langue en paramètre
lang = pas une string, mais un type spécifique pour la lisibilité
commençons par un switch, on verra plus tard que ce n’est pas scalable
locale = 2 caractères, restons simple
greet va retourner la string Hello World dans la langue demandée : on ajoute le type de retour, c’est donc le main qui est testé avec Example.
Ecrire un vrai TU simple sur la fonction greet, vérifier qu’il est rouge, corriger, vérifier qu’il est vert - ajouter le test sur le français, ajouter un test sur le default par exemple avec le Swahili, et aussi avec la chaine vide au cas où
Utilisation des ` pour ne pas avoir à échapper toutes les “
-- commit
Ce switch n’est pas viable, écrire une map. Évitez les variables globales dans la majorité des cas, au-delà du Hello
Commentez vos variables et vos méthodes
Remplacer le switch par la lecture de la map. Les valeurs par défaut, c’est la vie !
-- commit
msg, ok // expliquer la syntaxe - on peut omettre le ok
modifier le return par défaut: unsupported language
mettre à jour les tests
-- commit
Le test est illisible et pas scalable - faire une map de cas de test
t.Run()
Vérifier que c’est vert

https://pkg.go.dev/fmt#Println


---- 

- Réorganisation du chapitre : mettre un seul _Let's test_, à la fin.