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

Microsoft open source projects website : https://opensource.microsoft.com <br />

**_Go programming language projects_** <br />
GoBridge : https://github.com/gobridge <br />

**_Go programming language related articles_** <br />
Why did Google develop Go? by Quora : https://www.quora.com/Why-did-Google-develop-Go <br />

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

Then exit the vim text editor by typing pressing **[ Esc ]** on your keyboard, and then press **[ Shift ]** + **[ : ]**, then press **[ W ]** and then press **[ Q ]** on your keyboard, finally press **[ Enter ]** on your keyboard to write the command you added into the file.

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

Then press **[ 3 ]** or **[ PgDn ]** on your keyboard until you reach to the end of the file, and press **[ I ]** to edit the `.bashrc` file, 


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
