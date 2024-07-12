# Learn Golang

Le langage de programmation Go a été développé par Google.
Il est l'un des plus en vue grâce à sa rapidité et à sa simplicité.

Dans cette apprentissage, je découvre son environnement de travail,
la syntaxe utilisée, les types de données, aussi bien  basiques que composites,
et les différentes fonctions du langage.

### Les domaines d'utilisation du langage Go sont les suivants:
* Développement web
* Jeux vidéo
* Logiciels et applications mobiles
* Serveurs et base de données
* Machine Learning
* IOT
* Blockchain
* etc

### Installer l'environnement de travail
Aller à l'url [https://go.dev/doc/install](https://go.dev/doc/install)

### Syntaxe du langage Go
Un fichier `.go` se compose généralement des éléments suivants:
* Déclaration de packages
* Importation de packages
* Des fonctions
* Déclarations et expressions

```shell
Example: définir dans src/main.go

package main // Package src permet de définir l'appartenance au programme au package src
import "fmt" // Importation des fonctions dans le package `fmt`

func main() { // déclaration de fonction
	fmt.Println("Hello World")
}
```

Exécution du programme avec la commande: `go run src/main.go`

### Identifiant
Élément important à prendre en compte est l'**`identifiant`**

Il s'agit du nom désigné pour identifier une variable, une fonction ou
tout autre élément définir par l'utilisateur. Il permet de nommer les
entités du programme.

Un identifiant peut-être:
* constituer d'une séquence de lettres et de chiffres
* le premier caractère doit commencer par une lettre ou un underscore (_)
* les caractères de ponctuation tels que `@`, `$` et `%`sont non autorisés
* Go est un langage sensible à la casse, ie `Flavien` est de `flavien`.

Les mots-clés réservés en `Go` sont:

```shell
break, case, chan, const, continu, default, defer, else, fallthrough, for
func, Go, Goto,if, import, interface, map, package, range, return
select, Struct, Switch, Type, Var
```

### Les types de données basiques du langage Go

* Les types numériques:
Il comprend

```shell
les types entiers

int: stocke les nombres entiers positifs et négatifs
uint: stocke les nombres uniquement positifs
uint8: stocke les entiers non signés de 8 bits (0 à 255)
int8: stocke les entiers signés de 8 bits (-128 à 127)
uint16, uint32 et uint64
int16, int32, int64
```

```shell
les flottants

float32: permet une précision à six chiffres
float64: permet une précison à quinze chiffres
```

```shell
les complexes sont exprimés sous la forme: a + bi
avec a et b des nombres réels et i une solution de l'équation i^2 = -1
Go offre deux types de nombres complexes:
  - Complex64: avec une précision pour la partie réelle et imaginaire égale à six chiffres
  - Complex128: avec une précison à quinze équivalants à deux float64
```

* Les types chaînes de caractères:

1. Le type chaîne de caractères pré-déclarer en Go est `string`.
2. Les chaîne de caractères sont `immuables`, ie qu'on fois créer, on ne peut modifier le contenu 
3. La fonction `len` permet de retourner la longueur de la chaîne

* Les boucles en Go:

La boucle `for`, permet de faire des itérations
Go ne possède pas de boucle while

* Les pointeurs:

Ce sont des objets contenant des adresses mémoires, il permet de partager des
données entre les fonctions Go. Il permet également de faire la différence
entre des valeurs, par exemple différencier la valeur zéro et une valeur non
définie.
