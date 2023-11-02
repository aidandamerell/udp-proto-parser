package main
	
import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
)

func check(e error) {
	if e != nil {
    panic(e)
  }
}

func main() {

  pathPtr := flag.String("path", "", "Path to UDP proto scanner text file")

  // Parse the flag above
  flag.Parse()

  // Open the provided file
  data, err := os.Open(*pathPtr)

  // Ensure opening the file hasn't raised any errors
  check(err)

  // Build a buffer for the file reading
  fileScanner := bufio.NewScanner(data)
  fileScanner.Split(bufio.ScanLines)
  // define an array for the lines
  var fileLines []string

  // Put each line of the file into the array
  for fileScanner.Scan() {
    fileLines = append(fileLines, fileScanner.Text())
  }

  // Close the file back up
  data.Close()

  // Define a regex to match the relevant lines, with match groups
  receivedRegex, _ := regexp.Compile(`^Received.+port (.+)\) from (.+?)\:`)


  // Print a header to start
  fmt.Println("IP Address\tProtocol\tPort\tState")

  // Iterate over the lines, matching the regex and then
  // print the output if required.
  for _, line := range fileLines {

    matches := receivedRegex.FindStringSubmatch(line)

    if len(matches) == 0 {
      continue
    }

    fmt.Printf("%v\tUDP\t%v\topen\n", matches[2], matches[1])

  }
}
