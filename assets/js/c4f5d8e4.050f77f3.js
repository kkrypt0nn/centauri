"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[195],{3612:(e,t,n)=>{n.d(t,{Z:()=>d});var i=n(7294),a=n(6010),o=n(5281),l=n(5999);const r={admonition:"admonition_LlT9",admonitionHeading:"admonitionHeading_tbUL",admonitionIcon:"admonitionIcon_kALy",admonitionContent:"admonitionContent_S0QG"};const c={note:{infimaClassName:"secondary",iconComponent:function(){return i.createElement("svg",{viewBox:"0 0 14 16"},i.createElement("path",{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))},label:i.createElement(l.Z,{id:"theme.admonition.note",description:"The default label used for the Note admonition (:::note)"},"note")},tip:{infimaClassName:"success",iconComponent:function(){return i.createElement("svg",{viewBox:"0 0 12 16"},i.createElement("path",{fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))},label:i.createElement(l.Z,{id:"theme.admonition.tip",description:"The default label used for the Tip admonition (:::tip)"},"tip")},danger:{infimaClassName:"danger",iconComponent:function(){return i.createElement("svg",{viewBox:"0 0 12 16"},i.createElement("path",{fillRule:"evenodd",d:"M5.05.31c.81 2.17.41 3.38-.52 4.31C3.55 5.67 1.98 6.45.9 7.98c-1.45 2.05-1.7 6.53 3.53 7.7-2.2-1.16-2.67-4.52-.3-6.61-.61 2.03.53 3.33 1.94 2.86 1.39-.47 2.3.53 2.27 1.67-.02.78-.31 1.44-1.13 1.81 3.42-.59 4.78-3.42 4.78-5.56 0-2.84-2.53-3.22-1.25-5.61-1.52.13-2.03 1.13-1.89 2.75.09 1.08-1.02 1.8-1.86 1.33-.67-.41-.66-1.19-.06-1.78C8.18 5.31 8.68 2.45 5.05.32L5.03.3l.02.01z"}))},label:i.createElement(l.Z,{id:"theme.admonition.danger",description:"The default label used for the Danger admonition (:::danger)"},"danger")},info:{infimaClassName:"info",iconComponent:function(){return i.createElement("svg",{viewBox:"0 0 14 16"},i.createElement("path",{fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"}))},label:i.createElement(l.Z,{id:"theme.admonition.info",description:"The default label used for the Info admonition (:::info)"},"info")},caution:{infimaClassName:"warning",iconComponent:function(){return i.createElement("svg",{viewBox:"0 0 16 16"},i.createElement("path",{fillRule:"evenodd",d:"M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"}))},label:i.createElement(l.Z,{id:"theme.admonition.caution",description:"The default label used for the Caution admonition (:::caution)"},"caution")}},m={secondary:"note",important:"info",success:"tip",warning:"danger"};function s(e){const{mdxAdmonitionTitle:t,rest:n}=function(e){const t=i.Children.toArray(e),n=t.find((e=>i.isValidElement(e)&&"mdxAdmonitionTitle"===e.props?.mdxType)),a=i.createElement(i.Fragment,null,t.filter((e=>e!==n)));return{mdxAdmonitionTitle:n,rest:a}}(e.children);return{...e,title:e.title??t,children:n}}function d(e){const{children:t,type:n,title:l,icon:d}=s(e),u=function(e){const t=m[e]??e,n=c[t];return n||(console.warn(`No admonition config found for admonition type "${t}". Using Info as fallback.`),c.info)}(n),h=l??u.label,{iconComponent:p}=u,f=d??i.createElement(p,null);return i.createElement("div",{className:(0,a.Z)(o.k.common.admonition,o.k.common.admonitionType(e.type),"alert",`alert--${u.infimaClassName}`,r.admonition)},i.createElement("div",{className:r.admonitionHeading},i.createElement("span",{className:r.admonitionIcon},f),h),i.createElement("div",{className:r.admonitionContent},t))}},2841:(e,t,n)=>{n.r(t),n.d(t,{default:()=>r});var i=n(2263),a=n(3612),o=n(7961),l=n(7294);function r(){const{siteConfig:e}=(0,i.Z)();return l.createElement(o.Z,{title:"Hello from Centauri",description:"Centauri is a basic and simple Discord API wrapper written in Go."},l.createElement("main",null,l.createElement("div",{class:"container padding-top--md padding-bottom--lg"},l.createElement("div",{class:"markdown"},l.createElement("h1",null,"Centauri"),l.createElement("p",{className:"spaced"},l.createElement("a",{href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri"},l.createElement("img",{src:"https://pkg.go.dev/badge/github.com/kkrypt0nn/centauri.svg",alt:"Go Reference Badge"}))," ",l.createElement("a",{href:"https://github.com/kkrypt0nn/centauri/actions"},l.createElement("img",{src:"https://github.com/kkrypt0nn/centauri/actions/workflows/ci.yml/badge.svg",alt:"CI Badge"}))," ",l.createElement("a",{href:"https://discord.gg/feA6ZGRgpw"},l.createElement("img",{src:"https://img.shields.io/discord/1095412499416891482?logo=discord",alt:"Discord Server Badge"}))," ",l.createElement("a",{href:"https://github.com/kkrypt0nn/centauri/commits/main"},l.createElement("img",{src:"https://img.shields.io/github/last-commit/kkrypt0nn/centauri",alt:"Last Commit Badge"}))," ",l.createElement("a",{href:"https://conventionalcommits.org/en/v1.0.0/"},l.createElement("img",{src:"https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white",alt:"Conventional Commits Badge"}))),l.createElement(a.Z,{type:"note",title:"Note"},l.createElement("p",null,"Due to my military service from ",l.createElement("strong",null,"July 2023")," until"," ",l.createElement("strong",null,"April 2024")," I will not be able to work on this project a lot, most likely only during weekends if I have time and motivation ^-^")),l.createElement(a.Z,{type:"caution",title:"Warning"},l.createElement("p",null,"This is a WIP library that may be ",l.createElement("strong",null,"very unstable")," ","and eventually ",l.createElement("strong",null,"not fully optimised"),", use at your own care! This page will also be reworked.")),l.createElement("p",null,"Centauri is a basic and simple Discord API wrapper written in Go."),l.createElement("h2",{id:"features"},"Features"),l.createElement("p",null,"The plan for Centauri would be for it to cover Discord's APIs in their entirety. A small list of features that will be in Centauri are:"),l.createElement("ul",null,l.createElement("li",null,"REST API"),l.createElement("li",null,"Gateway API"),l.createElement("li",null,"Interactions (Buttons, Modals, etc.)"),l.createElement("li",null,"Tasks"),l.createElement("li",null,"ANSI Colors")),l.createElement("h2",{id:"examples"},"Examples"),l.createElement("p",null,"Examples are available"," ",l.createElement("a",{href:"https://github.com/kkrypt0nn/centauri/tree/main/_examples"},"here"),"."),l.createElement("h2",{id:"contributing"},"Contributing"),l.createElement("p",null,"Before the version ",l.createElement("strong",null,"1.0.0")," of the project, contributions are mostly closed as I'd like to lay down the foundations of the library. Afterward contributions will be more than welcome if following the"," ",l.createElement("a",{href:"https://centauri.krypton.ninja/community/contributing-guidelines/"},"Contributing Guidelines"),"."),l.createElement("h2",{id:"license"},"License"),l.createElement("p",null,"This library was made with \ud83d\udc9c by Krypton and is under the"," ",l.createElement("a",{href:"https://github.com/kkrypt0nn/centauri/blob/main/LICENSE.md"},"GPL v3")," ","license.")))))}}}]);