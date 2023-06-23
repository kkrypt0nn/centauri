"use strict";(self.webpackChunkcentauri_documentation=self.webpackChunkcentauri_documentation||[]).push([[74],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>y});var a=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,a,i=function(e,t){if(null==e)return{};var n,a,i={},l=Object.keys(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var s=a.createContext({}),c=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},u=function(e){var t=c(e.components);return a.createElement(s.Provider,{value:t},e.children)},p="mdxType",g={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},d=a.forwardRef((function(e,t){var n=e.components,i=e.mdxType,l=e.originalType,s=e.parentName,u=o(e,["components","mdxType","originalType","parentName"]),p=c(n),d=i,y=p["".concat(s,".").concat(d)]||p[d]||g[d]||l;return n?a.createElement(y,r(r({ref:t},u),{},{components:n})):a.createElement(y,r({ref:t},u))}));function y(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var l=n.length,r=new Array(l);r[0]=d;var o={};for(var s in t)hasOwnProperty.call(t,s)&&(o[s]=t[s]);o.originalType=e,o[p]="string"==typeof e?e:i,r[1]=o;for(var c=2;c<l;c++)r[c]=n[c];return a.createElement.apply(null,r)}return a.createElement.apply(null,n)}d.displayName="MDXCreateElement"},3950:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>r,default:()=>g,frontMatter:()=>l,metadata:()=>o,toc:()=>c});var a=n(7462),i=(n(7294),n(3905));const l={title:"Gateway Client",description:"This guide will explain how to use the Gateway Client and listen to events."},r=void 0,o={unversionedId:"guides/gateway_client/README",id:"guides/gateway_client/README",title:"Gateway Client",description:"This guide will explain how to use the Gateway Client and listen to events.",source:"@site/docs/guides/gateway_client/README.md",sourceDirName:"guides/gateway_client",slug:"/guides/gateway_client/",permalink:"/docs/guides/gateway_client/",draft:!1,editUrl:"https://github.com/kkrypt0nn/centauri/tree/main/documentation/docs/guides/gateway_client/README.md",tags:[],version:"current",frontMatter:{title:"Gateway Client",description:"This guide will explain how to use the Gateway Client and listen to events."},sidebar:"sidebar",previous:{title:"Background Tasks",permalink:"/docs/guides/background_tasks/"},next:{title:"HTTP Interceptor",permalink:"/docs/guides/http_interceptor/"}},s={},c=[{value:"Creating a Gateway Client",id:"creating-a-gateway-client",level:2},{value:"Arguments",id:"arguments",level:2},{value:"Listening to events",id:"listening-to-events",level:2},{value:"Login and close",id:"login-and-close",level:2},{value:"Sharding",id:"sharding",level:2},{value:"Using REST Client with Gateway Client",id:"using-rest-client-with-gateway-client",level:2}],u={toc:c},p="wrapper";function g(e){let{components:t,...n}=e;return(0,i.kt)(p,(0,a.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("p",null,"When you want to get events coming from Discord, for example when a message is sent in a channel or when a user joined a guild, then you will need a so called ",(0,i.kt)("strong",{parentName:"p"},"Gateway Client"),"."),(0,i.kt)("p",null,"You can see an example of a Gateway Client ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/kkrypt0nn/centauri/blob/main/_examples/gateway_client/main.go"},"here"),"."),(0,i.kt)("h2",{id:"creating-a-gateway-client"},"Creating a Gateway Client"),(0,i.kt)("p",null,"Creating and then using a Gateway Client is not rocket science, no really - it isn't. Similarly to the ",(0,i.kt)("a",{parentName:"p",href:"/docs/guides/rest_client"},"REST Client")," you will just need to call the ",(0,i.kt)("inlineCode",{parentName:"p"},"NewGatewayClient")," function on the ",(0,i.kt)("inlineCode",{parentName:"p"},"centauri")," package like the following."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'botClient := centauri.NewGatewayClient("Bot BOT_TOKEN", 0)\n')),(0,i.kt)("h2",{id:"arguments"},"Arguments"),(0,i.kt)("p",null,"The arguments you ",(0,i.kt)("strong",{parentName:"p"},"must")," pass into the function are the token and the intents you want. If you don't know what intents are, check ",(0,i.kt)("a",{parentName:"p",href:"/docs/guides/intents"},"this guide")," that explains it."),(0,i.kt)("p",null,"Similarly to the REST Client, the token must be prefixed with ",(0,i.kt)("inlineCode",{parentName:"p"},"Bot"),"."),(0,i.kt)("h2",{id:"listening-to-events"},"Listening to events"),(0,i.kt)("p",null,"To listen to events you just need to call the ",(0,i.kt)("inlineCode",{parentName:"p"},"On()")," method on the Gateway Client variable you've created."),(0,i.kt)("p",null,"As first argument you pass the ",(0,i.kt)("strong",{parentName:"p"},"type")," of the event, which you can easily list with IntelliSense and typing ",(0,i.kt)("inlineCode",{parentName:"p"},"gateway.EventType")," - one example is ",(0,i.kt)("inlineCode",{parentName:"p"},"EventTypeMessageCreate"),"."),(0,i.kt)("p",null,"As second argument you will pass the callback function that will get executed when the event is dispatched by Discord. It will always be like the following:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func(c *gateway.Client, x *gateway.<EventType>) {\n    // ...\n}\n")),(0,i.kt)("p",null,"Obviously you have to replace ",(0,i.kt)("inlineCode",{parentName:"p"},"<EventType>")," with the event type you have, for example ",(0,i.kt)("inlineCode",{parentName:"p"},"gateway.MessageCreate"),"."),(0,i.kt)("p",null,"The entire nonsense above results in something like"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'botClient.On(gateway.EventTypeMessageCreate, func(c *gateway.Client, message *gateway.MessageCreate) {\n    botClient.Logger.Info(fmt.Sprintf("Got a new message from %s: %s", message.Author.Username, message.Content))\n})\n')),(0,i.kt)("p",null,"to listen to events when messages are created (sent)."),(0,i.kt)("h2",{id:"login-and-close"},"Login and close"),(0,i.kt)("p",null,"After creating all of the above, you will have to call the ",(0,i.kt)("inlineCode",{parentName:"p"},"Login()")," method, otherwise your bot will never login and come online. If there are some errors or signals coming that the program has ended, you will want to call the ",(0,i.kt)("inlineCode",{parentName:"p"},"Close()")," method. For example"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"err := botClient.Login()\nif err != nil {\n    botClient.Logger.Error(err.Error())\n    botClient.Close()\n    return\n}\n\nsc := make(chan os.Signal, 1)\nsignal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)\n<-sc\nbotClient.Close()\n")),(0,i.kt)("h2",{id:"sharding"},"Sharding"),(0,i.kt)("p",null,"Sharding is something that is coming soon, so stay tuned!"),(0,i.kt)("h2",{id:"using-rest-client-with-gateway-client"},"Using REST Client with Gateway Client"),(0,i.kt)("p",null,"Yes, you can [use the REST Client","[(/docs/guides/rest_client/#using-the-client)]"," when just creating a Gateway Client and you won't need additional steps. All you have to do is call the ",(0,i.kt)("inlineCode",{parentName:"p"},"Rest()")," method on the Gateway Client variable and you will then have access to everything needed."))}g.isMDXComponent=!0}}]);