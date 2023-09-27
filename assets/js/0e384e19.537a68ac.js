"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[671],{3905:(e,i,t)=>{t.d(i,{Zo:()=>b,kt:()=>G});var l=t(7294);function I(e,i,t){return i in e?Object.defineProperty(e,i,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[i]=t,e}function c(e,i){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);i&&(l=l.filter((function(i){return Object.getOwnPropertyDescriptor(e,i).enumerable}))),t.push.apply(t,l)}return t}function d(e){for(var i=1;i<arguments.length;i++){var t=null!=arguments[i]?arguments[i]:{};i%2?c(Object(t),!0).forEach((function(i){I(e,i,t[i])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):c(Object(t)).forEach((function(i){Object.defineProperty(e,i,Object.getOwnPropertyDescriptor(t,i))}))}return e}function a(e,i){if(null==e)return{};var t,l,I=function(e,i){if(null==e)return{};var t,l,I={},c=Object.keys(e);for(l=0;l<c.length;l++)t=c[l],i.indexOf(t)>=0||(I[t]=e[t]);return I}(e,i);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(l=0;l<c.length;l++)t=c[l],i.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(I[t]=e[t])}return I}var n=l.createContext({}),Z=function(e){var i=l.useContext(n),t=i;return e&&(t="function"==typeof e?e(i):d(d({},i),e)),t},b=function(e){var i=Z(e.components);return l.createElement(n.Provider,{value:i},e.children)},s="mdxType",m={inlineCode:"code",wrapper:function(e){var i=e.children;return l.createElement(l.Fragment,{},i)}},o=l.forwardRef((function(e,i){var t=e.components,I=e.mdxType,c=e.originalType,n=e.parentName,b=a(e,["components","mdxType","originalType","parentName"]),s=Z(t),o=I,G=s["".concat(n,".").concat(o)]||s[o]||m[o]||c;return t?l.createElement(G,d(d({ref:i},b),{},{components:t})):l.createElement(G,d({ref:i},b))}));function G(e,i){var t=arguments,I=i&&i.mdxType;if("string"==typeof e||I){var c=t.length,d=new Array(c);d[0]=o;var a={};for(var n in i)hasOwnProperty.call(i,n)&&(a[n]=i[n]);a.originalType=e,a[s]="string"==typeof e?e:I,d[1]=a;for(var Z=2;Z<c;Z++)d[Z]=t[Z];return l.createElement.apply(null,d)}return l.createElement.apply(null,t)}o.displayName="MDXCreateElement"},9881:(e,i,t)=>{t.r(i),t.d(i,{assets:()=>n,contentTitle:()=>d,default:()=>m,frontMatter:()=>c,metadata:()=>a,toc:()=>Z});var l=t(7462),I=(t(7294),t(3905));const c={sidebar_position:1,title:"Introduction",slug:"/"},d="Welcome to TruStacks",a={unversionedId:"intro",id:"intro",title:"Introduction",description:"Welcome to the TruStacks docs! Click below to get started.",source:"@site/docs/intro.md",sourceDirName:".",slug:"/",permalink:"/",draft:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/intro.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1,title:"Introduction",slug:"/"},sidebar:"tutorialSidebar",next:{title:"Quickstart",permalink:"/quickstart"}},n={},Z=[{value:"What is TruStacks?",id:"what-is-trustacks",level:2},{value:"Generative Sofware Delivery",id:"generative-sofware-delivery",level:3},{value:"Generation Flow",id:"generation-flow",level:3},{value:"Facts",id:"facts",level:4},{value:"Rules &amp; Actions",id:"rules--actions",level:4},{value:"Action Plan",id:"action-plan",level:4},{value:"Schedule",id:"schedule",level:4}],b={toc:Z},s="wrapper";function m(e){let{components:i,...c}=e;return(0,I.kt)(s,(0,l.Z)({},b,c,{components:i,mdxType:"MDXLayout"}),(0,I.kt)("h1",{id:"welcome-to-trustacks"},"Welcome to TruStacks"),(0,I.kt)("div",{className:"TrustacksWelcome"},(0,I.kt)("p",null,"Welcome to the TruStacks docs! Click below to get started."),(0,I.kt)("a",{href:"quickstart",className:"TrustacksButtonLink"},"Quickstart")),(0,I.kt)("h2",{id:"what-is-trustacks"},"What is TruStacks?"),(0,I.kt)("p",null,"TruStacks is a generative software delivery engine that removes the need to build pipelines."),(0,I.kt)("h3",{id:"generative-sofware-delivery"},"Generative Sofware Delivery"),(0,I.kt)("p",null,(0,I.kt)("em",{parentName:"p"},"Intermediary Wiring")," is the pipeline code that sits between the source and the delivered product that is typically implemented with yaml or other no-code DSL's."),(0,I.kt)("p",null,"TruStacks uses a rule engine to generate actions plans that contain actions based on discovered facts in the application sources."),(0,I.kt)("h3",{id:"generation-flow"},"Generation Flow"),(0,I.kt)("p",null,"TruStacks uses the following flow to generate and execute action plans."),(0,I.kt)("p",null,(0,I.kt)("img",{alt:"Sonar Create Project",src:t(6769).Z,width:"761",height:"61"})),(0,I.kt)("h4",{id:"facts"},"Facts"),(0,I.kt)("p",null,"Fact collection is the first step in the generation flow. During fact collection the engine sets attributes about sources such as language, frameworks, and granular facts such as multi-stage docker builds or test script discovery."),(0,I.kt)("h4",{id:"rules--actions"},"Rules & Actions"),(0,I.kt)("p",null,"After collecting facts the engine applies matching rules against the fact set. If rules are matched then the appropriate actions will be admitted into the action plan. Actions contain common steps in a CI/CD pipeline such as linting or unit testing."),(0,I.kt)("h4",{id:"action-plan"},"Action Plan"),(0,I.kt)("p",null,"The action plan contains the list of matched actions and their associated inputs. Inputs are parameters and credentials that exists outside of the application source. "),(0,I.kt)("p",null,"Inputs must be populated before executing the action plan."),(0,I.kt)("h4",{id:"schedule"},"Schedule"),(0,I.kt)("p",null,"Actions admitted into an action plan are naive with no specific order. The scheduler places rules in appropriate order based on action classification and artifacts. "),(0,I.kt)("p",null,'Rules can be classified in a fixed stage, or selected for execution in a stage at runtime by the scheduler as "feeder" actions. Feeder actions exist only to provide inputs to a downstream action such as a container build action that "feeds" the output image to a vulnerability scan or image publish action.'),(0,I.kt)("p",null,"The scheduler ensures that actions between stages and inner stage are executed in the order of the required inputs. If no input is required by a given action the scheduler will run the action at whatever point it is introduced into the schedule."))}m.isMDXComponent=!0},6769:(e,i,t)=>{t.d(i,{Z:()=>l});const l="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPCEtLSBEbyBub3QgZWRpdCB0aGlzIGZpbGUgd2l0aCBlZGl0b3JzIG90aGVyIHRoYW4gZGlhZ3JhbXMubmV0IC0tPgo8IURPQ1RZUEUgc3ZnIFBVQkxJQyAiLS8vVzNDLy9EVEQgU1ZHIDEuMS8vRU4iICJodHRwOi8vd3d3LnczLm9yZy9HcmFwaGljcy9TVkcvMS4xL0RURC9zdmcxMS5kdGQiPgo8c3ZnIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHZlcnNpb249IjEuMSIgd2lkdGg9Ijc2MXB4IiBoZWlnaHQ9IjYxcHgiIHZpZXdCb3g9Ii0wLjUgLTAuNSA3NjEgNjEiIGNvbnRlbnQ9IiZsdDtteGZpbGUgaG9zdD0mcXVvdDthcHAuZGlhZ3JhbXMubmV0JnF1b3Q7IG1vZGlmaWVkPSZxdW90OzIwMjMtMDktMTNUMjM6MzY6NTMuMTU1WiZxdW90OyBhZ2VudD0mcXVvdDs1LjAgKFdpbmRvd3MpJnF1b3Q7IGV0YWc9JnF1b3Q7bklocTdBM01oNWhLOGp4NTc4TTgmcXVvdDsgdmVyc2lvbj0mcXVvdDsyMC43LjMmcXVvdDsmZ3Q7Jmx0O2RpYWdyYW0gaWQ9JnF1b3Q7R1hITXJlNDVNcTl6VmVvREUtYjAmcXVvdDsgbmFtZT0mcXVvdDtQYWdlLTEmcXVvdDsmZ3Q7N1pqYmJxTXdFSWFmSnBlTkFHK1N6V1UyaDdiU0hodHAyMTVhTUFGYXd5QXo1TkNuWHdOMkFMR2hYU25iNWdJdUV1YjNHSWI1L1NHTEFadEgrMnZKaytBYmVpQUdqdVh0QjJ3eGNCeDdOSGJVWDY0Y1NtVTYrVlFLdmd3OW5WUUo2L0FGdEdocE5RczlTQnVKaENnb1RKcWlpM0VNTGpVMExpWHVtbWtiRk0yN0p0eUhsckIydVdpcjk2RkhRYWwrSGxtVmZnT2hINWc3MjVZZWliaEoxa0lhY0E5M05Za3RCMnd1RWFrOGkvWnpFSG56VEYvS2Vhc1RvOGZDSk1UMGxnbVR4N3Q3dXRvOGIyOS9SL1F5KzM1ei9lUFhsYTNkU09sZ25oZzgxUUFkb3FRQWZZeTVXRmJxRjRsWjdFRitXVXRGVmM1WHhFU0p0aEtmZ09pZzNlUVpvWklDaW9RZVZSWEx3NE9lWHdTUGVUQWNtWEN4cnc4dURqcEtTZUl6ekZHZ0xHcGwwK0pRSXh1TXlSU3RPNU0veDhsV2FTbkZUTHJRMForcFhuSmMra0JkZmF3Y1ZTZ0FScURxVmhNbENFN2h0bGtJMTJ2U1ArWlZ0cWtUN2R3L3VLaXIzSEtSNlR1dHVFdHB5OXJLdU55RlhSQVNyQk5lTkdDbjhHMmExR3AyZWVobXIzZ1VpdHlYU0FVcFNNbkoyR0I2R2hBcFNKMFJtNmtmOVh6NVQ1NlFEbjFFWHdCUHduVG9ZbFFNdUdtUnV0cVUxMVduZjdueTBXQWRsK3RMcmVFeXJ0VTZMZzZsNjg2QUpOaDNyNGEyZVhxQ1lWMi93Wmg1TmUycTk0RnRJQTlxNzRLeDlaL3N0a2M5dEYzUUdvTmVwL2FFOGU5RHJTbXpodTFkSnFESDlqellIcmNRbDhQdHVPZTJrMXV6d1h1VlcrZER1YlZiM001Y0NqSHV5VDBQdVd4eWNlUk9lbkk3eVhYZVNpNzdVSEtkRStRcTdhZmdjYy92ZVRiTTdPTDRaUzNqMTI0QW50cHU5WjZmeGZQeDlQMDhWMkgxMmFRWXEzMThZc3MvJmx0Oy9kaWFncmFtJmd0OyZsdDsvbXhmaWxlJmd0OyI+PGRlZnM+PHN0eWxlIHR5cGU9InRleHQvY3NzIj5AaW1wb3J0IHVybChodHRwczovL2ZvbnRzLmdvb2dsZWFwaXMuY29tL2Nzcz9mYW1pbHk9bW9udHNlcnJhdCk7JiN4YTs8L3N0eWxlPjwvZGVmcz48Zz48cGF0aCBkPSJNIDEyMCAzMCBMIDE1My42MyAzMCIgZmlsbD0ibm9uZSIgc3Ryb2tlPSIjOTk5OTk5IiBzdHJva2UtbWl0ZXJsaW1pdD0iMTAiIHBvaW50ZXItZXZlbnRzPSJzdHJva2UiLz48cGF0aCBkPSJNIDE1OC44OCAzMCBMIDE1MS44OCAzMy41IEwgMTUzLjYzIDMwIEwgMTUxLjg4IDI2LjUgWiIgZmlsbD0iIzk5OTk5OSIgc3Ryb2tlPSIjOTk5OTk5IiBzdHJva2UtbWl0ZXJsaW1pdD0iMTAiIHBvaW50ZXItZXZlbnRzPSJhbGwiLz48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTIwIiBoZWlnaHQ9IjYwIiByeD0iOSIgcnk9IjkiIGZpbGw9InJnYigyNTUsIDI1NSwgMjU1KSIgc3Ryb2tlPSIjMzMzMzMzIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuNSAtMC41KSI+PHN3aXRjaD48Zm9yZWlnbk9iamVjdCBzdHlsZT0ib3ZlcmZsb3c6IHZpc2libGU7IHRleHQtYWxpZ246IGxlZnQ7IiBwb2ludGVyLWV2ZW50cz0ibm9uZSIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgcmVxdWlyZWRGZWF0dXJlcz0iaHR0cDovL3d3dy53My5vcmcvVFIvU1ZHMTEvZmVhdHVyZSNFeHRlbnNpYmlsaXR5Ij48ZGl2IHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hodG1sIiBzdHlsZT0iZGlzcGxheTogZmxleDsgYWxpZ24taXRlbXM6IHVuc2FmZSBjZW50ZXI7IGp1c3RpZnktY29udGVudDogdW5zYWZlIGNlbnRlcjsgd2lkdGg6IDExOHB4OyBoZWlnaHQ6IDFweDsgcGFkZGluZy10b3A6IDMwcHg7IG1hcmdpbi1sZWZ0OiAxcHg7Ij48ZGl2IHN0eWxlPSJib3gtc2l6aW5nOiBib3JkZXItYm94OyBmb250LXNpemU6IDBweDsgdGV4dC1hbGlnbjogY2VudGVyOyIgZGF0YS1kcmF3aW8tY29sb3JzPSJjb2xvcjogIzY2NjY2NjsgIj48ZGl2IHN0eWxlPSJkaXNwbGF5OiBpbmxpbmUtYmxvY2s7IGZvbnQtc2l6ZTogMTRweDsgZm9udC1mYW1pbHk6IG1vbnRzZXJyYXQ7IGNvbG9yOiByZ2IoMTAyLCAxMDIsIDEwMik7IGxpbmUtaGVpZ2h0OiAxLjI7IHBvaW50ZXItZXZlbnRzOiBhbGw7IGZvbnQtd2VpZ2h0OiBib2xkOyB3aGl0ZS1zcGFjZTogbm9ybWFsOyBvdmVyZmxvdy13cmFwOiBub3JtYWw7Ij5GYWN0czwvZGl2PjwvZGl2PjwvZGl2PjwvZm9yZWlnbk9iamVjdD48dGV4dCB4PSI2MCIgeT0iMzQiIGZpbGw9IiM2NjY2NjYiIGZvbnQtZmFtaWx5PSJtb250c2VycmF0IiBmb250LXNpemU9IjE0cHgiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtd2VpZ2h0PSJib2xkIj5GYWN0czwvdGV4dD48L3N3aXRjaD48L2c+PHBhdGggZD0iTSAyODAgMzAgTCAzMTMuNjMgMzAiIGZpbGw9Im5vbmUiIHN0cm9rZT0iIzk5OTk5OSIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBwb2ludGVyLWV2ZW50cz0ic3Ryb2tlIi8+PHBhdGggZD0iTSAzMTguODggMzAgTCAzMTEuODggMzMuNSBMIDMxMy42MyAzMCBMIDMxMS44OCAyNi41IFoiIGZpbGw9IiM5OTk5OTkiIHN0cm9rZT0iIzk5OTk5OSIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PHJlY3QgeD0iMTYwIiB5PSIwIiB3aWR0aD0iMTIwIiBoZWlnaHQ9IjYwIiByeD0iOSIgcnk9IjkiIGZpbGw9InJnYigyNTUsIDI1NSwgMjU1KSIgc3Ryb2tlPSIjMzMzMzMzIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuNSAtMC41KSI+PHN3aXRjaD48Zm9yZWlnbk9iamVjdCBzdHlsZT0ib3ZlcmZsb3c6IHZpc2libGU7IHRleHQtYWxpZ246IGxlZnQ7IiBwb2ludGVyLWV2ZW50cz0ibm9uZSIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgcmVxdWlyZWRGZWF0dXJlcz0iaHR0cDovL3d3dy53My5vcmcvVFIvU1ZHMTEvZmVhdHVyZSNFeHRlbnNpYmlsaXR5Ij48ZGl2IHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hodG1sIiBzdHlsZT0iZGlzcGxheTogZmxleDsgYWxpZ24taXRlbXM6IHVuc2FmZSBjZW50ZXI7IGp1c3RpZnktY29udGVudDogdW5zYWZlIGNlbnRlcjsgd2lkdGg6IDExOHB4OyBoZWlnaHQ6IDFweDsgcGFkZGluZy10b3A6IDMwcHg7IG1hcmdpbi1sZWZ0OiAxNjFweDsiPjxkaXYgc3R5bGU9ImJveC1zaXppbmc6IGJvcmRlci1ib3g7IGZvbnQtc2l6ZTogMHB4OyB0ZXh0LWFsaWduOiBjZW50ZXI7IiBkYXRhLWRyYXdpby1jb2xvcnM9ImNvbG9yOiAjNjY2NjY2OyAiPjxkaXYgc3R5bGU9ImRpc3BsYXk6IGlubGluZS1ibG9jazsgZm9udC1zaXplOiAxNHB4OyBmb250LWZhbWlseTogbW9udHNlcnJhdDsgY29sb3I6IHJnYigxMDIsIDEwMiwgMTAyKTsgbGluZS1oZWlnaHQ6IDEuMjsgcG9pbnRlci1ldmVudHM6IGFsbDsgZm9udC13ZWlnaHQ6IGJvbGQ7IHdoaXRlLXNwYWNlOiBub3JtYWw7IG92ZXJmbG93LXdyYXA6IG5vcm1hbDsiPlJ1bGVzPC9kaXY+PC9kaXY+PC9kaXY+PC9mb3JlaWduT2JqZWN0Pjx0ZXh0IHg9IjIyMCIgeT0iMzQiIGZpbGw9IiM2NjY2NjYiIGZvbnQtZmFtaWx5PSJtb250c2VycmF0IiBmb250LXNpemU9IjE0cHgiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtd2VpZ2h0PSJib2xkIj5SdWxlczwvdGV4dD48L3N3aXRjaD48L2c+PHBhdGggZD0iTSA0NDAgMzAgTCA0NzMuNjMgMzAiIGZpbGw9Im5vbmUiIHN0cm9rZT0iIzk5OTk5OSIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBwb2ludGVyLWV2ZW50cz0ic3Ryb2tlIi8+PHBhdGggZD0iTSA0NzguODggMzAgTCA0NzEuODggMzMuNSBMIDQ3My42MyAzMCBMIDQ3MS44OCAyNi41IFoiIGZpbGw9IiM5OTk5OTkiIHN0cm9rZT0iIzk5OTk5OSIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PHJlY3QgeD0iMzIwIiB5PSIwIiB3aWR0aD0iMTIwIiBoZWlnaHQ9IjYwIiByeD0iOSIgcnk9IjkiIGZpbGw9InJnYigyNTUsIDI1NSwgMjU1KSIgc3Ryb2tlPSIjMzMzMzMzIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuNSAtMC41KSI+PHN3aXRjaD48Zm9yZWlnbk9iamVjdCBzdHlsZT0ib3ZlcmZsb3c6IHZpc2libGU7IHRleHQtYWxpZ246IGxlZnQ7IiBwb2ludGVyLWV2ZW50cz0ibm9uZSIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgcmVxdWlyZWRGZWF0dXJlcz0iaHR0cDovL3d3dy53My5vcmcvVFIvU1ZHMTEvZmVhdHVyZSNFeHRlbnNpYmlsaXR5Ij48ZGl2IHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hodG1sIiBzdHlsZT0iZGlzcGxheTogZmxleDsgYWxpZ24taXRlbXM6IHVuc2FmZSBjZW50ZXI7IGp1c3RpZnktY29udGVudDogdW5zYWZlIGNlbnRlcjsgd2lkdGg6IDExOHB4OyBoZWlnaHQ6IDFweDsgcGFkZGluZy10b3A6IDMwcHg7IG1hcmdpbi1sZWZ0OiAzMjFweDsiPjxkaXYgc3R5bGU9ImJveC1zaXppbmc6IGJvcmRlci1ib3g7IGZvbnQtc2l6ZTogMHB4OyB0ZXh0LWFsaWduOiBjZW50ZXI7IiBkYXRhLWRyYXdpby1jb2xvcnM9ImNvbG9yOiAjNjY2NjY2OyAiPjxkaXYgc3R5bGU9ImRpc3BsYXk6IGlubGluZS1ibG9jazsgZm9udC1zaXplOiAxNHB4OyBmb250LWZhbWlseTogbW9udHNlcnJhdDsgY29sb3I6IHJnYigxMDIsIDEwMiwgMTAyKTsgbGluZS1oZWlnaHQ6IDEuMjsgcG9pbnRlci1ldmVudHM6IGFsbDsgZm9udC13ZWlnaHQ6IGJvbGQ7IHdoaXRlLXNwYWNlOiBub3JtYWw7IG92ZXJmbG93LXdyYXA6IG5vcm1hbDsiPkFjdGlvbnM8L2Rpdj48L2Rpdj48L2Rpdj48L2ZvcmVpZ25PYmplY3Q+PHRleHQgeD0iMzgwIiB5PSIzNCIgZmlsbD0iIzY2NjY2NiIgZm9udC1mYW1pbHk9Im1vbnRzZXJyYXQiIGZvbnQtc2l6ZT0iMTRweCIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZm9udC13ZWlnaHQ9ImJvbGQiPkFjdGlvbnM8L3RleHQ+PC9zd2l0Y2g+PC9nPjxwYXRoIGQ9Ik0gNjAwIDMwIEwgNjMzLjYzIDMwIiBmaWxsPSJub25lIiBzdHJva2U9IiM5OTk5OTkiIHN0cm9rZS1taXRlcmxpbWl0PSIxMCIgcG9pbnRlci1ldmVudHM9InN0cm9rZSIvPjxwYXRoIGQ9Ik0gNjM4Ljg4IDMwIEwgNjMxLjg4IDMzLjUgTCA2MzMuNjMgMzAgTCA2MzEuODggMjYuNSBaIiBmaWxsPSIjOTk5OTk5IiBzdHJva2U9IiM5OTk5OTkiIHN0cm9rZS1taXRlcmxpbWl0PSIxMCIgcG9pbnRlci1ldmVudHM9ImFsbCIvPjxyZWN0IHg9IjQ4MCIgeT0iMCIgd2lkdGg9IjEyMCIgaGVpZ2h0PSI2MCIgcng9IjkiIHJ5PSI5IiBmaWxsPSJyZ2IoMjU1LCAyNTUsIDI1NSkiIHN0cm9rZT0iIzMzMzMzMyIgcG9pbnRlci1ldmVudHM9ImFsbCIvPjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKC0wLjUgLTAuNSkiPjxzd2l0Y2g+PGZvcmVpZ25PYmplY3Qgc3R5bGU9Im92ZXJmbG93OiB2aXNpYmxlOyB0ZXh0LWFsaWduOiBsZWZ0OyIgcG9pbnRlci1ldmVudHM9Im5vbmUiIHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIHJlcXVpcmVkRmVhdHVyZXM9Imh0dHA6Ly93d3cudzMub3JnL1RSL1NWRzExL2ZlYXR1cmUjRXh0ZW5zaWJpbGl0eSI+PGRpdiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94aHRtbCIgc3R5bGU9ImRpc3BsYXk6IGZsZXg7IGFsaWduLWl0ZW1zOiB1bnNhZmUgY2VudGVyOyBqdXN0aWZ5LWNvbnRlbnQ6IHVuc2FmZSBjZW50ZXI7IHdpZHRoOiAxMThweDsgaGVpZ2h0OiAxcHg7IHBhZGRpbmctdG9wOiAzMHB4OyBtYXJnaW4tbGVmdDogNDgxcHg7Ij48ZGl2IHN0eWxlPSJib3gtc2l6aW5nOiBib3JkZXItYm94OyBmb250LXNpemU6IDBweDsgdGV4dC1hbGlnbjogY2VudGVyOyIgZGF0YS1kcmF3aW8tY29sb3JzPSJjb2xvcjogIzY2NjY2NjsgIj48ZGl2IHN0eWxlPSJkaXNwbGF5OiBpbmxpbmUtYmxvY2s7IGZvbnQtc2l6ZTogMTRweDsgZm9udC1mYW1pbHk6IG1vbnRzZXJyYXQ7IGNvbG9yOiByZ2IoMTAyLCAxMDIsIDEwMik7IGxpbmUtaGVpZ2h0OiAxLjI7IHBvaW50ZXItZXZlbnRzOiBhbGw7IGZvbnQtd2VpZ2h0OiBib2xkOyB3aGl0ZS1zcGFjZTogbm9ybWFsOyBvdmVyZmxvdy13cmFwOiBub3JtYWw7Ij5BY3Rpb24gUGxhbjwvZGl2PjwvZGl2PjwvZGl2PjwvZm9yZWlnbk9iamVjdD48dGV4dCB4PSI1NDAiIHk9IjM0IiBmaWxsPSIjNjY2NjY2IiBmb250LWZhbWlseT0ibW9udHNlcnJhdCIgZm9udC1zaXplPSIxNHB4IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LXdlaWdodD0iYm9sZCI+QWN0aW9uIFBsYW48L3RleHQ+PC9zd2l0Y2g+PC9nPjxyZWN0IHg9IjY0MCIgeT0iMCIgd2lkdGg9IjEyMCIgaGVpZ2h0PSI2MCIgcng9IjkiIHJ5PSI5IiBmaWxsPSJyZ2IoMjU1LCAyNTUsIDI1NSkiIHN0cm9rZT0iIzMzMzMzMyIgcG9pbnRlci1ldmVudHM9ImFsbCIvPjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKC0wLjUgLTAuNSkiPjxzd2l0Y2g+PGZvcmVpZ25PYmplY3Qgc3R5bGU9Im92ZXJmbG93OiB2aXNpYmxlOyB0ZXh0LWFsaWduOiBsZWZ0OyIgcG9pbnRlci1ldmVudHM9Im5vbmUiIHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIHJlcXVpcmVkRmVhdHVyZXM9Imh0dHA6Ly93d3cudzMub3JnL1RSL1NWRzExL2ZlYXR1cmUjRXh0ZW5zaWJpbGl0eSI+PGRpdiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94aHRtbCIgc3R5bGU9ImRpc3BsYXk6IGZsZXg7IGFsaWduLWl0ZW1zOiB1bnNhZmUgY2VudGVyOyBqdXN0aWZ5LWNvbnRlbnQ6IHVuc2FmZSBjZW50ZXI7IHdpZHRoOiAxMThweDsgaGVpZ2h0OiAxcHg7IHBhZGRpbmctdG9wOiAzMHB4OyBtYXJnaW4tbGVmdDogNjQxcHg7Ij48ZGl2IHN0eWxlPSJib3gtc2l6aW5nOiBib3JkZXItYm94OyBmb250LXNpemU6IDBweDsgdGV4dC1hbGlnbjogY2VudGVyOyIgZGF0YS1kcmF3aW8tY29sb3JzPSJjb2xvcjogIzY2NjY2NjsgIj48ZGl2IHN0eWxlPSJkaXNwbGF5OiBpbmxpbmUtYmxvY2s7IGZvbnQtc2l6ZTogMTRweDsgZm9udC1mYW1pbHk6IG1vbnRzZXJyYXQ7IGNvbG9yOiByZ2IoMTAyLCAxMDIsIDEwMik7IGxpbmUtaGVpZ2h0OiAxLjI7IHBvaW50ZXItZXZlbnRzOiBhbGw7IGZvbnQtd2VpZ2h0OiBib2xkOyB3aGl0ZS1zcGFjZTogbm9ybWFsOyBvdmVyZmxvdy13cmFwOiBub3JtYWw7Ij5TY2hlZHVsZTwvZGl2PjwvZGl2PjwvZGl2PjwvZm9yZWlnbk9iamVjdD48dGV4dCB4PSI3MDAiIHk9IjM0IiBmaWxsPSIjNjY2NjY2IiBmb250LWZhbWlseT0ibW9udHNlcnJhdCIgZm9udC1zaXplPSIxNHB4IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LXdlaWdodD0iYm9sZCI+U2NoZWR1bGU8L3RleHQ+PC9zd2l0Y2g+PC9nPjwvZz48c3dpdGNoPjxnIHJlcXVpcmVkRmVhdHVyZXM9Imh0dHA6Ly93d3cudzMub3JnL1RSL1NWRzExL2ZlYXR1cmUjRXh0ZW5zaWJpbGl0eSIvPjxhIHRyYW5zZm9ybT0idHJhbnNsYXRlKDAsLTUpIiB4bGluazpocmVmPSJodHRwczovL3d3dy5kaWFncmFtcy5uZXQvZG9jL2ZhcS9zdmctZXhwb3J0LXRleHQtcHJvYmxlbXMiIHRhcmdldD0iX2JsYW5rIj48dGV4dCB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LXNpemU9IjEwcHgiIHg9IjUwJSIgeT0iMTAwJSI+VGV4dCBpcyBub3QgU1ZHIC0gY2Fubm90IGRpc3BsYXk8L3RleHQ+PC9hPjwvc3dpdGNoPjwvc3ZnPg=="}}]);