# Spec: a greet command

## What this is for

The canary program has one command, `add`. This adds a second one, `greet`, so
that there is a small, self-contained change for the Fleet engine to carry all
the way from an issue to a merged pull request.

## What it must do

Running `canary greet Ben` prints exactly one line:

```
Hello, Ben!
```

Running `canary greet` with no name prints exactly one line:

```
Hello, world!
```

Both cases exit with status 0.

A name made only of spaces counts as no name at all, so `canary greet "   "`
prints `Hello, world!`.

## How it must be built

- Add a function `Greet(name string) string` in the main package. It returns
  the line to print, without a trailing newline. It must not print anything
  itself and must not exit.
- Wire `greet` into the switch in `main`, taking the name from the first
  argument after the command, if there is one.
- Leave the `add` command exactly as it is.

## How it will be checked

Add tests in `main_test.go` covering all three cases above: a name, no name,
and a name that is only spaces. The existing test for `Add` must still pass.
The repository's test workflow runs `go vet ./...` and `go test ./...`, and
both must be clean.

## Out of scope

- Any change to how the program handles unknown commands.
- Any new dependency. The standard library is enough.
