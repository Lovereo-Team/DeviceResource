<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
<!--                <el-form-item label="员工编号" prop="member_id">-->
<!--                    <el-input v-model="formData.member_id" placeholder="请输入员工编号" />-->
<!--                </el-form-item>-->
<!--                <el-form-item label="设备编号" prop="device_code">-->
<!--                    <el-input v-model="formData.device_code" placeholder="请输入设备编号" />-->
<!--                </el-form-item>-->
<!--                <el-form-item label="扫码日期" prop="date">-->
<!--                    <el-date-picker-->
<!--                        class="flex-1 !flex"-->
<!--                        v-model="formData.date"-->
<!--                        type="datetime"-->
<!--                        clearable-->
<!--                        value-format="YYYY-MM-DD hh:mm:ss"-->
<!--                        placeholder="请选择扫码日期"-->
<!--                    />-->
<!--                </el-form-item>-->
                <el-form-item label="图片_上" prop="imgTop">
                    <material-picker v-model="formData.imgTop" :limit="1" />
                </el-form-item>
                <el-form-item label="图片_前" prop="imgFront">
                    <material-picker v-model="formData.imgFront" :limit="1" />
                </el-form-item>
                <el-form-item label="图片_后" prop="imgBehind">
                    <material-picker v-model="formData.imgBehind" :limit="1" />
                </el-form-item>
                <el-form-item label="图片_左" prop="imgLeft">
                    <material-picker v-model="formData.imgLeft" :limit="1" />
                </el-form-item>
                <el-form-item label="图片_右" prop="imgRight">
                    <material-picker v-model="formData.imgRight" :limit="1" />
                </el-form-item>
<!--                <el-form-item label="图片集合" prop="img_s">
                    <el-input
                        v-model="formData.img_s"
                        placeholder="请输入图片集合"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>-->
<!--                <el-form-item label="视频" prop="video">-->
<!--                    <el-input v-model="formData.video" placeholder="请输入视频" />-->
<!--                </el-form-item>-->
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {  resourceEdit, resourceAdd, resourceDetail } from '@/api/resource'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import type { PropType } from 'vue'
defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    }
})
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑溯源列' : '新增溯源列'
})

const formData = reactive({
    id: '',
    // member_id: '',
    // device_code: '',
    date: '',
    imgTop: '',
    imgFront: '',
    imgBehind: '',
    imgLeft: '',
    imgRight: '',
    // img_s: '',
    // video: '',
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ],
    member_id: [
        {
            required: true,
            message: '请输入员工编号',
            trigger: ['blur']
        }
    ],
    device_code: [
        {
            required: true,
            message: '请输入设备编号',
            trigger: ['blur']
        }
    ],
    date: [
        {
            required: true,
            message: '请选择扫码日期',
            trigger: ['blur']
        }
    ],
    img_top: [
        {
            required: true,
            message: '请输入图片_上',
            trigger: ['blur']
        }
    ],
    img_front: [
        {
            required: true,
            message: '请输入图片_前',
            trigger: ['blur']
        }
    ],
    img_behind: [
        {
            required: true,
            message: '请输入图片_后',
            trigger: ['blur']
        }
    ],
    img_left: [
        {
            required: true,
            message: '请输入图片_左',
            trigger: ['blur']
        }
    ],
    img_right: [
        {
            required: true,
            message: '请输入图片_右',
            trigger: ['blur']
        }
    ],
    img_s: [
        {
            required: true,
            message: '请输入图片集合',
            trigger: ['blur']
        }
    ],
    video: [
        {
            required: true,
            message: '请输入视频',
            trigger: ['blur']
        }
    ],
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    mode.value == 'edit' ? await resourceEdit(data) : await resourceAdd(data)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
        }
    }
}

const getDetail = async (row: Record<string, any>) => {
    const data = await resourceDetail({
        id: row.id
    })
    setFormData(data)
}

const handleClose = () => {
    emit('close')
}

defineExpose({
    open,
    setFormData,
    getDetail
})
</script>
