"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[85],{2503:(e,n,t)=>{t.d(n,{Z:()=>m});var a=t(7462),l=t(7294),r=t(6010),i=t(5999),c=t(6668),s=t(9960);const o={anchorWithStickyNavbar:"anchorWithStickyNavbar_LWe7",anchorWithHideOnScrollNavbar:"anchorWithHideOnScrollNavbar_WYt5"};function m(e){let{as:n,id:t,...m}=e;const{navbar:{hideOnScroll:d}}=(0,c.L)();if("h1"===n||!t)return l.createElement(n,(0,a.Z)({},m,{id:void 0}));const u=(0,i.I)({id:"theme.common.headingLinkTitle",message:"Direct link to {heading}",description:"Title for link to heading"},{heading:"string"==typeof m.children?m.children:t});return l.createElement(n,(0,a.Z)({},m,{className:(0,r.Z)("anchor",d?o.anchorWithHideOnScrollNavbar:o.anchorWithStickyNavbar,m.className),id:t}),m.children,l.createElement(s.Z,{className:"hash-link",to:`#${t}`,"aria-label":u,title:u},"\u200b"))}},76:(e,n,t)=>{t.d(n,{Z:()=>_});var a=t(7294),l=t(3905),r=t(7462),i=t(5742);var c=t(814);var s=t(9960);var o=t(6010),m=t(2389),d=t(6043);const u={details:"details_lb9f",isBrowser:"isBrowser_bmU9",collapsibleContent:"collapsibleContent_i85q"};function f(e){return!!e&&("SUMMARY"===e.tagName||f(e.parentElement))}function h(e,n){return!!e&&(e===n||h(e.parentElement,n))}function p(e){let{summary:n,children:t,...l}=e;const i=(0,m.Z)(),c=(0,a.useRef)(null),{collapsed:s,setCollapsed:p}=(0,d.u)({initialState:!l.open}),[v,g]=(0,a.useState)(l.open),E=a.isValidElement(n)?n:a.createElement("summary",null,n??"Details");return a.createElement("details",(0,r.Z)({},l,{ref:c,open:v,"data-collapsed":s,className:(0,o.Z)(u.details,i&&u.isBrowser,l.className),onMouseDown:e=>{f(e.target)&&e.detail>1&&e.preventDefault()},onClick:e=>{e.stopPropagation();const n=e.target;f(n)&&h(n,c.current)&&(e.preventDefault(),s?(p(!1),g(!0)):p(!0))}}),E,a.createElement(d.z,{lazy:!1,collapsed:s,disableSSRStyle:!0,onCollapseTransitionEnd:e=>{p(e),g(!e)}},a.createElement("div",{className:u.collapsibleContent},t)))}const v={details:"details_b_Ee"},g="alert alert--info";function E(e){let{...n}=e;return a.createElement(p,(0,r.Z)({},n,{className:(0,o.Z)(g,v.details,n.className)}))}var b=t(2503);function N(e){return a.createElement(b.Z,e)}const Z={containsTaskList:"containsTaskList_mC6p"};const L={img:"img_ev3q"};var C=t(3612),k=t(1875);const y={head:function(e){const n=a.Children.map(e.children,(e=>a.isValidElement(e)?function(e){if(e.props?.mdxType&&e.props.originalType){const{mdxType:n,originalType:t,...l}=e.props;return a.createElement(e.props.originalType,l)}return e}(e):e));return a.createElement(i.Z,e,n)},code:function(e){const n=["a","abbr","b","br","button","cite","code","del","dfn","em","i","img","input","ins","kbd","label","object","output","q","ruby","s","small","span","strong","sub","sup","time","u","var","wbr"];return a.Children.toArray(e.children).every((e=>"string"==typeof e&&!e.includes("\n")||(0,a.isValidElement)(e)&&n.includes(e.props?.mdxType)))?a.createElement("code",e):a.createElement(c.Z,e)},a:function(e){return a.createElement(s.Z,e)},pre:function(e){return a.createElement(c.Z,(0,a.isValidElement)(e.children)&&"code"===e.children.props?.originalType?e.children.props:{...e})},details:function(e){const n=a.Children.toArray(e.children),t=n.find((e=>a.isValidElement(e)&&"summary"===e.props?.mdxType)),l=a.createElement(a.Fragment,null,n.filter((e=>e!==t)));return a.createElement(E,(0,r.Z)({},e,{summary:t}),l)},ul:function(e){return a.createElement("ul",(0,r.Z)({},e,{className:(n=e.className,(0,o.Z)(n,n?.includes("contains-task-list")&&Z.containsTaskList))}));var n},img:function(e){return a.createElement("img",(0,r.Z)({loading:"lazy"},e,{className:(n=e.className,(0,o.Z)(n,L.img))}));var n},h1:e=>a.createElement(N,(0,r.Z)({as:"h1"},e)),h2:e=>a.createElement(N,(0,r.Z)({as:"h2"},e)),h3:e=>a.createElement(N,(0,r.Z)({as:"h3"},e)),h4:e=>a.createElement(N,(0,r.Z)({as:"h4"},e)),h5:e=>a.createElement(N,(0,r.Z)({as:"h5"},e)),h6:e=>a.createElement(N,(0,r.Z)({as:"h6"},e)),admonition:C.Z,mermaid:k.Z};function _(e){let{children:n}=e;return a.createElement(l.Zo,{components:y},n)}},4247:(e,n,t)=>{t.r(n),t.d(n,{default:()=>d});var a=t(7294),l=t(6010),r=t(1944),i=t(5281),c=t(7961),s=t(76),o=t(9407);const m={mdxPageWrapper:"mdxPageWrapper_j9I6"};function d(e){const{content:n}=e,{metadata:{title:t,description:d,frontMatter:u}}=n,{wrapperClassName:f,hide_table_of_contents:h}=u;return a.createElement(r.FG,{className:(0,l.Z)(f??i.k.wrapper.mdxPages,i.k.page.mdxPage)},a.createElement(r.d,{title:t,description:d}),a.createElement(c.Z,null,a.createElement("main",{className:"container container--fluid margin-vert--lg"},a.createElement("div",{className:(0,l.Z)("row",m.mdxPageWrapper)},a.createElement("div",{className:(0,l.Z)("col",!h&&"col--8")},a.createElement("article",null,a.createElement(s.Z,null,a.createElement(n,null)))),!h&&n.toc.length>0&&a.createElement("div",{className:"col col--2"},a.createElement(o.Z,{toc:n.toc,minHeadingLevel:u.toc_min_heading_level,maxHeadingLevel:u.toc_max_heading_level}))))))}},9407:(e,n,t)=>{t.d(n,{Z:()=>m});var a=t(7462),l=t(7294),r=t(6010),i=t(3743);const c={tableOfContents:"tableOfContents_bqdL",docItemContainer:"docItemContainer_F8PC"},s="table-of-contents__link toc-highlight",o="table-of-contents__link--active";function m(e){let{className:n,...t}=e;return l.createElement("div",{className:(0,r.Z)(c.tableOfContents,"thin-scrollbar",n)},l.createElement(i.Z,(0,a.Z)({},t,{linkClassName:s,linkActiveClassName:o})))}},3743:(e,n,t)=>{t.d(n,{Z:()=>h});var a=t(7462),l=t(7294),r=t(6668);function i(e){const n=e.map((e=>({...e,parentIndex:-1,children:[]}))),t=Array(7).fill(-1);n.forEach(((e,n)=>{const a=t.slice(2,e.level);e.parentIndex=Math.max(...a),t[e.level]=n}));const a=[];return n.forEach((e=>{const{parentIndex:t,...l}=e;t>=0?n[t].children.push(l):a.push(l)})),a}function c(e){let{toc:n,minHeadingLevel:t,maxHeadingLevel:a}=e;return n.flatMap((e=>{const n=c({toc:e.children,minHeadingLevel:t,maxHeadingLevel:a});return function(e){return e.level>=t&&e.level<=a}(e)?[{...e,children:n}]:n}))}function s(e){const n=e.getBoundingClientRect();return n.top===n.bottom?s(e.parentNode):n}function o(e,n){let{anchorTopOffset:t}=n;const a=e.find((e=>s(e).top>=t));if(a){return function(e){return e.top>0&&e.bottom<window.innerHeight/2}(s(a))?a:e[e.indexOf(a)-1]??null}return e[e.length-1]??null}function m(){const e=(0,l.useRef)(0),{navbar:{hideOnScroll:n}}=(0,r.L)();return(0,l.useEffect)((()=>{e.current=n?0:document.querySelector(".navbar").clientHeight}),[n]),e}function d(e){const n=(0,l.useRef)(void 0),t=m();(0,l.useEffect)((()=>{if(!e)return()=>{};const{linkClassName:a,linkActiveClassName:l,minHeadingLevel:r,maxHeadingLevel:i}=e;function c(){const e=function(e){return Array.from(document.getElementsByClassName(e))}(a),c=function(e){let{minHeadingLevel:n,maxHeadingLevel:t}=e;const a=[];for(let l=n;l<=t;l+=1)a.push(`h${l}.anchor`);return Array.from(document.querySelectorAll(a.join()))}({minHeadingLevel:r,maxHeadingLevel:i}),s=o(c,{anchorTopOffset:t.current}),m=e.find((e=>s&&s.id===function(e){return decodeURIComponent(e.href.substring(e.href.indexOf("#")+1))}(e)));e.forEach((e=>{!function(e,t){t?(n.current&&n.current!==e&&n.current.classList.remove(l),e.classList.add(l),n.current=e):e.classList.remove(l)}(e,e===m)}))}return document.addEventListener("scroll",c),document.addEventListener("resize",c),c(),()=>{document.removeEventListener("scroll",c),document.removeEventListener("resize",c)}}),[e,t])}function u(e){let{toc:n,className:t,linkClassName:a,isChild:r}=e;return n.length?l.createElement("ul",{className:r?void 0:t},n.map((e=>l.createElement("li",{key:e.id},l.createElement("a",{href:`#${e.id}`,className:a??void 0,dangerouslySetInnerHTML:{__html:e.value}}),l.createElement(u,{isChild:!0,toc:e.children,className:t,linkClassName:a}))))):null}const f=l.memo(u);function h(e){let{toc:n,className:t="table-of-contents table-of-contents__left-border",linkClassName:s="table-of-contents__link",linkActiveClassName:o,minHeadingLevel:m,maxHeadingLevel:u,...h}=e;const p=(0,r.L)(),v=m??p.tableOfContents.minHeadingLevel,g=u??p.tableOfContents.maxHeadingLevel,E=function(e){let{toc:n,minHeadingLevel:t,maxHeadingLevel:a}=e;return(0,l.useMemo)((()=>c({toc:i(n),minHeadingLevel:t,maxHeadingLevel:a})),[n,t,a])}({toc:n,minHeadingLevel:v,maxHeadingLevel:g});return d((0,l.useMemo)((()=>{if(s&&o)return{linkClassName:s,linkActiveClassName:o,minHeadingLevel:v,maxHeadingLevel:g}}),[s,o,v,g])),l.createElement(f,(0,a.Z)({toc:E,className:t,linkClassName:s},h))}}}]);