"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[426],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>h});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=a.createContext({}),p=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},c=function(e){var t=p(e.components);return a.createElement(s.Provider,{value:t},e.children)},u="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},k=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,s=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),u=p(n),k=r,h=u["".concat(s,".").concat(k)]||u[k]||d[k]||o;return n?a.createElement(h,l(l({ref:t},c),{},{components:n})):a.createElement(h,l({ref:t},c))}));function h(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,l=new Array(o);l[0]=k;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[u]="string"==typeof e?e:r,l[1]=i;for(var p=2;p<o;p++)l[p]=n[p];return a.createElement.apply(null,l)}return a.createElement.apply(null,n)}k.displayName="MDXCreateElement"},1878:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>l,default:()=>d,frontMatter:()=>o,metadata:()=>i,toc:()=>p});var a=n(7462),r=(n(7294),n(3905));const o={title:"Setup",sidebar_position:2,slug:"/get-started/setup"},l="Setup",i={unversionedId:"get-started/setup",id:"get-started/setup",title:"Setup",description:"This part of the guide will get your environment set up for running an action plan.",source:"@site/docs/get-started/setup.md",sourceDirName:"get-started",slug:"/get-started/setup",permalink:"/get-started/setup",draft:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/get-started/setup.md",tags:[],version:"current",sidebarPosition:2,frontMatter:{title:"Setup",sidebar_position:2,slug:"/get-started/setup"},sidebar:"tutorialSidebar",previous:{title:"Overview",permalink:"/get-started/intro"},next:{title:"Plan",permalink:"/get-started/plan"}},s={},p=[{value:"TruStacks Setup",id:"trustacks-setup",level:2},{value:"Docker",id:"docker",level:3},{value:"Validate docker",id:"validate-docker",level:4},{value:"Age",id:"age",level:3},{value:"Validate age",id:"validate-age",level:4},{value:"Action Plan Setup",id:"action-plan-setup",level:2},{value:"k3d",id:"k3d",level:3},{value:"Install",id:"install",level:4},{value:"Cluster Creation",id:"cluster-creation",level:4},{value:"SonarCloud",id:"sonarcloud",level:3}],c={toc:p},u="wrapper";function d(e){let{components:t,...o}=e;return(0,r.kt)(u,(0,a.Z)({},c,o,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"setup"},"Setup"),(0,r.kt)("p",null,"This part of the guide will get your environment set up for running an action plan."),(0,r.kt)("admonition",{type:"info"},(0,r.kt)("p",{parentName:"admonition"},"This guide uses hosted TruStacks to get you started quickly. Register for a trial account ",(0,r.kt)("a",{parentName:"p",href:"https://trustacks-website.web.app/#pricing"},"here")," if you haven't already.")),(0,r.kt)("h2",{id:"trustacks-setup"},"TruStacks Setup"),(0,r.kt)("p",null,"TruStacks requires ",(0,r.kt)("a",{parentName:"p",href:"https://www.docker.com/"},"docker")," to run action plans and ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/FiloSottile/age/releases"},"age")," for input encryption at rest."),(0,r.kt)("p",null,"Follow the instructions below to get those tools installed and configured."),(0,r.kt)("h3",{id:"docker"},"Docker"),(0,r.kt)("p",null,"Follow the instructions ",(0,r.kt)("a",{parentName:"p",href:"https://docs.docker.com/engine/install/"},"here")," to install docker on your machine."),(0,r.kt)("h4",{id:"validate-docker"},"Validate docker"),(0,r.kt)("p",null,"Run the following command to validate your docker installation."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"docker run -it quay.io/trustacks/tsctl:0.1.0 -h\n")),(0,r.kt)("p",null,"The output should display the TruStacks cli usage:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},'Trustacks software delivery engine\n\nUsage:\n  tsctl [command]\n\nAvailable Commands:\n  completion  Generate the autocompletion script for the specified shell\n  help        Help about any command\n  login       login to trustacks\n  plan        Generate an action plan\n  run         Run an action plan\n  server      start the api server\n  stack       manage input stacks\n  version     Show the cli version\n\nFlags:\n  -h, --help            help for tsctl\n      --server string   rpc server host\n\nUse "tsctl [command] --help" for more information about a command.\n')),(0,r.kt)("admonition",{type:"caution"},(0,r.kt)("p",{parentName:"admonition"},"TruStacks is not tested with other OCI runtimes such as podman or runc. They are not likely to work without additional modifications.")),(0,r.kt)("h3",{id:"age"},"Age"),(0,r.kt)("p",null,(0,r.kt)("a",{parentName:"p",href:"https://github.com/FiloSottile/age/releases"},"Age")," keys are used to encrypt ",(0,r.kt)("a",{parentName:"p",href:"/stacks"},"Stack Inputs")," at rest. Get the latest age release for your machine ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/FiloSottile/age/releases"},"here")," and are not recommended."),(0,r.kt)("h4",{id:"validate-age"},"Validate age"),(0,r.kt)("p",null,"Run the following command to validate age:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"age-keygen -h\n")),(0,r.kt)("p",null,"The output should the age cli usage:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"Usage:\n    age-keygen [-o OUTPUT]\n    age-keygen -y [-o OUTPUT] [INPUT]\n\nOptions:\n    -o, --output OUTPUT       Write the result to the file at path OUTPUT.\n    -y                        Convert an identity file to a recipients file.\n\nage-keygen generates a new native X25519 key pair, and outputs it to\nstandard output or to the OUTPUT file.\n\nIf an OUTPUT file is specified, the public key is printed to standard error.\nIf OUTPUT already exists, it is not overwritten.\n\nIn -y mode, age-keygen reads an identity file from INPUT or from standard\ninput and writes the corresponding recipient(s) to OUTPUT or to standard\noutput, one per line, with no comments.\n...\n")),(0,r.kt)("h2",{id:"action-plan-setup"},"Action Plan Setup"),(0,r.kt)("p",null,"The sample app is a react application that will be deployed on kubernetes."),(0,r.kt)("p",null,"Follow the instructions below to configure the local k8s cluster and additional services for the action plan."),(0,r.kt)("h3",{id:"k3d"},"k3d"),(0,r.kt)("h4",{id:"install"},"Install"),(0,r.kt)("p",null,"K3d deploys a rancher k3s cluster into a docker container, so you will need to have docker installed. Follow the ",(0,r.kt)("a",{parentName:"p",href:"https://k3d.io/#installation"},"installation guide")," to get k3d installed on your local machine."),(0,r.kt)("h4",{id:"cluster-creation"},"Cluster Creation"),(0,r.kt)("p",null,"Once k3d is installed, use the following command to create a cluster:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},'k3d cluster create -p "50123:80@loadbalancer" trustacks\n')),(0,r.kt)("p",null,"This command will create a new k3d cluster named ",(0,r.kt)("inlineCode",{parentName:"p"},"trustacks"),". The ",(0,r.kt)("inlineCode",{parentName:"p"},"-p")," option will create a loadbalancer on ",(0,r.kt)("inlineCode",{parentName:"p"},"8081"),". This loadbalancer will be used later in the guide to access the toolchain components."),(0,r.kt)("admonition",{type:"tip"},(0,r.kt)("p",{parentName:"admonition"},"If port ",(0,r.kt)("inlineCode",{parentName:"p"},"8081")," is already in use on your machine then feel free to use a different port.")),(0,r.kt)("p",null,"Once the cluster is created check the output of ",(0,r.kt)("inlineCode",{parentName:"p"},"docker ps")," and confirm that you have the ",(0,r.kt)("inlineCode",{parentName:"p"},"k3d-trustacks-serverlb")," and ",(0,r.kt)("inlineCode",{parentName:"p"},"k3d-trustacks-server-0")," containers."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},'CONTAINER ID IMAGE                            COMMAND                  CREATED          STATUS          PORTS                                           NAMES\n3e0600614d9f ghcr.io/k3d-io/k3d-tools:5.4.3   "/app/k3d-tools noop"    27 seconds ago   Up 26 seconds                                                   k3d-trustacks-tools\n6b4aaee146ba ghcr.io/k3d-io/k3d-proxy:5.4.3   "/bin/sh -c nginx-pr\u2026"   28 seconds ago   Up 19 seconds   0.0.0.0:8081->80/tcp, 0.0.0.0:44345->6443/tcp   k3d-trustacks-serverlb\n9eda9a6fc566 rancher/k3s:v1.23.6-k3s1         "/bin/k3s server --t\u2026"   28 seconds ago   Up 24 seconds                                                   k3d-trustacks-server-0\n')),(0,r.kt)("p",null,"Run the following command to confirm that the kubectl client was installed and that the cluster api is healthy:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"kubectl get ns\n")),(0,r.kt)("p",null,"The command should return the following output:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"NAME              STATUS   AGE\nkube-system       Active   20s\ndefault           Active   20s\nkube-public       Active   20s\nkube-node-lease   Active   20s\n")),(0,r.kt)("h3",{id:"sonarcloud"},"SonarCloud"),(0,r.kt)("p",null,"Sign up for a ",(0,r.kt)("a",{parentName:"p",href:"https://sonarcloud.io"},"SonarCloud")," account."),(0,r.kt)("p",null,"After signing up, click the ",(0,r.kt)("button",{className:"TrustacksSonarcloudNewButton"},"+")," at the top right of the page and select ",(0,r.kt)("strong",{parentName:"p"},"Analyze new project")),(0,r.kt)("p",null,"On the Analyze projects page click ",(0,r.kt)("strong",{parentName:"p"},"create a project manually")," in the second to the right."),(0,r.kt)("p",null,(0,r.kt)("img",{alt:"Sonar Create Project",src:n(6683).Z,width:"421",height:"83"})),(0,r.kt)("p",null,"If you have a new SonarCloud and you have not created an organization, under the ",(0,r.kt)("inlineCode",{parentName:"p"},"Organization")," drop-down, click ",(0,r.kt)("strong",{parentName:"p"},"Create another organization"),". Create an organization with your desired name."),(0,r.kt)("p",null,"After being redirected back to the Analyze projects page, enter the ",(0,r.kt)("strong",{parentName:"p"},"Display Name")," ",(0,r.kt)("inlineCode",{parentName:"p"},"TruStacks React Sample"),". "),(0,r.kt)("p",null,"This will automatically generate a project key with the format ",(0,r.kt)("inlineCode",{parentName:"p"},"<organization>_trustacks-react-sample"),"."),(0,r.kt)("p",null,"Click ",(0,r.kt)("strong",{parentName:"p"},"Next")," to proceed."),(0,r.kt)("p",null,"On the next page select ",(0,r.kt)("strong",{parentName:"p"},"Previous version")," and then click ",(0,r.kt)("strong",{parentName:"p"},"Create project")))}d.isMDXComponent=!0},6683:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/sonar-create-project-manually-dcc086e0415c5022c97eaed1c91e5a6c.png"}}]);