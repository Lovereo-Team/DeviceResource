<template>
  <div v-loading="loading" class="common-layout">
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
            <view class="boxer">
              <p>日期</p>
              <span>{{ currentDate }}</span>
            </view>
            <img src="../../assets/Calendar.svg">
          </div>
          <div class="block">
            <view class="boxer">
              <p>时间</p>
              <span>{{ currentTime }}</span>
            </view>
            <img src="../../assets/clock.svg">
          </div>
          <div class="block">
            <view class="boxer">
              <p>当前设备ID</p>
              <span>{{ code }}</span>
            </view>
            <img src="../../assets/tag.svg">
          </div>
          <div class="block">
            <view class="boxer">
              <p>今日扫描次数</p>
              <span>{{ scanNumber }}</span>
            </view>
            <img src="../../assets/printer.svg">
          </div>
          <div class="block">
            <view class="boxer">
              <p>今日拍摄台数</p>
              <span>{{ cameraNumber }}</span>
            </view>
            <img src="../../assets/camera_header.svg">
          </div>
        </div>
      </el-header>
      <el-main class="main">
        <p>设备图像</p>
        <view class="image-boxer">
          <div class="box first-box">
            <el-image :initial-index="0" :max-scale="7" :min-scale="0.2" :preview-src-list="img_top"
                      :src="selectedImgTop" :zoom-rate="1.2" fit="fill" style="width: 100%; height: 55%" />
            <span class="image-text">上面</span>
            <el-select v-model="selectedTopIndex" placeholder="Select" size="large"
                       style="width: 70%;">
              <el-option v-for="(item, index) in img_top" :key="index" :label="`Image ${index + 1}`"
                         :value="index" />
            </el-select>
          </div>
          <div class="box">
            <el-image :initial-index="0" :max-scale="7" :min-scale="0.2" :preview-src-list="img_front"
                      :src="selectedImgFront" :zoom-rate="1.2" fit="fill" style="width: 100%; height: 55%" />
            <span class="image-text">前面</span>
            <el-select v-model="selectedFrontIndex" placeholder="Select" size="large"
                       style="width: 70%;">
              <el-option v-for="(item, index) in img_front" :key="index" :label="`Image ${index + 1}`"
                         :value="index" />
            </el-select>
          </div>
          <div class="box">
            <el-image :initial-index="0" :max-scale="7" :min-scale="0.2" :preview-src-list="img_behind"
                      :src="selectedImgBehend" :zoom-rate="1.2" fit="cover" style="width: 100%; height: 55%" />
            <span class="image-text">后面</span>
            <el-select v-model="selectedBehendIndex" placeholder="Select" size="large"
                       style="width: 70%;">
              <el-option v-for="(item, index) in img_behind" :key="index" :label="`Image ${index + 1}`"
                         :value="index" />
            </el-select>
          </div>
          <div class="box">
            <el-image :fit="fill" :src="img_left[0]" style="width: 100%; height: 55%" />
            <span class="image-text">左面</span>
            <el-select v-model="value" placeholder="Select" size="large"
                       style="width: 70%;">
              <el-option v-for="item in img_left" :key="item.value" :label="item.label"
                         :value="item.value" />
            </el-select>
          </div>
          <div class="box">
            <el-image :fit="fill" :src="img_right[0]" style="width: 100%; height: 55%" />
            <span class="image-text">右面</span>
            <el-select v-model="value" placeholder="Select" size="large"
                       style="width: 70%;">
              <el-option v-for="item in img_right" :key="item.value" :label="item.label"
                         :value="item.value" />
            </el-select>
          </div>
        </view>
      </el-main>
      <el-footer class="footer">
        <p>设备图像</p>
        <div class="footer-content">
          <div class="button-group">
            <el-button class="button" type="warning">重拍</el-button>
            <el-button class="button" type="primary" @click="upload">确定
            </el-button>
          </div>
        </div>
      </el-footer>
    </el-container>
  </div>
</template>
<script>
import axios from 'axios';
// import Api from "@/api/api"

export default {
  data() {
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
      token: '99966b4201c8b51cdb89e741772dc46eE0Gxoq',
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
        }, 10 * 1000);
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
        const headers = {
          "Token": this.token
        };
        const response = await axios.post("http://127.0.0.1:8000/api/resource/add", {
          deviceCode: code,
          Date: this.currentDate,
        }, {
          headers
        });
        console.log(response.data);
      } catch (error) {
        console.error(error);
      }
    },
    async getInfo() {
      try {
        const response = await axios.get("http://127.0.0.1:8000/api/member/info", {
          params: {
            createDate: this.currentDate
          },
          headers: {
            "Token": this.token
          }
        }, );
        this.scanNumber = response.data.data.scanNumber;
        this.cameraNumber = response.data.data.cameraNumber;
        // console.log(response.data);
      } catch (error) {
        console.log(error)
      }
    },
    async getImg() {

      try {
        const response = await axios.get("http://127.0.0.1:8000/api/member/img", {
          params: {
            deviceCode: this.code
          },
          headers: {
            "Token": this.token
          }
        }, );
        console.log(response.data.data.lists);
        let img = response.data.data.lists
        this.img_top = img.imgTop;
        this.img_front = img.imgFront;
        this.img_behind = img.imgBehind;
        this.img_left = img.imgLeft;
        this.img_right = img.imgRight;
        console.log(this.img_top)
        var that = this
        if (response.data.code == 200) {
          that.loading = false
        }

        //
      } catch (error) {
        console.log(error)
      }
    },
    async updateInfo() {
      try {
        const response = await axios.get("http://127.0.0.1:8000/api/member/update", {
          params: {
            createDate: this.currentDate
          },
          headers: {
            "Token": this.token
          }
        }, );
        console.log(response.data);
        //
      } catch (error) {
        console.log(error)
      }
    },
    async upload() {
      try {
        const res = await axios.post("http://127.0.0.1:8000/api/pic/add", {
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
          headers: {
            "Token": this.token
          }
        });
        console.log(res.data)
      } catch (error) {
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
      return this.img_top[this.selectedTopIndex] + "?t=" + Math.random() || ''; // 根据选中的索引获取对应的图片
    },
    selectedImgFront() {
      return this.img_front[this.selectedFrontIndex] + "?t=" + Math.random() || ''; // 根据选中的索引获取对应的图片
    },
    selectedImgBehend() {
      return this.img_behind[this.selectedBehendIndex] + "?t=" + Math.random() || ''; // 根据选中的索引获取对应的图片
    },
    selectedImgLeft() {
      return this.img_left[this.selectedLeftIndex] + "?t=" + Math.random() || ''; // 根据选中的索引获取对应的图片
    },
    selectedImgRight() {
      return this.img_right[this.selectedRightIndex] + "?t=" + Math.random() || ''; // 根据选中的索引获取对应的图片
    },

  },
}
</script>
<style scoped>
.common-layout {
  background-color: #EEF1F3;
  /* 设置整个网页的背景颜色 */
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
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
}

.header_1 div {
  width: 218px;
  display: flex;
  /* 将 .header_1 div 设置为 Flex 容器 */
  align-items: center;
  /* 垂直居中 */
}

.header_1 div span {
  font-size: 14px;
  color: white;
}

.header_1 div img {
  margin-right: 10px;
  /* 设置图像与文字之间的间距 */
}

.boxer {
  display: flex;
  flex-wrap: nowrap;
  align-items: center;
  flex-direction: column;
  /* padding: 5px; */
  width: 120px;
}

.image-boxer {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-around;
}
.image-text {
  width: 100%;
  text-align: center;
  height: 90px;
  align-content: center;
}
.header_2 {
  display: flex;
  flex-wrap: nowrap;
  justify-content: center;
  align-items: center;
  height: 205px;
  border-style: solid;
  border-color: #E4E4E4;
  border-width: 1px;
}

.block {
  position: relative;
  width: 13%;
  height: 124px;
  border-style: solid;
  border-color: #E4E4E4;
  border-width: 1px;
  margin-right: 60px;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
}

.first-block {
  margin-left: 127px;
  /* 设置第一个块的左边距为127px */
}

.block p {
  margin: 21px 17px;
  color: #CCCCCC;
  font-size: 14px;
}

.block img {
  position: absolute;
  /* 将图片设置为绝对定位 */
  bottom: 14px;
  /* 将图片定位到底部 */
  right: 21px;
  /* 将图片定位到右边 */
  width: 49px;
  height: 49px;
}

.main {
  display: flex;
  flex-wrap: nowrap;
  background: white;
  width: 100%;
  height: 50%;
  /* margin-top: 21px; */
  flex-direction: column;
}

.main p {
  color: #333333;
  font-size: 14px;
  font-weight: bold;
}

.box {
  display: flex;
  flex-wrap: wrap;
  width: 14%;
  height: 380px;
  border-style: solid;
  border-color: #E4E4E4;
  border-width: 1px;
  margin-top: 5px;
  margin-right: 60px;
  flex-direction: column;
  align-items: center;
}

.box p {
  font-size: 16px;
  color: black;
}

.first-box {
  margin-left: 80px;
}

.footer {
  background: white;
  width: 100%;
  height: 23%;
  //margin-top: 29px;
  //display: flex;
  //justify-content: center;
  //align-items: center;
}

.footer-content {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 25px;
}

.footer p {
  color: #333333;
  font-size: 14px;
  font-weight: bold;
}

.button-group {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 95%;
  flex-direction: row;
  flex-wrap: nowrap;
}

.button{
  width: 18%;
  height: 70px;
  font-size: 24px;
  text-align: center;
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