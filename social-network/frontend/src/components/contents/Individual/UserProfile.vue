<template>
    <div class="user-profile">
      <h2>个人信息</h2>
      <div class="profile-info">
        <p>用户名: {{ user.username }}</p>
        <p>邮箱: {{ user.email }}</p>
        <!-- 其他个人信息字段 -->
      </div>
    </div>
</template>
  
  <script>
  export default {
    data() {
      return {
        user: {
          username: '用户名',
          email: '邮箱',
          // 其他个人信息字段
        }
      };
    }
  };
  </script>
  
  <style scoped>
  /* 样式可以根据您的需求进行自定义 */
  .user-profile {
    margin-bottom: 20px;
  }
  .profile-info {
    border: 1px solid #ccc;
    padding: 10px;
  }
  </style>
  