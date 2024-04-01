<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="操作人:" prop="userName">
          <el-input v-model="formData.userName" :clearable="true"  placeholder="请输入操作人" />
       </el-form-item>
        <el-form-item label="操作pathname:" prop="operPath">
          <el-input v-model="formData.operPath" :clearable="true"  placeholder="请输入操作pathname" />
       </el-form-item>
        <el-form-item label="提交数据:" prop="operPost">
          <el-input v-model="formData.operPost" :clearable="true"  placeholder="请输入提交数据" />
       </el-form-item>
        <el-form-item label="操作说明:" prop="content">
          <el-input v-model="formData.content" :clearable="true"  placeholder="请输入操作说明" />
       </el-form-item>
        <el-form-item label="操作ip:" prop="userIp">
          <el-input v-model="formData.userIp" :clearable="true"  placeholder="请输入操作ip" />
       </el-form-item>
        <el-form-item label="操作时间:" prop="operTime">
          <el-input v-model.number="formData.operTime" :clearable="true" placeholder="请输入" />
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
  createRobAdminLog,
  updateRobAdminLog,
  findRobAdminLog
} from '@/api/robAdminLog'

defineOptions({
    name: 'RobAdminLogForm'
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
            userName: '',
            operPath: '',
            operPost: '',
            content: '',
            userIp: '',
            operTime: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRobAdminLog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rerobAdminLog
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
               res = await createRobAdminLog(formData.value)
               break
             case 'update':
               res = await updateRobAdminLog(formData.value)
               break
             default:
               res = await createRobAdminLog(formData.value)
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
