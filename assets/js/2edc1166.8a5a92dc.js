"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[633],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>m});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=a.createContext({}),u=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},c=function(e){var t=u(e.components);return a.createElement(s.Provider,{value:t},e.children)},p="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},h=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,s=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),p=u(n),h=r,m=p["".concat(s,".").concat(h)]||p[h]||d[h]||o;return n?a.createElement(m,l(l({ref:t},c),{},{components:n})):a.createElement(m,l({ref:t},c))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,l=new Array(o);l[0]=h;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[p]="string"==typeof e?e:r,l[1]=i;for(var u=2;u<o;u++)l[u]=n[u];return a.createElement.apply(null,l)}return a.createElement.apply(null,n)}h.displayName="MDXCreateElement"},8971:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>S,contentTitle:()=>I,default:()=>C,frontMatter:()=>O,metadata:()=>P,toc:()=>x});var a=n(7462),r=n(7294),o=n(3905),l=n(6010),i=n(2466),s=n(6550),u=n(1980),c=n(7392),p=n(12);function d(e){return function(e){return r.Children.map(e,(e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}(e).map((e=>{let{props:{value:t,label:n,attributes:a,default:r}}=e;return{value:t,label:n,attributes:a,default:r}}))}function h(e){const{values:t,children:n}=e;return(0,r.useMemo)((()=>{const e=t??d(n);return function(e){const t=(0,c.l)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,n])}function m(e){let{value:t,tabValues:n}=e;return n.some((e=>e.value===t))}function f(e){let{queryString:t=!1,groupId:n}=e;const a=(0,s.k6)(),o=function(e){let{queryString:t=!1,groupId:n}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:t,groupId:n});return[(0,u._X)(o),(0,r.useCallback)((e=>{if(!o)return;const t=new URLSearchParams(a.location.search);t.set(o,e),a.replace({...a.location,search:t.toString()})}),[o,a])]}function b(e){const{defaultValue:t,queryString:n=!1,groupId:a}=e,o=h(e),[l,i]=(0,r.useState)((()=>function(e){let{defaultValue:t,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!m({value:t,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const a=n.find((e=>e.default))??n[0];if(!a)throw new Error("Unexpected error: 0 tabValues");return a.value}({defaultValue:t,tabValues:o}))),[s,u]=f({queryString:n,groupId:a}),[c,d]=function(e){let{groupId:t}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(t),[a,o]=(0,p.Nk)(n);return[a,(0,r.useCallback)((e=>{n&&o.set(e)}),[n,o])]}({groupId:a}),b=(()=>{const e=s??c;return m({value:e,tabValues:o})?e:null})();(0,r.useLayoutEffect)((()=>{b&&i(b)}),[b]);return{selectedValue:l,selectValue:(0,r.useCallback)((e=>{if(!m({value:e,tabValues:o}))throw new Error(`Can't select invalid tab value=${e}`);i(e),u(e),d(e)}),[u,d,o]),tabValues:o}}var g=n(2389);const y={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};function k(e){let{className:t,block:n,selectedValue:o,selectValue:s,tabValues:u}=e;const c=[],{blockElementScrollPositionUntilNextRender:p}=(0,i.o5)(),d=e=>{const t=e.currentTarget,n=c.indexOf(t),a=u[n].value;a!==o&&(p(t),s(a))},h=e=>{let t=null;switch(e.key){case"Enter":d(e);break;case"ArrowRight":{const n=c.indexOf(e.currentTarget)+1;t=c[n]??c[0];break}case"ArrowLeft":{const n=c.indexOf(e.currentTarget)-1;t=c[n]??c[c.length-1];break}}t?.focus()};return r.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,l.Z)("tabs",{"tabs--block":n},t)},u.map((e=>{let{value:t,label:n,attributes:i}=e;return r.createElement("li",(0,a.Z)({role:"tab",tabIndex:o===t?0:-1,"aria-selected":o===t,key:t,ref:e=>c.push(e),onKeyDown:h,onClick:d},i,{className:(0,l.Z)("tabs__item",y.tabItem,i?.className,{"tabs__item--active":o===t})}),n??t)})))}function v(e){let{lazy:t,children:n,selectedValue:a}=e;const o=(Array.isArray(n)?n:[n]).filter(Boolean);if(t){const e=o.find((e=>e.props.value===a));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return r.createElement("div",{className:"margin-top--md"},o.map(((e,t)=>(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==a}))))}function w(e){const t=b(e);return r.createElement("div",{className:(0,l.Z)("tabs-container",y.tabList)},r.createElement(k,(0,a.Z)({},e,t)),r.createElement(v,(0,a.Z)({},e,t)))}function T(e){const t=(0,g.Z)();return r.createElement(w,(0,a.Z)({key:String(t)},e))}const E={tabItem:"tabItem_Ymn6"};function N(e){let{children:t,hidden:n,className:a}=e;return r.createElement("div",{role:"tabpanel",className:(0,l.Z)(E.tabItem,a),hidden:n},t)}const O={title:"Plan",sidebar_position:3,slug:"/get-started/plan"},I="Plan",P={unversionedId:"get-started/plan",id:"get-started/plan",title:"Plan",description:"This part of the guide will get you started with your first action plan.",source:"@site/docs/get-started/plan.mdx",sourceDirName:"get-started",slug:"/get-started/plan",permalink:"/get-started/plan",draft:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/get-started/plan.mdx",tags:[],version:"current",sidebarPosition:3,frontMatter:{title:"Plan",sidebar_position:3,slug:"/get-started/plan"},sidebar:"tutorialSidebar",previous:{title:"Setup",permalink:"/get-started/setup"},next:{title:"Input",permalink:"/get-started/input"}},S={},x=[{value:"What is an action plan?",id:"what-is-an-action-plan",level:3},{value:"Your First Action Plan",id:"your-first-action-plan",level:2},{value:"Generate the action plan",id:"generate-the-action-plan",level:3},{value:"&quot;There Is No Spoon&quot;",id:"there-is-no-spoon",level:3}],V={toc:x},j="wrapper";function C(e){let{components:t,...n}=e;return(0,o.kt)(j,(0,a.Z)({},V,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"plan"},"Plan"),(0,o.kt)("p",null,"This part of the guide will get you started with your first action plan."),(0,o.kt)("h3",{id:"what-is-an-action-plan"},"What is an action plan?"),(0,o.kt)("p",null,"An action plan is a set of actions that are selected based on the contents of an applications's source code."),(0,o.kt)("h2",{id:"your-first-action-plan"},"Your First Action Plan"),(0,o.kt)("p",null,"To get started, ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/TruStacks/react-sample/fork"},"fork")," the sample application repo."),(0,o.kt)("h3",{id:"generate-the-action-plan"},"Generate the action plan"),(0,o.kt)("p",null,"In the TruStacks web console, click the ",(0,o.kt)("strong",{parentName:"p"},"New Action Plan")," button in the option bar. Enter the url of your repo fork."),(0,o.kt)(T,{defaultValue:"https",values:[{label:"https",value:"https"},{label:"ssh",value:"ssh"}],mdxType:"Tabs"},(0,o.kt)(N,{value:"https",mdxType:"TabItem"},"If you are using https with a private repo, enter your username and password."),(0,o.kt)(N,{value:"ssh",mdxType:"TabItem"},"If you are using ssh, you will prompted to enter a deploy key before generating the action plan.")),(0,o.kt)("p",null,"Leave the sub-path input empty since are using the project root to generate the action plan."),(0,o.kt)("p",null,"Click ",(0,o.kt)("strong",{parentName:"p"},"Generate Action Plan")),(0,o.kt)("admonition",{type:"tip"},(0,o.kt)("p",{parentName:"admonition"},"SSH urls will prompt for a key. "),(0,o.kt)("p",{parentName:"admonition"},"After entering the key, click ",(0,o.kt)("strong",{parentName:"p"},"Generate Action Plan"),".")),(0,o.kt)("admonition",{title:"monorepos",type:"info"},(0,o.kt)("p",{parentName:"admonition"},"The sub-path input allows for a single repo to have multiple action plan targets, with the limitation that all sub-paths must be self-contained. "),(0,o.kt)("p",{parentName:"admonition"},(0,o.kt)("em",{parentName:"p"},"Monorepos will be covered in detail in the core concepts"),".")),(0,o.kt)("p",null,"After a few seconds, the action plan will be generated and you should now have an action plan in your action plans list named ",(0,o.kt)("inlineCode",{parentName:"p"},"trustacks-react-sample")," (unless you renamed your fork)."),(0,o.kt)("p",null,"Click the action plan list item to display the discovered actions. Click on an action to get a brief description of what the action does."),(0,o.kt)("h3",{id:"there-is-no-spoon"},'"There Is No Spoon"'),(0,o.kt)("p",null,"Where is the pipeline? You are looking it... sort of. "),(0,o.kt)("p",null,"TruStacks takes an orchestration through classification approach. Actions are not simply selected by the system, they are also understood. With ",(0,o.kt)("strong",{parentName:"p"},(0,o.kt)("em",{parentName:"strong"},"Source-to-Action")),' TruStacks "builds pipelines" with no pipeline code. '),(0,o.kt)("p",null,"This enables action plans to be changed on the fly as developers and application sources grow and mature."),(0,o.kt)("p",null,'"There is no pipeline."'))}C.isMDXComponent=!0}}]);