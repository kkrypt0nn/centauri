"use strict";(self.webpackChunkcentauri_documentation=self.webpackChunkcentauri_documentation||[]).push([[241],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>h});var r=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=r.createContext({}),u=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},c=function(e){var t=u(e.components);return r.createElement(l.Provider,{value:t},e.children)},p="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,a=e.originalType,l=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),p=u(n),d=i,h=p["".concat(l,".").concat(d)]||p[d]||m[d]||a;return n?r.createElement(h,o(o({ref:t},c),{},{components:n})):r.createElement(h,o({ref:t},c))}));function h(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var a=n.length,o=new Array(a);o[0]=d;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s[p]="string"==typeof e?e:i,o[1]=s;for(var u=2;u<a;u++)o[u]=n[u];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},9643:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>m,frontMatter:()=>a,metadata:()=>s,toc:()=>u});var r=n(7462),i=(n(7294),n(3905));const a={},o="Contributing Guidelines",s={unversionedId:"contributing-guidelines",id:"contributing-guidelines",title:"Contributing Guidelines",description:"Your input is amazing! Making contributing to this project as easy and transparent as possible is one of the most important side, this includes:",source:"@site/community/contributing-guidelines.md",sourceDirName:".",slug:"/contributing-guidelines",permalink:"/community/contributing-guidelines",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"community",previous:{title:"Code of Conduct",permalink:"/community/code-of-conduct"},next:{title:"Security Policy",permalink:"/community/security-policy"}},l={},u=[{value:"Wanted changes",id:"wanted-changes",level:2},{value:"Unwanted changes",id:"unwanted-changes",level:2},{value:"All code changes happen through pull requests",id:"all-code-changes-happen-through-pull-requests",level:2},{value:"Commit messages guidelines",id:"commit-messages-guidelines",level:2},{value:"Create a GitHub Issue and <strong>then</strong> a pull request",id:"create-a-github-issue-and-then-a-pull-request",level:2},{value:"License",id:"license",level:2}],c={toc:u},p="wrapper";function m(e){let{components:t,...n}=e;return(0,i.kt)(p,(0,r.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"contributing-guidelines"},"Contributing Guidelines"),(0,i.kt)("p",null,"Your input is amazing! Making contributing to this project as easy and transparent as possible is one of the most important side, this includes:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Reporting a bug"),(0,i.kt)("li",{parentName:"ul"},"Discussing the current state of the code"),(0,i.kt)("li",{parentName:"ul"},"Submitting a fix"),(0,i.kt)("li",{parentName:"ul"},"Proposing new features"),(0,i.kt)("li",{parentName:"ul"},"Becoming a maintainer")),(0,i.kt)("h2",{id:"wanted-changes"},"Wanted changes"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"New features"),(0,i.kt)("li",{parentName:"ul"},"Bug fixing"),(0,i.kt)("li",{parentName:"ul"},"Better documentation"),(0,i.kt)("li",{parentName:"ul"},"Fixing of spelling and grammatical issues")),(0,i.kt)("h2",{id:"unwanted-changes"},"Unwanted changes"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Whitespaces and punctuation changes"),(0,i.kt)("li",{parentName:"ul"},"Word changes using synonyms"),(0,i.kt)("li",{parentName:"ul"},"Entire rewrites of the project, or parts of the project - unless first approved by a maintainer")),(0,i.kt)("h2",{id:"all-code-changes-happen-through-pull-requests"},"All code changes happen through pull requests"),(0,i.kt)("p",null,"Pull requests are the best way to propose changes to the codebase. We actively welcome your pull requests:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"Fork the repo and create your branch from ",(0,i.kt)("inlineCode",{parentName:"li"},"main"),"."),(0,i.kt)("li",{parentName:"ol"},"Keep consistency with the current state of the codebase, this includes but is not limited to naming convention, Discord embeds, etc."),(0,i.kt)("li",{parentName:"ol"},"Format the code of the files properly."),(0,i.kt)("li",{parentName:"ol"},"Issue that pull request!")),(0,i.kt)("h2",{id:"commit-messages-guidelines"},"Commit messages guidelines"),(0,i.kt)("p",null,"This project uses ",(0,i.kt)("a",{parentName:"p",href:"https://conventionalcommits.org/en/v1.0.0/"},(0,i.kt)("inlineCode",{parentName:"a"},"Conventional Commits 1.0.0"))," hence your commit messages ",(0,i.kt)("strong",{parentName:"p"},"must")," follow the same convention or your contributions will be ignored, refused or assigned to another user or maintainer."),(0,i.kt)("p",null,"It would be more than welcome to keep your contributions as a single commit rather than, for examples, 20 ",(0,i.kt)("inlineCode",{parentName:"p"},'"fix: Stuff"')," commits in-between. You may use multiple commits if you believe the changes made in these commits have nothing, or close to nothing, in common - feel free to ask a maintainer on whether it should be a single commit or not."),(0,i.kt)("h2",{id:"create-a-github-issue-and-then-a-pull-request"},"Create a GitHub ",(0,i.kt)("a",{parentName:"h2",href:"https://github.com/kkrypt0nn/centauri/issues"},"Issue")," and ",(0,i.kt)("strong",{parentName:"h2"},"then")," a pull request"),(0,i.kt)("p",null,"Start contributing by first ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/kkrypt0nn/centauri/issues/new/choose"},"opening a new issue"),". Once that is done, you can create a pull request for the issue."),(0,i.kt)("h2",{id:"license"},"License"),(0,i.kt)("p",null,"Your submissions are understood to be under the same ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/kkrypt0nn/centauri/blob/main/LICENSE.md"},"MIT License")," that covers the project."))}m.isMDXComponent=!0}}]);