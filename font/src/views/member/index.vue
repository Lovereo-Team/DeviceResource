<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="用户名" prop="username">
                    <el-input class="w-[280px]" v-model="queryParams.username" />
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input class="w-[280px]" v-model="queryParams.password" />
                </el-form-item>
                <el-form-item label="真实姓名" prop="RealName">
                    <el-input class="w-[280px]" v-model="queryParams.RealName" />
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input class="w-[280px]" v-model="queryParams.email" />
                </el-form-item>
                <el-form-item label="用户状态" prop="status">
                    <el-select
                        v-model="queryParams.status"
                        class="w-[280px]"
                        clearable
                    >
                        <el-option v-for="(item,index) in statusList" :label="item.label" :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="登录时间" prop="login_time">
                    <el-input class="w-[280px]" v-model="queryParams.login_time" />
                </el-form-item>
                <el-form-item label="用户凭证" prop="token">
                    <el-input class="w-[280px]" v-model="queryParams.token" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button v-perms="['member:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <el-table
                class="mt-4"
                size="large"
                v-loading="pager.loading"
                :data="pager.lists"
            >
                <el-table-column label="用户名" prop="username" min-width="100" />
                <el-table-column label="密码" prop="password" min-width="100" />
                <el-table-column label="真实姓名" prop="realName" min-width="100" />
                <el-table-column label="邮箱" prop="email" min-width="100" />
                <el-table-column label="用户状态" prop="status" min-width="100" >
                    <template #default="{ row }">
                        <el-tag v-if="row.status == 1">正常</el-tag>
                        <el-tag type="danger" v-else>禁用</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="登录时间" prop="login_time" min-width="100" />
                <el-table-column label="创建时间" prop="create_time" min-width="100" />
                <el-table-column label="更新时间" prop="update_time" min-width="100" />
                <el-table-column label="用户凭证" prop="token" min-width="100" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['member:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['member:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
                        >
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <edit-popup
            v-if="showEdit"
            ref="editRef"
            :status-list="statusList"
            @success="getLists"
            @close="showEdit = false"

        />
    </div>
</template>
<script lang="ts" setup name="member">
import { memberDelete, memberLists } from '@/api/member'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './components/edit.vue'
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    username: '',
    password: '',
    RealName: '',
    email: '',
    status: '',
    login_time: '',
    token: '',
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: memberLists,
    params: queryParams
})

const statusList = [
    {
        label:"正常",
        value:1
    },
    {
        label:"禁用",
        value:2
    },
]
const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await memberDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

getLists()
</script>
