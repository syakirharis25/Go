# Go
My works related to Go programming language.

## Table of Contents
1. [Introduction.](#introduction)
2. [Official references websites.](#references)
3. [Communicating sequential processes.](#CSP)
4. [Installing Go programming language for Ubuntu 18.04 LTS manual.](#ubuntugo) 
5. [Go programming jobs.](#jobs)
6. [GitHub notes.](#github)

<a name="introduction"></a>
## 1. Introduction.
<img src="Golang.png" height="150"> 
Go, also known as Golang, is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

<a name="references"></a>
## 2. Official references websites. <br />
Go programming language official website : https://golang.org <br />
Go programming language documentation : https://golang.org/doc/ <br />
Gopher Academy Blog : https://blog.gopheracademy.com/gophers-slack-community/
Go programming language playground : https://play.golang.org/

**_ Go programming questions and answers_**
Stack Overflow questions and answers website : https://stackoverflow.com <br />
Quora questions and answers website : https://www.quora.com/ <br />

**_Go programming language documentation by golang.org_** <br />
Go Package fmt by golang.org : https://golang.org/pkg/fmt/ <br />
The Go Programming Language Specification by golang.org : https://golang.org/ref/spec

**_Go programming language text editor_** <br />
Visual Studio Code text editor : https://code.visualstudio.com <br />

**_Go programming language related articles_** <br />
Why did Google develop Go? by Quora : https://www.quora.com/Why-did-Google-develop-Go <br />
Go in Visual Studio Code : https://code.visualstudio.com/docs/languages/go <br />
Q: How do you Display Chinese characters in command prompt by Stack Overflow: https://stackoverflow.com/questions/39551549/q-how-do-you-display-chinese-characters-in-command-prompt <br />
Go: Meaning of the 'fmt' package acronym : https://stackoverflow.com/questions/23597165/go-meaning-of-the-fmt-package-acronym <br />
Whats the difference between uint and int in golang? by Quora : https://www.quora.com/Whats-the-difference-between-uint-and-int-in-golang <br />
golang is slow in Windows [closed] by Quora : https://stackoverflow.com/questions/32062493/golang-is-slow-in-windows <br />
Constants by Rob Pike : https://blog.golang.org/constants <br />
Loops in Go Language by GeeksforGeeks : https://www.geeksforgeeks.org/loops-in-go-language/ <br />
Difference between := and = operators in Go by Stack Overflow : https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go <br />
Difference between fmt.Println() and println() in Go by Stack Overflow : https://stackoverflow.com/questions/14680255/difference-between-fmt-println-and-println-in-go <br />
Functions by golang-book.com : https://www.golang-book.com/books/intro/7
Package math by golang.org : https://golang.org/pkg/math/ <br />

**_Go programming language projects_** <br />
GoBridge : https://github.com/gobridge <br />
Microsoft open source projects website : https://opensource.microsoft.com <br />

**_Go programming language developers in GitHub_** <br />
Robert Griesemer : https://github.com/griesemer <br />
Rob Pike : https://github.com/robpike <br />
Ken Thompson : https://github.com/ken <br />
Russ Cox : https://github.com/rsc <br />
Brandon Minnick : https://github.com/brminnick <br />
Francesc Campoy : https://github.com/campoy <br />

<a name="CSP"></a>
## 3. Communicating sequential processes.
In computer science, communicating sequential processes (CSP) is a formal language for describing patterns of interaction in concurrent systems. It is a member of the family of mathematical theories of concurrency known as process algebras, or process calculi, based on message passing via channels. CSP was highly influential in the design of the occam programming language, and also influenced the design of programming languages such as Limbo, RaftLib, Go, Crystal, and Clojure's core.async.

CSP was first described in a 1978 paper by Tony Hoare, but has since evolved substantially. CSP has been practically applied in industry as a tool for specifying and verifying the concurrent aspects of a variety of different systems, such as the T9000 Transputer, as well as a secure ecommerce system. The theory of CSP itself is also still the subject of active research, including work to increase its range of practical applicability (e.g., increasing the scale of the systems that can be tractably analyzed).

<a name="ubuntugo"></a>
## 4. Installing Go programming language for Ubuntu 18.04 LTS manual.
For Ubuntu 18.04 LTS user, copy and paste this address `https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz` on Mozilla Firefox browser and hit **[ Enter ]** on your keyboard. Select `Save File` when the menu appears asking `What should Firefox do with this file?`, and **[ Right Click ]** the `OK` button on the bottom right of the menu shown. Then open the terminal by pressing **[ Ctrl ]** + **[ Alt ]** + **[ T ]**.

Follow the below commands to properly setup Go programming language environment in the local machine, this example is using `go1.13.8.linux-amd64.tar.gz` Go file.
```
$ cd Downloads
$ ls
$ sudo tar -C /usr/local -xzf go1.13.8.linux-amd64.tar.gz
$ /usr/local/go/bin/go version
$ sudo apt install vim
$ sudo vim ~/.profile
```

Then add the following contents, the global variable, on the bottom of the `.profile` file, press **[ 3 ]** or **[ PgDn ]** until you reach to the bottom of the file, then pressing **[ I ]** on your keyboard, copy and paste the below command on the end of the file.
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

Then exit the vim text editor by typing pressing **[ Esc ]** on your keyboard, and then press **[ Shift ]** + **[ : ]**, then press **[ W ]** and then press **[ Q ]** on your keyboard, finally press **[ Enter ]** on your keyboard to write the command you added into the `.profile` file.

To review back what you already edited in home `.profile` file.
```
$ cat ~/.profile
```

Then update the current home shell session.
```
$ source ~/.profile
```

Then check your installed Go version in your local machine.
```
$ go version
```

To check your Go tools in the local machine.
```
$ cd /usr/local/go
$ ls
```

Then make some changes in your home `.bashrc` file
```
$ cd ~
$ vim .bashrc
```

Then press **[ 3 ]** or **[ PgDn ]** on your keyboard until you reach to the end of the file, and then press **[ -> ]** and **[ -> ]** again, then press **[ ↓ ]** and press **[ ↓ ]** again, then press **[ I ]** to edit the `.bashrc` file, but before setting the GOROOT, read the post by `Dave` on the effect of modifying the file if you are working with multiple version of Go in your system, on his blog : https://dave.cheney.net/2013/06/14/you-dont-need-to-set-goroot-really, and then add the following text into the bottom of the `.bashrc` file.
```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

Then press **[ Esc ]**, then press **[ Shift ]** + **[ : ]**, then press **[ W ]** and then press **[ Q ]** on your keyboard, finally press **[ Enter ]** on your keyboard to write the command you added into the `.bashrc` file.

Review back what you already edited in home `.bashrc` file.
```
$ cat ~/.bashrc
```

Then update the current home shell session.
```
$ source ~/.bashrc
```

Then check your installed Go version in your local machine.
```
$ go version
```

Then go back into your `.bashrc` file to add the GOPATH to set the environment properly, to tell Go where your working files are located. 
```
$ vim .bashrc
```

Then press **[ 3 ]** or **[ PgDn ]** on your keyboard until you reach to the end of the file, and then keep on press **[ -> ]** and **[ -> ]** again until it is end of the line, then press **[ ↓ ]** and press **[ ↓ ]** again, then press **[ I ]** to edit the `.bashrc` file, and then press **[ Enter ]** and  **[ Enter ]** again, and then add the following text into your `.bashrc` file.
```
export GOPATH=/home/username/golib
export PATH=$PATH:GOPATH/bin
```
Then press **[ Esc ]**, then press **[ Shift ]** + **[ : ]**, then press **[ W ]** and then press **[ Q ]** on your keyboard, finally press **[ Enter ]** on your keyboard to write the command you added into the `.bashrc` file.

Review back what you already edited in home `.bashrc` file.
```
$ cat ~/.bashrc
```

Then update the current home shell session.
```
$ source ~/.bashrc
```

Then go back into your `.bashrc` file to add the GOPATH to set the environment properly, to tell Go where your next working files are located. 
```
$ vim .bashrc
```

Then make a new folder called `golib`
```
$ mkdir golib
```

Then check the content inside the `golib`
```
$ cd golib
$ ls
```

Then go back to your home directory.
```
$ cd ..
```

Then install `git` into Ubuntu machine. 
```
$ sudo apt-get install git
```
Enter your password when it asking for it, press **[ Y ]** and then **[ Enter ]** on your keyboard when it asking to continue.

Check whether the content from GitHub was downloaded into `golib` folder.
```
$ go get github.com/nsf/gocode
$ cd golib
$ ls
$ cd bin 
$ ls
$ cd ../src
$ ls
$ cd github.com/nsf/
$ ls
$ cd gocode
$ ls
$ cd ~
```

Then go back into your `.bashrc` file to add the GOPATH to set the environment properly, to tell Go where your next working files are located. 
```
$ vim .bashrc
```

Then press **[ 3 ]** or **[ PgDn ]** on your keyboard until you reach to the end of the file, and then keep on press **[ -> ]** and **[ -> ]** again until it is end of the line, then press **[ ↓ ]**, then press **[ I ]** to edit the `.bashrc` file, and then press **[ Enter ]**, and then add the following text into your `.bashrc` file. 
```
export GOPATH=/home/username/code
```
This will act as your working space location. Then press **[ Esc ]**, then press **[ Shift ]** + **[ : ]**, then press **[ W ]** and then press **[ Q ]** on your keyboard, finally press **[ Enter ]** on your keyboard to write the command you added into the `.bashrc` file.

Review back what you already edited in home `.bashrc` file.
```
$ cat ~/.bashrc
```

Then update the current home shell session.
```
$ source ~/.bashrc
```
Then make a new folder called `code`
```
$ mkdir code
```

Then check the content inside the `code`
```
$ cd code
$ ls
```

Then go back to your home directory.
```
$ cd ..
```

Then remove everything from golib.
```
$ cd golib
$ rm -rf *
$ ls
```

Check whether the content from GitHub was downloaded into `golib` folder.
```
$ cd ..
$ go get github.com/nsf/gocode
$ cd golib
$ ls
```

Check whether the content from GitHub was downloaded into `code` folder.
```
$ cd ../code
$ ls
```

In order to have Go workspace, you need to have `src` folder in it
```
$ mkdir src
$ ls
```

The other 2 folders that you need is `bin` and `pkg` folders
```
$ mkdir bin
$ mkdir pkg
$ ls
```

<a name="jobs"></a>
## 5. Go programming jobs.
Go / Golang Jobs & Developers by golangprojects : https://www.golangprojects.com <br />
Go / Golang Jobs & Developers by golangprojects by Totaljobs : https://www.totaljobs.com/jobs/go-programmer <br />
Go / Golang Jobs & Developers by golangprojects by SimplyHired : https://www.simplyhired.com/search?q=go+programming&job=yzD7SgAyVlbXVk56Z1e7EUJnm1CSOvVfNNgGES9xpACHOiI5Y8bUBw <br />

<a name="github"></a>
## 6. GitHub notes.
Clone the current GitHub remote repository contents into local machine.
```
$ git clone https://github.com/syakirharis25/Go.git
$ cd Go/
$ git remote -v
$ git status
