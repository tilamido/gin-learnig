(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-user-index"],{"101c":function(t,e,i){"use strict";i.r(e);var n=i("f0a2"),a=i.n(n);for(var o in n)["default"].indexOf(o)<0&&function(t){i.d(e,t,(function(){return n[t]}))}(o);e["default"]=a.a},"286f":function(t,e,i){"use strict";var n=i("946e"),a=i.n(n);a.a},"37e7":function(t,e,i){"use strict";i.r(e);var n=i("c49b"),a=i.n(n);for(var o in n)["default"].indexOf(o)<0&&function(t){i.d(e,t,(function(){return n[t]}))}(o);e["default"]=a.a},"49da":function(t,e,i){"use strict";var n=i("a616"),a=i.n(n);a.a},5408:function(t,e,i){var n=i("24fb");e=n(!1),e.push([t.i,".uhead[data-v-05d65882]{display:flex;padding:%?22?%;background-color:#fff;margin-bottom:%?11?%;align-items:center}.uhead-img[data-v-05d65882]{width:%?172?%;height:%?172?%;margin-right:%?22?%;display:block;border-radius:50%}.uhead-box[data-v-05d65882]{flex:1}.uhead-nick[data-v-05d65882]{margin-top:%?10?%;margin-bottom:%?10?%;font-size:%?40?%;font-weight:700}.uhead-rnum[data-v-05d65882]{color:#999;margin-bottom:%?32?%;line-height:%?29?%;display:flex;font-size:%?29?%}.row-item-icon[data-v-05d65882]{font-size:%?42?%}.cl-u[data-v-05d65882], .cl-u[data-v-05d65882]:before{color:#ed6d53}",""]),t.exports=e},"57b3":function(t,e,i){t.exports=i.p+"static/user_head.jpg"},"67b1":function(t,e,i){"use strict";i.r(e);var n=i("deea"),a=i("101c");for(var o in a)["default"].indexOf(o)<0&&function(t){i.d(e,t,(function(){return a[t]}))}(o);i("286f");var s=i("f0c5"),u=Object(s["a"])(a["default"],n["b"],n["c"],!1,null,"05d65882",null,!1,n["a"],void 0);e["default"]=u.exports},"68d4":function(t,e,i){"use strict";i.r(e);var n=i("71fd"),a=i("37e7");for(var o in a)["default"].indexOf(o)<0&&function(t){i.d(e,t,(function(){return a[t]}))}(o);i("49da");var s=i("f0c5"),u=Object(s["a"])(a["default"],n["b"],n["c"],!1,null,"3270db79",null,!1,n["a"],void 0);e["default"]=u.exports},"71fd":function(t,e,i){"use strict";i.d(e,"b",(function(){return n})),i.d(e,"c",(function(){return a})),i.d(e,"a",(function(){}));var n=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",[i("v-uni-view",{staticClass:"footer-row"}),i("v-uni-view",{staticClass:"footer"},[i("v-uni-view",{staticClass:"footer-item icon-home",class:{"footer-active":"index"==t.tab},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.goHome()}}},[t._v("首页")]),i("v-uni-view",{staticClass:"footer-item footer-add icon-edit_light",class:{"footer-active":"add"==t.tab},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.goAdd()}}},[t._v("发布")]),i("v-uni-view",{staticClass:"footer-item icon-my_light",class:{"footer-active":"my"==t.tab},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.goUser()}}},[t._v("我的")])],1)],1)},a=[]},"946e":function(t,e,i){var n=i("5408");n.__esModule&&(n=n.default),"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var a=i("4f06").default;a("2e5b0c3d",n,!0,{sourceMap:!1,shadowMode:!1})},a616:function(t,e,i){var n=i("f6cf");n.__esModule&&(n=n.default),"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var a=i("4f06").default;a("35e7a973",n,!0,{sourceMap:!1,shadowMode:!1})},c49b:function(t,e,i){"use strict";i("7a82"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n={props:{tab:""},data:function(){return{}},methods:{goHome:function(){console.log("goHone"),uni.reLaunch({url:"/pages/index/index"})},goAdd:function(){console.log("add"),uni.navigateTo({url:"/bbs/topic/add",fail:function(t){console.log(t)}})},goUser:function(){uni.reLaunch({url:"/pages/user/index"})}}};e.default=n},deea:function(t,e,i){"use strict";i.d(e,"b",(function(){return n})),i.d(e,"c",(function(){return a})),i.d(e,"a",(function(){}));var n=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",[t.unLogin?n("un-login"):n("v-uni-view",[n("v-uni-view",{staticClass:"main-body"},[n("v-uni-view",{staticClass:"uhead"},[n("v-uni-view",[n("v-uni-image",{staticClass:"uhead-img",attrs:{src:i("57b3")}})],1),n("v-uni-view",{staticClass:"uhead-box"},[n("v-uni-view",{staticClass:"uhead-nick"},[t._v(t._s(t.ssusername))])],1)],1),n("v-uni-view",{staticClass:"row-box"},[n("v-uni-view",{staticClass:"row-item",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.gourl("../../bbs/topic/my")}}},[n("v-uni-view",{staticClass:"row-item-title"},[t._v("我的帖子")])],1),n("v-uni-view",{staticClass:"row-item",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.gourl("../../bbs/topic/ranked")}}},[n("v-uni-view",{staticClass:"row-item-title"},[t._v("点赞排行")])],1),n("v-uni-view",{staticClass:"row-item",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.gourl("/pages/user/password")}}},[n("v-uni-view",{staticClass:"row-item-title"},[t._v("修改密码")])],1),n("v-uni-view",{staticClass:"row-item",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.logOut()}}},[n("v-uni-view",{staticClass:"row-item-title"},[t._v("退出登录")])],1)],1)],1)],1),n("mt-footer",{attrs:{tab:"user"}})],1)},a=[]},f0a2:function(t,e,i){"use strict";i("7a82");var n=i("ee27").default;Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=n(i("68d4")),o={components:{mtFooter:a.default},data:function(){return{pageLoad:!0,pageHide:!1,user:{},unLogin:!0,ssuserid:"",ssusername:""}},onLoad:function(t){if(uni.setNavigationBarTitle({title:"个人中心"}),this.ssuserid=getApp().globalData.ssuserid,this.ssusername=getApp().globalData.ssusername,this.unLogin=getApp().globalData.unLogin,""==this.ssuserid)return!1;this.getPage()},onShow:function(){if(this.ssuserid=getApp().globalData.ssuserid,this.ssusername=getApp().globalData.ssusername,this.unLogin=getApp().globalData.unLogin,""==this.ssuserid)return!1;this.getPage()},onHide:function(){},methods:{gourl:function(t){uni.navigateTo({url:t})},getPage:function(){},logOut:function(){var t=this;uni.showModal({title:"提示",content:"确定要退出登录吗？",success:function(e){e.confirm?(uni.setStorageSync("uniIdToken",""),getApp().globalData.ssuserid="",getApp().globalData.ssusername="",getApp().globalData.unLogin=!0,t.gourl("/pages/login/index")):e.cancel&&console.log("用户点击取消")}})}}};e.default=o},f6cf:function(t,e,i){var n=i("24fb");e=n(!1),e.push([t.i,".footer-active[data-v-3270db79]:after{color:#fff;background-color:#007aff}",""]),t.exports=e}}]);