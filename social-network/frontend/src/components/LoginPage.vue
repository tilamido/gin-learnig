<template>
  <div class="login-page">
    <div class="form">
      <form class="register-form" v-show="!showLoginForm" @submit.prevent="register">
        <input v-model="username" type="text" placeholder="username" />
        <input v-model="password" type="password" placeholder="password" />
        <input v-model="confirmPassword" type="password" placeholder="enter password again" />
        <button type="submit">create</button>
        <p class="message">Already registered? <a href="#" @click.prevent="toggleForm">Sign In</a></p>
      </form>
      <form class="login-form" v-show="showLoginForm" @submit.prevent="login">
        <input v-model="username" type="text" placeholder="username" />
        <input v-model="password" type="password" placeholder="password" />
        <button type="submit">login</button>
        <p class="message">Not registered? <a href="#" @click.prevent="toggleForm">Create an account</a></p>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showLoginForm: true,
      username: '',
      password: '',
      confirmPassword: '', 
      loginError: false,
      errorMessage: ''
    };
  },
  methods: {
    toggleForm() {
      this.showLoginForm = !this.showLoginForm;
    },
    async login() {
      try {
        const response = await fetch('http://127.0.0.1:8888/user/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            username: this.username,
            password: this.password,
          }),
        });

        if (!response.ok) {
          throw new Error('Login failed');
        }

        // 处理登录成功逻辑
        alert('Login successful');
      } catch (error) {
        this.loginError = true;
        this.errorMessage = error.message;
      }
    },
    async register() {
      // Here you will need to implement the registration logic
      // Similar to the login method, but likely to a different endpoint
      // Don't forget to validate the inputs, especially to check if password and confirmPassword match
    },
  },
};
</script>
  
  <style>
  @import url('https://fonts.googleapis.com/css?family=Roboto:300');
  
  .login-page {
    width: 360px;
    padding: 8% 0 0;
    margin: auto;
  }
  .form {
    position: relative;
    z-index: 1;
    background: #FFFFFF;
    max-width: 360px;
    margin: 0 auto 100px;
    padding: 45px;
    text-align: center;
    box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);
  }
  .form input {
    font-family: 'Roboto', sans-serif;
    outline: 0;
    background: #f2f2f2;
    width: 100%;
    border: 0;
    margin: 0 0 15px;
    padding: 15px;
    box-sizing: border-box;
    font-size: 14px;
  }
  .form button {
    font-family: 'Roboto', sans-serif;
    text-transform: uppercase;
    outline: 0;
    background: #4CAF50;
    width: 100%;
    border: 0;
    padding: 15px;
    color: #FFFFFF;
    font-size: 14px;
    -webkit-transition: all 0.3 ease;
    transition: all 0.3 ease;
    cursor: pointer;
  }
  .form button:hover,.form button:active,.form button:focus {
    background: #43A047;
  }
  .form .message {
    margin: 15px 0 0;
    color: #b3b3b3;
    font-size: 12px;
  }
  .form .message a {
    color: #4CAF50;
    text-decoration: none;
  }
  body {
    background: #76b852; /* fallback for old browsers */
    background: -webkit-linear-gradient(right, #76b852, #8DC26F);
    background: -moz-linear-gradient(right, #76b852, #8DC26F);
    background: -o-linear-gradient(right, #76b852, #8DC26F);
    background: linear-gradient(to left, #76b852, #8DC26F);
    font-family: 'Roboto', sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;      
  }
  </style>
  