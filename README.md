kirk
====

![](http://www.startrek.com/legacy_media/images/200307/kirk01/320x240.jpg)

### Be the captain of your enterprise

#### What is kirk?

kirk is a command line utility for deploying scripts to your servers/enterprise. It is still under development and has a long way to go.

#### Install

`go get github.com/acmacalister/kirk`

`go install github.com/acmacalister/kirk`

#### Usage

Currently kirk has two commands:

`kirk version` to get the current version of kirk.

`kirk deploy [environment]` To run the shell scripts you put in your kirk.yml file.

With the environment being whatever you put in your kirk.yml to run. Kirk looks for a kirk.yml file in the current directory it is being run from. Take a look at the kirk.yml for an example of how it should look.

#### License

Kirk is license under the Apache License.

#### Authors

* [acmacalister](http://twitter.com/acmacalister)
* you

#### ToDo

- [] Move into a library with binary option?
- [] Build releases for users to test with.
- [] Support Pemfiles for auth.
- [] Need a better way to handle reading kirk files.
- [] Add Unit Tests.
- [] Better docs.

If you would like to help or have a feature suggestion, feel free to open an issue.