ipfmt - ip converter
====================

Convert IP address between 5 standard formats:

  * Dotted Hexadecimal
  * Dotted Octal
  * Undotted Integer
  * Undotted Hexadecimal
  * Combo in format Hexadecimal.Octal.Integer.Hexadecimal


Usage
-----

    # ipfmt 127.0.0.1
    ip              : 127.0.0.1
    hex             : 0x7f.0x0.0x0.0x1
    octal           : 0177.0.0.01
    integer         : 2130706433
    hex (no dotted) : 0x7f000001
    combo (h.d.o.h) : 0x7f.0.0.0x1

