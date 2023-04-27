"use strict";(self.webpackChunkcentauri_documentation=self.webpackChunkcentauri_documentation||[]).push([[237],{9754:(e,t,n)=>{n.r(t),n.d(t,{default:()=>o});var l=n(3612),a=n(814),r=n(7961),i=n(7294);function o(){return i.createElement(r.Z,{title:"Hello from Centauri",description:"Centauri is a basic and simple Discord API wrapper written in Go."},i.createElement("main",null,i.createElement("div",{className:"container padding-top--md padding-bottom--lg"},i.createElement("div",{className:"markdown"},i.createElement("h1",null,"Centauri"),i.createElement("p",{className:"spaced"},i.createElement("a",{href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri"},i.createElement("img",{src:"https://pkg.go.dev/badge/github.com/kkrypt0nn/centauri.svg",alt:"Go Reference Badge"}))," ",i.createElement("a",{href:"https://github.com/kkrypt0nn/centauri/actions"},i.createElement("img",{src:"https://github.com/kkrypt0nn/centauri/actions/workflows/ci.yml/badge.svg",alt:"CI Badge"}))," ",i.createElement("a",{href:"https://discord.gg/feA6ZGRgpw"},i.createElement("img",{src:"https://img.shields.io/discord/1095412499416891482?logo=discord",alt:"Discord Server Badge"}))," ",i.createElement("a",{href:"https://github.com/kkrypt0nn/centauri/commits/main"},i.createElement("img",{src:"https://img.shields.io/github/last-commit/kkrypt0nn/centauri",alt:"Last Commit Badge"}))," ",i.createElement("a",{href:"https://conventionalcommits.org/en/v1.0.0/"},i.createElement("img",{src:"https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white",alt:"Conventional Commits Badge"}))),i.createElement(l.Z,{type:"note",title:"Note"},i.createElement("p",null,"Due to my military service from ",i.createElement("strong",null,"July 2023")," until"," ",i.createElement("strong",null,"April 2024")," I will not be able to work on this project a lot, most likely only during weekends if I have time and motivation ^-^")),i.createElement(l.Z,{type:"caution",title:"Warning"},i.createElement("p",null,"This is a WIP library that may be ",i.createElement("strong",null,"very unstable")," ","and eventually ",i.createElement("strong",null,"not fully optimised"),", use at your own care! This page will also be reworked.")),i.createElement("p",null,"Centauri is a Discord API wrapper written in Go with the goal of being easily understandable and simple to use, even for newcomers."),i.createElement("h2",{id:"features"},"Features"),i.createElement("p",null,"The plan for Centauri would be for it to cover Discord's APIs in their entirety. A small list of features that will be, or are already, in Centauri are:"),i.createElement("ul",null,i.createElement("li",null,"ANSI Colors"),i.createElement("li",null,"Background Tasks"),i.createElement("li",null,"Caching"),i.createElement("li",null,"Easy Commands Creation"),i.createElement("li",null,"Gateway API"),i.createElement("li",null,"Interactions (Buttons, Modals, etc.)"),i.createElement("li",null,"OAuth 2.0"),i.createElement("li",null,"Rate Limiter"),i.createElement("li",null,"REST API"),i.createElement("li",null,"RPC"),i.createElement("li",null,"Sharding"),i.createElement("li",null,"Webhooks")),i.createElement("h2",{id:"getting-started"},"Getting Started"),i.createElement("h3",{id:"installation"},"Installation"),i.createElement("p",null,"To get started you will simply need to install the library in your project by executing the following command."),i.createElement(a.Z,{language:"bash"},"go get github.com/kkrypt0nn/centauri"),i.createElement("h3",{id:"example-usage"},"Example Usage"),i.createElement("h4",{id:"rest-client"},"REST Client"),i.createElement("p",null,"If you just want to interact with the REST API, you may use the following code."),i.createElement(a.Z,{language:"go",showLineNumbers:!0},'package main\n\nimport (\n\t"fmt"\n\t"github.com/kkrypt0nn/centauri"\n)\n\nfunc main() {\n\tbotClient := centauri.NewRestClient("Bot BOT_TOKEN")\n\tbot, err := botClient.GetCurrentUser()\n\tif err != nil {\n\t\tfmt.Println(err)\n\t} else {\n\t\tfmt.Println("Got Bot:", bot.Username+"#"+bot.Discriminator)\n\t}\n}'),i.createElement("p",null,"More examples are available"," ",i.createElement("a",{href:"https://github.com/kkrypt0nn/centauri/tree/main/_examples"},"here"),"."),i.createElement("h2",{id:"documentation"},"Documentation"),i.createElement("p",null,"Other than this documentation, there is also the Go reference page available"," ",i.createElement("a",{target:"_blank",href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri"},"here")),i.createElement("h2",{id:"troubleshooting"},"Troubleshooting"),i.createElement("p",null,"Some troubleshooting articles and guides may be found"," ",i.createElement("a",{href:"/docs/troubleshooting/"},"here"),"."),i.createElement("h2",{id:"contributing"},"Contributing"),i.createElement("p",null,"Before the version ",i.createElement("strong",null,"1.0.0")," of the project, contributions are mostly closed as I'd like to lay down the foundations of the library. Afterward contributions will be more than welcome if following the"," ",i.createElement("a",{href:"https://centauri.krypton.ninja/community/contributing-guidelines/"},"Contributing Guidelines"),"."),i.createElement("h2",{id:"license"},"License"),i.createElement("p",null,"This library was made with \ud83d\udc9c by Krypton and is under the"," ",i.createElement("a",{href:"https://github.com/kkrypt0nn/centauri/blob/main/LICENSE.md"},"MIT License"),".")))))}}}]);