"use strict";(self.webpackChunkcompass=self.webpackChunkcompass||[]).push([[915],{3905:function(e,t,a){a.d(t,{Zo:function(){return d},kt:function(){return g}});var n=a(7294);function r(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function i(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function s(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?i(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):i(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function o(e,t){if(null==e)return{};var a,n,r=function(e,t){if(null==e)return{};var a,n,r={},i=Object.keys(e);for(n=0;n<i.length;n++)a=i[n],t.indexOf(a)>=0||(r[a]=e[a]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)a=i[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var l=n.createContext({}),p=function(e){var t=n.useContext(l),a=t;return e&&(a="function"==typeof e?e(t):s(s({},t),e)),a},d=function(e){var t=p(e.components);return n.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},c=n.forwardRef((function(e,t){var a=e.components,r=e.mdxType,i=e.originalType,l=e.parentName,d=o(e,["components","mdxType","originalType","parentName"]),c=p(a),g=r,m=c["".concat(l,".").concat(g)]||c[g]||u[g]||i;return a?n.createElement(m,s(s({ref:t},d),{},{components:a})):n.createElement(m,s({ref:t},d))}));function g(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=a.length,s=new Array(i);s[0]=c;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o.mdxType="string"==typeof e?e:r,s[1]=o;for(var p=2;p<i;p++)s[p]=a[p];return n.createElement.apply(null,s)}return n.createElement.apply(null,a)}c.displayName="MDXCreateElement"},1067:function(e,t,a){a.r(t),a.d(t,{assets:function(){return d},contentTitle:function(){return l},default:function(){return g},frontMatter:function(){return o},metadata:function(){return p},toc:function(){return u}});var n=a(7462),r=a(3366),i=(a(7294),a(3905)),s=["components"],o={},l="Querying  metadata",p={unversionedId:"guides/querying",id:"guides/querying",title:"Querying  metadata",description:"Prerequisites",source:"@site/docs/guides/querying.md",sourceDirName:"guides",slug:"/guides/querying",permalink:"/compass/guides/querying",draft:!1,editUrl:"https://github.com/odpf/compass/edit/master/docs/docs/guides/querying.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Ingesting metadata",permalink:"/compass/guides/ingestion"},next:{title:"Starring",permalink:"/compass/guides/starring"}},d={},u=[{value:"Prerequisites",id:"prerequisites",level:2},{value:"Using the Search API",id:"using-the-search-api",level:2},{value:"Filter",id:"filter",level:3},{value:"Query",id:"query",level:3},{value:"Ranking Results",id:"ranking-results",level:3},{value:"Size",id:"size",level:3},{value:"Using the Suggest API",id:"using-the-suggest-api",level:2},{value:"Using the Get Assets API",id:"using-the-get-assets-api",level:2},{value:"Using the Lineage API",id:"using-the-lineage-api",level:2}],c={toc:u};function g(e){var t=e.components,a=(0,r.Z)(e,s);return(0,i.kt)("wrapper",(0,n.Z)({},c,a,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"querying--metadata"},"Querying  metadata"),(0,i.kt)("h2",{id:"prerequisites"},"Prerequisites"),(0,i.kt)("p",null,"This guide assumes that you have a local instance of compass running and listening on ",(0,i.kt)("inlineCode",{parentName:"p"},"localhost:8080"),". See ",(0,i.kt)("a",{parentName:"p",href:"/compass/installation"},"Installation")," guide for information on how to run Compass."),(0,i.kt)("h2",{id:"using-the-search-api"},"Using the Search API"),(0,i.kt)("p",null,"The API contract is available ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/odpf/compass/blob/main/third_party/OpenAPI/compass.swagger.json"},"here"),"."),(0,i.kt)("p",null,"To demonstrate how to use compass, we\u2019re going to query it for resources that contain the word \u2018booking\u2019."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ curl 'http://localhost:8080/v1beta1/search?text=booking' \\\n--header 'Compass-User-UUID:odpf@email.com' \n")),(0,i.kt)("p",null,"This will return a list of search results. Here\u2019s a sample response:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'{\n    "data": [\n        {\n            "id": "00c06ef7-badb-4236-9d9e-889697cbda46",\n            "urn": "kafka::g-godata-id-playground/ g-godata-id-seg-enriched-booking-dagger",\n            "type": "topic",\n            "service": "kafka",\n            "name": "g-godata-id-seg-enriched-booking-dagger",\n            "description": "",\n            "labels": {\n                "flink_name": "g-godata-id-playground",\n                "sink_type": "kafka"\n            }\n        },\n        {\n            "id": "9e69c08a-c3c2-4e04-957f-c8010c1e6515",\n            "urn": "kafka::g-godata-id-playground/ g-godata-id-booking-bach-test-dagger",\n            "type": "topic",\n            "service": "kafka",\n            "name": "g-godata-id-booking-bach-test-dagger",\n            "description": "",\n            "labels": {\n                "flink_name": "g-godata-id-playground",\n                "sink_type": "kafka"\n            }\n        },\n        {\n            "id": "ff597a0f-8062-4370-a54c-fd6f6c12d2a0",\n            "urn": "kafka::g-godata-id-playground/ g-godata-id-booking-bach-test-3-dagger",\n            "type": "topic",\n            "service": "kafka",\n            "title": "g-godata-id-booking-bach-test-3-dagger",\n            "description": "",\n            "labels": {\n                "flink_name": "g-godata-id-playground",\n                "sink_type": "kafka"\n            }\n        }\n    ]\n}\n')),(0,i.kt)("p",null,"Compass decouple identifier from external system with the one that is being used internally. ID is the internally auto-generated unique identifier. URN is the external identifier of the asset, while Name is the human friendly name for it. See the complete API spec to learn more about what the rest of the fields mean."),(0,i.kt)("h3",{id:"filter"},"Filter"),(0,i.kt)("p",null,"Compass search API also supports restricting search results via filter by passing it in query params.\nFilter query params format is ",(0,i.kt)("inlineCode",{parentName:"p"},"filter[{field_key}]={value}")," where ",(0,i.kt)("inlineCode",{parentName:"p"},"field_key")," is the field name that we want to restrict and ",(0,i.kt)("inlineCode",{parentName:"p"},"value")," is what value that should be matched. Filter could also support nested field by chaining key ",(0,i.kt)("inlineCode",{parentName:"p"},"field_key")," with ",(0,i.kt)("inlineCode",{parentName:"p"},".")," ","(","dot",")"," such as ",(0,i.kt)("inlineCode",{parentName:"p"},"filter[{field_key}.{nested_field_key}]={value}"),". For instance, to restrict search results to the \u2018id\u2019 landscape for \u2018odpf\u2019 organisation, run:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ curl 'http://localhost:8080/v1beta1/search?text=booking&filter[labels.landscape]=vn&filter[labels.entity]=odpf' \\\n--header 'Compass-User-UUID:odpf@email.com'\n")),(0,i.kt)("p",null,"Under the hood, filter's work by checking whether the matching document's contain the filter key and checking if their values match. Filters can be specified multiple times to specify a set of filter criteria. For example, to search for \u2018booking\u2019 in both \u2018vn\u2019 and \u2018th\u2019 landscape, run:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ curl 'http://localhost:8080/v1beta1/search?text=booking&filter[labels.landscape]=id&filter[labels.landscape]=th' \\\n--header 'Compass-User-UUID:odpf@email.com' \n")),(0,i.kt)("h3",{id:"query"},"Query"),(0,i.kt)("p",null,"Apart from filters, Compass search API also supports fuzzy restriction in its query params. The difference of filter and query are, filter is for exact match on a specific field in asset while query is for fuzzy match."),(0,i.kt)("p",null,"Query format is not different with filter ",(0,i.kt)("inlineCode",{parentName:"p"},"query[{field_key}]={value}")," where ",(0,i.kt)("inlineCode",{parentName:"p"},"field_key")," is the field name that we want to query and ",(0,i.kt)("inlineCode",{parentName:"p"},"value")," is what value that should be fuzzy matched. Query could also support nested field by chaining key ",(0,i.kt)("inlineCode",{parentName:"p"},"field_key")," with ",(0,i.kt)("inlineCode",{parentName:"p"},".")," ","(","dot",")"," such as ",(0,i.kt)("inlineCode",{parentName:"p"},"query[{field_key}.{nested_field_key}]={value}"),". For instance, to search results that has a name ",(0,i.kt)("inlineCode",{parentName:"p"},"kafka")," and belongs to the team ",(0,i.kt)("inlineCode",{parentName:"p"},"data_engineering"),", run:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ curl 'http://localhost:8080/v1beta1/search?text=booking&query[name]=kafka&query[labels.team]=data_eng' \\\n--header 'Compass-User-UUID:odpf@email.com' \n")),(0,i.kt)("h3",{id:"ranking-results"},"Ranking Results"),(0,i.kt)("p",null,"Compass allows user to rank the results based on a numeric field in the asset. It supports nested field by using the ",(0,i.kt)("inlineCode",{parentName:"p"},".")," ","(","dot",")"," to point to the nested field. For instance, to rank the search results based on ",(0,i.kt)("inlineCode",{parentName:"p"},"usage_count")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"data")," field, run:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ curl 'http://localhost:8080/v1beta1/search?text=booking&rankby=data.usage_count' \\\n--header 'Compass-User-UUID:odpf@email.com' \n")),(0,i.kt)("h3",{id:"size"},"Size"),(0,i.kt)("p",null,"You can also specify the number of maximum results you want compass to return using the \u2018size\u2019 parameter"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ curl 'http://localhost:8080/v1beta1/search?text=booking&size=5' \\\n--header 'Compass-User-UUID:odpf@email.com' \n")),(0,i.kt)("h2",{id:"using-the-suggest-api"},"Using the Suggest API"),(0,i.kt)("p",null,"The Suggest API gives a number of suggestion based on asset's name. There are 5 suggestions by default return by this API."),(0,i.kt)("p",null,"The API contract is available ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/odpf/compass/blob/main/third_party/OpenAPI/compass.swagger.json"},"here"),"."),(0,i.kt)("p",null,"Example of searching assets suggestion that has a name \u2018booking\u2019."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ curl 'http://localhost:8080/v1beta1/search/suggest?text=booking' \\\n--header 'Compass-User-UUID:odpf@email.com' \n")),(0,i.kt)("p",null,"This will return a list of suggestions. Here\u2019s a sample response:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'{\n    "data": [\n        "booking-daily-test-962ZFY",\n        "booking-daily-test-c7OUZv",\n        "booking-weekly-test-fmDeUf",\n        "booking-daily-test-jkQS2b",\n        "booking-daily-test-m6Oe9M"\n    ]\n}\n')),(0,i.kt)("h2",{id:"using-the-get-assets-api"},"Using the Get Assets API"),(0,i.kt)("p",null,"The Get Assets API returns assets from Compass' main storage (PostgreSQL) while the Search API returns assets from Elasticsearch. The Get Assets API has several options (filters, size, offset, etc...) in its query params."),(0,i.kt)("table",null,(0,i.kt)("thead",{parentName:"table"},(0,i.kt)("tr",{parentName:"thead"},(0,i.kt)("th",{parentName:"tr",align:null},"Query Params"),(0,i.kt)("th",{parentName:"tr",align:null},"Description"))),(0,i.kt)("tbody",{parentName:"table"},(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},(0,i.kt)("inlineCode",{parentName:"td"},"types=topic,table")),(0,i.kt)("td",{parentName:"tr",align:null},"filter by types")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},(0,i.kt)("inlineCode",{parentName:"td"},"services=kafka,postgres")),(0,i.kt)("td",{parentName:"tr",align:null},"filter by services")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},(0,i.kt)("inlineCode",{parentName:"td"},"data[dataset]=booking&data[project]=p-godata-id")),(0,i.kt)("td",{parentName:"tr",align:null},"filter by field in asset.data")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},(0,i.kt)("inlineCode",{parentName:"td"},"q=internal&q_fields=name,urn,description,services")),(0,i.kt)("td",{parentName:"tr",align:null},"querying by field")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},(0,i.kt)("inlineCode",{parentName:"td"},"sort=created_at")),(0,i.kt)("td",{parentName:"tr",align:null},"sort by certain fields")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},(0,i.kt)("inlineCode",{parentName:"td"},"direction=desc")),(0,i.kt)("td",{parentName:"tr",align:null},"sorting direction (asc / desc)")))),(0,i.kt)("p",null,"The API contract is available ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/odpf/compass/blob/main/third_party/OpenAPI/compass.swagger.json"},"here"),"."),(0,i.kt)("h2",{id:"using-the-lineage-api"},"Using the Lineage API"),(0,i.kt)("p",null,"The Lineage API allows the clients to query the data flow relationship between different assets managed by Compass."),(0,i.kt)("p",null,"See the swagger definition of ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/odpf/compass/blob/main/third_party/OpenAPI/compass.swagger.json"},"Lineage API"),") for more information."),(0,i.kt)("p",null,"Lineage API returns a list of directed edges. For each edge, there are ",(0,i.kt)("inlineCode",{parentName:"p"},"source")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"target")," fields that represent nodes to indicate the direction of the edge. Each edge could have an optional property in the ",(0,i.kt)("inlineCode",{parentName:"p"},"props")," field."),(0,i.kt)("p",null,"Here's a sample API call:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'$ curl \'http://localhost:8080/v1beta1/lineage/data-project%3Adatalake.events\' \\\n--header \'Compass-User-UUID:odpf@email.com\' \n\n{\n    data: [\n        {\n            "source": {\n                "urn": "data-project:datalake.events",\n                "type": "table",\n                "service": "bigquery",\n            },\n            "target": {\n                "urn": "events-transform-dwh",\n                "type": "csv",\n                "service": "s3",\n            },\n            "props": nil\n        },\n        {\n            "source": {\n                "urn": "events-ingestion",\n                "type": "topic",\n                "service": "beast",\n            },\n            "target": {\n                "urn": "data-project:datalake.events",\n                "type": "table",\n                "service": "bigquery",\n            },\n            "props": nil\n        },\n    ]\n}\n')),(0,i.kt)("p",null,"The lineage is fetched from the perspective of an asset. The response shows it has a list of upstreams and downstreams assets of the requested asset.\nNotice that in the URL, we are using ",(0,i.kt)("inlineCode",{parentName:"p"},"urn")," instead of ",(0,i.kt)("inlineCode",{parentName:"p"},"id"),". The reason is because we use ",(0,i.kt)("inlineCode",{parentName:"p"},"urn")," as a main identifier in our lineage storage. We don't use ",(0,i.kt)("inlineCode",{parentName:"p"},"id")," to store the lineage as a main identifier, because ",(0,i.kt)("inlineCode",{parentName:"p"},"id")," is internally auto generated and in lineage, there might be some assets that we don't store in our Compass' storage yet."))}g.isMDXComponent=!0}}]);