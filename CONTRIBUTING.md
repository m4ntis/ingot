# Contributing to ingot

Thanks for contributing!

## Development

You can view and edit the source code by cloning this repository:

```bash
git clone https://github.com/ingotmc/ingot.git
```

## Pull Requests

### How to Send Pull Requests

Everyone is welcome to contribute code to `ingot` via
GitHub pull requests (PRs).

To create a new PR, fork the project in GitHub and clone the upstream
repo:

```sh
$ git clone https://github.com/ingotmc/ingot.git
```

This would put the project in the `ingot` directory in
current working directory.

Enter the newly created directory and add your fork as a new remote:

```sh
$ git remote add <YOUR_FORK> git@github.com:<YOUR_GITHUB_USERNAME>/ingot
```

Check out a new branch, make modifications, run linters and tests, and
push the branch to your fork:

```sh
$ git checkout -b <YOUR_BRANCH_NAME>
# edit files
$ go test -v ./...
$ git add <FILES>
$ git commit -m "<commit message>"
$ git push <YOUR_FORK> <YOUR_BRANCH_NAME>
```

Open a pull request against the main `ingotmc/ingot` repo.

### How to Receive Comments

* If the PR is not ready for review, please put `[WIP]` in the title,
  tag it as `work-in-progress`, or mark it as
  [`draft`](https://github.blog/2019-02-14-introducing-draft-pull-requests/).

* Need to address comments/feedback on your PR? Be sure to amend your commits as this keeps a clean, traceable repo! (example below)
  
  ```sh
  # edit files that address comments/feedback on PR
  $ git add <MODIFIED/AMENDED FILES>
  $ git commit --amend
  $ git push <YOUR_FORK> <YOUR_BRANCH_NAME> -f
  ```

* Worse comes to worse, commits can be squashed before merge

### How to Get PRs Merged

A PR is considered to be **ready to merge** when:

* It has received two approvals from Approvers/Maintainers.
* Major feedbacks are resolved.
* Consider rebasing your branch on master if your branch falls behind origin/master during approval process. Example:

  ```sh
  # rebase on origin/master in case any new features have been merged to origin/master
  $ git checkout master
  $ git pull
  $ git checkout <YOUR_BRANCH_NAME>
  $ git rebase master
  # this may cause conflicts so be sure to resolve before pushing
  $ git push <YOUR_FORK> <YOUR_BRANCH_NAME>
  ```

Any Approver/Maintainer can merge the PR once it is **ready to
merge**.

## Maintainers

- [guglicap](https://github.com/guglicap)

## Approvers

- [dcrodman](https://github.com/dcrodman)
- [ethanent](https://github.com/ethanent)
- [fuchstim](https://github.com/fuchstim)
- [liad344](https://github.com/liad344)
- [m4ntis](https://github.com/m4ntis)
- [qaisjp](https://github.com/qaisjp)
- [rizkybiz](https://github.com/rizkybiz)
- [wilfrid-askins](https://github.com/wilfrid-askins)