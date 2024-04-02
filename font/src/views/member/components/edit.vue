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
<!--                <el-form-item label="" prop="id">-->
<!--                    <el-input v-model="formData.id" placeholder="请输入" />-->
<!--                </el-form-item>-->
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="formData.username" placeholder="请输入用户名" />
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="formData.password" placeholder="请输入密码" />
                </el-form-item>
                <el-form-item label="真实姓名" prop="realName">
                    <el-input v-model="formData.realName" placeholder="请输入真实姓名" />
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model="formData.email" placeholder="请输入邮箱" />
                </el-form-item>
                <el-form-item label="用户状态" prop="status">
                    <el-radio-group v-model="formData.status" placeholder="状态">
                        <el-radio :label="item.value" v-for="(item,index) in statusList">{{ item.label }}</el-radio>
                    </el-radio-group>
                </el-form-item>
<!--                <el-form-item label="登录时间" prop="LoginTime">-->
<!--                    <el-input v-model="formData.login_time" placeholder="请输入登录时间" />-->
<!--                </el-form-item>-->
<!--                <el-form-item label="用户凭证" prop="token">-->
<!--                    <el-input v-model="formData.token" placeholder="请输入用户凭证" />-->
<!--                </el-form-item>-->
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {  memberEdit, memberAdd, memberDetail } from '@/api/member'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import type { PropType } from 'vue'
defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    },
    statusList:{
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    },
})
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑员工列' : '新增员工列'
})

const formData = reactive({
    id: '',
    username: '',
    password: '',
    realName: '',
    email: '',
    status: 1,
    // login_time: '',
    // token: '',
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ],
    username: [
        {
            required: true,
            message: '请输入用户名',
            trigger: ['blur']
        }
    ],
    password: [
        {
            required: true,
            message: '请输入密码',
            trigger: ['blur']
        }
    ],
    RealName: [
        {
            required: true,
            message: '请输入真实姓名',
            trigger: ['blur']
        }
    ],
    email: [
        {
            required: true,
            message: '请输入邮箱',
            trigger: ['blur']
        }
    ],
    status: [
        {
            required: true,
            message: '请选择用户状态 1：正常 2：禁用',
            trigger: ['blur']
        }
    ],
    login_time: [
        {
            required: true,
            message: '请输入登录时间',
            trigger: ['blur']
        }
    ],
    token: [
        {
            required: true,
            message: '请输入用户凭证',
            trigger: ['blur']
        }
    ],
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    mode.value == 'edit' ? await memberEdit(data) : await memberAdd(data)
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
    const data = await memberDetail({
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
