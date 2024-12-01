# Build Your Own wc Tool

### This challenge is to build your own version of the Unix command line tool wc!


- The Unix command line tools are a great metaphor for good software engineering and they follow the Unix Philosophies of:

- Writing simple parts connected by clean interfaces - each tool does just one thing and provides a simple CLI that handles text input from either files or file streams.
Design programs to be connected to other programs - each tool can be easily connected to other tools to create incredibly powerful compositions.

## Step One:
In this step your goal is to write a simple version of wc that takes the command line option **-c** and outputs the number of bytes in a file.
### Example:
```sh
gowc -c test.txt
342190 test.txt
```

## Step Two:
In this step your goal is to support the command line option **-l** that outputs the number of lines in a file.
### Example:
```sh
gowc -l test.txt
7145 test.txt
```

## Step Three:
In this step your goal is to support the command line option **-w** that outputs the number of words in a file:
### Example:
```sh
gowc -w test.txt
58164 test.txt
```
## Step Four:
In this step your goal is to support the command line option **-m** that outputs the number of characters in a file. If the current locale does not support multibyte characters this will match the **-c** option
### Example:
```sh
gowc -m test.txt
339292 test.txt
```

## Step Four:
In this step your goal is to support the default option - i.e. no options are provided, which is the equivalent to the **-c**, **-l** and **-w** options:
```sh
gowc test.txt
7145   58164  342190 test.txt
```

## Final Four:
In this step your goal is to support being able to read from standard input if no filename is specified:
```sh
cat test.txt | gowc -l
7145
```
