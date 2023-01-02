"use strict";(self.webpackChunkcompass=self.webpackChunkcompass||[]).push([[737],{3905:function(e,t,n){n.d(t,{Zo:function(){return u},kt:function(){return d}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var p=r.createContext({}),c=function(e){var t=r.useContext(p),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},u=function(e){var t=c(e.components);return r.createElement(p.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,p=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),m=c(n),d=a,f=m["".concat(p,".").concat(d)]||m[d]||l[d]||i;return n?r.createElement(f,o(o({ref:t},u),{},{components:n})):r.createElement(f,o({ref:t},u))}));function d(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=m;var s={};for(var p in t)hasOwnProperty.call(t,p)&&(s[p]=t[p]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var c=2;c<i;c++)o[c]=n[c];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},748:function(e,t,n){n.r(t),n.d(t,{assets:function(){return u},contentTitle:function(){return p},default:function(){return d},frontMatter:function(){return s},metadata:function(){return c},toc:function(){return l}});var r=n(7462),a=n(3366),i=(n(7294),n(3905)),o=["components"],s={},p="User",c={unversionedId:"concepts/user",id:"concepts/user",title:"User",description:"The current version of Compass does not have user management. Compass expects there is an external instance that manages user. Compass consumes user information from the configurable identity uuid header in every API call. The default name of the header is Compass-User-UUID.",source:"@site/docs/concepts/user.md",sourceDirName:"concepts",slug:"/concepts/user",permalink:"/compass/concepts/user",draft:!1,editUrl:"https://github.com/odpf/compass/edit/master/docs/docs/concepts/user.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Type",permalink:"/compass/concepts/type"},next:{title:"Architecture",permalink:"/compass/concepts/architecture"}},u={},l=[{value:"Phantom User",id:"phantom-user",level:2},{value:"Linking User",id:"linking-user",level:2},{value:"User Provider",id:"user-provider",level:2}],m={toc:l};function d(e){var t=e.components,n=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"user"},"User"),(0,i.kt)("p",null,"The current version of Compass does not have user management. Compass expects there is an external instance that manages user. Compass consumes user information from the configurable identity uuid header in every API call. The default name of the header is ",(0,i.kt)("inlineCode",{parentName:"p"},"Compass-User-UUID"),".\nCompass does not make any assumption of what kind of identity format that is being used. The ",(0,i.kt)("inlineCode",{parentName:"p"},"uuid")," indicates that it could be in any form (e.g. email, UUIDv4, etc) as long as it is universally unique.\nThe current behaviour is, Compass will add a new user if the user information consumed from the header does not exist in Compass' database. "),(0,i.kt)("h2",{id:"phantom-user"},"Phantom User"),(0,i.kt)("p",null,"In Compass ingestion API, Compass allows asset to mentioned who is its own owners. During the ingestion, if the ",(0,i.kt)("inlineCode",{parentName:"p"},"email")," field in the list of ",(0,i.kt)("inlineCode",{parentName:"p"},"owners")," field in the asset is not empty, Compass will create a new ",(0,i.kt)("inlineCode",{parentName:"p"},"'Phantom User'")," with the email but with empty UUID.\nA ",(0,i.kt)("inlineCode",{parentName:"p"},"'Phantom User'")," is a user that is written in the storage but with empty UUID. The ",(0,i.kt)("inlineCode",{parentName:"p"},"'Phantom User'")," cannot do any user-related interaction (e.g. Starring, Discussion) in Compass."),(0,i.kt)("h2",{id:"linking-user"},"Linking User"),(0,i.kt)("p",null,"There is another configurable optional email header that Compass expect. The default name is ",(0,i.kt)("inlineCode",{parentName:"p"},"Compass-User-Email"),". In case there is already an existing ",(0,i.kt)("inlineCode",{parentName:"p"},"'Phantom User'"),", if there is a request coming to Compass with completed user information in its header (uuid header and email header are not empty), Compass will register the UUID to the existing ",(0,i.kt)("inlineCode",{parentName:"p"},"'Phantom User'")," and the ",(0,i.kt)("inlineCode",{parentName:"p"},"'Phantom User'")," becomes a normal user. By doing so, assets ownership of that new user will immediately reflected."),(0,i.kt)("h2",{id:"user-provider"},"User Provider"),(0,i.kt)("p",null,"Since Compass expects that there is an external instance that manages user, it is possible for Compass to consume user information from multiple external instances. Compass distinguishes the source of user by marking it in the ",(0,i.kt)("inlineCode",{parentName:"p"},"provider")," field. The default ",(0,i.kt)("inlineCode",{parentName:"p"},"provider")," field value can be configured via config."))}d.isMDXComponent=!0}}]);