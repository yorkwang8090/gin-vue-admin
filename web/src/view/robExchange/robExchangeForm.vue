<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易所名称:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入交易所名称" />
       </el-form-item>
        <el-form-item label="api接口地址:" prop="apiUrl">
          <el-input v-model="formData.apiUrl" :clearable="true"  placeholder="请输入api接口地址" />
       </el-form-item>
        <el-form-item label="扩展参数，如comexvip的uid：156:" prop="extend">
          <el-input v-model="formData.extend" :clearable="true"  placeholder="请输入扩展参数，如comexvip的uid：156" />
       </el-form-item>
        <el-form-item label="是否为跟随交易所:" prop="isfollow">
          <el-switch v-model="formData.isfollow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="添加时间:" prop="addtime">
          <el-input v-model.number="formData.addtime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createRobExchange,
  updateRobExchange,
  findRobExchange
} from '@/api/robExchange'

defineOptions({
    name: 'RobExchangeForm'
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
            title: '',
            apiUrl: '',
            extend: '',
            isfollow: false,
            addtime: 0,
            status: false,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRobExchange({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rerobExchange
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
               res = await createRobExchange(formData.value)
               break
             case 'update':
               res = await updateRobExchange(formData.value)
               break
             default:
               res = await createRobExchange(formData.value)
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
