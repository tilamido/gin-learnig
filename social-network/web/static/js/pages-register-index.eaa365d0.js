(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-register-index"],{"2e3c":function(t,e,a){var n=a("24fb");e=n(!1),e.push([t.i,".yzmDisable[data-v-462b852d]{background-color:#eee;color:#999}.regBox[data-v-462b852d]{position:absolute;top:50%;left:%?30?%;right:%?30?%;margin-top:%?-390?%;padding:%?30?% %?20?%;background-color:#fff;border-radius:%?20?%}.regBg[data-v-462b852d]{background:linear-gradient(#29cee8,#619ad6);position:absolute;top:%?0?%;bottom:%?0?%;left:%?0?%;right:%?0?%}",""]),t.exports=e},"43db":function(t,e,a){"use strict";a.d(e,"b",(function(){return n})),a.d(e,"c",(function(){return i})),a.d(e,"a",(function(){}));var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-uni-view",{staticClass:"regBg"},[t.pageLoad?a("v-uni-view",{staticClass:"regBox"},[a("v-uni-form",{on:{submit:function(e){arguments[0]=e=t.$handleEvent(e),t.formSubmit.apply(void 0,arguments)}}},[a("v-uni-view",{staticClass:"input-flex"},[a("v-uni-view",{staticClass:"input-flex-label"},[t._v("昵称")]),a("v-uni-input",{staticClass:"input-flex-text",attrs:{name:"nickname",type:"text",placeholder:"请填写昵称"}})],1),a("v-uni-view",{staticClass:"input-flex"},[a("v-uni-view",{staticClass:"input-flex-label"},[t._v("密码")]),a("v-uni-input",{staticClass:"input-flex-text",attrs:{name:"password",type:"text",password:"true",placeholder:"请填写密码"}})],1),a("v-uni-view",{staticClass:"input-flex"},[a("v-uni-view",{staticClass:"input-flex-label"},[t._v("确认密码")]),a("v-uni-input",{staticClass:"input-flex-text",attrs:{name:"password2",type:"text",password:"true",placeholder:"请填写密码"}})],1),a("v-uni-button",{staticClass:"btn-row-submit btn-success",attrs:{type:"primary","form-type":"submit"}},[t._v("立即注册")])],1)],1):t._e()],1)},i=[]},6182:function(t,e,a){"use strict";var n=a("99bb"),i=a.n(n);i.a},"99bb":function(t,e,a){var n=a("2e3c");n.__esModule&&(n=n.default),"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var i=a("4f06").default;i("3a40d888",n,!0,{sourceMap:!1,shadowMode:!1})},a3a5:function(t,e,a){"use strict";a.r(e);var n=a("43db"),i=a("d678");for(var s in i)["default"].indexOf(s)<0&&function(t){a.d(e,t,(function(){return i[t]}))}(s);a("6182");var o=a("f0c5"),u=Object(o["a"])(i["default"],n["b"],n["c"],!1,null,"462b852d",null,!1,n["a"],void 0);e["default"]=u.exports},cf39:function(t,e,a){"use strict";a("7a82"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n={data:function(){return{pageLoad:!1,pageData:{}}},onLoad:function(t){this.pageLoad=!0},methods:{formSubmit:function(t){if(t.detail.value.username=t.detail.value.nickname,t.detail.value.password2!=t.detail.value.password)return uni.showToast({icon:"none",title:"两次密码不一样"}),!1;uni.request({url:this.app.apiHost+"/user/register",method:"POST",data:{username:t.detail.value.username,password:t.detail.value.password,confirmpassword:t.detail.value.password2},success:function(t){console.log(111,t),0==t.data.code?uni.showModal({showCancel:!1,content:"注册成功",success:function(t){t.confirm&&uni.redirectTo({url:"/pages/login/index"})}}):uni.showModal({showCancel:!1,content:"注册失败:"+t.data.msg})},fail:function(t){console.log(t)}})}}};e.default=n},d678:function(t,e,a){"use strict";a.r(e);var n=a("cf39"),i=a.n(n);for(var s in n)["default"].indexOf(s)<0&&function(t){a.d(e,t,(function(){return n[t]}))}(s);e["default"]=i.a}}]);