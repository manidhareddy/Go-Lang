# Go Installation and Environment Setup

---

## After Installing

```bash
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.25.3.linux-amd64.tar.gz
```

This will extract Go to the following directory:

```
/usr/local/go/
```

---

## Setting Up Environment Variables

### GoRoot

`GOROOT` can sometimes cause problems when working with different versions of Go on your system.

For more information:  
🔗 [You don’t need to set GOROOT — really](https://dave.cheney.net/2013/06/14/you-dont-need-to-set-goroot-really)

---

## 1️⃣ First Environment Variable Setup (Go Installation Path)

We need to tell the environment where Go is installed and where to find its binaries.

Edit your `~/.bashrc` file and add:

```bash
# Tell the environment where to find the Go binaries
export GOROOT=/usr/local/go

# Set PATH variable to include Go binaries
export PATH=$PATH:$GOROOT/bin
```

Then apply the changes:

```bash
source ~/.bashrc
```

Verify the installation:

```bash
go version
```

---

## 2️⃣ Setting Up GOPATH

- We often download packages and libraries published by others to build our own Go applications.
- `GOPATH` specifies where Go projects and libraries are located.

Add this to your `~/.bashrc`:

```bash
export GOPATH=/home/manidhar/golib
export PATH=$PATH:$GOPATH/bin
```

Apply the changes:

```bash
source ~/.bashrc
```

Create the Go workspace directory:

```bash
mkdir ~/golib
```

---

### Example: Installing a Common Go Library

Install a library that provides autocompletion functionality in Go applications:

```bash
go get github.com/nsf/gocode
```

This will be stored inside your `GOPATH`.  
Go maintains all downloaded libraries in a **single monolithic repository**, simplifying dependency management.

---

## 3️⃣ Workspace Setup

If you want to add an additional workspace for your code:

```bash
export GOPATH=$GOPATH:/home/manidhar/code
source ~/.bashrc
```

Now create the standard Go workspace structure:

```bash
mkdir -p ~/code/{src,bin,pkg}
```

### Folder Explanation
- **src/** → Holds your Go source code and all internal packages created for this modules.
- **bin/** → Stores compiled binaries for your projects.  
- **pkg/** → Contains intermediate compiled packages (not final executables).  

> The `pkg` directory helps speed up builds. When you recompile, Go checks if any source files have changed — if not, it skips recompiling and reuses the existing binaries.

---

- **go run** Compile temporarily and  run that for me. I also compile any third party libraries. // it will take main file path
- **go build** It will take actual package path. It finds main package with main function that it's going to compile that as an executable.
- **go install** genrates executable in bin folder.


✅ **You’re now ready to start developing with Go!**

