<template>
   <nav class="sidebar close">
  <header>
    <div class="image-text"> 
      <!-- <a href="#" class="text-logo-a" target="_blank"> -->
        <span class="image">
          <img src="#" alt="">
        </span>
        <div class="text logo-text">
          <span class="name">Social-network</span>
        </div>
      <!-- </a> -->
    </div>
    <i class='bx bx-chevron-right toggle'></i>
  </header>
  <div class="menu-bar">
    <div class="menu">
      <!-- <li class="search-box">
        <i class='bx bx-search icon'></i>
        <input type="text" placeholder="Search...">
      </li> -->
      <ul class="menu-links">
        <li v-for="(item, index) in navLinks" :key="index" class="nav-link">
        <!-- <a :href="item.link"> -->
          <router-link :to="{name:item.link}"> 
            <i :class="item.iconClass"></i>
            <span class="text nav-text">{{ item.text }}</span>
          </router-link> 
        <!-- </a> -->
        </li>
      </ul>
    </div>
    <div class="bottom-content">
      <li class="">
        <a href="#" target="_blank">
          <i class='bx bx-log-out icon'></i>
          <span class="text nav-text">Logout</span>
        </a>
      </li>
      <li class="mode">
        <div class="sun-moon">
          <i class='bx bx-moon icon moon'></i>
          <i class='bx bx-sun icon sun'></i>
        </div>
        <span class="mode-text text">Dark mode</span>
        <div class="toggle-switch">
          <span class="switch"></span>
        </div>
      </li>
    </div>
  </div>
</nav>
</template>
<script>

export default {
  name: 'SideBar',
  data() {
    return {
      navLinks: [
        { iconClass: 'bx bx-home-alt icon', text: 'Home', link: 'index' },
        { iconClass: 'bx bx-bar-chart-alt-2 icon', text: 'Hotest', link: 'hotest' },
        { iconClass: 'bx bx-plus-circle icon', text: 'Publish', link: 'publish' },
        { iconClass: 'bx bx-user icon', text: 'Individual', link: 'individual' },
        // { iconClass: 'bx bx-heart icon', text: 'Likes', link: 'individual' },
        // { iconClass: 'bx bx-wallet icon', text: 'Wallets', link: 'individual' }
      ]
    };
  },
  methods: {
    // increment() {
    //   this.count++;
    // }
    toggleClose:function() {
      const body = document.querySelector("body");
      const sidebar = body.querySelector("nav");
      sidebar.classList.toggle("close");
      this.$emit('toggleClose'); // 发送事件通知父组件切换展开和收起状态
    },
    toggleDark: function () {
      const body = document.querySelector("body"),
      modeText = body.querySelector(".mode-text");
      body.classList.toggle("dark");
      if (body.classList.contains("dark")) {
        modeText.innerText = "Light mode";
      } else {
        modeText.innerText = "Dark mode";
      }
      this.$emit('toggleDark');
    }
  },
  // 在组件中引入需要的外部样式表和脚本
  mounted() {
    // 添加外部 CSS 样式表
    const boxiconsCss = document.createElement('link');
    boxiconsCss.rel = 'stylesheet';
    boxiconsCss.href = 'https://unpkg.com/boxicons@2.1.1/css/boxicons.min.css';
    document.head.appendChild(boxiconsCss);

    const fontawesomeCss = document.createElement('link');
    fontawesomeCss.rel = 'stylesheet';
    fontawesomeCss.href = 'https://pro.fontawesome.com/releases/v6.0.0-beta3/css/all.css';
    document.head.appendChild(fontawesomeCss);

    // 添加 Google Fonts
    const googleFontsCss = document.createElement('link');
    googleFontsCss.rel = 'stylesheet';
    googleFontsCss.href = 'https://fonts.googleapis.com/css2?family=Outfit:wght@100;200;300;400;500;600;700;800;900&amp;display=swap';
      document.head.appendChild(googleFontsCss);

    const body = document.querySelector("body"),
    toggle = body.querySelector(".toggle"),
    // searchBtn = body.querySelector(".search-box"),
    modeSwitch = body.querySelector(".toggle-switch");
    
    toggle.addEventListener("click", this.toggleClose,false);
    
    // searchBtn.addEventListener("click", () => {
    //   sidebar.classList.remove("close");
    // });
    
    modeSwitch.addEventListener("click", this.toggleDark,false);
  }
  
};
</script>
<style scoped>
.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  width: 250px;
  padding: 10px 14px;
  background: var(--sidebar-color);
  transition: var(--tran-05);
  z-index: 100;
}
.sidebar.close {
  width: 88px;
}
.sidebar li {
  height: 50px;
  list-style: none;
  display: flex;
  align-items: center;
  margin-top: 10px;
}
.sidebar header .image,
.sidebar .icon {
  min-width: 60px;
  border-radius: 6px;
}
.sidebar .icon {
  min-width: 60px;
  border-radius: 6px;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}
.sidebar .text,
.sidebar .icon {
  color: var(--text-color);
  transition: var(--tran-03);
}
.sidebar .text {
  font-size: 17px;
  font-weight: 500;
  white-space: nowrap;
  opacity: 1;
}
.sidebar.close .text {
  opacity: 0;
}
.sidebar header {
  position: relative;
}
.sidebar header .image-text {
  display: flex;
  align-items: center;
}
.sidebar header .logo-text {
  display: flex;
  flex-direction: column;
}
header .image-text .name {
  margin-top: 2px;
  font-size: 18px;
  font-weight: 600;
}
header .image-text .profession {
  font-size: 16px;
  margin-top: -2px;
  display: block;
}
.sidebar header .image {
  display: flex;
  align-items: center;
  justify-content: center;
}
.sidebar header .image img {
  width: 40px;
  border-radius: 6px;
}
.sidebar header .toggle {
  position: absolute;
  top: 50%;
  right: -25px;
  transform: translateY(-50%) rotate(180deg);
  height: 25px;
  width: 25px;
  background-color: var(--primary-color);
  color: var(--sidebar-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  cursor: pointer;
  transition: var(--tran-05);
}

.sidebar.close .toggle {
  transform: translateY(-50%) rotate(0deg);
}
.sidebar .menu {
  margin-top: 40px;
}
/* .sidebar li.search-box {
  border-radius: 6px;
  background-color: var(--primary-color-light);
  cursor: pointer;
  border: 1px solid #dadce0;
  transition: var(--tran-05);
}
.sidebar li.search-box input {
  height: 100%;
  width: 100%;
  outline: none;
  border: none;
  background-color: var(--primary-color-light);
  color: var(--text-color);
  border-radius: 6px;
  font-size: 17px;
  font-weight: 500;
  transition: var(--tran-05);
} */
.sidebar li a {
  list-style: none;
  height: 100%;
  background-color: transparent;
  display: flex;
  align-items: center;
  height: 100%;
  width: 100%;
  border-radius: 6px;
  text-decoration: none;
  transition: var(--tran-03);
}
.sidebar li a:hover {
  background-color: var(--primary-color);
}
.sidebar li a:hover .icon,
.sidebar li a:hover .text {
  color: var(--sidebar-color);
}
.sidebar .menu-bar {
  height: calc(100% - 55px);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  overflow-y: scroll;
}
.menu-bar::-webkit-scrollbar {
  display: none;
}
.sidebar .menu-bar .mode {
  border-radius: 6px;
  background-color: var(--primary-color-light);
  position: relative;
  transition: var(--tran-05);
}
.menu-bar .mode .sun-moon {
  height: 50px;
  width: 60px;
}
.mode .sun-moon i {
  position: absolute;
}
.mode .sun-moon i.sun {
  opacity: 0;
}
body.dark .mode .sun-moon i.sun {
  opacity: 1;
}
body.dark .mode .sun-moon i.moon {
  opacity: 0;
}
.menu-bar .bottom-content .toggle-switch {
  position: absolute;
  right: 0;
  height: 100%;
  min-width: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 6px;
  cursor: pointer;
}
.toggle-switch .switch {
  position: relative;
  height: 22px;
  width: 40px;
  border-radius: 25px;
  background-color: var(--toggle-color);
  transition: var(--tran-05);
}
.switch::before {
  content: "";
  position: absolute;
  height: 15px;
  width: 15px;
  border-radius: 50%;
  top: 50%;
  left: 5px;
  transform: translateY(-50%);
  background-color: var(--sidebar-color);
  transition: var(--tran-04);
}
body.dark .switch::before {
  left: 20px;
}
.home {
  position: absolute;
  top: 0;
  top: 0;
  left: 250px;
  height: 100vh;
  width: calc(100% - 250px);
  background-color: var(--body-color);
  transition: var(--tran-05);
}
.home .text {
  font-size: 30px;
  font-weight: 500;
  color: var(--text-color);
  padding: 12px 60px;
}
.sidebar.close ~ .home {
  left: 78px;
  height: 100vh;
  width: calc(100% - 78px);
}
body.dark .home .text {
  color: var(--text-color);
}
.text-logo-a {
  text-decoration: none !important;
  color: #1b6ff3;
}

</style>