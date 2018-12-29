#Simple Web Server

##kill process on given port
`netstat -vanp tcp | grep <portnumber>` second last column is the PID
`kill -9 <PID>`  to end the process