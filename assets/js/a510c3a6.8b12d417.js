"use strict";(self.webpackChunkcompass=self.webpackChunkcompass||[]).push([[300],{3905:function(e,n,t){t.d(n,{Zo:function(){return u},kt:function(){return m}});var r=t(7294);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function s(e,n){if(null==e)return{};var t,r,a=function(e,n){if(null==e)return{};var t,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(a[t]=e[t]);return a}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var l=r.createContext({}),c=function(e){var n=r.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):o(o({},n),e)),t},u=function(e){var n=c(e.components);return r.createElement(l.Provider,{value:n},e.children)},p={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},d=r.forwardRef((function(e,n){var t=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),d=c(t),m=a,h=d["".concat(l,".").concat(m)]||d[m]||p[m]||i;return t?r.createElement(h,o(o({ref:n},u),{},{components:t})):r.createElement(h,o({ref:n},u))}));function m(e,n){var t=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var i=t.length,o=new Array(i);o[0]=d;var s={};for(var l in n)hasOwnProperty.call(n,l)&&(s[l]=n[l]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var c=2;c<i;c++)o[c]=t[c];return r.createElement.apply(null,o)}return r.createElement.apply(null,t)}d.displayName="MDXCreateElement"},4980:function(e,n,t){t.r(n),t.d(n,{assets:function(){return u},contentTitle:function(){return l},default:function(){return m},frontMatter:function(){return s},metadata:function(){return c},toc:function(){return p}});var r=t(7462),a=t(3366),i=(t(7294),t(3905)),o=["components"],s={},l="Internals",c={unversionedId:"concepts/internals",id:"concepts/internals",title:"Internals",description:"This document details information about how Compass interfaces with elasticsearch. It is meant to give an overview of how some concepts work internally, to help streamline understanding of how things work under the hood.",source:"@site/docs/concepts/internals.md",sourceDirName:"concepts",slug:"/concepts/internals",permalink:"/compass/concepts/internals",draft:!1,editUrl:"https://github.com/odpf/compass/edit/master/docs/docs/concepts/internals.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Architecture",permalink:"/compass/concepts/architecture"},next:{title:"Compass",permalink:"/compass/reference/api"}},u={},p=[{value:"Index Setup",id:"index-setup",level:2},{value:"Search",id:"search",level:2}],d={toc:p};function m(e){var n=e.components,t=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,r.Z)({},d,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"internals"},"Internals"),(0,i.kt)("p",null,"This document details information about how Compass interfaces with elasticsearch. It is meant to give an overview of how some concepts work internally, to help streamline understanding of how things work under the hood."),(0,i.kt)("h2",{id:"index-setup"},"Index Setup"),(0,i.kt)("p",null,"There is a migration command in compass to setup all storages. Once the migration is executed, all types are being created (if does not exist). When a type is created, an index is created in elasticsearch by it's name. All created indices are aliased to the ",(0,i.kt)("inlineCode",{parentName:"p"},"universe")," index, which is used to run the search when all types need to be searched, or when ",(0,i.kt)("inlineCode",{parentName:"p"},"filter[type]")," is not specifed in the Search API."),(0,i.kt)("p",null,"The indices are also configured with a camel case tokenizer, to support proper lexing of some resources that use camel case in their nomenclature ","(","protobuf names for instance",")",". Given below is a sample of the index settings that are used:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-javascript"},'// PUT http://${ES_HOST}/{index}\n{\n        "mappings": {},         // used for boost\n        "aliases": {            // all indices are aliased to the "universe" index\n            "universe": {} \n        },\n        "settings": {           // configuration for handling camel case text\n            "analysis": {\n                "analyzer": {\n                    "default": {\n                        "type": "pattern",\n                        "pattern": "([^\\\\p{L}\\\\d]+)|(?<=\\\\D)(?=\\\\d)|(?<=\\\\d)(?=\\\\D)|(?<=[\\\\p{L}&&[^\\\\p{Lu}]])(?=\\\\p{Lu})|(?<=\\\\p{Lu})(?=\\\\p{Lu}[\\\\p{L}&&[^\\\\p{Lu}]])"\n                    }\n                }\n            }\n        }\n    }\n')),(0,i.kt)("h2",{id:"search"},"Search"),(0,i.kt)("p",null,"We use elasticsearch's ",(0,i.kt)("inlineCode",{parentName:"p"},"multi_match")," search for running our queries. Depending on whether there are additional filter's specified during search, we augment the query with a custom script query that filter's the result set."),(0,i.kt)("p",null,"The script filter is designed to match a document if:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"the document contains the filter key and it's value matches the filter value OR"),(0,i.kt)("li",{parentName:"ul"},"the document doesn't contain the filter key at all")),(0,i.kt)("p",null,"To demonstrate, the following API call:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},"$ curl http://localhost:8080/v1beta1/search?text=log&filter[landscape]=id\n")),(0,i.kt)("p",null,"is internally translated to the following elasticsearch query"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-javascript"},'{\n    "query": {\n        "bool": {\n            "must": {\n                "multi_match": {\n                    "query": "log"\n                }\n            },\n            "filter": [{\n                "script": {\n                    "script": {\n                        "source": "doc.containsKey(\\"landscape.keyword\\") == false || doc[\\"landscape.keyword\\"].value == \\"id\\""\n                    }\n                }\n            }]\n        }\n    }\n}\n')),(0,i.kt)("p",null,"Compass also supports filter with fuzzy match with ",(0,i.kt)("inlineCode",{parentName:"p"},"query")," query params. The script query is designed to match a document if:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"the document contains the filter key and it's value is fuzzily matches the ",(0,i.kt)("inlineCode",{parentName:"li"},"query")," value")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},"$ curl http://localhost:8080/v1beta1/search?text=log&filter[landscape]=id\n")),(0,i.kt)("p",null,"is internally translated to the following elasticsearch query"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-javascript"},'{\n   "query":{\n      "bool":{\n         "filter":{\n            "match":{\n               "description":{\n                  "fuzziness":"AUTO",\n                  "query":"test"\n               }\n            }\n         },\n         "should":{\n            "bool":{\n               "should":[\n                  {\n                     "multi_match":{\n                        "fields":[\n                           "urn^10",\n                           "name^5"\n                        ],\n                        "query":"log"\n                     }\n                  },\n                  {\n                     "multi_match":{\n                        "fields":[\n                           "urn^10",\n                           "name^5"\n                        ],\n                        "fuzziness":"AUTO",\n                        "query":"log"\n                     }\n                  },\n                  {\n                     "multi_match":{\n                        "fields":[\n                           \n                        ],\n                        "fuzziness":"AUTO",\n                        "query":"log"\n                     }\n                  }\n               ]\n            }\n         }\n      }\n   },\n   "min_score":0.01\n}\n')))}m.isMDXComponent=!0}}]);