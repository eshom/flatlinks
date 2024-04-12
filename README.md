# What Is This About
This is a small command line tool when you want to create links to all files in some directory recursively.
Useful when you want to explore all files in one directory without having to walk through the directory tree manually.
I use it when I backup media from my mobile phone and want to quickly see what videos or photos were backed up.

*This is very alpha. Use at your own risk.*
Tested on Linux only, but might work on Windows. Let me know if you try.

You'd probably better off just running a `find` one-liner that does the same thing with more control
```bash
cd <where-to-generate-links>
find <dir> -type f -exec ln -s {} \;
```

# Install
Build executable:
```bash
go build -ldflags="-s"
```

Install in $GOPATH (if exists) or $HOME/bin:
```bash
go install -ldflags="-s"
```

# How to use
First argument is the directory name to walk.
Second argument is where to generate the links.

```bash
flatlinks <directory-with-files> <target-directory>
```

- Max directory name is 100 characters.
- If you need more than that for your directories you probably have a problem :-)
- Already existing files will be skipped.
- Both absolute and relative paths are supported, but this was not tested thoroughly.
