<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="所属用户:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易所:" prop="exid">
          <el-input v-model.number="formData.exid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易对:" prop="pid">
          <el-input v-model.number="formData.pid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易对:" prop="symbol">
          <el-input v-model="formData.symbol" :clearable="true"  placeholder="请输入交易对" />
       </el-form-item>
        <el-form-item label="类型 对敲  趋势  跟随:" prop="type">
          <el-input v-model="formData.type" :clearable="true"  placeholder="请输入类型 对敲  趋势  跟随" />
       </el-form-item>
        <el-form-item label="内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true"  placeholder="请输入内容" />
       </el-form-item>
        <el-form-item label="addtime字段:" prop="addtime">
          <el-input v-model.number="formData.addtime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createRobCronLog,
  updateRobCronLog,
  findRobCronLog
} from '@/api/robCronLog'

defineOptions({
    name: 'RobCronLogForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: 0,
            uid: 0,
            exid: 0,
            pid: 0,
            symbol: '',
            type: '',
            content: '',
            addtime: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRobCronLog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rerobCronLog
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createRobCronLog(formData.value)
               break
             case 'update':
               res = await updateRobCronLog(formData.value)
               break
             default:
               res = await createRobCronLog(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
