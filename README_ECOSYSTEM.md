going gopathless -- with other tools
====================================

A project-local $GOPATH works absolutely fine with almost all tools.



Go build, install, test, etc
----------------------------

*It's just $GOPATH.*  Of *course* all the core tools work.



Go get
------

No, seriously.  *It's just $GOPATH.*  `go get` will work fine.



Other people go-getting you
---------------------------

It's fine.

The self-symlink in your project-local gopath means all your project's
code uses the full public URL for your project, just like any outside
code would when importing your project.  It just works.  No fuss.

[Here's another repo to prove it](https://github.com/warpfork/go-gopathless-pt2).  Try it!



Dependency managers
-------------------

Dependency managers have *long* been a bugaboo in the Go community.

Fortunately, a project-local $GOPATH doesn't make this more complicated.

All dependency managers will work with your project local $GOPATH.
Just make sure $GOPATH is exported before you run them.

### Dependency managers with vendor

Setting a project-local $GOPATH does not change the behavior of "vendor" at all
(nor did it ever depend on "vendor").

The behavior is exactly the same on all versions of Go, and exactly the same
whether or not you or your other tools use "vendor".



Vendoring without Vendor
------------------------

You can commit the content inside your project-local $GOPATH into git.  It's fine.

This is effectively "vendor" behavior, but without a directory called "vendor".

It turns out you don't need "vendor" and you don't need to learn any new rules;
you just commit the things you want committed, and they're on your $GOPATH,
and they act normal, and it's fine.

Bonus: unlike the "vendor" behavior, you can never *accidentally* depend on your
global $GOPATH, because you don't have one.  It turns out simpler is better sometimes.



Git submodules?!
----------------

You can use git submodules in your project-local $GOPATH if you want.  It's fine.

You don't have to.  It's just $GOPATH.  But it turns out having other git repos in
your $GOPATH is fine, and having git repos tracked by a parent git repo is called
a submodule, and using submodules gives you total control of versions, recursively,
which is... pretty fine.  Maybe you should try it.  It even automatically shows you
diffs, makes it easy to fetch updates, and can easily pin any source version you want.

Wow.



Editors and IDEs
----------------

It's fine.

- VScode [supports configuring a gopath per project with a simple json config file](https://github.com/Microsoft/vscode-go/wiki/GOPATH-in-the-VS-Code-Go-extension#gopath-from-gogopath-setting)
- Intellij golang plugins [support per-project gopath](https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/867)
- Atom go-plus plugins [support gopath from env](https://github.com/joefitzgerald/go-plus/issues/405)

In general: you know the refrain: *it's just $GOPATH*.
The hartest this can possibly be is if if your editor doesn't have explicit support for per-project gopath,
then you just launch your editor with the GOPATH env var set; done!



It's just GOPATH
----------------

The true joy of "gopathless" project-local gopath is that it's just $GOPATH.

It doesn't need buy-in and it doesn't need ecosystem support
because *it's already universal*.
$GOPATH has always been here,
$GOPATH will always be here,
and we're just using $GOPATH.

It's fine.
