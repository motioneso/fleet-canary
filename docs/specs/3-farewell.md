# Spec: a farewell command

## What this is for

The canary program has `add` and `greet`. This adds a third command,
`farewell`, so that there is a second small, self-contained change for the
Fleet engine to carry from an issue to a merged pull request. This one exists
to prove that the engine stops at the merge and waits for a person.

## What it must do

Running `canary farewell Ben` prints exactly one line:

```
Goodbye, Ben!
```

Running `canary farewell` with no name prints exactly one line:

```
Goodbye, world!
```

Both cases exit with status 0.

A name made only of spaces counts as no name at all, so `canary farewell "   "`
prints `Goodbye, world!`.

## How it must be built

- Add a function `Farewell(name string) string` in the main package. It returns
  the line to print, without a trailing newline. It must not print anything
  itself and must not exit.
- Wire `farewell` into the switch in `main`, taking the name from the first
  argument after the command, if there is one.
- Leave the `add` and `greet` commands exactly as they are.

## How it will be checked

Add tests in `main_test.go` covering all three cases above: a name, no name,
and a name that is only spaces. The existing tests must still pass.
