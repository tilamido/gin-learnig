<template>
    <sidebar :isClose="isClose" @toggleClose="toggleClose" :isDark="isDark" @toggleDark="toggleDark"></sidebar>
    <!-- <contentcomponent></contentcomponent> -->
    <div :style="{ marginLeft: contentMargin }">
        <router-view :class="{ 'close': isClose,'open':!isClose,'dark':isDark }"></router-view>
    </div>
    <!-- 通过isClose设置宽度 -->
</template>
<script>
import sidebar from '../components/SideBar.vue';

export default {
    name: 'HomePage',
    components: {
        sidebar,
    },
    data() {
        return {
            // msg: 'this is message',
            // condition:true
            isClose: true,
            isDark:false,
        };
    },
    computed: {
        contentMargin() {
            return this.isClose ? '88px': '250px';
        }
    },
    methods: {
        // handleClick:function () {
        //     this.condition = !this.condition;
        // }
        toggleClose() {
            this.isClose = !this.isClose; // 切换展开和收起状态
        },
        toggleDark() {
            console.log(this.isDark);
            this.isDark = !this.isDark; // 切换日间和黑夜状态
        }
    },
};
</script>
<style scoped>
.content{
    transition: var(--tran-05);
}
.close {
  /* 默认样式 */
  width: calc(100% - 88px);
}

.open {
  /* 关闭状态下的样式 */
  width: calc(100% - 250px);
}

.dark{
    color: var(--text-color);
}

</style>