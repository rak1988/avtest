# Prequisites
go runtime needs to be installed

# SETUP

cd into go src home directory and run:
git clone https://github.com/rak1988/avtest.git

cd into cloned directory

go get github.com/mattn/go-sqlite3



### How to Run

## export binpath for binary target
export GOBIN=\<projectpath\>/bin
## load the json data into sqlite file
./loadDataIntoSqlite -f \<inputfile in json format\>
## start TCP server
./server
## Open another terminal and run
./show_avg_marks -studentids \<studentids\>
  
### Example

cd into \<projectpath\>/bin folder
./loadDataIntoSqlite -f \.\./inputfiles/student_marks.json
./server
in another terminal
./show_avg_marks -studentids 456,567,3456,1267

## Output Format:
 Avg Marks of 456: 55.33
 Avg Marks of 567: 52.67
 Avg Marks of 3456: 34.67
 Avg Marks of 126: 65.67


