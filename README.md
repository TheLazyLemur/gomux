# Gomux

Gomux is a command-line interface application that acts as a tmux session manager. It is written in Go.
> **Disclaimer: This is unfinished software and should not be used in production environments. Use at your own risk.**


## Usage

To use Gomux, you need to have tmux installed on your system. Once you have tmux installed, you can run Gomux by providing it with a configuration file:

```sh
$ gomux -c config-file
```

The configuration file is a plain text file that specifies the layout of your tmux session.
Here's an example configuration file:

```
session-name sentinel
root-dir /home/dan/Workspace/VATIT/sentinel-service

new-session

split-pane v 10
select-pane 1
split-pane h 50

select-pane 0
send-keys 0 nvim

select-pane 2
send-keys 0 ./gradlew test

select-pane 0
attach-session
```

The above configuration file creates a new tmux session named "my-session" with a horizontal split pane that takes up 50% of the window. It then sends the commands "echo 'Hello, world!'" and "ls -la" to the first and second panes, respectively.

[Example config](/media/demo.gif)

You can find more information on the different commands in the source code.
