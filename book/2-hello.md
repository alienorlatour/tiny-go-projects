
## Hello vibrating stars

Hello world in a random human language

### All space travel begins on the ground

During the setup of your environment, you wrote a minimalistic Hello World. Let's understand it.

```go
package main
```

The file starts with the name of te package, `main`, which tells the compiler that the `main()` function will be found here and that's what should be executed when the program is run.

```go
import "fmt"
```

`fmt` is the standard-library package for formatting and outputting stuff. A very useful function for cheap debugging!

Non-standard libraries are identified by the URL to their repository. More on this later. For the moment, know that any import that does not look like a URL is from the standard lib.

Finally, we have the `main()` function itself. It does not take any argument and does not return anything. Simple. Go is a simple language.

```go
func main() {
	fmt.Println("Hello, world!")
}
```

Inside the `fmt` package, `Println` _formats using the default formats for its operands and writes to standard output._

##### Why the long letter, my friend?

The long chapter about scope and visibility:

> Any symbol starting with a capital is exposed to external users of the package. Anything that starts with a lower-case letter stays inside.

This applies to variables, constants, functions, types, etc.

And that's it. Really.

That's why the function `Println()` starts with a capital.

### Let's test!



<hr>

On a un hello wolrd, on est content.
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
