# Prequisites
go runtime needs to be installed

# SETUP

cd into go src home directory and run:
git clone <>

cd into cloned directory

go get github.com/mattn/go-sqlite3



### How to Run

## export binpath for binary target
export GOBIN=<projectpath>/bin
## load the json data into sqlite file
./loadDataIntoSqlite -f <inputfile in json format>
## start TCP server
./server
## Open another terminal and run
./show_avg_marks -studentids <studentids>


