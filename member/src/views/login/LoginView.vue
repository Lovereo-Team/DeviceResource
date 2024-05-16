<template>
  <view class="bj">
    <view class="box-card">
      <view class="box">
        <img class="company-logo" src="./staic/Company_icon.jpg" alt="Company Logo">
        <h2 class="title">设备拍照检测系统</h2>
      </view>
      <form ref="loginForm" @submit.prevent="login" class="login-form">
        <view class="form-group">
          <label for="uname">用户名</label>
          <input class="input" v-model="uname" type="text" id="uname" placeholder="请输入用户名" required>
        </view>
        <view class="form-group">
          <label for="password">密码</label>
          <input class="input" v-model="password" type="password" id="password" placeholder="请输入登录密码" required>
        </view>
        <div class="button-container">
          <span class="mas">登录</span>
          <button :disabled="loading" type="submit">登录</button>
        </div>
      </form>
    </view>
  </view>
</template>

<script>
import { ref } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  setup() {
    const uname = ref('');
    const password = ref('');
    const loading = ref(false);
    const router = useRouter();

    const login = () => {
      loading.value = true;
      axios.post('http://127.0.0.1:8000/api/member/login', {
        username: uname.value,
        password: password.value
      })
          .then(response => {
            if (response.data.code === 200) {
              console.log('登录成功', response.data);
              const token = response.data.data.token;
              localStorage.setItem('token', token);
              if (router) {
                router.push({ path: '/dashboard' });
              } else {
                console.error('未找到Router对象');
              }
            } else {
              console.error('登录失败', response.data.msg);
            }
          })
          .catch(error => {
            console.error('登录错误', error);
          })
          .finally(() => {
            loading.value = false;
          });
    };

    return {
      uname,
      password,
      loading,
      login
    };
  }
};
</script>


<style scoped>
.bj {
  background-image: linear-gradient(135deg, #7A88FF 5%, #ffffff 75%, #d1d6ff 100%);
  width: 100%;
  height: calc(100vh);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  flex-wrap: nowrap;
}
.box{
  display: flex;
  flex-direction: column;
  align-items: center;
}
.box-card {
  width: 500px;
  height: 404px;
  text-align: -webkit-center;
  border-radius: 20px;
  background-color: #fff;
}

.box-card el-input {
  width: 438px !important;
  height: 52px !important;
}

.company-logo {
  background-image: url(http://localhost:8081/img/Company_icon.11ac893d.jpg);
  background-size: contain;
  background-repeat: no-repeat;
  /* background-position: center; */
  width: auto;
  height: 85px;
  margin-top: 70px;
  margin-right: auto;
  margin-bottom: 10px;
  margin-left: auto;
}
.login-form{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
  align-content: space-around;
  flex-wrap: wrap;
  width: 100%;
  height: 135px;
}
.form-group{
  width: 70%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-direction: row;
}
.input{
  outline-style: none;
  border: 1px solid #ccc;
  border-radius: 3px;
  font-size: 14px;
  font-weight: 700;
  font-family: "Microsoft soft";
  margin-left: 15px;
  width: 70%;
}
.title {
  font-size: 26px;
  font-weight: bolder;
  margin-right: 12px;
  margin-bottom: 22px;
  margin-left: 12px;
}

.button_content {
  align-items: center;
  display: flex;
  flex-wrap: wrap;
  font-size: var(--font-size);
  line-height: 32px;
  position: relative;
  width: 70%;
  justify-content: center;
}

@import url("https://fonts.googleapis.com/css?family=Lato:100,300,400");

.button-container {
  position: relative;
  width: 70%;
  margin-bottom: 7px;
  overflow: hidden;
  /* border: 1px solid; */
  font-family: 'Lato', sans-serif;
  font-weight: bold;
  font-size: 16px;
  transition: 0.5s;
  letter-spacing: 1px;
  background-color: #1f7dff;
  border-radius: 5px;
  height: 45px;
}

.button-container button {
  width: 100%;
  height: 100%;
  font-family: 'Lato', sans-serif;
  font-weight: bold;
  font-size: 16px;
  letter-spacing: 1px;
  background: #f5f5f5;
  -webkit-mask: url("./staic/nature-sprite.png");
  mask: url("./staic/nature-sprite.png");
  -webkit-mask-size: 2300% 100%;
  mask-size: 2300% 100%;
  border: none;
  color: #000000;
  cursor: pointer;
  -webkit-animation: ani2 0.7s steps(22) forwards;
  animation: ani2 0.7s steps(22) forwards;
}

.button-container button:hover {
  -webkit-animation: ani 0.7s steps(22) forwards;
  animation: ani 0.7s steps(22) forwards;
}

.mas {
  position: absolute;
  color: #ffffff;
  background-image: linear-gradient(135deg, #7A88FF 0%, #d1d6ff 100%);
  text-align: center;
  width: 100%;
  height: 100%;
  align-content: center;
  font-family: 'Lato', sans-serif;
  font-weight: bold;
  font-size: 18px;
  overflow: hidden;
}

@-webkit-keyframes ani {
  from {
    -webkit-mask-position: 0 0;
    mask-position: 0 0;
  }

  to {
    -webkit-mask-position: 100% 0;
    mask-position: 100% 0;
  }
}

@keyframes ani {
  from {
    -webkit-mask-position: 0 0;
    mask-position: 0 0;
  }

  to {
    -webkit-mask-position: 100% 0;
    mask-position: 100% 0;
  }
}
</style>