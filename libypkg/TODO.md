# TODO

- [x] Create internal.Package from v3.Package
    - [x] Add a function for Bumping
    - [x] Add a function for creating a Default package
    - [x] Add a function for auto-creating a package
    - [x] Add a function for updating Sources
- [x] Create a new PackageSpec interface
    - [x] Move Lint() to internal.Package
    - [x] Add a function for converting from internal.Package to a versions format
- [x] Create a generic function for loading any PackageSpec
- [x] Implement conversions for the different PackageSpecs
- [x] Rename `yml` package to `spec`

## CLI

- [ ] ypkg auto
    Given a list of sources:
    - [x] Fail if package.yml exists
    - [x] Create a Default Package
    - [ ] Automatically add Sources to the Default package
    - [ ] Scan first source, directory name, etc. to fill out package fields
    - [x] Convert internal.Package to the current version of the ypkg spec (v2)
    - [x] Write out a new package.yml
- [x] ypkg bump
    Given an existing package.yml:
    - [x] Fail if package.yml does not exist
    - [x] Load the package.yml
    - [x] Convert it to internal.Package
    - [x] Bump the internal.Package
    - [x] Convert it back to the original version of the ypkg spec
    - [x] Write out the updated package.yml
- [x] ypkg convert
    Given an existing package.yml:
    - [x] Fail if package.yml does not exist
    - [x] Load the package.yml
    - [x] Convert it to internal.Package
    - [x] Bump the internal.Package
    - [x] Convert it to the specified version of the ypkg spec
    - [x] Write out the updated package.yml
- [x] ypkg init
    - [x] Fail if package.yml exists
    - [x] Create a Default Package
    - [x] Convert internal.Package to the current version of the ypkg spec (v2)
    - [x] Write out a new package.yml
- [ ] ypkg lint
    Given an existing package.yml:
    - [x] Fail if package.yml does not exist
    - [x] Load the package.yml
    - [x] Convert it to internal.Package
    - [ ] Lint() the internal.Package
- [ ] ypkg update
    Given a list of sources and an existing package.yml:
    - [x] Fail if package.yml does not exist
    - [x] Load the package.yml
    - [x] Convert it to internal.Package
    - [x] Bump the internal.Package
    - [ ] Update the internal.Package Sources using the list
    - [x] Convert it to the current version of the ypkg spec
    - [x] Write out the update package.yml
