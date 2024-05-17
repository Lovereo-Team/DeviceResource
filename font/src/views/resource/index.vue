<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" :inline="true" :model="queryParams" class="mb-[-16px]">
                <el-form-item label="员工编号" prop="member_id">
                    <el-input v-model="queryParams.member_id" class="w-[280px]" />
                </el-form-item>
                <el-form-item label="设备编号" prop="device_code">
                    <el-input v-model="queryParams.device_code" class="w-[280px]" />
                </el-form-item>
                <el-form-item label="扫码日期" prop="date">
                    <el-date-picker
                        v-model="queryParams.date"
                        placeholder="选择日期"
                        type="date">
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
                v-loading="pager.loading"
                :data="pager.lists"
                class="mt-4"
                size="large"
            >
                <el-table-column label="员工编号" min-width="60" prop="memberId" />
                <el-table-column label="设备编号" min-width="100" prop="deviceCode" />
                <el-table-column label="扫码日期" min-width="100" prop="date" />
                <el-table-column label="图片_上" min-width="100" prop="imgTop" >
                    <template #default="scope" >
                        <view class="box">
                            <el-image
                                :hide-on-click-modal="true"
                                :preview-src-list="[scope.row.imgTop]"
                                :preview-teleported="true"
                                :src="scope.row.imgTop"
                                fit="cover"
                                style="width: 40px; height: 40px"
                            >
                                <template #error>
                                    <div class="h-full flex justify-center items-center iconStyle">
                                        <el-icon :size="30"><Warning /></el-icon>
                                    </div>
                                </template>
                            </el-image>
                            <el-button
                                v-perms="['resource:video']"
                                link
                                type="primary"
                                @click="handleVideo(scope.row.VideoTop)"
                            >
                                视频
                            </el-button>
                        </view>
                    </template>
                </el-table-column>
                <el-table-column label="图片_前" min-width="100" prop="imgFront" >
                    <template #default="scope">
                        <view class="box">
                            <el-image
                                :hide-on-click-modal="true"
                                :preview-src-list="[scope.row.imgFront]"
                                :preview-teleported="true"
                                :src="scope.row.imgFront"
                                fit="cover"
                                style="width: 40px; height: 40px"
                            >
                                <template #error>
                                    <div class="h-full flex justify-center items-center iconStyle">
                                        <el-icon :size="30"><Warning /></el-icon>
                                    </div>
                                </template>
                            </el-image>
                            <el-button
                                v-perms="['resource:video']"
                                link
                                type="primary"
                                @click="handleVideo(scope.row.VideoFront)"
                            >
                                视频
                            </el-button>
                        </view>
                    </template>
                </el-table-column>
                <el-table-column label="图片_后" min-width="100" prop="imgBehind" >
                    <template #default="scope">
                        <view class="box">
                            <el-image
                                :hide-on-click-modal="true"
                                :preview-src-list="[scope.row.imgBehind]"
                                :preview-teleported="true"
                                :src="scope.row.imgBehind"
                                fit="cover"
                                style="width: 40px; height: 40px"
                            >
                                <template #error>
                                    <div class="h-full flex justify-center items-center iconStyle">
                                        <el-icon :size="30"><Warning /></el-icon>
                                    </div>
                                </template>
                            </el-image>
                        </view>
                    </template>
                </el-table-column>
                <el-table-column label="图片_左" min-width="100" prop="imgLeft" >
                    <template #default="scope">
                        <view class="box">
                            <el-image
                                :hide-on-click-modal="true"
                                :preview-src-list="[scope.row.imgLeft]"
                                :preview-teleported="true"
                                :src="scope.row.imgLeft"
                                fit="cover"
                                style="width: 40px; height: 40px"
                            >
                                <template #error>
                                    <div class="h-full flex justify-center items-center iconStyle">
                                        <el-icon :size="30"><Warning /></el-icon>
                                    </div>
                                </template>
                            </el-image>
                        </view>
                    </template>
                </el-table-column>
                <el-table-column label="图片_右" min-width="100" prop="imgRight" >
                    <template #default="scope">
                        <view class="box">
                            <el-image
                                :hide-on-click-modal="true"
                                :preview-src-list="[scope.row.imgRight]"
                                :preview-teleported="true"
                                :src="scope.row.imgRight"
                                fit="cover"
                                style="width: 40px; height: 40px"
                            >
                                <template #error>
                                    <div class="h-full flex justify-center items-center iconStyle">
                                        <el-icon :size="30"><Warning /></el-icon>
                                    </div>
                                </template>
                            </el-image>
                        </view>
                    </template>
                </el-table-column>
<!--                <el-table-column label="图片集合" prop="imgS" min-width="100" />-->
<!--                <el-table-column label="视频" prop="video" min-width="100" />-->
                <el-table-column label="创建时间" min-width="150" prop="createTime" />
                <el-table-column fixed="right" label="操作" width="120">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['resource:edit']"
                            link
                            type="primary"
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['resource:del']"
                            link
                            type="danger"
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
            @close="showEdit = false"
            @success="getLists"
        />
        <video-group
            v-if="showVideo"
            ref="videoRef"
            @close="showVideo = false"
            @success="getLists"
        />
    </div>
</template>
<script lang="ts" name="resource" setup>
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

const handleVideo = async (device) => {
    console.log(device)
    showVideo.value = true
    await nextTick()

    videoRef.value?.open(device)
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
<style scoped>
.box{
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    height: 70px;
    flex-wrap: nowrap;
    justify-content: space-around;
}
</style>