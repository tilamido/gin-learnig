(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["bbs-love-my"],{"041d":function(t,i,s){"use strict";var e=s("320c"),n=s.n(e);n.a},"2d83":function(t,i,s){var e=s("24fb");i=e(!1),i.push([t.i,".container[data-v-5f124fb9]{padding:20px;font-size:14px;line-height:24px}",""]),t.exports=i},"320c":function(t,i,s){var e=s("2d83");e.__esModule&&(e=e.default),"string"===typeof e&&(e=[[t.i,e,""]]),e.locals&&(t.exports=e.locals);var n=s("4f06").default;n("7e07217e",e,!0,{sourceMap:!1,shadowMode:!1})},6038:function(t,i,s){"use strict";s.d(i,"b",(function(){return e})),s.d(i,"c",(function(){return n})),s.d(i,"a",(function(){}));var e=function(){var t=this,i=t.$createElement,s=t._self._c||i;return s("v-uni-view",{},[t.unLogin?s("un-login"):s("v-uni-view",[0==Object.keys(t.list).length?s("v-uni-view",{staticClass:"emptyData"},[t._v("暂无数据")]):t._e(),s("v-uni-view",{staticClass:"sglist"},t._l(t.list,(function(i,e){return s("v-uni-view",{key:e,staticClass:"sglist-item"},[s("v-uni-view",{staticClass:"sglist-title",on:{click:function(s){arguments[0]=s=t.$handleEvent(s),t.goDetail(i._id)}}},[t._v(t._s(i.title))]),s("v-uni-view",{staticClass:"sglist-imglist",on:{click:function(s){arguments[0]=s=t.$handleEvent(s),t.goDetail(i._id)}}},t._l(i.imgList,(function(t,i){return s("v-uni-image",{key:i,staticClass:"sglist-imglist-img",attrs:{mode:"widthFix",src:t.imgurl+"?x-oss-process=image/resize,m_fixed,h_100,w_100"}})})),1),s("v-uni-view",{staticClass:"sglist-desc",on:{click:function(s){arguments[0]=s=t.$handleEvent(s),t.goDetail(i._id)}}},[t._v(t._s(i.content))])],1)})),1)],1)],1)},n=[]},7776:function(t,i,s){"use strict";s.r(i);var e=s("d9cf"),n=s.n(e);for(var a in e)["default"].indexOf(a)<0&&function(t){s.d(i,t,(function(){return e[t]}))}(a);i["default"]=n.a},a710:function(t,i,s){"use strict";s.r(i);var e=s("6038"),n=s("7776");for(var a in n)["default"].indexOf(a)<0&&function(t){s.d(i,t,(function(){return n[t]}))}(a);s("041d");var u=s("f0c5"),r=Object(u["a"])(n["default"],e["b"],e["c"],!1,null,"5f124fb9",null,!1,e["a"],void 0);i["default"]=r.exports},d9cf:function(t,i,s){"use strict";(function(t){s("7a82");var e=s("ee27").default;Object.defineProperty(i,"__esModule",{value:!0}),i.default=void 0,s("14d9");var n=e(s("f07e")),a=e(s("c964")),u={data:function(){return{pageLoad:!1,list:[],isFirst:!0,start:0,limit:4,unLogin:!0,ssuserid:""}},onLoad:function(){this.getPage()},onPullDownRefresh:function(){this.getPage(),uni.stopPullDownRefresh()},onReachBottom:function(){this.getList()},methods:{goDetail:function(t){uni.navigateTo({url:"../topic/show?id="+t})},getPage:function(){var i=(0,a.default)((0,n.default)().mark((function i(){var s;return(0,n.default)().wrap((function(i){while(1)switch(i.prev=i.next){case 0:if(s=this,this.ssuserid=getApp().globalData.ssuserid,this.unLogin=getApp().globalData.unLogin,""!=this.ssuserid){i.next=5;break}return i.abrupt("return",!1);case 5:t.callFunction({name:"bbs_love",data:{cloudAction:"my",params:{ssuserid:this.ssuserid,start:0,limit:this.limit}}}).then((function(t){var i=t.result.data;s.list=i.list,s.isFirst=!1,s.start=s.start+s.limit,s.pageLoad=!0}));case 6:case"end":return i.stop()}}),i,this)})));return function(){return i.apply(this,arguments)}}(),getList:function(){var i=this;if(0==i.start&&!i.isFirst)return!1;t.callFunction({name:"bbs_love",data:{cloudAction:"my",params:{ssuserid:this.ssuserid,start:this.start,limit:this.limit}}}).then((function(t){var s=t.result.data;if(i.isFirst)i.list=s.list,i.isFirst=!1;else{if(console.log("getList"),console.log(i.start,s.list),0==s.list.length)return i.start=0,!1;for(var e in s.list)i.list.push(s.list[e])}i.start=i.start+i.limit}))}}};i.default=u}).call(this,s("a9ff")["default"])}}]);