#!/bin/bash

echo ""
go run . "" | cat -e
echo '-------------------------'
echo "\n"
go run . "\n" | cat -e
echo '-------------------------'
echo "Hello\n"
go run . "Hello\n" | cat -e
echo '-------------------------'
echo "hello"
go run . "hello" | cat -e
echo '-------------------------'
echo "HeLlO"
go run . "HeLlO" | cat -e
echo '-------------------------'
echo "Hello There"
go run . "Hello There" | cat -e
echo '-------------------------'
echo "1Hello 2There"
go run . "1Hello 2There" | cat -e
echo '-------------------------'
echo "{Hello There}"
go run . "{Hello There}" | cat -e
echo '-------------------------'
echo "Hello\nThere"
go run . "Hello\nThere" | cat -e
echo '-------------------------'
echo "Hello\n\nThere"
go run . "Hello\n\nThere" | cat -e
echo '-------------------------'

echo 'Audit'
echo '*******'

echo  "hello"
go run .  "hello" | cat -e
echo '-------------------------'
echo  "HELLO"
go run .  "HELLO" | cat -e
echo '-------------------------'
echo  "HeLlo HuMaN"
go run .  "HeLlo HuMaN" | cat -e
echo '-------------------------'
echo  "1Hello 2There"
go run .  "1Hello 2There" | cat -e
echo '-------------------------'
echo  "Hello\nThere"
go run .  "Hello\\nThere" | cat -e
echo '-------------------------'
echo  "Hello\n\nThere"
go run .  "Hello\\n\\nThere" | cat -e
echo '-------------------------'
echo  "{Hello & There #}"
go run .  "{Hello & There #}" | cat -e
echo '-------------------------'
echo  'hello There 1 to 2!'
go run .  "hello There 1 to 2!" | cat -e
echo '-------------------------'
echo  "MaD3IrA&LiSboN"
go run .  "MaD3IrA&LiSboN" | cat -e
echo '-------------------------'
echo  "1a\"#FdwHywR&/()="
go run .  "1a\"#FdwHywR&/()=" | cat -e
echo '-------------------------'
echo  "{|}~"
go run .  "{|}~" | cat -e
echo '-------------------------'
echo  "[\]^_ 'a"
go run .  "[\\]^_ 'a" | cat -e
echo '-------------------------'
echo  "RGB"
go run .  "RGB" | cat -e
echo '-------------------------'
echo  ":;<=>?@"
go run .  ":;<=>?@" | cat -e
echo '-------------------------'
echo  '\!" #$%&'"'"'()*+,-./'
go run .  "\\!\" #$%&'()*+,-./" | cat -e
echo '-------------------------'
echo  "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
go run .  "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e
echo '-------------------------'
echo  "abcdefghijklmnopqrstuvwxyz"
go run .  "abcdefghijklmnopqrstuvwxyz" | cat -e
echo '-------------------------'


echo 'Random'
echo '******'

echo 'Try passing <a random string> with at least four lower case letters and three upper case letters.'
echo "AaBbCcDdEe"
go run . "AaBbCcDdEe" | cat -e
echo '-------------------------'

echo 'Try passing <a random string> with at least five lower case letters, a space and two numbers.'
echo "zxcvb 33"
go run . "zxcvb 33" | cat -e
echo '-------------------------'

echo 'Try passing <a random string> with at least one upper case letters and 3 special characters.'
echo "A%@#"
go run . "A%@#" | cat -e
echo '-------------------------'

echo 'Try passing <a random string> with at least two lower case letters, two spaces, one number, two special characters and three upper case letters.'
echo "xc  3%^WSE"
go run . "xc  3%^WSE" | cat -e
echo '-------------------------'

