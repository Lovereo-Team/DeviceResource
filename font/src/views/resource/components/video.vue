<template>
    <div class="video_group"  v-show="visible">
        <popup
            ref="popupRef"
            :async="true"
            width="550px"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <!-- 视频播放器 -->
            <video controls :src="videoUrl"></video>
        </popup>
    </div>
</template>

<script setup lang="ts">
import Popup from '@/components/popup/index.vue'
import { ref } from 'vue'

const visible = ref(false) // 控制弹窗显示状态
const videoUrl = ref('') // 视频链接
const mode = ref('add')
const popupRef = shallowRef<InstanceType<typeof Popup>>()
// 处理弹窗提交
const handleSubmit = () => {

}

// 处理弹窗关闭
const handleClose = () => {
    // 处理弹窗关闭逻辑
}

// 定义 open 方法，用于显示视频弹窗并传入视频数据
const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}
defineExpose({
    open
})
</script>

<style scoped lang="scss">
/* 样式 */
</style>
