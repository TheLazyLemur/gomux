# Gomux

Gomux is a command-line interface application that acts as a tmux session manager. It is written in Go.
> **Disclaimer: This is unfinished software and should not be used in production environments. Use at your own risk.**


## Usage

To use Gomux, you need to have tmux installed on your system. Once you have tmux installed, you can run Gomux by providing it with a configuration file:

```sh
$ gomux -c config-file
```

The configuration file is a plain text file that specifies the layout of your tmux session. The format of the file is as follows:

```
new-session session-name root-dir
split-pane session-name root-dir h|v percentage
select-pane pane-index
send-keys session-name window-index command
attach-session session-name
```

Here's an example configuration file:

```
new-session dbot /home/dan/Workspace/DBot
split-pane dbot /home/dan/Workspace/DBot v 10
select-pane 3
split-pane dbot /home/dan/Workspace/DBot h 50

select-pane 0
send-keys dbot 0 nvim

select-pane 1
send-keys dbot 0 cat .env

select-pane 2
send-keys dbot 0 go run .

select-pane 0
attach-session dbot
```

The above configuration file creates a new tmux session named "my-session" with a horizontal split pane that takes up 50% of the window. It then sends the commands "echo 'Hello, world!'" and "ls -la" to the first and second panes, respectively.

You can find more information on the different commands in the source code.
