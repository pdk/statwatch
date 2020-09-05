# statwatch

A stupid simple file watcher to help with restarting/rebuilding stuff when any
file changes.

Invoke with a directory name, and then a series of filename patterns (glob
style). Example:

    statwatch . '*.go' '*.html'

statwatch will find all files under the named directory that match any of the
patterns, and will then check every 500 milliseconds (2x per second) to see if
anything was modified. If/when any file is modified, the program will exit.

statwatch will skip any directory named like `.*` (other than `.`, the current
directly).

Example script:

    OK=0
    while [ $OK -eq 0 ]
    do
        go run cmd/path/prog.go args &
        PID=$!

        statwatch . '*.go' '*.html' '*.json'
        OK=$?

        pkill -9 -P $PID
    done
