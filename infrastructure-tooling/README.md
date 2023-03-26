# Infrastructure Tooling

## Setting up infrastructure tooling

Clone `infrastructure-tooling` into the same build environment as your `common` folder:

```
git clone ssh://vcs@dev.getsol.us:2222/source/infrastructure-tooling.git
```

Setup SSH to make committing to Phab work. Example `~/.ssh/config` entry:

```
Host dev.getsol.us
    User vcs
    PreferredAuthentications publickey
    Port 2222
    IdentityFile ~/.ssh/private/solus-phab
```

Setup SSH to make pushing to the build server work. Example `~/.ssh/config` entry:

```
Host getsol.us
    User build
    Port 798
    PreferredAuthentications publickey
    IdentityFile ~/.ssh/private/solus-build
```

## Working with Differential

Every `Diff` follows a similar set of stages:

1. Review
2. Revision (Sometimes Optional)
3. Acceptance
4. Landing
5. Publishing

### Request Changes

While reviewing a `Diff`, you can request a change by choosing the Action:

- Request Changes

Then, add information about what you would like to be fixed and click the `Submit` button.

When the author has made revisions, they will use the Action `Request Review` to restart the Review process.

### Accept Changes

Is everything correct with the `Diff`, you can mark it as reading to be Landed by choosing the action: 

- Accept Revision

**NOTE:**

> Should be there more than one reviewer, all reviewers must accept the patch before it can be landed. In some cases, the other reviewer(s) may be removed from the diff before landing it. Removing an additional reviewer can be done via the `Edit` button or from the action:
>
> -  `Change Reviewers`

## Landing and Publishing a Diff

### Merge the Diff

First, you need to clone the repository:

```
make REPONAME.clone
```

If you already have a local copy, update it:

```
make pull
```

To apply the Diff to the local copy use Arcanist:

```
arc patch Dxxxxx
```


**NOTE:**

>It is always suggested to run a build locally first before merging and pushing it to the build server to ensure it will build.

### Landing the Diff

Once you are ready to merge in the Diff and push those changes back up to Phabricator, use Arcanist again to land the Diff:

```
arc land
```

**NOTE:**

> If the Diff has not been accepted by one or more reviewers, `arc land` will complain, and ask if you'd like to landed or not. You can force land it by accepting it as is, but we recommend that you either remove reviewers or go back and accept the patch as requested.


### Publishing to the Build Server (aka Pushing)

After you landed a patch with `arc land` or modified a package yourself (without a Diff) you can push it to the build server:

**NOTE:**
> If you directly push a package, you need to add the changelog inside the commit message as you would have done with a diff

```
make publish
```

This will create a new tag for this release and then push the tag and new commits to the Phab repo. Afterwards, a Job will be submitted to the build server.

**NOTE:**

> If your repo is out of sync with the repo on Phab, this step will fail.

If the Job cannot be sent to the build server or needs to be retried, the job can be resubmitted by running:

```
make republish
```

**NOTE:**

> If the build fails on the build server and any changes to the build recipe (package.yml) itself are needed to ensure it will build a new tag will need to be created and `make republish` will be insufficient.

> Make sure you are on the git **master branch** and not inside the **arcpatch-Dxxxxx** branch. You will piss off the dragon and building will fail. The package has to then be bumped because the old tag is no longer useful.

## Working with Repos

### Recommended Aliases

The following Bash aliases are useful for managing repos:

```
alias create-repo='${PATHBUILDDIR}/infrastructure-tooling/create-repo' 
alias init-repo='${PATHBUILDDIR}/infrastructure-tooling/init-repo' 
alias new-repo-from-differential='${PATHBUILDDIR}/infrastructure-tooling/new-repo-from-differential' 
```

### Creating a new Repo with a DIFF

If you need to create a new repo, you can simply run:

```
create-repo REPONAME
make REPONAME.clone
cd REPONAME
```

You may choose to initialize the new repo in order to apply a Diff to it:

```
init-repo
```

You can shortcut this whole process by running:

```
new-repo-from-differential REPONAME Dxxxxxx`
```

**NOTE:**

> You will still need to `arc land` and `make publish` as usual.

### Creating a new Repo without a DIFF

Run the commands above till `init-repo` after you this you can copy the needed files of the new package into your empty folder but the `.git` folder and land it as any other package directly to the buildserver.

# Building a Test-ISO

Choose the relevant image (`budgie`, `gnome`, `plasma`, `mate`) and substitute `TEST_ISO` below to begin building the ISO.

```bash
    export TEST_ISO="budgie"

    # Ensure you have the correct ISO build dependencies
    sudo eopkg it syslinux libisoburn squashfs-tools sbsigntools

    # Clone an ISO repo
    git clone https://dev.getsol.us/source/solus-image-$TEST_ISO.git
    pushd solus-image-$TEST_ISO

    # Build the ISO. You must have the `common/` tree set up per the Packaging introduction.
    make
```
