## ascii-art-color

### Objectives

- Ascii-art-color is a program which consists in receiving a `string` as an argument and outputting the `string` in a graphic representation using ASCIIб you can manipulate the color.

- What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:



![Hello](/1.png)

- supports standard colors

### Instructions

- Download Repository
- Open in code editor (for example: [VS Code](https://code.visualstudio.com))
- Type in terminal

```console
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <letters to be colored> "something"
EX: go run . --color=<color> "something"
```

### Usage

```console
go run . --color=orange GuYs "HeY GuYs"
```

![Hello](/2.png)

```console
go run . --color=green "HeY GuYs"
```
![Hello](/3.png)

```console
go run . --color=blue B "RGB"
```
![Hello](/4.png)


### Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed

