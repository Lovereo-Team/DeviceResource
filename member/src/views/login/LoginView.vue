<template>
  <div>
    <el-card class="box-card">
      <el-form
          :model="ruleForm"
          status-icon
          :rules="rules"
          ref="ruleForm"
          label-position="left"
      >
        <el-form-item prop="uname">
          <el-input v-model="ruleForm.uname"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
              type="password"
              v-model="ruleForm.password"
              autocomplete="off"
          ></el-input>
        </el-form-item>
      </el-form>
      <div class="btnGroup">
        <el-button type="primary" @click="submitForm('ruleForm')"
        >登录</el-button
        >
        <el-button @click="resetForm('ruleForm')">重置</el-button>
        <router-link to="/register">
          <el-button style="margin-left:10px">注册</el-button>
        </router-link>
      </div>
    </el-card>
  </div>
</template>

<script>
import axios from 'axios';
import { useRouter } from 'vue-router';
export default {
  data() {
    return {
      ruleForm: {
        uname: "",
        password: "",
      },
      rules: {
        uname: [
          { required: true, message: "用户名不能为空！", trigger: "blur" },
        ],
        password: [
          { required: true, message: "密码不能为空！", trigger: "blur" },
        ],
      },
    };
  },
  methods: {
// 在 submitForm 方法中
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          // 构造请求数据
          const requestData = {
            username: this.ruleForm.uname,
            password: this.ruleForm.password
          };

          // 发送 POST 请求到登录接口
          axios.post('http://127.0.0.1:8000/api/member/login', requestData)
              .then(response => {
                // 处理登录成功的响应
                console.log('登录成功', response.data);
                useRouter().push('/index');
                // 可以在这里进行页面跳转或者其他操作
              })
              .catch(error => {
                // 处理登录失败的响应
                console.error('登录失败', error.response.data);
                // 可以在这里给用户显示错误信息或者执行其他操作
              });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
  }
};
</script>

<style scoped>
.box-card {
  //margin: auto auto;
  width: 500px;
  height: 404px;
  margin-top: 337px;
  margin-left: 710px;
}
.box-card el-input {
  width: 438px !important;
  height: 52px !important;
}
</style>
