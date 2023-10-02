"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[183],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>f});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function c(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var i=r.createContext({}),p=function(e){var t=r.useContext(i),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},u=function(e){var t=p(e.components);return r.createElement(i.Provider,{value:t},e.children)},l="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,i=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),l=p(n),m=a,f=l["".concat(i,".").concat(m)]||l[m]||d[m]||o;return n?r.createElement(f,s(s({ref:t},u),{},{components:n})):r.createElement(f,s({ref:t},u))}));function f(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,s=new Array(o);s[0]=m;var c={};for(var i in t)hasOwnProperty.call(t,i)&&(c[i]=t[i]);c.originalType=e,c[l]="string"==typeof e?e:a,s[1]=c;for(var p=2;p<o;p++)s[p]=n[p];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},407:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>p,default:()=>b,frontMatter:()=>i,metadata:()=>u,toc:()=>d});var r=n(7462),a=(n(7294),n(3905)),o=n(941);const s=n.p+"assets/images/sonar-dark-c69ef43ead1c1b5048d3051a9505ad6e.png",c=n.p+"assets/images/sonar-light-9bc40d370d048bfb9da1f3a0cabb8038.png",i={title:"Scan",hide_title:!0,slug:"/actions/sonar-scanner/scan"},p=void 0,u={unversionedId:"actions/sonarqube/scan",id:"actions/sonarqube/scan",title:"Scan",description:"<ThemedImage",source:"@site/docs/actions/sonarqube/scan.mdx",sourceDirName:"actions/sonarqube",slug:"/actions/sonar-scanner/scan",permalink:"/actions/sonar-scanner/scan",draft:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/actions/sonarqube/scan.mdx",tags:[],version:"current",frontMatter:{title:"Scan",hide_title:!0,slug:"/actions/sonar-scanner/scan"},sidebar:"tutorialSidebar",previous:{title:"Run",permalink:"/actions/pytest/run"},next:{title:"Run",permalink:"/actions/tox/run"}},l={},d=[{value:"Input Variables",id:"input-variables",level:3}],m={toc:d},f="wrapper";function b(e){let{components:t,...n}=e;return(0,a.kt)(f,(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)(o.Z,{style:{height:"75px",margin:"20px 0 20px 0",marginLeft:"-15px"},sources:{light:s,dark:c},mdxType:"ThemedImage"}),(0,a.kt)("h1",{id:"sonar-scanner---scan"},"Sonar Scanner - Scan"),(0,a.kt)("p",null,"The scan action runs the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/SonarSource/sonar-scanner-cli"},"sonar-scanner-cli")," to generate reports for a ",(0,a.kt)("a",{parentName:"p",href:"https://www.sonarsource.com/products/sonarqube/"},"SonarQube")," code quality server."),(0,a.kt)("admonition",{type:"tip"},(0,a.kt)("p",{parentName:"admonition"},"This action uses the ",(0,a.kt)("a",{parentName:"p",href:"https://docs.sonarsource.com/sonarqube/9.9/analyzing-source-code/scanners/sonarscanner/#configuring-your-project"},"sonar-project.properties")," in the project root.")),(0,a.kt)("h3",{id:"input-variables"},"Input Variables"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/inputs#sonarqube"},"SONARQUBE_TOKEN"))))}b.isMDXComponent=!0}}]);