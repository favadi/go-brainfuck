# go-brainfuck

## Brainfuck Specs

Brainfuck is represented by an array with 30,000 cells initialized to zero
and a data pointer pointing at the current cell.

There are eight commands:

<dl>
  <dt>+</dt>
  <dd>Increments the value at the current cell by one.</dd>

  <dt>-</dt>
  <dd>Decrements the value at the current cell by one.</dd>

  <dt>></dt>
  <dd>Moves the data pointer to the next cell (cell on the right).</dd>

  <dt><</dt>
  <dd>Moves the data pointer to the previous cell (cell on the left).</dd>

  <dt>.</dt>
  <dd>Prints the ASCII value at the current cell (i.e. 65 = 'A').</dd>

  <dt>,</dt>
  <dd>Reads a single input character into the current cell.</dd>

  <dt>[</dt>
  <dd>
    If the value at the current cell is zero, skips to the corresponding ] .
    Otherwise, move to the next instruction.
  </dd>

  <dt>]</dt>
  <dd>
    If the value at the current cell is zero, move to the next instruction.
    Otherwise, move backwards in the instructions to the corresponding [ .
  </dd>
</dl>
