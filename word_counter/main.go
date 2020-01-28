package main
/*os allows us to use system resources, bufio allows us to read files*/
import(
"fmt"
"io"
"bufio"
"os"
"flag"
)
func main (){
//Define a boolean flag that will enable the program to count lines when set to 'true'
lines := flag.Bool("l", false, "Count lines")
//Another flag to count the number of bytes when set to true
bytes := flag.Bool("b", false, "Count bytes")

//parse the flags provided by the user
flag.Parse()
/*Call the function that counts and prints the number of words, bytes or lines parsed to the program
This function takes Stdin, lines and bytesas it's arguments
Stdin is assigned to 'r io.Reader', 'lines' is assigned to 'countLines, and bytes is assigned to countBytes*/
fmt.Println(count(os.Stdin, *lines, *bytes))}

func count (r io.Reader, countLines bool, countBytes bool) int{
//Use a scanner to read the words from a text file
scanner := bufio.NewScanner(r)
//Call the Split function on our Scanner object and define the split type as words
//default split type is lines
//Only call this if the 'lines' and 'countBytes' variables are set to false
if (countLines(false)  && countBytes(false)){
scanner.Split(bufio.ScanWords)
} else if countBytes(true) {
scanner.Split(bufio.ScanBytes)
} else if countLines(true) {
scanner.Split(bufio.ScanWords)
}

//Define a counter
wc := 0
for scanner.Scan() {
wc++
}
return wc
}
