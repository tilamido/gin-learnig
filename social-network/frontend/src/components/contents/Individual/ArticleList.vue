<template>
    <div class="article-list">
      <h2>文章列表</h2>
      <ul>
        <li v-for="article in articles" :key="article.id">
          <router-link :to="'/article/' + article.id">{{ article.title }}</router-link>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        articles: [
          { id: 1, title: '文章标题1' },
          { id: 2, title: '文章标题2' },
          { id: 3, title: '文章标题3' },
          // 其他文章
        ]
      };
    }
  };
  </script>
  
  <style scoped>
  /* 样式可以根据您的需求进行自定义 */
  .article-list {
    margin-bottom: 20px;
  }
  .article-list ul {
    list-style: none;
    padding: 0;
  }
  .article-list li {
    margin-bottom: 10px;
  }
  </style>
  