# TODOs
- `log` devrait lancer une go routine (on ne veut pas etre bloquant sur le .Write())
- ajouter un formatteur ? (json, csv, html, ...)
- une part non nulle de loggers ont une {key,value} (pratique pour elasticsearch) - je propose de ne pas l'inclure
- `logger.Log(logger.Level, format string, args ...interface {})` ? Permet aux autres libs de mapper un niveau de log sur le notre


## I wanted to be a lumberjack

In the previous chapter, we saw how to write to the standard output. It's important not to belittle this activity, as it is still used in a way we do not condone.

The night is dark. You've been working on trying to fix this bug for 4 hours straight. You don't understand what's happening with a `uri` variable that you expect to be non-empty when you download a resource.

And this happens. You decide to add this line in the code, and rerun everything, to get a better insight as to what's going on.  
```go
func download(uri string) {
	...
	fmt.Println("in download, step 2, uri is %q", uri)
	...
}
```

We, human beings, are used to writing - and reading - in non-machine languages. Having the code write us something we can understand is our easiest way of following the program as it executes.
 
Unfortunately, sometimes, replaying the scenario is not an option - maybe it removes a file you don't want to remove, or it crashes the database. A smart move is, in these cases, to write to the user that _"Everything is going extremely well"_.

Keeping track of current state or events is called _logging_. Every piece of tracked information is _a log_, and _to log_ is the associated action. 

Etymologically, a _log_ relates to a ship's _logbook_, a document in which were written records of the speed and progress of the ship. 
The _logbook_'s name deriving, itself, from a _chip log_, a piece of wood attached to a string that was tossed into the water to measure the speed of the vessel.

Every application has a _logger_, whose task is to write messages at specific moments in its execution so that they could be read and analysed later, if need be. Sometimes, we want these messages to be written to a file. Sometimes, to the standard output. Sometimes, to a printer, or streamed through the network.

However, not all messages bear the same amount of information. _"Everything is going extremely well"_ is very different from _"I just picked up a fault in the AE-35 Unit"_. We might want to emphasise critical messages, or discard those of lesser matter. 
Acknowledging that there are different degrees of importance was already performed by scribes in Ancient Egypt, when they'd write some sections in red (this is where the word _rubric_ originates), to highlight them.   

The goal of this chapter is to implement such a logger, and to explain why we can't simply use `fmt.Println()` for such a task. 

### Levels

It is necessary, when using a logger, to assign an importance level to a message. This is the task of the developer, who has to think about the criticalness of the information that is about to be recorded.
Loggers around the world have a wide variety of levels, which always follow the same pattern: the one with the lowest importance (usually "Trace" or "Debug"), and then it goes up to (usually) "Fatal".  

Let's dig into our example: we decided to offer 3 levels of logging. They can be found in the `level.go` file, and they are our introduction to enumerations, in Go:
```go
// Level describes the criticalness of a log message
type Level int

const (
    // LevelDebug messages are used to debug the application
    LevelDebug Level = iota
    // LevelInfo messages are used to log meaningful information about the processes going on
    LevelInfo
    // LevelError messages are used to highlight unexpected behaviours caught by the application
    LevelError
)
```

We start by declaring a _named type_ `Level`, of _underlying type_ `int`. `Level` can be used by other packages, as it is exposed. 

Then, we declare constants of the type `Level`. The syntax here is to use `= iota` to let the compiler know that we are starting an enumeration - a list of entities of the same kind.
We don't need to assign explicit values to these constants, the compiler does it automatically for us thanks to the `iota` syntax. If we decide to add a `Warning` level later, we will only need to add a line and not worry about renumbering everything. 

`= iota` can be used on any type that can be equal to an `int`.

We now have 3 log levels, each one will have its own purpose.

### A logger in Go

Let's define the `logger` package. Here are the exposed structure and functions in the `logger.go` file:
```go
// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(level Level) *Logger

// WithOutput sets the output of the logger, and returns it.
// You can call logger.WithOutput(os.StdOut).Info().
// The thread safety is left to the implementation of output.
func (l Logger) WithOutput(output io.Writer) Logger

// Debug formats and prints a message if the log level is debug or higher
func (l Logger) Debug(format string, args ...any)

// Info formats and prints a message if the log level is info or higher
func (l Logger) Info(format string, args ...any)

// Error formats and prints a message if the log level is error or higher
func (l Logger) Error(format string, args ...any)
```

- Input

an input can be anything we want to save

- Output

the output should be anywhere we want to save

- 