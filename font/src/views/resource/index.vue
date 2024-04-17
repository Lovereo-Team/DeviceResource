<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="员工编号" prop="member_id">
                    <el-input class="w-[280px]" v-model="queryParams.member_id" />
                </el-form-item>
                <el-form-item label="设备编号" prop="device_code">
                    <el-input class="w-[280px]" v-model="queryParams.device_code" />
                </el-form-item>
                <el-form-item label="扫码日期" prop="date">
                    <el-date-picker
                        v-model="queryParams.date"
                        type="date"
                        placeholder="选择日期">
                    </el-date-picker>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button v-perms="['resource:add']" type="primary" @click="handleAdd()">
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
                <el-table-column label="员工编号" prop="memberId" min-width="60" />
                <el-table-column label="设备编号" prop="deviceCode" min-width="100" />
                <el-table-column label="扫码日期" prop="date" min-width="100" />
                <el-table-column label="图片_上" prop="imgTop" min-width="100" >
                    <template #default="scope">
                        <el-image
                            style="width: 40px; height: 40px"
                            :src="scope.row.imgTop"
                            :preview-src-list="[scope.row.imgTop]"
                            :hide-on-click-modal="true"
                            :preview-teleported="true"
                            fit="cover"
                        >
                            <template #error>
                                <div class="h-full flex justify-center items-center iconStyle">
                                    <el-icon :size="30"><Warning /></el-icon>
                                </div>
                            </template>
                        </el-image>
                        <el-button
                            v-perms="['resource:video']"
                            type="primary"
                            link
                            @click="handleVideo()"
                        >
                            视频
                        </el-button>
                    </template>
                </el-table-column>
                <el-table-column label="图片_前" prop="imgFront" min-width="100" >
                    <template #default="scope">
                        <el-image
                            style="width: 40px; height: 40px"
                            :src="scope.row.imgFront"
                            :preview-src-list="[scope.row.imgFront]"
                            :hide-on-click-modal="true"
                            :preview-teleported="true"
                            fit="cover"
                        >
                            <template #error>
                                <div class="h-full flex justify-center items-center iconStyle">
                                    <el-icon :size="30"><Warning /></el-icon>
                                </div>
                            </template>
                        </el-image>
                    </template>
                </el-table-column>
                <el-table-column label="图片_后" prop="imgBehind" min-width="100" >
                    <template #default="scope">
                        <el-image
                            style="width: 40px; height: 40px"
                            :src="scope.row.imgBehind"
                            :preview-src-list="[scope.row.imgBehind]"
                            :hide-on-click-modal="true"
                            :preview-teleported="true"
                            fit="cover"
                        >
                            <template #error>
                                <div class="h-full flex justify-center items-center iconStyle">
                                    <el-icon :size="30"><Warning /></el-icon>
                                </div>
                            </template>
                        </el-image>
                    </template>
                </el-table-column>
                <el-table-column label="图片_左" prop="imgLeft" min-width="100" >
                    <template #default="scope">
                        <el-image
                            style="width: 40px; height: 40px"
                            :src="scope.row.imgLeft"
                            :preview-src-list="[scope.row.imgLeft]"
                            :hide-on-click-modal="true"
                            :preview-teleported="true"
                            fit="cover"
                        >
                            <template #error>
                                <div class="h-full flex justify-center items-center iconStyle">
                                    <el-icon :size="30"><Warning /></el-icon>
                                </div>
                            </template>
                        </el-image>
                    </template>
                </el-table-column>
                <el-table-column label="图片_右" prop="imgRight" min-width="100" >
                    <template #default="scope">
                        <el-image
                            style="width: 40px; height: 40px"
                            :src="scope.row.imgRight"
                            :preview-src-list="[scope.row.imgRight]"
                            :hide-on-click-modal="true"
                            :preview-teleported="true"
                            fit="cover"
                        >
                            <template #error>
                                <div class="h-full flex justify-center items-center iconStyle">
                                    <el-icon :size="30"><Warning /></el-icon>
                                </div>
                            </template>
                        </el-image>

                    </template>
                </el-table-column>
<!--                <el-table-column label="图片集合" prop="imgS" min-width="100" />-->
                <el-table-column label="视频" prop="video" min-width="100" />
                <el-table-column label="创建时间" prop="createTime" min-width="150" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['resource:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['resource:del']"
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
            @success="getLists"
            @close="showEdit = false"
        />
        <video-group
            v-if="showVideo"
            ref="videoRef"
            @success="getLists"
            @close="showVideo = false"
        />
    </div>
</template>
<script lang="ts" setup name="resource">
import { resourceDelete, resourceLists } from '@/api/resource'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './components/edit.vue'
import VideoGroup from './components/video.vue'
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const videoRef = shallowRef<InstanceType<typeof VideoGroup>>()
const showEdit = ref(false)
const showVideo = ref(false)
const queryParams = reactive({
    member_id: '',
    device_code: '',
    date: "",
    // img_top: '',
    // img_front: '',
    // img_behind: '',
    // img_left: '',
    // img_right: '',
    // img_s: '',
    // video: '',
})


const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: resourceLists,
    params: queryParams
})


const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    console.log(data)
    await nextTick()

    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleVideo = async () => {
    showVideo.value = true
    await nextTick()

    videoRef.value?.open('video')
    // videoRef.value?.getDetail(data)
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await resourceDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

getLists()
</script>
