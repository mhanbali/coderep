# CodeRep

A great way to learn and remember what you learned is through repition.

CodeRep helps you remember what you learned by giving you a barebones editor to put in the reps and run your practice code directly from the app and receive the output.

It takes in a json file with a list of exercises and lets you select which one to work on. It currently doesn't check for correctness, but this is for your own personal use anyway, so don't cheat yourself :P

So far I've put some exercises in Go under the exercises folder. You can easily create your own, or check back for more that I'll be creating and adding.

It runs the command directly to bash and takes the output to display it in the column to the right.

## Screenshots

![App Screenshot](https://github.com/mhanbali/coderep/blob/master/screenshots/coderep-ss2.png?raw=true)

## Run It

Clone the project

```bash
  git clone https://github.com/mhanbali/coderep
```

Go to the project directory

```bash
  cd coderep
```

Running it requires you pass in the exercise file in the exercises folder.

```shell
go run . exercises/go.json
```

If you want to build a standalone executable just run

```shell
go build
```

then to run it

```shell
./coderep exercises/go.json
```

## Using and creating exercises

To create your own exercises you can create a file.json and pass that file into the app as an argument.

```json
{
  "ext": ".go",
  "command": "go run code/code.go",
  "auto_command": false,
  "exercises": [
    {
      "title": "Variables",
      "instructions": "Assign variables for integers, strings, and booleans.\n\nUse both the shorthand and long form methods.\n\nThen print each variable."
    },
    {
      "title": "Write for loops",
      "instructions": "Write while loops"
    }
  ]
}
```

- "ext" is the type of code you'll be running
- "command": if you decide to use a button to run your commands then this is the command that the button will run.
- "auto_command":
  - false = enables the text field to run your own commands
  - true = enables the button that runs the "command"

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Acknowledgements

The tview library is what makes the UI so easy to make

- [https://github.com/rivo/tview](https://github.com/rivo/tview)
