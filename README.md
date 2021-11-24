Step 1) 

`git clone https://github.com/0xj1mmy/minimal-mempool-logger.git`

Step 2) 

`cd minimal-mempool-logger && go build`

Step 3) 

`./minimal-mempool-logger > txpool.log` to store results in a log file

* Ideally, do this in a tmux shell to keep it running

Step 4) 

Stream logs with `tail -f txpool.log`
