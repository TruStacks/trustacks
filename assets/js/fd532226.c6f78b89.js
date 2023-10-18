"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[88],{3905:(e,t,i)=>{i.d(t,{Zo:()=>s,kt:()=>m});var l=i(7294);function n(e,t,i){return t in e?Object.defineProperty(e,t,{value:i,enumerable:!0,configurable:!0,writable:!0}):e[t]=i,e}function a(e,t){var i=Object.keys(e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);t&&(l=l.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),i.push.apply(i,l)}return i}function c(e){for(var t=1;t<arguments.length;t++){var i=null!=arguments[t]?arguments[t]:{};t%2?a(Object(i),!0).forEach((function(t){n(e,t,i[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(i)):a(Object(i)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(i,t))}))}return e}function d(e,t){if(null==e)return{};var i,l,n=function(e,t){if(null==e)return{};var i,l,n={},a=Object.keys(e);for(l=0;l<a.length;l++)i=a[l],t.indexOf(i)>=0||(n[i]=e[i]);return n}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(l=0;l<a.length;l++)i=a[l],t.indexOf(i)>=0||Object.prototype.propertyIsEnumerable.call(e,i)&&(n[i]=e[i])}return n}var I=l.createContext({}),o=function(e){var t=l.useContext(I),i=t;return e&&(i="function"==typeof e?e(t):c(c({},t),e)),i},s=function(e){var t=o(e.components);return l.createElement(I.Provider,{value:t},e.children)},r="mdxType",Z={inlineCode:"code",wrapper:function(e){var t=e.children;return l.createElement(l.Fragment,{},t)}},b=l.forwardRef((function(e,t){var i=e.components,n=e.mdxType,a=e.originalType,I=e.parentName,s=d(e,["components","mdxType","originalType","parentName"]),r=o(i),b=n,m=r["".concat(I,".").concat(b)]||r[b]||Z[b]||a;return i?l.createElement(m,c(c({ref:t},s),{},{components:i})):l.createElement(m,c({ref:t},s))}));function m(e,t){var i=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var a=i.length,c=new Array(a);c[0]=b;var d={};for(var I in t)hasOwnProperty.call(t,I)&&(d[I]=t[I]);d.originalType=e,d[r]="string"==typeof e?e:n,c[1]=d;for(var o=2;o<a;o++)c[o]=i[o];return l.createElement.apply(null,c)}return l.createElement.apply(null,i)}b.displayName="MDXCreateElement"},8252:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>r,contentTitle:()=>o,default:()=>u,frontMatter:()=>I,metadata:()=>s,toc:()=>Z});var l=i(7462),n=(i(7294),i(3905)),a=i(941),c=i(2426);const d=i.p+"assets/images/architecture-dark-25d045a3cf0095af28547cd2398534d4.png",I={slug:"/architecture",sidebar_position:2,title:"Architecture"},o=void 0,s={unversionedId:"architecture",id:"architecture",title:"Architecture",description:"Developer Intent",source:"@site/docs/architecture.mdx",sourceDirName:".",slug:"/architecture",permalink:"/architecture",draft:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/architecture.mdx",tags:[],version:"current",sidebarPosition:2,frontMatter:{slug:"/architecture",sidebar_position:2,title:"Architecture"},sidebar:"tutorialSidebar",previous:{title:"Introduction",permalink:"/"},next:{title:"Overview",permalink:"/tutorial"}},r={},Z=[{value:"Developer Intent",id:"developer-intent",level:2},{value:"Architecture",id:"architecture",level:2},{value:"Action Plans",id:"action-plans",level:3},{value:"Inputs",id:"inputs",level:4},{value:"Guidelines",id:"guidelines",level:3},{value:"Single Source of Truth",id:"single-source-of-truth",level:4},{value:"Not Available Not Applicable",id:"not-available-not-applicable",level:4},{value:"Engine Flow",id:"engine-flow",level:2},{value:"Facts",id:"facts",level:4},{value:"Rules &amp; Actions",id:"rules--actions",level:4},{value:"Action Plan",id:"action-plan",level:4},{value:"Schedule",id:"schedule",level:4}],b={toc:Z},m="wrapper";function u(e){let{components:t,...I}=e;return(0,n.kt)(m,(0,l.Z)({},b,I,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h2",{id:"developer-intent"},"Developer Intent"),(0,n.kt)("p",null,(0,n.kt)("strong",{parentName:"p"},"Developer Intent")," is the desired process for the delivery of a software product. The desired software delivery process includes activities such as linting, testing, vulnerability scanning, building, and deploying."),(0,n.kt)("h2",{id:"architecture"},"Architecture"),(0,n.kt)("h3",{id:"action-plans"},"Action Plans"),(0,n.kt)("p",null,"The goal of the TruStacks engine is to accurately determine developer intent in the form an action plan. An action plan contains a list of actions and a list of the inputs that are required for the actions to be performed."),(0,n.kt)("h4",{id:"inputs"},"Inputs"),(0,n.kt)("p",null,"Inputs are required parameters that are necessary for running the actions in an action plan. Inputs are different from configuration because they do not change the behavior of actions."),(0,n.kt)("h3",{id:"guidelines"},"Guidelines"),(0,n.kt)("p",null,"TruStacks follows a set of guidelines to increase the efficacy of capturing developer intent."),(0,n.kt)("h4",{id:"single-source-of-truth"},"Single Source of Truth"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("p",{parentName:"li"},"There is one and only one source of truth. In general the source will be a single git repository with the contents of a single micro or monolithic application.")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("p",{parentName:"li"},"Monorepos and other full-stack sources can be considered a single source of truth if there is enough isolation between the components for them to be independently deployed. In this case the action plan would be derived from the root of a given monorepo path."))),(0,n.kt)("h4",{id:"not-available-not-applicable"},"Not Available Not Applicable"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("p",{parentName:"li"},"NANA reinforces immutable sources in that all source artifacts (ie. configs, scripts, tests, etc.) must be present during action plan generation in order to be used. ")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("p",{parentName:"li"},"Due to the high degree of fragmentation between project sources TruStacks does not admit actions into action plans where artifacts cannot be reliably predicted."))),(0,n.kt)("admonition",{type:"tip"},(0,n.kt)("p",{parentName:"admonition"},"NANA does not apply to source artifacts that do not impact action plan generation.")),(0,n.kt)("h2",{id:"engine-flow"},"Engine Flow"),(0,n.kt)("p",null,"TruStacks uses the following flow to generate and execute action plans."),(0,n.kt)("p",null,(0,n.kt)("img",{alt:"Engine Flow Diagram",src:i(6769).Z,width:"761",height:"61"})),(0,n.kt)("h4",{id:"facts"},"Facts"),(0,n.kt)("p",null,"Fact collection is the first step in the Engine flow where the engine gathers information about languages, frameworks, tool artifacts such as .rc or yaml, source files and tests, tool configurations and any other source facts that could be used for determining ",(0,n.kt)("a",{parentName:"p",href:"#developer-intent"},(0,n.kt)("inlineCode",{parentName:"a"},"Developer Intent")),"."),(0,n.kt)("h4",{id:"rules--actions"},"Rules & Actions"),(0,n.kt)("p",null,"After collecting facts the engine applies matching rules against the fact set. If rules are matched then the appropriate actions will be admitted into the action plan."),(0,n.kt)("p",null,"Actions are individual functions written in ",(0,n.kt)("a",{parentName:"p",href:"https://dagger.io/"},"Dagger")," that perform steps in a CI/CD pipeline such as linting or unit testing."),(0,n.kt)("h4",{id:"action-plan"},"Action Plan"),(0,n.kt)("p",null,"The action plan contains the list of matched actions and their associated inputs. Inputs are configuration parameters or credentials used by actions that exists outside of the application source. "),(0,n.kt)("admonition",{type:"tip"},(0,n.kt)("p",{parentName:"admonition"},"Inputs must be populated before executing an action plan.")),(0,n.kt)("h4",{id:"schedule"},"Schedule"),(0,n.kt)("p",null,"Actions admitted into an action plan are naive with no specific order. The scheduler places rules in appropriate order based on action classification and artifacts. "),(0,n.kt)("p",null,'Rules can be classified in a fixed stage, or selected for execution in a stage at runtime by the scheduler as "feeder" actions. Feeder actions exist only to provide inputs to a downstream action such as a container build action that "feeds" the output image to a vulnerability scan or image publish action.'),(0,n.kt)("p",null,"The scheduler ensures that actions between stages and inner stage are executed in the order of the required inputs. If no input is required by a given action the scheduler will run the action at whatever point it is introduced into the schedule."),(0,n.kt)(a.Z,{sources:{light:c.Z,dark:d},mdxType:"ThemedImage"}))}u.isMDXComponent=!0},2426:(e,t,i)=>{i.d(t,{Z:()=>l});const l=i.p+"assets/images/architecture-7b55079dc1673191d36e811bc4e1a23c.png"},6769:(e,t,i)=>{i.d(t,{Z:()=>l});const l="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPCEtLSBEbyBub3QgZWRpdCB0aGlzIGZpbGUgd2l0aCBlZGl0b3JzIG90aGVyIHRoYW4gZGlhZ3JhbXMubmV0IC0tPgo8IURPQ1RZUEUgc3ZnIFBVQkxJQyAiLS8vVzNDLy9EVEQgU1ZHIDEuMS8vRU4iICJodHRwOi8vd3d3LnczLm9yZy9HcmFwaGljcy9TVkcvMS4xL0RURC9zdmcxMS5kdGQiPgo8c3ZnIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHZlcnNpb249IjEuMSIgd2lkdGg9Ijc2MXB4IiBoZWlnaHQ9IjYxcHgiIHZpZXdCb3g9Ii0wLjUgLTAuNSA3NjEgNjEiIGNvbnRlbnQ9IiZsdDtteGZpbGUgaG9zdD0mcXVvdDthcHAuZGlhZ3JhbXMubmV0JnF1b3Q7IG1vZGlmaWVkPSZxdW90OzIwMjMtMDktMTNUMjM6MzY6NTMuMTU1WiZxdW90OyBhZ2VudD0mcXVvdDs1LjAgKFdpbmRvd3MpJnF1b3Q7IGV0YWc9JnF1b3Q7bklocTdBM01oNWhLOGp4NTc4TTgmcXVvdDsgdmVyc2lvbj0mcXVvdDsyMC43LjMmcXVvdDsmZ3Q7Jmx0O2RpYWdyYW0gaWQ9JnF1b3Q7R1hITXJlNDVNcTl6VmVvREUtYjAmcXVvdDsgbmFtZT0mcXVvdDtQYWdlLTEmcXVvdDsmZ3Q7N1pqYmJxTXdFSWFmSnBlTkFHK1N6V1UyaDdiU0hodHAyMTVhTUFGYXd5QXo1TkNuWHdOMkFMR2hYU25iNWdJdUV1YjNHSWI1L1NHTEFadEgrMnZKaytBYmVpQUdqdVh0QjJ3eGNCeDdOSGJVWDY0Y1NtVTYrVlFLdmd3OW5WUUo2L0FGdEdocE5RczlTQnVKaENnb1RKcWlpM0VNTGpVMExpWHVtbWtiRk0yN0p0eUhsckIydVdpcjk2RkhRYWwrSGxtVmZnT2hINWc3MjVZZWliaEoxa0lhY0E5M05Za3RCMnd1RWFrOGkvWnpFSG56VEYvS2Vhc1RvOGZDSk1UMGxnbVR4N3Q3dXRvOGIyOS9SL1F5KzM1ei9lUFhsYTNkU09sZ25oZzgxUUFkb3FRQWZZeTVXRmJxRjRsWjdFRitXVXRGVmM1WHhFU0p0aEtmZ09pZzNlUVpvWklDaW9RZVZSWEx3NE9lWHdTUGVUQWNtWEN4cnc4dURqcEtTZUl6ekZHZ0xHcGwwK0pRSXh1TXlSU3RPNU0veDhsV2FTbkZUTHJRMForcFhuSmMra0JkZmF3Y1ZTZ0FScURxVmhNbENFN2h0bGtJMTJ2U1ArWlZ0cWtUN2R3L3VLaXIzSEtSNlR1dHVFdHB5OXJLdU55RlhSQVNyQk5lTkdDbjhHMmExR3AyZWVobXIzZ1VpdHlYU0FVcFNNbkoyR0I2R2hBcFNKMFJtNmtmOVh6NVQ1NlFEbjFFWHdCUHduVG9ZbFFNdUdtUnV0cVUxMVduZjdueTBXQWRsK3RMcmVFeXJ0VTZMZzZsNjg2QUpOaDNyNGEyZVhxQ1lWMi93Wmg1TmUycTk0RnRJQTlxNzRLeDlaL3N0a2M5dEYzUUdvTmVwL2FFOGU5RHJTbXpodTFkSnFESDlqellIcmNRbDhQdHVPZTJrMXV6d1h1VlcrZER1YlZiM001Y0NqSHV5VDBQdVd4eWNlUk9lbkk3eVhYZVNpNzdVSEtkRStRcTdhZmdjYy92ZVRiTTdPTDRaUzNqMTI0QW50cHU5WjZmeGZQeDlQMDhWMkgxMmFRWXEzMThZc3MvJmx0Oy9kaWFncmFtJmd0OyZsdDsvbXhmaWxlJmd0OyI+PGRlZnM+PHN0eWxlIHR5cGU9InRleHQvY3NzIj5AaW1wb3J0IHVybChodHRwczovL2ZvbnRzLmdvb2dsZWFwaXMuY29tL2Nzcz9mYW1pbHk9bW9udHNlcnJhdCk7JiN4YTs8L3N0eWxlPjwvZGVmcz48Zz48cGF0aCBkPSJNIDEyMCAzMCBMIDE1My42MyAzMCIgZmlsbD0ibm9uZSIgc3Ryb2tlPSIjOTk5OTk5IiBzdHJva2UtbWl0ZXJsaW1pdD0iMTAiIHBvaW50ZXItZXZlbnRzPSJzdHJva2UiLz48cGF0aCBkPSJNIDE1OC44OCAzMCBMIDE1MS44OCAzMy41IEwgMTUzLjYzIDMwIEwgMTUxLjg4IDI2LjUgWiIgZmlsbD0iIzk5OTk5OSIgc3Ryb2tlPSIjOTk5OTk5IiBzdHJva2UtbWl0ZXJsaW1pdD0iMTAiIHBvaW50ZXItZXZlbnRzPSJhbGwiLz48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTIwIiBoZWlnaHQ9IjYwIiByeD0iOSIgcnk9IjkiIGZpbGw9InJnYigyNTUsIDI1NSwgMjU1KSIgc3Ryb2tlPSIjMzMzMzMzIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuNSAtMC41KSI+PHN3aXRjaD48Zm9yZWlnbk9iamVjdCBzdHlsZT0ib3ZlcmZsb3c6IHZpc2libGU7IHRleHQtYWxpZ246IGxlZnQ7IiBwb2ludGVyLWV2ZW50cz0ibm9uZSIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgcmVxdWlyZWRGZWF0dXJlcz0iaHR0cDovL3d3dy53My5vcmcvVFIvU1ZHMTEvZmVhdHVyZSNFeHRlbnNpYmlsaXR5Ij48ZGl2IHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hodG1sIiBzdHlsZT0iZGlzcGxheTogZmxleDsgYWxpZ24taXRlbXM6IHVuc2FmZSBjZW50ZXI7IGp1c3RpZnktY29udGVudDogdW5zYWZlIGNlbnRlcjsgd2lkdGg6IDExOHB4OyBoZWlnaHQ6IDFweDsgcGFkZGluZy10b3A6IDMwcHg7IG1hcmdpbi1sZWZ0OiAxcHg7Ij48ZGl2IHN0eWxlPSJib3gtc2l6aW5nOiBib3JkZXItYm94OyBmb250LXNpemU6IDBweDsgdGV4dC1hbGlnbjogY2VudGVyOyIgZGF0YS1kcmF3aW8tY29sb3JzPSJjb2xvcjogIzY2NjY2NjsgIj48ZGl2IHN0eWxlPSJkaXNwbGF5OiBpbmxpbmUtYmxvY2s7IGZvbnQtc2l6ZTogMTRweDsgZm9udC1mYW1pbHk6IG1vbnRzZXJyYXQ7IGNvbG9yOiByZ2IoMTAyLCAxMDIsIDEwMik7IGxpbmUtaGVpZ2h0OiAxLjI7IHBvaW50ZXItZXZlbnRzOiBhbGw7IGZvbnQtd2VpZ2h0OiBib2xkOyB3aGl0ZS1zcGFjZTogbm9ybWFsOyBvdmVyZmxvdy13cmFwOiBub3JtYWw7Ij5GYWN0czwvZGl2PjwvZGl2PjwvZGl2PjwvZm9yZWlnbk9iamVjdD48dGV4dCB4PSI2MCIgeT0iMzQiIGZpbGw9IiM2NjY2NjYiIGZvbnQtZmFtaWx5PSJtb250c2VycmF0IiBmb250LXNpemU9IjE0cHgiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtd2VpZ2h0PSJib2xkIj5GYWN0czwvdGV4dD48L3N3aXRjaD48L2c+PHBhdGggZD0iTSAyODAgMzAgTCAzMTMuNjMgMzAiIGZpbGw9Im5vbmUiIHN0cm9rZT0iIzk5OTk5OSIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBwb2ludGVyLWV2ZW50cz0ic3Ryb2tlIi8+PHBhdGggZD0iTSAzMTguODggMzAgTCAzMTEuODggMzMuNSBMIDMxMy42MyAzMCBMIDMxMS44OCAyNi41IFoiIGZpbGw9IiM5OTk5OTkiIHN0cm9rZT0iIzk5OTk5OSIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PHJlY3QgeD0iMTYwIiB5PSIwIiB3aWR0aD0iMTIwIiBoZWlnaHQ9IjYwIiByeD0iOSIgcnk9IjkiIGZpbGw9InJnYigyNTUsIDI1NSwgMjU1KSIgc3Ryb2tlPSIjMzMzMzMzIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuNSAtMC41KSI+PHN3aXRjaD48Zm9yZWlnbk9iamVjdCBzdHlsZT0ib3ZlcmZsb3c6IHZpc2libGU7IHRleHQtYWxpZ246IGxlZnQ7IiBwb2ludGVyLWV2ZW50cz0ibm9uZSIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgcmVxdWlyZWRGZWF0dXJlcz0iaHR0cDovL3d3dy53My5vcmcvVFIvU1ZHMTEvZmVhdHVyZSNFeHRlbnNpYmlsaXR5Ij48ZGl2IHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hodG1sIiBzdHlsZT0iZGlzcGxheTogZmxleDsgYWxpZ24taXRlbXM6IHVuc2FmZSBjZW50ZXI7IGp1c3RpZnktY29udGVudDogdW5zYWZlIGNlbnRlcjsgd2lkdGg6IDExOHB4OyBoZWlnaHQ6IDFweDsgcGFkZGluZy10b3A6IDMwcHg7IG1hcmdpbi1sZWZ0OiAxNjFweDsiPjxkaXYgc3R5bGU9ImJveC1zaXppbmc6IGJvcmRlci1ib3g7IGZvbnQtc2l6ZTogMHB4OyB0ZXh0LWFsaWduOiBjZW50ZXI7IiBkYXRhLWRyYXdpby1jb2xvcnM9ImNvbG9yOiAjNjY2NjY2OyAiPjxkaXYgc3R5bGU9ImRpc3BsYXk6IGlubGluZS1ibG9jazsgZm9udC1zaXplOiAxNHB4OyBmb250LWZhbWlseTogbW9udHNlcnJhdDsgY29sb3I6IHJnYigxMDIsIDEwMiwgMTAyKTsgbGluZS1oZWlnaHQ6IDEuMjsgcG9pbnRlci1ldmVudHM6IGFsbDsgZm9udC13ZWlnaHQ6IGJvbGQ7IHdoaXRlLXNwYWNlOiBub3JtYWw7IG92ZXJmbG93LXdyYXA6IG5vcm1hbDsiPlJ1bGVzPC9kaXY+PC9kaXY+PC9kaXY+PC9mb3JlaWduT2JqZWN0Pjx0ZXh0IHg9IjIyMCIgeT0iMzQiIGZpbGw9IiM2NjY2NjYiIGZvbnQtZmFtaWx5PSJtb250c2VycmF0IiBmb250LXNpemU9IjE0cHgiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtd2VpZ2h0PSJib2xkIj5SdWxlczwvdGV4dD48L3N3aXRjaD48L2c+PHBhdGggZD0iTSA0NDAgMzAgTCA0NzMuNjMgMzAiIGZpbGw9Im5vbmUiIHN0cm9rZT0iIzk5OTk5OSIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBwb2ludGVyLWV2ZW50cz0ic3Ryb2tlIi8+PHBhdGggZD0iTSA0NzguODggMzAgTCA0NzEuODggMzMuNSBMIDQ3My42MyAzMCBMIDQ3MS44OCAyNi41IFoiIGZpbGw9IiM5OTk5OTkiIHN0cm9rZT0iIzk5OTk5OSIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PHJlY3QgeD0iMzIwIiB5PSIwIiB3aWR0aD0iMTIwIiBoZWlnaHQ9IjYwIiByeD0iOSIgcnk9IjkiIGZpbGw9InJnYigyNTUsIDI1NSwgMjU1KSIgc3Ryb2tlPSIjMzMzMzMzIiBwb2ludGVyLWV2ZW50cz0iYWxsIi8+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuNSAtMC41KSI+PHN3aXRjaD48Zm9yZWlnbk9iamVjdCBzdHlsZT0ib3ZlcmZsb3c6IHZpc2libGU7IHRleHQtYWxpZ246IGxlZnQ7IiBwb2ludGVyLWV2ZW50cz0ibm9uZSIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgcmVxdWlyZWRGZWF0dXJlcz0iaHR0cDovL3d3dy53My5vcmcvVFIvU1ZHMTEvZmVhdHVyZSNFeHRlbnNpYmlsaXR5Ij48ZGl2IHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hodG1sIiBzdHlsZT0iZGlzcGxheTogZmxleDsgYWxpZ24taXRlbXM6IHVuc2FmZSBjZW50ZXI7IGp1c3RpZnktY29udGVudDogdW5zYWZlIGNlbnRlcjsgd2lkdGg6IDExOHB4OyBoZWlnaHQ6IDFweDsgcGFkZGluZy10b3A6IDMwcHg7IG1hcmdpbi1sZWZ0OiAzMjFweDsiPjxkaXYgc3R5bGU9ImJveC1zaXppbmc6IGJvcmRlci1ib3g7IGZvbnQtc2l6ZTogMHB4OyB0ZXh0LWFsaWduOiBjZW50ZXI7IiBkYXRhLWRyYXdpby1jb2xvcnM9ImNvbG9yOiAjNjY2NjY2OyAiPjxkaXYgc3R5bGU9ImRpc3BsYXk6IGlubGluZS1ibG9jazsgZm9udC1zaXplOiAxNHB4OyBmb250LWZhbWlseTogbW9udHNlcnJhdDsgY29sb3I6IHJnYigxMDIsIDEwMiwgMTAyKTsgbGluZS1oZWlnaHQ6IDEuMjsgcG9pbnRlci1ldmVudHM6IGFsbDsgZm9udC13ZWlnaHQ6IGJvbGQ7IHdoaXRlLXNwYWNlOiBub3JtYWw7IG92ZXJmbG93LXdyYXA6IG5vcm1hbDsiPkFjdGlvbnM8L2Rpdj48L2Rpdj48L2Rpdj48L2ZvcmVpZ25PYmplY3Q+PHRleHQgeD0iMzgwIiB5PSIzNCIgZmlsbD0iIzY2NjY2NiIgZm9udC1mYW1pbHk9Im1vbnRzZXJyYXQiIGZvbnQtc2l6ZT0iMTRweCIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZm9udC13ZWlnaHQ9ImJvbGQiPkFjdGlvbnM8L3RleHQ+PC9zd2l0Y2g+PC9nPjxwYXRoIGQ9Ik0gNjAwIDMwIEwgNjMzLjYzIDMwIiBmaWxsPSJub25lIiBzdHJva2U9IiM5OTk5OTkiIHN0cm9rZS1taXRlcmxpbWl0PSIxMCIgcG9pbnRlci1ldmVudHM9InN0cm9rZSIvPjxwYXRoIGQ9Ik0gNjM4Ljg4IDMwIEwgNjMxLjg4IDMzLjUgTCA2MzMuNjMgMzAgTCA2MzEuODggMjYuNSBaIiBmaWxsPSIjOTk5OTk5IiBzdHJva2U9IiM5OTk5OTkiIHN0cm9rZS1taXRlcmxpbWl0PSIxMCIgcG9pbnRlci1ldmVudHM9ImFsbCIvPjxyZWN0IHg9IjQ4MCIgeT0iMCIgd2lkdGg9IjEyMCIgaGVpZ2h0PSI2MCIgcng9IjkiIHJ5PSI5IiBmaWxsPSJyZ2IoMjU1LCAyNTUsIDI1NSkiIHN0cm9rZT0iIzMzMzMzMyIgcG9pbnRlci1ldmVudHM9ImFsbCIvPjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKC0wLjUgLTAuNSkiPjxzd2l0Y2g+PGZvcmVpZ25PYmplY3Qgc3R5bGU9Im92ZXJmbG93OiB2aXNpYmxlOyB0ZXh0LWFsaWduOiBsZWZ0OyIgcG9pbnRlci1ldmVudHM9Im5vbmUiIHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIHJlcXVpcmVkRmVhdHVyZXM9Imh0dHA6Ly93d3cudzMub3JnL1RSL1NWRzExL2ZlYXR1cmUjRXh0ZW5zaWJpbGl0eSI+PGRpdiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94aHRtbCIgc3R5bGU9ImRpc3BsYXk6IGZsZXg7IGFsaWduLWl0ZW1zOiB1bnNhZmUgY2VudGVyOyBqdXN0aWZ5LWNvbnRlbnQ6IHVuc2FmZSBjZW50ZXI7IHdpZHRoOiAxMThweDsgaGVpZ2h0OiAxcHg7IHBhZGRpbmctdG9wOiAzMHB4OyBtYXJnaW4tbGVmdDogNDgxcHg7Ij48ZGl2IHN0eWxlPSJib3gtc2l6aW5nOiBib3JkZXItYm94OyBmb250LXNpemU6IDBweDsgdGV4dC1hbGlnbjogY2VudGVyOyIgZGF0YS1kcmF3aW8tY29sb3JzPSJjb2xvcjogIzY2NjY2NjsgIj48ZGl2IHN0eWxlPSJkaXNwbGF5OiBpbmxpbmUtYmxvY2s7IGZvbnQtc2l6ZTogMTRweDsgZm9udC1mYW1pbHk6IG1vbnRzZXJyYXQ7IGNvbG9yOiByZ2IoMTAyLCAxMDIsIDEwMik7IGxpbmUtaGVpZ2h0OiAxLjI7IHBvaW50ZXItZXZlbnRzOiBhbGw7IGZvbnQtd2VpZ2h0OiBib2xkOyB3aGl0ZS1zcGFjZTogbm9ybWFsOyBvdmVyZmxvdy13cmFwOiBub3JtYWw7Ij5BY3Rpb24gUGxhbjwvZGl2PjwvZGl2PjwvZGl2PjwvZm9yZWlnbk9iamVjdD48dGV4dCB4PSI1NDAiIHk9IjM0IiBmaWxsPSIjNjY2NjY2IiBmb250LWZhbWlseT0ibW9udHNlcnJhdCIgZm9udC1zaXplPSIxNHB4IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LXdlaWdodD0iYm9sZCI+QWN0aW9uIFBsYW48L3RleHQ+PC9zd2l0Y2g+PC9nPjxyZWN0IHg9IjY0MCIgeT0iMCIgd2lkdGg9IjEyMCIgaGVpZ2h0PSI2MCIgcng9IjkiIHJ5PSI5IiBmaWxsPSJyZ2IoMjU1LCAyNTUsIDI1NSkiIHN0cm9rZT0iIzMzMzMzMyIgcG9pbnRlci1ldmVudHM9ImFsbCIvPjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKC0wLjUgLTAuNSkiPjxzd2l0Y2g+PGZvcmVpZ25PYmplY3Qgc3R5bGU9Im92ZXJmbG93OiB2aXNpYmxlOyB0ZXh0LWFsaWduOiBsZWZ0OyIgcG9pbnRlci1ldmVudHM9Im5vbmUiIHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIHJlcXVpcmVkRmVhdHVyZXM9Imh0dHA6Ly93d3cudzMub3JnL1RSL1NWRzExL2ZlYXR1cmUjRXh0ZW5zaWJpbGl0eSI+PGRpdiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94aHRtbCIgc3R5bGU9ImRpc3BsYXk6IGZsZXg7IGFsaWduLWl0ZW1zOiB1bnNhZmUgY2VudGVyOyBqdXN0aWZ5LWNvbnRlbnQ6IHVuc2FmZSBjZW50ZXI7IHdpZHRoOiAxMThweDsgaGVpZ2h0OiAxcHg7IHBhZGRpbmctdG9wOiAzMHB4OyBtYXJnaW4tbGVmdDogNjQxcHg7Ij48ZGl2IHN0eWxlPSJib3gtc2l6aW5nOiBib3JkZXItYm94OyBmb250LXNpemU6IDBweDsgdGV4dC1hbGlnbjogY2VudGVyOyIgZGF0YS1kcmF3aW8tY29sb3JzPSJjb2xvcjogIzY2NjY2NjsgIj48ZGl2IHN0eWxlPSJkaXNwbGF5OiBpbmxpbmUtYmxvY2s7IGZvbnQtc2l6ZTogMTRweDsgZm9udC1mYW1pbHk6IG1vbnRzZXJyYXQ7IGNvbG9yOiByZ2IoMTAyLCAxMDIsIDEwMik7IGxpbmUtaGVpZ2h0OiAxLjI7IHBvaW50ZXItZXZlbnRzOiBhbGw7IGZvbnQtd2VpZ2h0OiBib2xkOyB3aGl0ZS1zcGFjZTogbm9ybWFsOyBvdmVyZmxvdy13cmFwOiBub3JtYWw7Ij5TY2hlZHVsZTwvZGl2PjwvZGl2PjwvZGl2PjwvZm9yZWlnbk9iamVjdD48dGV4dCB4PSI3MDAiIHk9IjM0IiBmaWxsPSIjNjY2NjY2IiBmb250LWZhbWlseT0ibW9udHNlcnJhdCIgZm9udC1zaXplPSIxNHB4IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LXdlaWdodD0iYm9sZCI+U2NoZWR1bGU8L3RleHQ+PC9zd2l0Y2g+PC9nPjwvZz48c3dpdGNoPjxnIHJlcXVpcmVkRmVhdHVyZXM9Imh0dHA6Ly93d3cudzMub3JnL1RSL1NWRzExL2ZlYXR1cmUjRXh0ZW5zaWJpbGl0eSIvPjxhIHRyYW5zZm9ybT0idHJhbnNsYXRlKDAsLTUpIiB4bGluazpocmVmPSJodHRwczovL3d3dy5kaWFncmFtcy5uZXQvZG9jL2ZhcS9zdmctZXhwb3J0LXRleHQtcHJvYmxlbXMiIHRhcmdldD0iX2JsYW5rIj48dGV4dCB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LXNpemU9IjEwcHgiIHg9IjUwJSIgeT0iMTAwJSI+VGV4dCBpcyBub3QgU1ZHIC0gY2Fubm90IGRpc3BsYXk8L3RleHQ+PC9hPjwvc3dpdGNoPjwvc3ZnPg=="}}]);