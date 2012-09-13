# local go, lo(cal)go, logo

Restores sanity to go's package management for applications. Use it just like you use the `go` command and logo will sets GOPATH to the current directory. This allows you to install and upgrade go packages just within your application.

Heavily inspired by [Dave Cheney's gogo](https://github.com/davecheney/gogo), but with a minimalist bent.

## Installation

`go get github.com/corasaurus-hex/logo`

## Use

The command to use logo is `logo`. `logo` can replace any go command, so use it just like you would use `go`.

Instead of the default `go get` which would install in your global GOPATH:

`go get github.com/corasaurus-hex/someproject`

You would use `logo get` which will install it in your current directory:

`logo get github.com/corasaurus-hex/someproject`

## How it works

1. Set GOPATH to the current directory
2. Execute `go` with the arguments passed to `logo`

That's it.

## Assumptions

logo assumes you want to install all your packages for a project in your project directory. logo also assumes that you want complete isolation from the world.

## An Explanation

I'm going to make some generalizations here, most of which can be addressed by "just do X which isn't explicitly documented", "just manually manage GOPATH all over the place and have all your teammates do the same", or "I don't think that, but I guess the majority of the Go community does". Please keep this in mind.

The Go community's stance on packages is that you should have one location  where you keep all your packages and application source code. While technically you can have more than one location on the system (multiple entries in GOPATH), generally there is one location and your applications and packages should be stored in the same `src/` tree, commingled.

This means that all third party packages should be shared across the system and all applications should require the same version of any given package. The implication is that if one of your applications needs an updated package then all other applications must be simultaneously upgraded. Any collaborators on your application are then required to upgrade their packages as well, and if they have any applications that depend on the same package then they need to upgrade those applications as well. This is a monumental waste of effort when any given application may not need the upgrade. And this all assumes that the upgrade will be compatible across all the applications for a developer, which may not be the case. They may be intentionally staying at an older version because of semantic changes or features being removed.

A side-effect of this is that there is no blessed way to specify the versioned dependencies of your application. You must create ad-hoc methods to convey these dependencies instead of having a way to enforce them.

There are two hacks that I know of that work around this. The first is that you can create your own repository for a package at the desired version and install that globally, importing that into all your applications. The second is that you can have a different directory from your global go path where you install packages and application code. This is done via manipulation of the GOPATH environment variable.

The first hack is painful in spurts. Every time you need to change the version of a package where there is incompatibility between versions you need to create a repository and update your application. Repository management isn't terribly difficult, but you now have a new repository that isn't strictly necessary which you'll need to eventually remove, if you remember to and can track down who is using it. The bonus is that you don't need to remember when to prefix `go` with `GOPATH=/path/to/foo`.

The second hack is painful all the time. Every go command you execute that needs GOPATH needs to be prefixed with your custom GOPATH: `GOPATH=/path/to/foo go install ...`. You don't need to do any custom repository management and you can bundle all of your package dependencies in your application and keep them in version control. This also solves the problem of specifying dependencies and the correct versions of the dependencies by keeping the code of your dependencies in source control.

Either of these solutions would work much better if they were automated. The first is difficult to automate, the second is easy.

logo automates the manipulation of GOPATH to run go commands with the current directory as the GOPATH. That's it.
