# Installing version manager Gobrew on Ubuntu

First Step
```
curl -sLk https://git.io/gobrew | sh -
```

Second Step
```
export PATH="$HOME/.gobrew/current/bin:$HOME/.gobrew/bin:$PATH"
```

Testing if it's working
```
gobrew ls-remote
```

You will need set this config in your bash or zshrc file
```
#GOBREW
export PATH="$HOME/.gobrew/current/bin:$HOME/.gobrew/bin:$PATH"
export GOROOT="/home/leonrd/.gobrew/versions/1.18/go"
```

The GOROOT you can get with this command
```
go env GOROOT
```

Commands
- `gobrew ls-remote` # list availables go versions
- `gobrew use 1.16` # to use a specific version
- `gobrew ls` # list current and installed versions