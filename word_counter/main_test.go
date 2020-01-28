package main
import(
"bytes"
"testing"
)

/*Test the word-counter program, to make sure it correctly counts the number of words or from stdin*/

func TestLineCount(t *testing.T){
b := bytes.NewBufferString("word1 word2 word3 \nline2\nline3 word1")
//If the program received anything other than 3 lines, the test fails
exp := 3
res := count(b, true)
if res != exp {
t.Errorf("Expected %d, got %d instead.\n", exp, res)
}
}

func TestWordCount(t *testing.T) {
/*Create a buffer of bytes from a string containing 4 words.
Then, parse the buffer to the count function.
If this function returns anything but 4, the test fails and we display an error.
If it returns 4, the test parses*/
b := bytes.NewBufferString("word1 word2 word3 word4\n")
exp := 4
res := count(b, false)
if res != exp {
t.Errorf("Expected %d, got %d instead.\n", exp, res)
}
}
