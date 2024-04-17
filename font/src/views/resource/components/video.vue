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
            <video controls preload="load">
                <source :src="videoUrl" type="video/mp4">
                <source :src="videoUrl" type="video/ogg">
                Your browser does not support the video tag.
            </video>
        </popup>
    </div>
</template>

<script setup lang="ts">
import Popup from '@/components/popup/index.vue'
import { ref } from 'vue'
import { usePaging } from '@/hooks/usePaging'
import { memberGetVideo, resourceDetail, resourceLists } from '@/api/resource'

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
const open = async (device) => {
    videoUrl.value = device
    popupRef.value?.open()
}

defineExpose({
    open
})
</script>

<style scoped lang="scss">
/* 样式 */
</style>
