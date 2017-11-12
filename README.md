go-gopathless
=============

**gopathless** demonstrates:

- a project-local $GOPATH
- with no fuss.

That's it.


What?  Really?
--------------

Yeah.  This is way easy.

Vendor is not used.  Your gopath is just inside your project source dir.

Clone this repo wherever.  It'll work.

Try it:  `./go test ./...`, `./go install ./...`, and so on.  Even `./go list -f '{{join .Deps "\n"}}'`.


Does it work with $ECOSYSTEM?
-----------------------------

Yes.


Really??
--------

[*yes.*](./README_ECOSYSTEM.md)


How is it so simple??!
----------------------

There's no mystery; there's no secret sauce.

It's just $GOPATH... set to $PWD.

Read [the fine source](./go).


Why should I care?
------------------

Project-local $GOPATH saves you from
[many](https://github.com/polydawn/golink#why)
[serious](http://0x74696d.com/posts/go-get-considered-harmful/#convention-doesn-t-isolate-workspaces)
[issues](http://grokbase.com/t/gg/golang-nuts/135dvdrrrg/go-nuts-workspaces-and-versioned-dependencies),
[answers questions](https://groups.google.com/forum/?_escaped_fragment_=topic/golang-nuts/B1uvwHMclkE#!topic/golang-nuts/B1uvwHMclkE),
and [other major projects](https://github.com/git-lfs/git-lfs/pull/458) are
already adopting project-local GOPATH defaults as a key way to minimize barrier to entry for new developers.

Project-local $GOPATHs are the way forward.

- You *should* know what your dependencies are.  Global GOPATH obscures this.
- You *should* be able to build two different projects on your host which use
  different versions of a 3rd party library.  Global GOPATH prevents this;
  vendoring permits it, but is more complicated than simply having
  project-local GOPATH in the first place.
- You *should* be able to develop Library-A and Product-B which depends on it...
  while having your Product-B work off a stable release of Library-A.
  Again, global GOPATH makes this a landmine scenario, because you may have
  a seemingly-working state one day, but no version control over it;
  vendoring makes work possible, but it's still error-prone because you may
  *forget* to vendor properly, or be tempted to just skip it; and a
  project-local GOPATH avoids all of these issues by being clear and simple.

You like *clear* and *simple* right?

That's why you should care.

Clear.  Simple.  $GOPATH.  $PWD.  Done.


Wow
---

I know, right?

Also check out [the Part 2 repo](https://github.com/warpfork/go-gopathless-pt2) to see
how a repo like this behaves as a library for someone else (spoiler: it's fine).
