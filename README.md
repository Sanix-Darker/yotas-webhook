# Yotas-WebHook

This is the yotas webhook github app.

## Purpose

This app will watch PR events on a project and then do some interesting stuffs by communicating with the `yotas API`.

## Requirements

- Golang

## How to install

- Copy `.env.example` to `.env` and then configure the secret parameter.

- You need to install `go-github`, it's a webHook parser :
```
go get -u github.com/google/go-github/github
```
then you're all set


## How to start

There are two things when running it locally, you need to run it but also make it available to the internet !
For this, i used localhost.run with ssh !
- You just have to run : `make online`
- Then grap the generated link, add it /webhook and configure your githubApp with it
- To start the app itself, you can run `main run`


## Example of output

```
make run
go run main.go

2021/04/24 23:40:26 [+] Yotas-WebHook server started

# When a PR is created we got those informations :
2021/04/24 23:44:52 -----------------------------------------------------------
2021/04/24 23:44:52 [-] dev : 'Sanix-Darker'
2021/04/24 23:44:52 [-] project : 'osscameroon/cacho'
2021/04/24 23:44:52 [-] title : 'PR's title'
2021/04/24 23:44:52 [-] body : 'PR's body example !'
2021/04/24 23:44:52 [-] url : 'https://api.github.com/repos/osscameroon/cacho/pulls/10'
2021/04/24 23:44:52 -----------------------------------------------------------
2021/04/24 23:44:52 ===========================================================
2021/04/24 23:44:52    - PR-ID: 622659851
2021/04/24 23:44:52    - commits: 2
2021/04/24 23:44:52    - additions: 2
2021/04/24 23:44:52    - deletions: 1
2021/04/24 23:44:52    - changedfiles: 2
2021/04/24 23:44:52 ===========================================================

# When the PR got merged :
2021/04/24 23:44:52 [-] url : 'https://api.github.com/repos/osscameroon/cacho/pulls/10'
2021/04/24 23:44:52    - PR-ID: 622659851
2021/04/24 23:47:07    - Merged: true
2021/04/24 23:47:07    - Merged By: Sanix-Darker
2021/04/24 23:47:07 ===========================================================
2021/04/24 23:47:07
```

## Author

- [Sanix-darker](https://github.com/sanix-darker)

