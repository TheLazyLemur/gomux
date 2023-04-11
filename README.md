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
new-session my-session /path/to/my/project
split-pane my-session /path/to/my/project h 50
send-keys my-session 0 "echo 'Hello, world!'"
select-pane 1
send-keys my-session 1 "ls -la"
```

The above configuration file creates a new tmux session named "my-session" with a horizontal split pane that takes up 50% of the window. It then sends the commands "echo 'Hello, world!'" and "ls -la" to the first and second panes, respectively.

You can find more information on the different commands in the source code.
