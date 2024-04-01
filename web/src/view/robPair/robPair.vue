<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="id字段" prop="id" width="120" />
        <el-table-column align="left" label="所属用户" prop="uid" width="120" />
        <el-table-column align="left" label="所属交易所" prop="exid" width="120" />
        <el-table-column align="left" label="交易对   BTC_USDT" prop="symbol" width="120" />
        <el-table-column align="left" label="价格小数位数" prop="pricePrecision" width="120" />
        <el-table-column align="left" label="数量小数位数" prop="volumePrecision" width="120" />
        <el-table-column align="left" label="ak字段" prop="ak" width="120" />
        <el-table-column align="left" label="sk字段" prop="sk" width="120" />
        <el-table-column align="left" label="对敲" prop="beat" width="120" />
        <el-table-column align="left" label="趋势" prop="trend" width="120" />
        <el-table-column align="left" label="跟随" prop="follow" width="120" />
        <el-table-column align="left" label="盘口自动挂单撤单设置" prop="pankou" width="120" />
        <el-table-column align="left" label="添加时间" prop="addtime" width="120" />
        <el-table-column align="left" label="status字段" prop="status" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateRobPairFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="id字段:"  prop="id" >
              <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
            </el-form-item>
            <el-form-item label="所属用户:"  prop="uid" >
              <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入所属用户" />
            </el-form-item>
            <el-form-item label="所属交易所:"  prop="exid" >
              <el-input v-model.number="formData.exid" :clearable="true" placeholder="请输入所属交易所" />
            </el-form-item>
            <el-form-item label="交易对   BTC_USDT:"  prop="symbol" >
              <el-input v-model="formData.symbol" :clearable="true"  placeholder="请输入交易对   BTC_USDT" />
            </el-form-item>
            <el-form-item label="价格小数位数:"  prop="pricePrecision" >
              <el-input v-model.number="formData.pricePrecision" :clearable="true" placeholder="请输入价格小数位数" />
            </el-form-item>
            <el-form-item label="数量小数位数:"  prop="volumePrecision" >
              <el-input v-model.number="formData.volumePrecision" :clearable="true" placeholder="请输入数量小数位数" />
            </el-form-item>
            <el-form-item label="ak字段:"  prop="ak" >
              <el-input v-model="formData.ak" :clearable="true"  placeholder="请输入ak字段" />
            </el-form-item>
            <el-form-item label="sk字段:"  prop="sk" >
              <el-input v-model="formData.sk" :clearable="true"  placeholder="请输入sk字段" />
            </el-form-item>
            <el-form-item label="对敲:"  prop="beat" >
              <el-input v-model="formData.beat" :clearable="true"  placeholder="请输入对敲" />
            </el-form-item>
            <el-form-item label="趋势:"  prop="trend" >
              <el-input v-model="formData.trend" :clearable="true"  placeholder="请输入趋势" />
            </el-form-item>
            <el-form-item label="跟随:"  prop="follow" >
              <el-input v-model="formData.follow" :clearable="true"  placeholder="请输入跟随" />
            </el-form-item>
            <el-form-item label="盘口自动挂单撤单设置:"  prop="pankou" >
              <el-input v-model="formData.pankou" :clearable="true"  placeholder="请输入盘口自动挂单撤单设置" />
            </el-form-item>
            <el-form-item label="添加时间:"  prop="addtime" >
              <el-input v-model.number="formData.addtime" :clearable="true" placeholder="请输入添加时间" />
            </el-form-item>
            <el-form-item label="status字段:"  prop="status" >
              <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="id字段">
                        {{ formData.id }}
                </el-descriptions-item>
                <el-descriptions-item label="所属用户">
                        {{ formData.uid }}
                </el-descriptions-item>
                <el-descriptions-item label="所属交易所">
                        {{ formData.exid }}
                </el-descriptions-item>
                <el-descriptions-item label="交易对   BTC_USDT">
                        {{ formData.symbol }}
                </el-descriptions-item>
                <el-descriptions-item label="价格小数位数">
                        {{ formData.pricePrecision }}
                </el-descriptions-item>
                <el-descriptions-item label="数量小数位数">
                        {{ formData.volumePrecision }}
                </el-descriptions-item>
                <el-descriptions-item label="ak字段">
                        {{ formData.ak }}
                </el-descriptions-item>
                <el-descriptions-item label="sk字段">
                        {{ formData.sk }}
                </el-descriptions-item>
                <el-descriptions-item label="对敲">
                        {{ formData.beat }}
                </el-descriptions-item>
                <el-descriptions-item label="趋势">
                        {{ formData.trend }}
                </el-descriptions-item>
                <el-descriptions-item label="跟随">
                        {{ formData.follow }}
                </el-descriptions-item>
                <el-descriptions-item label="盘口自动挂单撤单设置">
                        {{ formData.pankou }}
                </el-descriptions-item>
                <el-descriptions-item label="添加时间">
                        {{ formData.addtime }}
                </el-descriptions-item>
                <el-descriptions-item label="status字段">
                    {{ formatBoolean(formData.status) }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createRobPair,
  deleteRobPair,
  deleteRobPairByIds,
  updateRobPair,
  findRobPair,
  getRobPairList
} from '@/api/robPair'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'RobPair'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        id: 0,
        uid: 0,
        exid: 0,
        symbol: '',
        pricePrecision: 0,
        volumePrecision: 0,
        ak: '',
        sk: '',
        beat: '',
        trend: '',
        follow: '',
        pankou: '',
        addtime: 0,
        status: false,
        })


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.status === ""){
        searchInfo.value.status=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getRobPairList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteRobPairFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteRobPairByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateRobPairFunc = async(row) => {
    const res = await findRobPair({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rerobPair
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteRobPairFunc = async (row) => {
    const res = await deleteRobPair({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findRobPair({ id: row.id })
  if (res.code === 0) {
    formData.value = res.data.rerobPair
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          id: 0,
          uid: 0,
          exid: 0,
          symbol: '',
          pricePrecision: 0,
          volumePrecision: 0,
          ak: '',
          sk: '',
          beat: '',
          trend: '',
          follow: '',
          pankou: '',
          addtime: 0,
          status: false,
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        id: 0,
        uid: 0,
        exid: 0,
        symbol: '',
        pricePrecision: 0,
        volumePrecision: 0,
        ak: '',
        sk: '',
        beat: '',
        trend: '',
        follow: '',
        pankou: '',
        addtime: 0,
        status: false,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createRobPair(formData.value)
                  break
                case 'update':
                  res = await updateRobPair(formData.value)
                  break
                default:
                  res = await createRobPair(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
