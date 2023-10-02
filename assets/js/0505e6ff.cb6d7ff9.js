"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[118],{3905:(e,t,n)=>{n.d(t,{Zo:()=>s,kt:()=>f});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),p=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},s=function(e){var t=p(e.components);return r.createElement(c.Provider,{value:t},e.children)},u="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,c=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),u=p(n),d=a,f=u["".concat(c,".").concat(d)]||u[d]||m[d]||i;return n?r.createElement(f,o(o({ref:t},s),{},{components:n})):r.createElement(f,o({ref:t},s))}));function f(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=d;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l[u]="string"==typeof e?e:a,o[1]=l;for(var p=2;p<i;p++)o[p]=n[p];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},396:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>l,default:()=>d,frontMatter:()=>o,metadata:()=>c,toc:()=>s});var r=n(7462),a=(n(7294),n(3905)),i=n(3724);const o={title:"Copy",hide_title:!0,slug:"/actions/container/copy"},l=void 0,c={unversionedId:"actions/container/copy",id:"actions/container/copy",title:"Copy",description:"The copy action publishes a container image to an image registry.",source:"@site/docs/actions/container/copy.mdx",sourceDirName:"actions/container",slug:"/actions/container/copy",permalink:"/actions/container/copy",draft:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/actions/container/copy.mdx",tags:[],version:"current",frontMatter:{title:"Copy",hide_title:!0,slug:"/actions/container/copy"},sidebar:"tutorialSidebar",previous:{title:"Build",permalink:"/actions/contianer/build"},next:{title:"Run",permalink:"/actions/eslint/run"}},p={},s=[{value:"Input Variables",id:"input-variables",level:3},{value:"Artifacts",id:"artifacts",level:3},{value:"Inputs:",id:"inputs",level:4}],u={toc:s},m="wrapper";function d(e){let{components:t,...n}=e;return(0,a.kt)(m,(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("img",{style:{height:"75px",margin:"20px 0 20px 0"},src:i.Z}),(0,a.kt)("h1",{id:"containerize---copy"},"Containerize - Copy"),(0,a.kt)("p",null,"The copy action publishes a container image to an image registry."),(0,a.kt)("h3",{id:"input-variables"},"Input Variables"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/inputs#container"},"CONTAINER_REGISTRY")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/inputs#container"},"CONTAINER_REGISTRY_USERNAME")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/inputs#container"},"CONTAINER_REGISTRY_PASSWORD"))),(0,a.kt)("h3",{id:"artifacts"},"Artifacts"),(0,a.kt)("h4",{id:"inputs"},"Inputs:"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"Name"),(0,a.kt)("th",{parentName:"tr",align:null},"Type"),(0,a.kt)("th",{parentName:"tr",align:null},"Description"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"image.tar"),(0,a.kt)("td",{parentName:"tr",align:null},"image"),(0,a.kt)("td",{parentName:"tr",align:null},"OCI compliant container image tar")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"version"),(0,a.kt)("td",{parentName:"tr",align:null},"file"),(0,a.kt)("td",{parentName:"tr",align:null},"The semantic version for the build that will be used as the container image tag")))))}d.isMDXComponent=!0},3724:(e,t,n)=>{n.d(t,{Z:()=>r});const r=n.p+"assets/images/oci-c57e1f440f0c853c5d2dbadc867dde12.png"}}]);