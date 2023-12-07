# Ducky Script Parser

<hr> 

Pre-requisites: golang 1.20.7+

Tired of having to manually put in all of the "ENTER"s and "STRING"s in a ducky script when converting a powershell function into a ducky script? Ducky Script parser will create a file based on the input file containing the powershell function, and it will export it into a new .txt file.

  -f string
<br>
        The file path of the file containing the desired script (.txt)
<br>
  -o string
<br>
        The file path of the desired output file (with no extension) 

Building for your platform:
1. Clone/download repo
2. Navigate to repo folder
3. Install Golang environment
4. Run go build
