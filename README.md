# goXercice

Greetings and welcome to goXercice. 
This is a small project that contains small exercices to introduce you to the Go programming language.

I recommend you to take a look at the examples from [A tour of Go](https://go.dev/tour/welcome/1).
This was a great inspiration for this project, and some exercices are taken from the site. 

## Getting Started

### Installing Go

Before installing goXercice, install the latest version of Go. Check out [https://go.dev/](https://go.dev/)

### Installing/Updating goXercice

#### Using Go

Make sure Go is installed on your machine. You can check it by running:
```bash
go version
```

Run the following command to install GoXercice:
```bash
go install github.com/Isines1611/goXercice/goxercice@latest
```

Verify that the installation was successful by running:
```bash
goxercice version
```

If you have errors downloading or launching the app, it might come from your `PATH`. Please follow the official instructions at [https://go.dev/doc/install](https://go.dev/doc/install) or try adding `export PATH=$PATH:$HOME/go/bin` to your `PATH`.

#### Manual Installation

If you prefer to build the CLI manually, follow these steps:
Clone the repository:
```bash
git clone https://github.com/Isines1611/goXercice.git
```

Change directory to the goxercice folder and build the CLI:
```bash
cd goXercice
go build -o goxercice
```

Now you can run the CLI from that directory with:

```bash
./goxercice version
```

### Initialization

Once Go and goXercice are installed, you can run `goxercice init [fullpath]` ([more details here](#terminal)) to initialize everything and start solving exercices. `TODO` will explain what is expected from you for the exercice. Please stick to what is asked and don't modify more. Also, it is worth mentioning that a majority of exercices can be validated by printing a certain string; however, you won't learn anything by doing that. 

## Working environment

### Organization

Once goXercice is installed, you can initialize it, which downloads the files needed in the chosen directory [more details below](#terminal). In this directory, there will be the `exercices` and `solutions` directories. Both directories are almost similar, unless `exercices` contains the files you will need to complete and `solutions` will contain the expected solutions. While goXercice has hints to help you, if you are still stuck on an exercice, even with the hints, you can check the associated solution and move on. Please try avoiding as much as possible to rely on the `solutions` directory.

### Errors

Exercices are designed to be able to be run on their own. This means that you can run any exercices alone (or solution) by running:

```go
go run [PATH-TO-FILE]
```

This design requires that every file has a `main` function. This causes some errors in an editor, but you can ignore them; they are expected. 

### Correction

By default, goXercice compares the output of your file with the associated solutions. 
This design allows ignoring custom habits like naming the variables. 
To trigger the correction, goXercice watch on the exercice file changes. 
When it detects a modification, then the correction mechanism is trigger.
You can also use command `v` or `verify` once in the CLI mode.

### Commands

goXercice offers several commands; here is a full explanation of them:

##### Terminal

- `goxercice help [command]`: Show a help message with the associated explanations.
- `goxercice init [fullpath]`: Initialize the program. It fills the directory `~/.goxercice` with the config files and downloads the exerices files in the given directory or the current directory. Then, it enters CLI mode to start solving exercices. Please provide a full path to avoid any issue.
- `goxercice hint`: Show the next hint of the pending exercice.
- `goxercice reset`: Needs confirmation. This will delete the directory with the config files (`~/.goxercice`).
- `goxercice resume`: Enter CLI mode to continue solving exercices.
- `goxercice verify`: Verifies the pending exercice. It updates progress if it is correct.
- `goxercice verifyAll`: Verifies all exercices until one is incorrect. It updates progress based on the number of correct exercices.
- `goxercice version`: Shows the current version. It should be the same as [here](#version)

##### CLI mode

In CLI mode, you just need to enter the letter or the full name. Here is what they do

- `(v) verify`: Trigger the same effect as `goxercice verify`.
- `(h) hint`: Trigger the same effect as `goxercice hint`.
- `(q) quit`: Quit the CLI mode.
- `(n) next`: After successfully doing an exercice, you can enter `(n) next` to move onto the next exercice.

## What's next ?

Once you've completed all the exercices, you have a good basic knowledge of the Go programming language. Continue practicing and improving by doing your own project or contributing to others' projects.

## Uninstall 

To remove goXercice from your system, start by running the following command:

```bash
goxercice reset
```

This removes all config files. 
You can then delete the directory with the exercices and solutions.

If you installed goXercice via `go install`, then you need to remove the binary file in your `$GOPATH/bin` directory.
You can verify it is deleted by running `goxercice version`. The output should be `command not found: goxercice`. 

If you installed goXercice manually, then delete all the files used to build the tool.

##### Version
Currently on version 1.0
