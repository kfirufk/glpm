# glpm
Go Lazy Package Manager

#install
`go get -u github.com/kfirufk/glpm`

# why
welp.. in general you can type `go get ./...` on your root project's directory and it will install all the required dependencies, 
but will complain about your own project directories with `go install: no install location for directory`. I wanted to have one json
file that includes all of my dependencies, and I wanted a nice method to manage it. so I wrote this tool.

# how to use ?
simply type `glpm -h` to see what options are available.

# glmp.json
for now it just includes a `packages` array that includes the names of the used packages. it doesn't include package versions. I wanted to make a simple and clean packages list, so for now I don't see the point of adding and depending on specific versions of packages.