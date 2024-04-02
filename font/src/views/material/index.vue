<template>
    <div class="material-index">
        <el-card class="!border-none" shadow="never">
            <el-tabs v-model="activeTab">
                <el-tab-pane
                    v-for="item in tabsMap"
                    :label="item.name"
                    :name="item.type"
                    :index="item.type"
                    :key="item.type"
                    lazy
                >
<!--                    Remove material Add Table By Lovereo-->
                    <material
                        :type="item.type"
                        mode="page"
                        file-size="120px"
                        :limit="-1"
                        :page-size="20"
                    />
<!--                    <el-table-->
<!--                        :data="tableData"-->
<!--                        border-->
<!--                        style="width: 100%">-->
<!--                        <el-table-column-->
<!--                            prop="productNumber"-->
<!--                            label="产品编号">-->
<!--                        </el-table-column>-->
<!--                        <el-table-column-->
<!--                            prop="fetchDate"-->
<!--                            label="抓取日期">-->
<!--                        </el-table-column>-->
<!--                        <el-table-column-->
<!--                            prop="fetchTime"-->
<!--                            label="抓取时间">-->
<!--                        </el-table-column>-->
<!--                        <el-table-column-->
<!--                            prop="fetchGallery"-->
<!--                            label="抓取图集">-->
<!--                        </el-table-column>-->
<!--                        <el-table-column-->
<!--                            label="操作">-->
<!--                            <template #default="{ row }">-->
<!--                                <el-button @click="viewDetails(row)">查看</el-button>-->
<!--                            </template>-->
<!--                        </el-table-column>-->
<!--                    </el-table>-->
                </el-tab-pane>
            </el-tabs>
        </el-card>
    </div>
</template>

<script lang="ts" setup name="materialCenter">
const tabsMap = [
    {
        type: 'image',
        name: '图片'
    },
    {
        type: 'video',
        name: '视频'
    }
]
const activeTab = ref('image')
// Edit Add By Lovereo →
const tableData = [{ productNumber: '001', fetchDate: '2024-04-01', fetchTime: '10:00', fetchGallery: 'gallery1', fetchExamine: '查看' }]
const viewDetails = (row: any) => {
    // 在这个方法中执行查看详情的操作，您可以根据需要自定义逻辑
    console.log('View details for row:', row);
    // 例如，您可以打开一个对话框显示详情
};
</script>

<style lang="scss" scoped>
.material-index {
    min-width: 700px;
    :deep(.el-tabs) {
        height: calc(100vh - 180px);
        display: flex;
        flex-direction: column;
        .el-tabs__header {
            margin-bottom: 0 !important;
        }
        .el-tabs__content,
        .el-tab-pane {
            min-height: 0;
            flex: 1;
            display: flex;
            flex-direction: column;
        }
    }
}
</style>
