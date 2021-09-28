# dcfs

dcfs is a cli tool which helps you work with the [decomposed fs](https://github.com/cs3org/reva/tree/master/pkg/storage/utils/decomposedfs) used in [reva](https://github.com/cs3org/reva/).

## Usage
```
NAME:
   dcfs - A tool for working with the decomposed fs

USAGE:
   dcfs [global options] command [command options] [arguments...]

COMMANDS:
   tree, tr  List the contents of the decomposed fs in a tree-like format
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

Currently only the tree command is implemented. You can execute that command in the nodes folder of the decomposed fs

```
storage/users/nodes
❯ dcfs tr
4c510ada-c86b-4815-8820-42cdf82c3d51
  Photos
    Memes
      OOP.jpg
      css.jpeg
  oc.png
932b4540-8d16-481e-8ef4-588e4b6b151c
  Work
    oauth2.drawio
    oauth2.png
  the_last_question_-_issac_asimov.pdf
```
Or pass the path to the decomposed fs root as a parameter

```
dcfs tr /var/tmp/ocis/storage/users/nodes/
4c510ada-c86b-4815-8820-42cdf82c3d51
  Photos
    Memes
      OOP.jpg
      css.jpeg
  oc.png
932b4540-8d16-481e-8ef4-588e4b6b151c
  Work
    oauth2.drawio
    oauth2.png
  the_last_question_-_issac_asimov.pdf
```
