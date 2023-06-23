"use strict";(self.webpackChunkcentauri_documentation=self.webpackChunkcentauri_documentation||[]).push([[613],{3905:(e,t,n)=>{n.d(t,{Zo:()=>l,kt:()=>y});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function a(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var c=r.createContext({}),u=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},l=function(e){var t=u(e.components);return r.createElement(c.Provider,{value:t},e.children)},p="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},v=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,c=e.parentName,l=a(e,["components","mdxType","originalType","parentName"]),p=u(n),v=o,y=p["".concat(c,".").concat(v)]||p[v]||d[v]||i;return n?r.createElement(y,s(s({ref:t},l),{},{components:n})):r.createElement(y,s({ref:t},l))}));function y(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,s=new Array(i);s[0]=v;var a={};for(var c in t)hasOwnProperty.call(t,c)&&(a[c]=t[c]);a.originalType=e,a[p]="string"==typeof e?e:o,s[1]=a;for(var u=2;u<i;u++)s[u]=n[u];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}v.displayName="MDXCreateElement"},2556:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>s,default:()=>d,frontMatter:()=>i,metadata:()=>a,toc:()=>u});var r=n(7462),o=(n(7294),n(3905));const i={title:"Not receiving events",description:"If you do not set the intents correctly when using the library, you will not receive events you expeted to receive."},s=void 0,a={unversionedId:"troubleshooting/not_receiving_events/README",id:"troubleshooting/not_receiving_events/README",title:"Not receiving events",description:"If you do not set the intents correctly when using the library, you will not receive events you expeted to receive.",source:"@site/docs/troubleshooting/not_receiving_events/README.md",sourceDirName:"troubleshooting/not_receiving_events",slug:"/troubleshooting/not_receiving_events/",permalink:"/docs/troubleshooting/not_receiving_events/",draft:!1,editUrl:"https://github.com/kkrypt0nn/centauri/tree/main/documentation/docs/troubleshooting/not_receiving_events/README.md",tags:[],version:"current",frontMatter:{title:"Not receiving events",description:"If you do not set the intents correctly when using the library, you will not receive events you expeted to receive."},sidebar:"sidebar",previous:{title:"Troubleshooting",permalink:"/docs/category/troubleshooting"},next:{title:"Token Issues",permalink:"/docs/troubleshooting/token_issues/"}},c={},u=[{value:"Missing intents",id:"missing-intents",level:2},{value:"More issues",id:"more-issues",level:2}],l={toc:u},p="wrapper";function d(e){let{components:t,...n}=e;return(0,o.kt)(p,(0,r.Z)({},l,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"missing-intents"},"Missing intents"),(0,o.kt)("p",null,"If you do not set the intents correctly when using the library, you will not receive events you expeted to receive."),(0,o.kt)("p",null,"Make sure that whenever you've created your Gateway Client you've specified the right intents. You can read about intents ",(0,o.kt)("a",{parentName:"p",href:"/docs/guides/intents"},"here")," and on how to use them properly. Here is an example to receive events when messages are sent in guilds and have access to their content:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'botClient := centauri.NewGatewayClient("Bot BOT_TOKEN", discord.IntentsGuildMessages|discord.IntentsMessageContent)\n')),(0,o.kt)("h2",{id:"more-issues"},"More issues"),(0,o.kt)("p",null,"If the explanations from this page do not help you with your issue(s), please join our ",(0,o.kt)("a",{parentName:"p",href:"https://discord.gg/feA6ZGRgpw"},"Discord server")," to get further assisted help."))}d.isMDXComponent=!0}}]);