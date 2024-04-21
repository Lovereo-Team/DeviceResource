<template>
  <div class="common-layout" v-loading="loading">
    <el-container>
      <el-header class="header">
        <div class="header_1">
          <div>
            <img src="../../assets/camera_fill.svg">
            <span>一体化智能工业相机视觉系统</span>
          </div>
        </div>
        <div class="header_2">
          <div class="block first-block">
            <p>日期</p>
            <span>{{currentDate}}</span>
            <img src="../../assets/Calendar.svg">
          </div>
          <div class="block">
            <p>时间</p>
            <span>{{currentTime}}</span>
            <img src="../../assets/clock.svg">
          </div>
          <div class="block">
            <p>当前设备ID</p>
            <span>{{code}}</span>
            <img src="../../assets/tag.svg">
          </div>
          <div class="block">
            <p>今日扫描次数</p>
            <span>{{scanNumber}}</span>
            <img src="../../assets/printer.svg">
          </div>
          <div class="block">
            <p>今日拍摄台数</p>
            <span>{{cameraNumber}}</span>
            <img src="../../assets/camera_header.svg">
          </div>
        </div>
      </el-header>
      <el-main class="main">
        <p>设备图像</p>
        <div class="box first-box">
          <el-image
              style="width: 100%; height: 55%"
              :src="selectedImgTop"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="img_top"
              :initial-index="0"
              fit="fill"
          />
          <span style="margin-left: 50%">上面</span>
          <el-select
              v-model="selectedTopIndex"
              placeholder="Select"
              size="large"
              style="width: 77%; margin-top: 42px; margin-left: 30px"
          >
            <el-option
                v-for="(item, index) in img_top"
                :key="index"
                :label="`Image ${index + 1}`"
                :value="index"
            />
          </el-select>
        </div>
        <div class="box">
          <el-image
              style="width: 100%; height: 55%"
              :src="selectedImgFront"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="img_front"
              :initial-index="0"
              fit="fill"
          />
          <span style="margin-left: 50%">前面</span>
          <el-select
              v-model="selectedFrontIndex"
              placeholder="Select"
              size="large"
              style="width: 77%; margin-top: 42px; margin-left: 30px"
          >
            <el-option
                v-for="(item, index) in img_front"
                :key="index"
                :label="`Image ${index + 1}`"
                :value="index"
            />
          </el-select>
        </div>
        <div class="box">
          <el-image
              style="width: 100%; height: 55%"
              :src="selectedImgBehend"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="img_behind"
              :initial-index="0"
              fit="cover"
          />
          <span style="margin-left: 50%">后面</span>
          <el-select
              v-model="selectedBehendIndex"
              placeholder="Select"
              size="large"
              style="width: 77%; margin-top: 42px; margin-left: 30px"
          >
            <el-option
                v-for="(item, index) in img_behind"
                :key="index"
                :label="`Image ${index + 1}`"
                :value="index"
            />
          </el-select>
        </div>
        <div class="box">
          <el-image style="width: 100%; height: 55%" :src="img_left[0]" :fit="fit" />
          <span style="margin-left: 50%">左面</span>
          <el-select
              v-model="value"
              placeholder="Select"
              size="large"
              style="width: 77%; margin-top: 42px; margin-left: 30px"
          >
            <el-option
                v-for="item in img_left"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </div>
        <div class="box">
          <el-image style="width: 100%; height: 55%" :src="img_right[0]" :fit="fit" />
          <span style="margin-left: 50%">右面</span>
          <el-select
              v-model="value"
              placeholder="Select"
              size="large"
              style="width: 77%; margin-top: 42px; margin-left: 30px"
          >
            <el-option
                v-for="item in img_right"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </div>
      </el-main>
      <el-footer class="footer">
        <p>设备图像</p>
        <div class="footer-content">

          <div class="button-group">
            <el-button type="warning" style="width: 30%; height: 100px; font-size: 26px">重拍</el-button>
            <el-button type="primary" style="width: 30%; height: 100px; font-size: 26px" @click="upload">确定</el-button>
          </div>
        </div>
      </el-footer>
    </el-container>
  </div>
</template>

<style scoped>
.common-layout {
  background-color: #EEF1F3; /* 设置整个网页的背景颜色 */
}
.header {
  background: white;
  width: 100%;
  height: 29%;
}
.header_1 {
  width: 275px;
  height: 100%;
  background: #0066FF;
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
}
.header_1 div {
  width: 218px;
  display: flex; /* 将 .header_1 div 设置为 Flex 容器 */
  align-items: center; /* 垂直居中 */
}
.header_1 div span {
  font-size: 14px;
  color: white;
}
.header_1 div img {
  margin-right: 10px; /* 设置图像与文字之间的间距 */
}
.header_2 {
  display: flex; /* 将 .header_2 设置为 Flex 容器 */
  flex-wrap: wrap; /* 允许换行 */
  justify-content: center;
  align-items: center;
  height: 204px;
  border-style: solid;
  border-color: #E4E4E4;
  border-width: 1px;
}
.block {
  position: relative; /* 将 .block 设置为相对定位的父容器 */
  width: 13%;
  height: 124px;
  border-style: solid;
  border-color: #E4E4E4;
  border-width: 1px;
  margin-right: 90px; /* 设置右边距为90px */
  //margin-bottom: 90px; /* 设置下边距为90px */
}
.first-block {
  margin-left: 127px; /* 设置第一个块的左边距为127px */
}
.block p {
  margin: 21px 17px;
  color: #CCCCCC;
  font-size: 14px;
}
.block img{
  position: absolute; /* 将图片设置为绝对定位 */
  bottom: 14px; /* 将图片定位到底部 */
  right: 21px; /* 将图片定位到右边 */
}
.main {
  display: flex; /* 将 .header_2 设置为 Flex 容器 */
  flex-wrap: wrap; /* 允许换行 */
  background: white;
  width: 100%;
  height: 48%;
  margin-top: 21px;
}
.main p {
  color: #333333;
  font-size: 14px;
}
.box {
  display: flex;
  flex-wrap: wrap; /* 允许换行 */
  //justify-content: center;
  width: 14%;
  height: 380px;
  border-style: solid;
  border-color: #E4E4E4;
  border-width: 1px;
  margin-top: 45px;
  margin-right: 60px; /* 设置右边距为90px */
}
.box p{
  font-size: 16px;
  color: black;
}
.first-box {
  margin-left: 137px;
}
.footer {
  background: white;
  width: 100%;
  height: 23%;
  margin-top: 29px;
  //display: flex;
  //justify-content: center;
  //align-items: center;
}
.footer-content {
  display: flex;
  justify-content: center;
  align-items: center;
}
.footer p {
  color: #333333;
  font-size: 14px;
}

.button-group {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 45%;
}
.demo-image__error .image-slot {
  font-size: 30px;
}
.demo-image__error .image-slot .el-icon {
  font-size: 30px;
}
.demo-image__error .el-image {
  width: 100%;
  height: 200px;
}
</style>
<script>
import axios from 'axios';
// import Api from "@/api/api"

export default {
  data(){
    return {
      loading: false,
      currentDate: '', // 存储当前计算机时间
      currentTime: '',
      code: "",
      selectedFrontIndex: 0,
      selectedTopIndex: 0,
      selectedBehendIndex: 0,
      selectedLeftIndex: 0,
      selectedRightIndex: 0,
      scanNumber: "0",
      cameraNumber: "0",
      img_top: [],
      img_front: [],
      img_behind: [],
      img_left: [],
      img_right: [],
      token: '06c0626c3659ac68358cb88b3deeed87JfMjE1',
    };
  },
  mounted() {
    this.getCurrentDateTime();
    // 设置定时器，每秒更新一次时间
    setInterval(() => {
      this.getCurrentDateTime();
    }, 1000);
    // 监听按键
    var code = ''
    var lastTime, nextTime // 上次时间、最新时间
    var lastCode, nextCode // 上次按键、最新按键

    document.onkeypress = (e) => {


      // 获取按键
      if (window.event) { // IE
        nextCode = e.keyCode
      } else if (e.which) { // Netscape/Firefox/Opera
        nextCode = e.which
      }
      console.log(nextCode)
      // 如果触发了回车事件(扫码结束时间)
      if (nextCode === 13) {
        if (code.length < 3) return // 手动输入的时间不会让code的长度大于2，所以这里只会对扫码枪有
        this.loading = true
        this.codeFind(code) // 获取到扫码枪输入的内容，做别的操作
        this.getInfo()
        this.updateInfo()
        setTimeout(() => {
          this.getImg();
        }, 5000);
        // this.getImg()
        code = ''
        lastCode = ''
        lastTime = ''

        return
      }

      nextTime = new Date().getTime() // 记录最新时间
      if (!lastTime && !lastCode) { // 如果上次时间和上次按键为空
        code += e.key // 执行叠加操作
      }
      // 如果有上次时间及上次按键
      if (lastCode && lastTime && nextTime - lastTime > 30) { // 当扫码前有keypress事件时,防止首字缺失
        code = e.key
      } else if (lastCode && lastTime) {
        code += e.key
      }
      lastCode = nextCode
      lastTime = nextTime

    }
    // 调用结束后10秒后调用this.getImg

  },
  methods: {
    getCurrentDateTime() {

      const now = new Date(); // 创建一个Date对象来获取当前时间
      const year = now.getFullYear(); // 获取年份
      const month = String(now.getMonth() + 1).padStart(2, '0'); // 获取月份
      const day = String(now.getDate()).padStart(2, '0'); // 获取日
      const hours = String(now.getHours()).padStart(2, '0'); // 获取小时
      const minutes = String(now.getMinutes()).padStart(2, '0'); // 获取分钟
      const seconds = String(now.getSeconds()).padStart(2, '0'); // 获取秒钟

      // 格式化日期和时间为字符串
      this.currentDate = `${year}-${month}-${day}`;
      this.currentTime = `${hours}:${minutes}:${seconds}`;
    },
    // 发送HTTP请求
    async codeFind(code) {
      var that = this
      that.loading = true
      this.code = code
      try {
        const headers = {"Token": this.token};
        const response = await axios.post("http://127.0.0.1:8000/api/resource/add", {
          deviceCode: code,
          Date: this.currentDate,
        }, {headers});
        console.log(response.data);
      } catch (error) {
        console.error(error);
      }
    },
    async getInfo(){
      try {
        const response = await axios.get("http://127.0.0.1:8000/api/member/info", {
          params:{
            createDate: this.currentDate
          },
          headers:{
            "Token": this.token
          }
        },);
        this.scanNumber = response.data.data.scanNumber;
        this.cameraNumber = response.data.data.cameraNumber;
        // console.log(response.data);
      } catch (error) {
        console.log(error)
      }
    },
    async getImg(){

      try {
        const response = await axios.get("http://127.0.0.1:8000/api/member/img", {
          params:{
            deviceCode: this.code
          },
          headers:{
            "Token": this.token
          }
        },);
        console.log(response.data.data.lists);
        let img = response.data.data.lists
        this.img_top = img.imgTop;
        this.img_front = img.imgFront;
        this.img_behind = img.imgBehind;
        this.img_left = img.imgLeft;
        this.img_right = img.imgRight;
        console.log(this.img_top)
        var that = this
        if (response.data.code == 200){
          that.loading = false
        }

        //
      } catch (error) {
        console.log(error)
      }
    },
    async updateInfo(){
      try {
        const response = await axios.get("http://127.0.0.1:8000/api/member/update", {
          params:{
            createDate: this.currentDate
          },
          headers:{
            "Token": this.token
          }
        },);
        console.log(response.data);
        //
      } catch (error) {
        console.log(error)
      }
    },
    async upload(){
      try {
        const res = await axios.post("http://127.0.0.1:8000/api/pic/add",{
          img_behind_camera_number: 3,
          img_front_camera_number: 2,
          img_left_camera_number: 4,
          img_right_camera_number: 5,
          img_top_camera_number: 1,
          imgBehind: this.img_behind,
          imgFront: this.img_front,
          imgTop: this.img_top,
          imgLeft: this.img_left,
          imgRight: this.img_right
        }, {
          headers:{
            "Token": this.token
          }
        });
        console.log(res.data)
      }catch (error){
        console.log(error)
      }
      this.img_top = [];
      this.img_front = [];
      this.img_behind = [];
      this.img_left = [];
      this.img_right = [];
    }
  },
  computed: {
    selectedImgTop() {
      return this.img_top[this.selectedTopIndex]+"?t="+Math.random() || ''; // 根据选中的索引获取对应的图片
    },
    selectedImgFront() {
      return this.img_front[this.selectedFrontIndex]+"?t="+Math.random() || ''; // 根据选中的索引获取对应的图片
    },
    selectedImgBehend() {
      return this.img_behind[this.selectedBehendIndex]+"?t="+Math.random() || ''; // 根据选中的索引获取对应的图片
    },
    selectedImgLeft() {
      return this.img_left[this.selectedLeftIndex]+"?t="+Math.random() || ''; // 根据选中的索引获取对应的图片
    },
    selectedImgRight() {
      return this.img_right[this.selectedRightIndex]+"?t="+Math.random() || ''; // 根据选中的索引获取对应的图片
    },

  },
}


</script>