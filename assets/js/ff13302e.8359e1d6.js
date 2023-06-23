"use strict";(self.webpackChunkcentauri_documentation=self.webpackChunkcentauri_documentation||[]).push([[468],{3905:(e,t,r)=>{r.d(t,{Zo:()=>c,kt:()=>f});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function o(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var u=n.createContext({}),p=function(e){var t=n.useContext(u),r=t;return e&&(r="function"==typeof e?e(t):s(s({},t),e)),r},c=function(e){var t=p(e.components);return n.createElement(u.Provider,{value:t},e.children)},l="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,u=e.parentName,c=o(e,["components","mdxType","originalType","parentName"]),l=p(r),d=a,f=l["".concat(u,".").concat(d)]||l[d]||m[d]||i;return r?n.createElement(f,s(s({ref:t},c),{},{components:r})):n.createElement(f,s({ref:t},c))}));function f(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,s=new Array(i);s[0]=d;var o={};for(var u in t)hasOwnProperty.call(t,u)&&(o[u]=t[u]);o.originalType=e,o[l]="string"==typeof e?e:a,s[1]=o;for(var p=2;p<i;p++)s[p]=r[p];return n.createElement.apply(null,s)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},5897:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>u,contentTitle:()=>s,default:()=>m,frontMatter:()=>i,metadata:()=>o,toc:()=>p});var n=r(7462),a=(r(7294),r(3905));const i={title:"Utils",description:"Centauri offers some useful utilities packages that is often used when coding bots."},s=void 0,o={unversionedId:"utils/README",id:"utils/README",title:"Utils",description:"Centauri offers some useful utilities packages that is often used when coding bots.",source:"@site/docs/utils/README.md",sourceDirName:"utils",slug:"/utils/",permalink:"/docs/utils/",draft:!1,editUrl:"https://github.com/kkrypt0nn/centauri/tree/main/documentation/docs/utils/README.md",tags:[],version:"current",frontMatter:{title:"Utils",description:"Centauri offers some useful utilities packages that is often used when coding bots."},sidebar:"sidebar",previous:{title:"Token Issues",permalink:"/docs/troubleshooting/token_issues/"}},u={},p=[{value:"Flags",id:"flags",level:2},{value:"Timestamp",id:"timestamp",level:2}],c={toc:p},l="wrapper";function m(e){let{components:t,...i}=e;return(0,a.kt)(l,(0,n.Z)({},c,i,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("p",null,"Centauri offers some useful utilities packages that is often used when coding bots."),(0,a.kt)("h2",{id:"flags"},"Flags"),(0,a.kt)("admonition",{type:"info"},(0,a.kt)("p",{parentName:"admonition"},"These functions are already implemented for the Discord structures. For example if you have a ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri/discord#User"},(0,a.kt)("inlineCode",{parentName:"a"},"User"))," structure the ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri/discord#UserFlags"},"user's flags")," structure has the various needed methods implemented. So you can directly use"),(0,a.kt)("pre",{parentName:"admonition"},(0,a.kt)("code",{parentName:"pre",className:"language-go"},"user.PublicFlags.Has(discord.UserFlagHypeSquadBravery) // True if the user is part of the HypeSquad Bravery\n"))),(0,a.kt)("p",null,"The ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri/utils/flags"},"flags package")," provides helpers to handle bitwise flags which are used for permissions, user flags, message flags, etc. The methods available can be seen in the ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri/utils/flags"},"Go Reference documentation"),"."),(0,a.kt)("h2",{id:"timestamp"},"Timestamp"),(0,a.kt)("p",null,"The ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri/utils/timestamp"},"timestamp package")," provides helpers to send messages in a time format on Discord, as shown here:"),(0,a.kt)("p",null,(0,a.kt)("img",{alt:"Timestamp format",src:r(9789).Z,width:"304",height:"80"})),(0,a.kt)("p",null,"The code used for the example above was the following."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"timestamp.FormatTime(time.Now(), timestamp.FormatTypeLongDate))\n")),(0,a.kt)("p",null,"The function and different formats available can be seen in the ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/github.com/kkrypt0nn/centauri/utils/timestamp"},"Go Reference documentation"),"."))}m.isMDXComponent=!0},9789:(e,t,r)=>{r.d(t,{Z:()=>n});const n=r.p+"assets/images/timestamp_format-bafc62ad00754ec3c510ac9ce1c8a0d3.png"}}]);