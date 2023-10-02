"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[840],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>f});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function c(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=r.createContext({}),u=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=u(e.components);return r.createElement(l.Provider,{value:t},e.children)},s="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,p=c(e,["components","mdxType","originalType","parentName"]),s=u(n),m=a,f=s["".concat(l,".").concat(m)]||s[m]||d[m]||i;return n?r.createElement(f,o(o({ref:t},p),{},{components:n})):r.createElement(f,o({ref:t},p))}));function f(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=m;var c={};for(var l in t)hasOwnProperty.call(t,l)&&(c[l]=t[l]);c.originalType=e,c[s]="string"==typeof e?e:a,o[1]=c;for(var u=2;u<i;u++)o[u]=n[u];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},1612:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>c,default:()=>m,frontMatter:()=>o,metadata:()=>l,toc:()=>p});var r=n(7462),a=(n(7294),n(3905)),i=n(3724);const o={title:"Build",hide_title:!0,slug:"/actions/contianer/build"},c=void 0,l={unversionedId:"actions/container/build",id:"actions/container/build",title:"Build",description:"The build action builds an OCI compliant image from the application Dockerfile or Containerfile.",source:"@site/docs/actions/container/build.mdx",sourceDirName:"actions/container",slug:"/actions/contianer/build",permalink:"/actions/contianer/build",draft:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/actions/container/build.mdx",tags:[],version:"current",frontMatter:{title:"Build",hide_title:!0,slug:"/actions/contianer/build"},sidebar:"tutorialSidebar",previous:{title:"Overview",permalink:"/actions"},next:{title:"Copy",permalink:"/actions/container/copy"}},u={},p=[{value:"Artifacts",id:"artifacts",level:3},{value:"Outputs:",id:"outputs",level:4}],s={toc:p},d="wrapper";function m(e){let{components:t,...n}=e;return(0,a.kt)(d,(0,r.Z)({},s,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("img",{style:{height:"75px",margin:"20px 0 20px 0"},src:i.Z}),(0,a.kt)("h1",{id:"containerize---build"},"Containerize - Build"),(0,a.kt)("p",null,"The build action builds an ",(0,a.kt)("a",{parentName:"p",href:"https://opencontainers.org/"},"OCI compliant")," image from the application Dockerfile or Containerfile."),(0,a.kt)("h3",{id:"artifacts"},"Artifacts"),(0,a.kt)("h4",{id:"outputs"},"Outputs:"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"Name"),(0,a.kt)("th",{parentName:"tr",align:null},"Type"),(0,a.kt)("th",{parentName:"tr",align:null},"Description"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"image.tar"),(0,a.kt)("td",{parentName:"tr",align:null},"image"),(0,a.kt)("td",{parentName:"tr",align:null},"OCI compliant container image tar")))))}m.isMDXComponent=!0},3724:(e,t,n)=>{n.d(t,{Z:()=>r});const r=n.p+"assets/images/oci-c57e1f440f0c853c5d2dbadc867dde12.png"}}]);