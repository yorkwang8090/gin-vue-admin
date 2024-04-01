<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主键:" prop="id">
          <el-select v-model="formData.id" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="登陆账号:" prop="userName">
          <el-input v-model="formData.userName" :clearable="true"  placeholder="请输入登陆账号" />
       </el-form-item>
        <el-form-item label="英文昵称:" prop="userNickName">
          <el-input v-model="formData.userNickName" :clearable="true"  placeholder="请输入英文昵称" />
       </el-form-item>
        <el-form-item label="登陆密码:" prop="userPwd">
          <el-input v-model="formData.userPwd" :clearable="true"  placeholder="请输入登陆密码" />
       </el-form-item>
        <el-form-item label="通知邮件:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入通知邮件" />
       </el-form-item>
        <el-form-item label="添加时间:" prop="addtime">
          <el-input v-model.number="formData.addtime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="exptime字段:" prop="exptime">
          <el-input v-model.number="formData.exptime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="status字段:" prop="status">
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
  createRobUser,
  updateRobUser,
  findRobUser
} from '@/api/robUser'

defineOptions({
    name: 'RobUserForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const intOptions = ref([])
const formData = ref({
            id: undefined,
            userName: '',
            userNickName: '',
            userPwd: '',
            mobile: '',
            addtime: 0,
            exptime: 0,
            status: false,
        })
// 验证规则
const rule = reactive({
               id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRobUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rerobuser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    intOptions.value = await getDictFunc('int')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createRobUser(formData.value)
               break
             case 'update':
               res = await updateRobUser(formData.value)
               break
             default:
               res = await createRobUser(formData.value)
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
