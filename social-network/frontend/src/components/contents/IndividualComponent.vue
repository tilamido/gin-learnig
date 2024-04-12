<template>
    <div class="user-center">
      <user-profile></user-profile>
      <article-list></article-list>
    </div>
  </template>
  
  <script>
  import UserProfile from './Individual/UserProfile.vue';
  import ArticleList from './Individual/ArticleList.vue';
  
  export default {
    components: {
      UserProfile,
      ArticleList,
    }
  };
  </script>
  
  <style scoped>
  /* 样式可以根据您的需求进行自定义 */
  .user-center {
    padding: 20px;
  }
  </style>
  